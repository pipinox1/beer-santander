package repository

import (
	"beer/internal/domain/model"
	"beer/internal/tools/restclient"
	"context"
	"fmt"
	"time"
)

var requestCurrentWeather = "https://weatherapi-com.p.rapidapi.com/forecast.json?q=%v,%v&dt=%s"

func NewRepository(restClient restclient.RestClient) Repository {
	return Repository{restClient}
}

type Repository struct {
	restclient.RestClient
}

func (r *Repository) GetWeather(ctx context.Context, latitude float64, longitude float64, date time.Time) (*model.Weather, error) {
	url := fmt.Sprintf(requestCurrentWeather, longitude, latitude, date.Format("2020-10-31"))
	weather := &model.Weather{}
	err := r.DoGet(ctx, url, weather,
		restclient.Header{Key: "x-rapidapi-key", Value: "5efcad2b8amsh00b118ca3e1a1b7p1ee70bjsn47d4746b606c"},
		restclient.Header{Key: "x-rapidapi-host", Value: "weatherapi-com.p.rapidapi.com"})
	if err != nil {
		return nil, err
	}
	return weather, nil
}
