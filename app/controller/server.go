package controller

import (
	carAdController "fm-scrapper-go/app/controller/carAd"
	searchQueryController "fm-scrapper-go/app/controller/query"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func initCarAdHandlers() {

	router.HandleFunc("/api/carAd/", carAdController.GetAllCarAds).Methods("GET")
	router.HandleFunc("/api/carAd/{id}", carAdController.GetCarAd).Methods("GET")
	router.HandleFunc("/api/carAd/", carAdController.CreateCarAd).Methods("POST")
	router.HandleFunc("/api/carAd/{id}", carAdController.DeleteCarAd).Methods("DELETE")
}

func initSearchQueryHandlers() {

	router.HandleFunc("/api/searchQuery/", searchQueryController.GetAllSearchQueries).Methods("GET")
	router.HandleFunc("/api/searchQuery/{id}", searchQueryController.GetSearchQuery).Methods("GET")
	router.HandleFunc("/api/searchQuery/", searchQueryController.CreateSearchQuery).Methods("POST")
	router.HandleFunc("/api/searchQuery/{id}", searchQueryController.DeleteSearchQuery).Methods("DELETE")
}

func Start() {
	router = mux.NewRouter()

	initCarAdHandlers()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
