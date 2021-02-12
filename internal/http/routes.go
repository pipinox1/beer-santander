package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func RegisterRoutes(r chi.Router, meetUpHandler *MeetupHandler) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/health"))
	r.Use(securityMiddleware)
	r.Route("/meetup", func(r chi.Router) {
		r.Get("/beer", meetUpHandler.calculateTotalBeers)
	})
}
