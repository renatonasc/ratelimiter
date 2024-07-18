package usecase

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type RateLimitUseCaseTestSuite struct {
	suite.Suite
	rdb     *redis.Client
	context context.Context
}

func (suite *RateLimitUseCaseTestSuite) SetupSuite() {

	suite.rdb = redis.NewClient(&redis.Options{Addr: "192.168.31.162:6379", Password: "", DB: 0})
	suite.context = context.Background()
}

func TestRateLimitUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(RateLimitUseCaseTestSuite))
}

func (suite *RateLimitUseCaseTestSuite) TestRateLimitUseCase_Execute() {
	rl := NewRateLimitUseCase(2, 2, suite.rdb)
	dto := RateLimitInputDTO{Key: "123", Context: suite.context}
	canAccess, err := rl.Execute(dto)
	suite.NoError(err)
	suite.True(canAccess)

	canAccess, err = rl.Execute(dto)
	suite.NoError(err)
	suite.True(canAccess)

	canAccess, err = rl.Execute(dto)
	suite.NoError(err)
	suite.False(canAccess)

}
