package service

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/paysuper/paysuper-billing-server/internal/config"
	"github.com/paysuper/paysuper-billing-server/internal/database"
	"github.com/paysuper/paysuper-billing-server/pkg"
	"github.com/paysuper/paysuper-billing-server/pkg/proto/billing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"testing"
)

type CountryTestSuite struct {
	suite.Suite
	service *Service
	log     *zap.Logger
	country *billing.Country
}

func Test_Country(t *testing.T) {
	suite.Run(t, new(CountryTestSuite))
}

func (suite *CountryTestSuite) SetupTest() {
	cfg, err := config.NewConfig()
	if err != nil {
		suite.FailNow("Config load failed", "%v", err)
	}
	cfg.AccountingCurrency = "RUB"

	settings := database.Connection{
		Host:     cfg.MongoHost,
		Database: cfg.MongoDatabase,
		User:     cfg.MongoUser,
		Password: cfg.MongoPassword,
	}

	db, err := database.NewDatabase(settings)

	if err != nil {
		suite.FailNow("Database connection failed", "%v", err)
	}

	rub := &billing.Currency{
		CodeInt:  643,
		CodeA3:   "RUB",
		Name:     &billing.Name{Ru: "Российский рубль", En: "Russian ruble"},
		IsActive: true,
	}
	currency := []interface{}{rub}
	err = db.Collection(pkg.CollectionCurrency).Insert(currency...)
	if err != nil {
		suite.FailNow("Insert currency test data failed", "%v", err)
	}

	suite.country = &billing.Country{
		CodeInt:  643,
		CodeA2:   "RU",
		CodeA3:   "RUS",
		Name:     &billing.Name{Ru: "Россия", En: "Russia (Russian Federation)"},
		IsActive: true,
	}
	err = db.Collection(pkg.CollectionCountry).Insert(suite.country)
	assert.NoError(suite.T(), err, "Insert country test data failed")

	suite.log, err = zap.NewProduction()

	if err != nil {
		suite.FailNow("Logger initialization failed", "%v", err)
	}

	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        cfg.CacheRedis.Address,
		Password:     cfg.CacheRedis.Password,
		MaxRetries:   cfg.CacheRedis.MaxRetries,
		MaxRedirects: cfg.CacheRedis.MaxRedirects,
		PoolSize:     cfg.CacheRedis.PoolSize,
	})

	suite.service = NewBillingService(db, cfg, make(chan bool, 1), nil, nil, nil, nil, nil, NewCacheRedis(redisdb))
	err = suite.service.Init()

	if err != nil {
		suite.FailNow("Billing service initialization failed", "%v", err)
	}
}

func (suite *CountryTestSuite) TearDownTest() {
	if err := suite.service.db.Drop(); err != nil {
		suite.FailNow("Database deletion failed", "%v", err)
	}

	suite.service.db.Close()
}

func (suite *CountryTestSuite) TestCountry_GetAll() {
	c := suite.service.country.GetAll()

	assert.NotNil(suite.T(), c)
}

func (suite *CountryTestSuite) TestCountry_GetCountryByCodeA2_Ok() {
	c, err := suite.service.country.GetCountryByCodeA2("RU")

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), c)
	assert.Equal(suite.T(), suite.country.CodeInt, c.CodeInt)
}

func (suite *CountryTestSuite) TestCountry_GetCountryByCodeA2_NotFound() {
	_, err := suite.service.country.GetCountryByCodeA2("AAA")

	assert.Error(suite.T(), err)
	assert.Errorf(suite.T(), err, fmt.Sprintf(errorNotFound, pkg.CollectionCountry))
}
