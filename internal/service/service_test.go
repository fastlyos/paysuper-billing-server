package service

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"github.com/ProtocolONE/rabbitmq/pkg"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/protobuf/ptypes"
	"github.com/paysuper/paysuper-billing-server/internal/config"
	"github.com/paysuper/paysuper-billing-server/internal/mock"
	"github.com/paysuper/paysuper-billing-server/pkg"
	"github.com/paysuper/paysuper-billing-server/pkg/proto/billing"
	"github.com/paysuper/paysuper-billing-server/pkg/proto/grpc"
	mongodb "github.com/paysuper/paysuper-database-mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"testing"
	"time"
)

type BillingServiceTestSuite struct {
	suite.Suite
	db      *mongodb.Source
	log     *zap.Logger
	cfg     *config.Config
	exCh    chan bool
	service *Service
	cache   CacheInterface

	project *billing.Project
}

func Test_BillingService(t *testing.T) {
	suite.Run(t, new(BillingServiceTestSuite))
}

func (suite *BillingServiceTestSuite) SetupTest() {
	cfg, err := config.NewConfig()
	if err != nil {
		suite.FailNow("Config load failed", "%v", err)
	}

	cfg.AccountingCurrency = "RUB"
	suite.cfg = cfg

	db, err := mongodb.NewDatabase()
	if err != nil {
		suite.FailNow("Database connection failed", "%v", err)
	}

	suite.db = db

	suite.log, err = zap.NewProduction()
	if err != nil {
		suite.FailNow("Logger initialization failed", "%v", err)
	}

	broker, err := rabbitmq.NewBroker(cfg.BrokerAddress)
	if err != nil {
		suite.FailNow("Creating RabbitMQ publisher failed", "%v", err)
	}

	rub := &billing.Currency{
		CodeInt:  643,
		CodeA3:   "RUB",
		Name:     &billing.Name{Ru: "Российский рубль", En: "Russian ruble"},
		IsActive: true,
	}

	usd := &billing.Currency{
		CodeInt:  840,
		CodeA3:   "USD",
		Name:     &billing.Name{Ru: "Доллар США", En: "US Dollar"},
		IsActive: true,
	}
	uah := &billing.Currency{
		CodeInt:  980,
		CodeA3:   "UAH",
		Name:     &billing.Name{Ru: "Украинская гривна", En: "Ukrainian Hryvnia"},
		IsActive: true,
	}
	if err := InitTestCurrency(db, []interface{}{rub, usd, uah}); err != nil {
		suite.FailNow("Insert currency test data failed", "%v", err)
	}

	redisdb := mock.NewTestRedis()
	suite.service = NewBillingService(
		db,
		cfg,
		mock.NewGeoIpServiceTestOk(),
		mock.NewRepositoryServiceOk(),
		mock.NewTaxServiceOkMock(),
		broker,
		nil,
		NewCacheRedis(redisdb),
		nil,
	)

	if err := suite.service.Init(); err != nil {
		suite.FailNow("Billing service initialization failed", "%v", err)
	}

	ps := &billing.PaymentSystem{
		Id:                 bson.NewObjectId().Hex(),
		Name:               "CardPay",
		AccountingCurrency: rub,
		AccountingPeriod:   "every-day",
		Country:            "",
		IsActive:           true,
	}

	country := &billing.Country{
		IsoCodeA2:       "RU",
		Region:          "Russia",
		Currency:        "RUB",
		PaymentsAllowed: true,
		ChangeAllowed:   true,
		VatEnabled:      true,
		PriceGroupId:    "",
		VatCurrency:     "RUB",
	}

	pmBankCard := &billing.PaymentMethod{
		Id:               bson.NewObjectId().Hex(),
		Name:             "Bank card",
		Group:            "BANKCARD",
		MinPaymentAmount: 100,
		MaxPaymentAmount: 15000,
		Currencies:       []int32{643, 840, 980},
		ExternalId:       "BANKCARD",
		TestSettings: &billing.PaymentMethodParams{
			TerminalId:     "15985",
			Secret:         "A1tph4I6BD0f",
			SecretCallback: "0V1rJ7t4jCRv",
		},
		Type:            "bank_card",
		IsActive:        true,
		PaymentSystemId: ps.Id,
	}

	date, err := ptypes.TimestampProto(time.Now().Add(time.Hour * -360))
	assert.NoError(suite.T(), err, "Generate merchant date failed")

	merchant := &billing.Merchant{
		Id:      bson.NewObjectId().Hex(),
		Name:    "Unit test",
		Country: country.IsoCodeA2,
		Zip:     "190000",
		City:    "St.Petersburg",
		Contacts: &billing.MerchantContact{
			Authorized: &billing.MerchantContactAuthorized{
				Name:     "Unit Test",
				Email:    "test@unit.test",
				Phone:    "123456789",
				Position: "Unit Test",
			},
			Technical: &billing.MerchantContactTechnical{
				Name:  "Unit Test",
				Email: "test@unit.test",
				Phone: "123456789",
			},
		},
		Banking: &billing.MerchantBanking{
			Currency: rub,
			Name:     "Bank name",
		},
		IsVatEnabled:              true,
		IsCommissionToUserEnabled: true,
		Status:                    pkg.MerchantStatusDraft,
		LastPayout: &billing.MerchantLastPayout{
			Date:   date,
			Amount: 999999,
		},
		IsSigned: true,
		PaymentMethods: map[string]*billing.MerchantPaymentMethod{
			pmBankCard.Id: {
				PaymentMethod: &billing.MerchantPaymentMethodIdentification{
					Id:   pmBankCard.Id,
					Name: pmBankCard.Name,
				},
				Commission: &billing.MerchantPaymentMethodCommissions{
					Fee: 2.5,
					PerTransaction: &billing.MerchantPaymentMethodPerTransactionCommission{
						Fee:      30,
						Currency: rub.CodeA3,
					},
				},
				Integration: &billing.MerchantPaymentMethodIntegration{
					TerminalId:       "1234567890",
					TerminalPassword: "0987654321",
					Integrated:       true,
				},
				IsActive: true,
			},
		},
	}

	projectDefault := &billing.Project{
		Id:                       bson.NewObjectId().Hex(),
		CallbackCurrency:         rub.CodeA3,
		CallbackProtocol:         "default",
		LimitsCurrency:           rub.CodeA3,
		MaxPaymentAmount:         15000,
		MinPaymentAmount:         1,
		Name:                     map[string]string{"en": "test project 1"},
		IsProductsCheckout:       true,
		AllowDynamicRedirectUrls: true,
		SecretKey:                "test project 1 secret key",
		Status:                   pkg.ProjectStatusInProduction,
		MerchantId:               merchant.Id,
	}
	projectXsolla := &billing.Project{
		Id:                 bson.NewObjectId().Hex(),
		MerchantId:         bson.NewObjectId().Hex(),
		CallbackCurrency:   rub.CodeA3,
		CallbackProtocol:   "xsolla",
		LimitsCurrency:     rub.CodeA3,
		MaxPaymentAmount:   15000,
		MinPaymentAmount:   0,
		Name:               map[string]string{"en": "test project 2"},
		IsProductsCheckout: true,
		SecretKey:          "test project 2 secret key",
		Status:             pkg.ProjectStatusInProduction,
	}
	projectCardpay := &billing.Project{
		Id:                 bson.NewObjectId().Hex(),
		MerchantId:         bson.NewObjectId().Hex(),
		CallbackCurrency:   rub.CodeA3,
		CallbackProtocol:   "cardpay",
		LimitsCurrency:     rub.CodeA3,
		MaxPaymentAmount:   15000,
		MinPaymentAmount:   0,
		Name:               map[string]string{"en": "test project 3"},
		IsProductsCheckout: true,
		SecretKey:          "test project 3 secret key",
		Status:             pkg.ProjectStatusInProduction,
	}

	rate := []*billing.CurrencyRate{
		{
			CurrencyFrom: 980,
			CurrencyTo:   643,
			Rate:         0.411128442,
			Date:         ptypes.TimestampNow(),
			IsActive:     true,
		},
		{
			CurrencyFrom: 980,
			CurrencyTo:   980,
			Rate:         27.13085922,
			Date:         ptypes.TimestampNow(),
			IsActive:     true,
		},
		{
			CurrencyFrom: 980,
			CurrencyTo:   978,
			Rate:         30.96446748,
			Date:         ptypes.TimestampNow(),
			IsActive:     true,
		},
		{
			CurrencyFrom: 840,
			CurrencyTo:   980,
			Rate:         0.034680066,
			Date:         ptypes.TimestampNow(),
			IsActive:     true,
		},
		{
			CurrencyFrom: 840,
			CurrencyTo:   643,
			Rate:         0.01469893,
			Date:         ptypes.TimestampNow(),
			IsActive:     true,
		},
		{
			CurrencyFrom: 840,
			CurrencyTo:   840,
			Rate:         1.00000000,
			Date:         ptypes.TimestampNow(),
			IsActive:     true,
		},
		{
			CurrencyFrom: 643,
			CurrencyTo:   840,
			Rate:         64.01146400,
			Date:         ptypes.TimestampNow(),
			IsActive:     true,
		},
		{
			CurrencyFrom: 643,
			CurrencyTo:   643,
			Rate:         1,
			Date:         ptypes.TimestampNow(),
			IsActive:     true,
		},
		{
			CurrencyFrom: 643,
			CurrencyTo:   980,
			Rate:         2.2885792,
			Date:         ptypes.TimestampNow(),
			IsActive:     true,
		},
	}

	pmQiwi := &billing.PaymentMethod{
		Id:               bson.NewObjectId().Hex(),
		Name:             "Qiwi",
		Group:            "QIWI",
		MinPaymentAmount: 0,
		MaxPaymentAmount: 0,
		Currencies:       []int32{643, 840, 980},
		ExternalId:       "QIWI",
		TestSettings: &billing.PaymentMethodParams{
			TerminalId: "15993",
		},
		Type:            "ewallet",
		IsActive:        true,
		PaymentSystemId: ps.Id,
	}
	pmBitcoin := &billing.PaymentMethod{
		Id:               bson.NewObjectId().Hex(),
		Name:             "Bitcoin",
		Group:            "BITCOIN",
		MinPaymentAmount: 0,
		MaxPaymentAmount: 0,
		Currencies:       []int32{643, 840, 980},
		ExternalId:       "BITCOIN",
		TestSettings: &billing.PaymentMethodParams{
			TerminalId: "16007",
		},
		Type:            "crypto",
		IsActive:        true,
		PaymentSystemId: ps.Id,
	}

	commissionStartDate, err := ptypes.TimestampProto(time.Now().Add(time.Minute * -10))

	if err != nil {
		suite.FailNow("Commission start date conversion failed", "%v", err)
	}

	commissions := []interface{}{
		&billing.Commission{
			PaymentMethodId:         pmBankCard.Id,
			ProjectId:               projectDefault.Id,
			PaymentMethodCommission: 1,
			PspCommission:           2,
			TotalCommissionToUser:   1,
			StartDate:               commissionStartDate,
		},
		&billing.Commission{
			PaymentMethodId:         pmQiwi.Id,
			ProjectId:               projectDefault.Id,
			PaymentMethodCommission: 1,
			PspCommission:           2,
			TotalCommissionToUser:   2,
			StartDate:               commissionStartDate,
		},
		&billing.Commission{
			PaymentMethodId:         pmBitcoin.Id,
			ProjectId:               projectDefault.Id,
			PaymentMethodCommission: 1,
			PspCommission:           2,
			TotalCommissionToUser:   3,
			StartDate:               commissionStartDate,
		},
		&billing.Commission{
			PaymentMethodId:         pmBankCard.Id,
			ProjectId:               projectXsolla.Id,
			PaymentMethodCommission: 1,
			PspCommission:           2,
			TotalCommissionToUser:   1,
			StartDate:               commissionStartDate,
		},
		&billing.Commission{
			PaymentMethodId:         pmQiwi.Id,
			ProjectId:               projectXsolla.Id,
			PaymentMethodCommission: 1,
			PspCommission:           2,
			TotalCommissionToUser:   2,
			StartDate:               commissionStartDate,
		},
		&billing.Commission{
			PaymentMethodId:         pmBitcoin.Id,
			ProjectId:               projectXsolla.Id,
			PaymentMethodCommission: 1,
			PspCommission:           2,
			TotalCommissionToUser:   3,
			StartDate:               commissionStartDate,
		},
		&billing.Commission{
			PaymentMethodId:         pmBankCard.Id,
			ProjectId:               projectCardpay.Id,
			PaymentMethodCommission: 1,
			PspCommission:           2,
			TotalCommissionToUser:   1,
			StartDate:               commissionStartDate,
		},
		&billing.Commission{
			PaymentMethodId:         pmQiwi.Id,
			ProjectId:               projectCardpay.Id,
			PaymentMethodCommission: 1,
			PspCommission:           2,
			TotalCommissionToUser:   2,
			StartDate:               commissionStartDate,
		},
		&billing.Commission{
			PaymentMethodId:         pmBitcoin.Id,
			ProjectId:               projectCardpay.Id,
			PaymentMethodCommission: 1,
			PspCommission:           2,
			TotalCommissionToUser:   3,
			StartDate:               commissionStartDate,
		},
	}

	err = suite.db.Collection(collectionCommission).Insert(commissions...)

	if err != nil {
		suite.FailNow("Insert commission test data failed", "%v", err)
	}

	projects := []*billing.Project{
		projectDefault,
		projectXsolla,
		projectCardpay,
	}

	pms := []*billing.PaymentMethod{pmBankCard, pmQiwi, pmBitcoin}
	if err := suite.service.paymentMethod.MultipleInsert(pms); err != nil {
		suite.FailNow("Insert payment methods test data failed", "%v", err)
	}

	if err := suite.service.merchant.Insert(merchant); err != nil {
		suite.FailNow("Insert merchant test data failed", "%v", err)
	}

	if err := suite.service.country.Insert(country); err != nil {
		suite.FailNow("Insert country test data failed", "%v", err)
	}

	if err = suite.service.currencyRate.MultipleInsert(rate); err != nil {
		suite.FailNow("Insert rates test data failed", "%v", err)
	}

	if err := suite.service.project.MultipleInsert(projects); err != nil {
		suite.FailNow("Insert project test data failed", "%v", err)
	}

	suite.exCh = make(chan bool, 1)
	suite.project = projectDefault
}

func (suite *BillingServiceTestSuite) TearDownTest() {
	if err := suite.db.Drop(); err != nil {
		suite.FailNow("Database deletion failed", "%v", err)
	}

	suite.db.Close()
}

func (suite *BillingServiceTestSuite) TestNewBillingService() {
	redisdb := mock.NewTestRedis()
	suite.cache = NewCacheRedis(redisdb)
	service := NewBillingService(suite.db, suite.cfg, nil, nil, nil, nil, nil, suite.cache, nil)

	err := service.Init()
	assert.Nil(suite.T(), err)
}

func (suite *BillingServiceTestSuite) TestBillingService_AccountingCurrencyInitError() {
	cfg, err := config.NewConfig()

	assert.NoError(suite.T(), err)

	cfg.AccountingCurrency = "AUD"
	suite.cache = NewCacheRedis(mock.NewTestRedis())
	service := NewBillingService(suite.db, cfg, nil, nil, nil, nil, nil, suite.cache, nil)

	err = service.Init()
	assert.Error(suite.T(), err)
}

func (suite *BillingServiceTestSuite) TestBillingService_IsProductionEnvironment() {
	redisdb := mock.NewTestRedis()
	suite.cache = NewCacheRedis(redisdb)
	service := NewBillingService(suite.db, suite.cfg, nil, nil, nil, nil, nil, suite.cache, nil)

	err := service.Init()
	assert.Nil(suite.T(), err)
}

func (suite *BillingServiceTestSuite) TestBillingService_CheckProjectRequestSignature_Ok() {
	req := &grpc.CheckProjectRequestSignatureRequest{
		Body:      `{"field1": "val1", "field2": "val2", "field3": "val3"}`,
		ProjectId: suite.project.Id,
	}
	rsp := &grpc.CheckProjectRequestSignatureResponse{}

	hashString := req.Body + suite.project.SecretKey
	h := sha512.New()
	h.Write([]byte(hashString))

	req.Signature = hex.EncodeToString(h.Sum(nil))

	err := suite.service.CheckProjectRequestSignature(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), pkg.ResponseStatusOk, rsp.Status)
}

func (suite *BillingServiceTestSuite) TestBillingService_CheckProjectRequestSignature_ProjectNotFound_Error() {
	req := &grpc.CheckProjectRequestSignatureRequest{
		Body:      `{"field1": "val1", "field2": "val2", "field3": "val3"}`,
		ProjectId: bson.NewObjectId().Hex(),
	}
	rsp := &grpc.CheckProjectRequestSignatureResponse{}

	err := suite.service.CheckProjectRequestSignature(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), pkg.ResponseStatusBadData, rsp.Status)
	assert.Equal(suite.T(), orderErrorProjectNotFound, rsp.Message)
}

func (suite *BillingServiceTestSuite) TestBillingService_CheckProjectRequestSignature_IncorrectSignature_Error() {
	req := &grpc.CheckProjectRequestSignatureRequest{
		Body:      `{"field1": "val1", "field2": "val2", "field3": "val3"}`,
		ProjectId: suite.project.Id,
	}
	rsp := &grpc.CheckProjectRequestSignatureResponse{}

	hashString := req.Body + "some_random_string"
	h := sha512.New()
	h.Write([]byte(hashString))

	req.Signature = hex.EncodeToString(h.Sum(nil))

	err := suite.service.CheckProjectRequestSignature(context.TODO(), req, rsp)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), pkg.ResponseStatusBadData, rsp.Status)
	assert.Equal(suite.T(), orderErrorSignatureInvalid, rsp.Message)
}
