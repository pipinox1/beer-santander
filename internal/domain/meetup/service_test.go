package meetup

import (
	"beer/internal/domain/model"
	"beer/internal/tools/customerror"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCalculateBeer(t *testing.T) {

	var inputs = []struct {
		beerPerson  float64
		totalGuest  int64
		expectedBox int64
	}{
		{beerPerson: 0.75, totalGuest: 100, expectedBox: 13},
		{beerPerson: 0.75, totalGuest: 79, expectedBox: 10},
		{beerPerson: 0.75, totalGuest: 480, expectedBox: 60},
		{beerPerson: 1, totalGuest: 59, expectedBox: 10},
		{beerPerson: 1, totalGuest: 60, expectedBox: 10},
		{beerPerson: 1, totalGuest: 61, expectedBox: 11},
		{beerPerson: 2, totalGuest: 119, expectedBox: 40},
		{beerPerson: 2, totalGuest: 120, expectedBox: 40},
		{beerPerson: 2, totalGuest: 121, expectedBox: 41},
	}
	for _, input := range inputs {
		assert.Equal(t, input.expectedBox, calculateTotalBeer(input.beerPerson, input.totalGuest))
	}
}

func TestServiceGetTotalBeerErrorGettingWeather(t *testing.T) {
	mockWeatherRepository := &mockWeatherRepository{}
	s := NewService(mockWeatherRepository)
	mockWeatherRepository.Mock.On("GetCurrentWeather", mock.Anything).Return(nil, customerror.NewExternalServiceError("Service Unavailable", "weather_api", 503))
	_, err := s.GetTotalBeer(context.Background(), &model.MeetUp{TotalGuests: 100, Location: model.Location{Latitude: -13, Longitude: 12}})
	assert.NotNil(t,err)
	assert.Equal(t,"Service Unavailable",err.Error())
}

func TestServiceGetTotalBeerLatitudeGreaterThan90(t *testing.T) {
	mockWeatherRepository := &mockWeatherRepository{}
	s := NewService(mockWeatherRepository)
	_, err := s.GetTotalBeer(context.Background(), &model.MeetUp{TotalGuests: 100, Location: model.Location{Latitude: 100, Longitude: 12}})
	assert.NotNil(t,err)
	assert.Equal(t,"invalid latitude",err.Error())
}

func TestServiceGetTotalBeerLatitudeLessThan90Negative(t *testing.T) {
	mockWeatherRepository := &mockWeatherRepository{}
	s := NewService(mockWeatherRepository)
	_, err := s.GetTotalBeer(context.Background(), &model.MeetUp{TotalGuests: 100, Location: model.Location{Latitude: -100, Longitude: 12}})
	assert.NotNil(t,err)
	assert.Equal(t,"invalid latitude",err.Error())
}

func TestServiceGetTotalBeerLongitudeGreaterThan90(t *testing.T) {
	mockWeatherRepository := &mockWeatherRepository{}
	s := NewService(mockWeatherRepository)
	_, err := s.GetTotalBeer(context.Background(), &model.MeetUp{TotalGuests: 100, Location: model.Location{Latitude: 12, Longitude: 100}})
	assert.NotNil(t,err)
	assert.Equal(t,"invalid longitude",err.Error())
}

func TestServiceGetTotalBeerLongitudeLessThan90Negative(t *testing.T) {
	mockWeatherRepository := &mockWeatherRepository{}
	s := NewService(mockWeatherRepository)
	_, err := s.GetTotalBeer(context.Background(), &model.MeetUp{TotalGuests: 100, Location: model.Location{Latitude: 12, Longitude: -100}})
	assert.NotNil(t,err)
	assert.Equal(t,"invalid longitude",err.Error())
}
