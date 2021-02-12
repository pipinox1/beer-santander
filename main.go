package main

import (
	"beer/internal/domain/meetup"
	http2 "beer/internal/http"
	"beer/internal/repository"
	"beer/internal/tools/restclient"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strings"
	"time"
)

func main()  {

	weatherRepo := repository.NewRepository(restclient.NewRestClient(2 * time.Second))
	meetupService := meetup.NewService(&weatherRepo)
	meetupHandler := http2.NewMeetupHandler(&meetupService)

	//Start App
	router := chi.NewRouter()
	http2.RegisterRoutes(router, meetupHandler)

	log.Print("Availability routes")
	for _, a := range router.Routes() {
		for _, b := range a.SubRoutes.Routes() {
			log.Print(fmt.Sprint(strings.ReplaceAll(a.Pattern, "/*", ""), b.Pattern))
		}

	}
	port := ":8020"
	log.Print(fmt.Sprint("Starting server at port", port))
	log.Fatal(http.ListenAndServe(port, router))
}
