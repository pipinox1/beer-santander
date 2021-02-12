package meetup

import (
	"beer/internal/domain/model"
	"beer/internal/tools/customerror"
	"context"
	"math"
	"time"
)

const unitOfBeer = 6

type WeatherRepository interface {
	GetWeather(ctx context.Context, latitude float64, longitude float64, date time.Time) (*model.Weather, error)
}

type Service struct {
	WeatherRepo WeatherRepository
}

func NewService(weatherRepository WeatherRepository) Service {
	return Service{WeatherRepo: weatherRepository}
}

func (s *Service) GetTotalBeer(ctx context.Context, event *model.MeetUp) (int64, error) {
	err := validateBeerEvent(event)
	if err != nil {
		return 0, err
	}
	weather, err := s.WeatherRepo.GetWeather(ctx, event.Location.Latitude, event.Location.Longitude,event.Date)
	if err != nil {
		return 0, err
	}

	if len(weather.Forecast.Forecastday) != 1 {
		return 0, customerror.NewBusinessError("couldn't calculate number of beer please specify other latitude-longitude")
	}

	day := weather.Forecast.Forecastday[0].Day
	var beerPerPerson float64
	if day.MaxtempC < 20 {
		beerPerPerson = 0.75
	} else if day.MaxtempC > 20 && day.MaxtempC < 24 {
		beerPerPerson = 1
	} else {
		beerPerPerson = 2
	}
	return calculateTotalBeer(beerPerPerson, event.TotalGuests), nil
}

func calculateTotalBeer(beerPerPerson float64, totalGuest int64) int64 {
	estimatedBeer := beerPerPerson * float64(totalGuest)
	boxOfBeers := estimatedBeer / unitOfBeer
	truncatedBoxBeers := math.Trunc(boxOfBeers)
	if boxOfBeers-truncatedBoxBeers > 0 {
		return int64(truncatedBoxBeers + 1)
	}
	return int64(truncatedBoxBeers)
}

func validateBeerEvent(event *model.MeetUp) error {
	if event == nil {
		return customerror.NewBusinessError("event needed")
	}
	if event.Location.Longitude > 90 || event.Location.Longitude < -90 {
		return customerror.NewBusinessError("invalid longitude")
	}
	if event.Location.Latitude > 90 || event.Location.Latitude < -90 {
		return customerror.NewBusinessError("invalid latitude")
	}
	if event.TotalGuests <= 0 {
		return customerror.NewBusinessError("total guest must be positive")
	}
	//should be between today and next 10 day.
	if (time.Now().Add(time.Hour * 24 * 10)).Sub(event.Date).Hours() < 0 {
		return customerror.NewBusinessError("total guest must be positive")
	}
	return nil
}
