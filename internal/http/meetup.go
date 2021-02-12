package http

import (
	"beer/internal/domain/model"
	"beer/internal/tools/customerror"
	"context"
	"net/http"
	"strconv"
	"time"
)

type Service interface {
	GetTotalBeer(ctx context.Context, event *model.MeetUp) (int64, error)
}

const (
	latitudeParam    = "latitude"
	longitudeParam   = "longitude"
	totalPersonParam = "guests"
	dateParam        = "date"
)

type MeetupHandler struct {
	service Service
}

func NewMeetupHandler(service Service) *MeetupHandler {
	return &MeetupHandler{
		service: service,
	}
}

func (sh *MeetupHandler) calculateTotalBeers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Header.Get("X-Is-Admin") != "true" {
		ErrorResponse(w, customerror.NewUnauthorizedError("invalid user"))
	}
	meetUp, err := buildMeetUp(r)
	if err != nil {
		return
	}
	totalInt, err := sh.service.GetTotalBeer(ctx, meetUp)
	if err != nil {
		ErrorResponse(w, err)
		return
	}
	WebResponse(w, 200, totalInt)
}

func buildMeetUp(r *http.Request) (*model.MeetUp, error) {
	meetUp := &model.MeetUp{}
	latitude, err := strconv.ParseFloat(r.URL.Query().Get(latitudeParam), 64)
	if err != nil {
		return nil, err
	}

	longitude, err := strconv.ParseFloat(r.URL.Query().Get(longitudeParam), 64)
	if err != nil {
		return nil, err
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get(dateParam))
	if err != nil {
		return nil, customerror.NewBusinessError("invalid_date_format")
	}
	totalGuest, err := strconv.ParseInt(r.URL.Query().Get(totalPersonParam), 10, 64)
	if err != nil {
		return nil, err
	}

	meetUp.Location = model.Location{Longitude: longitude, Latitude: latitude}
	meetUp.TotalGuests = totalGuest
	meetUp.Date = date
	return meetUp, nil
}
