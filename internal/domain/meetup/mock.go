package meetup

import (
	"beer/internal/domain/model"
	"context"
	"github.com/stretchr/testify/mock"
	"time"
)

type mockWeatherRepository struct {
	mock.Mock
}

func (r *mockWeatherRepository) GetWeather(ctx context.Context, latitude float64, longitude float64, date time.Time) (*model.Weather, error) {
	args := r.Mock.Called(mock.Anything)
	idempotencyKey, _ := args.Get(0).(*model.Weather)
	return idempotencyKey, args.Error(1)
}
