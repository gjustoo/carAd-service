package carAdController

import (
	"encoding/json"
	carAdModel "fm-scrapper-go/app/model/carAd"
	carAdService "fm-scrapper-go/app/service/carAd"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllCarAds(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var carAds []carAdModel.CarAd

	carAds = carAdService.GetAllCarAds()

	json.NewEncoder(w).Encode(carAds)

}

func GetCarAd(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var carAd carAdModel.CarAd
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	carAd = carAdService.GetCarAd(id)
	json.NewEncoder(w).Encode(carAd)
}

func CreateCarAd(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var carAd carAdModel.CarAd
	_ = json.NewDecoder(r.Body).Decode(&carAd)
	carAdService.CreateCarAd(carAd)
	json.NewEncoder(w).Encode("CarAd created")
}

func DeleteCarAd(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	carAdService.DeleteCarAd(id)
	json.NewEncoder(w).Encode("CarAd deleted")

}
