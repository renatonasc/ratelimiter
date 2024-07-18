package usecase

import (
	"context"
	"renatonasc/ratelimit/internal/infra/database"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RateLimitUseCaseTestSuite struct {
	suite.Suite
	rdb     database.DBClient
	context context.Context
}

func (suite *RateLimitUseCaseTestSuite) SetupSuite() {
	context := context.Background()

	mockRdb := database.NewRedisMock()
	// mockRdb.On("Incr", context, "123").Return(int64(1), nil).Once()
	// mockRdb.On("Incr", context, "123").Return(int64(2), nil).Once()
	// mockRdb.On("Incr", context, "123").Return(int64(3), nil).Once()

	// mockRdb.On("Expire", context, "123", mock.Anything).Return(nil).Once()

	suite.rdb = mockRdb //redis.NewClient(&redis.Options{Addr: "192.168.31.162:6379", Password: "", DB: 0})
	suite.context = context
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
