package searchQueryController

import (
	"encoding/json"
	searchQueryModel "fm-scrapper-go/app/model/query"
	searchQueryService "fm-scrapper-go/app/service/query"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllSearchQueries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var queries []searchQueryModel.SearchQuery = searchQueryService.GetAllSearchQueries()

	json.NewEncoder(w).Encode(queries)

}

func GetSearchQuery(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var query searchQueryModel.SearchQuery
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	query = searchQueryService.GetSearchQuery(id)
	json.NewEncoder(w).Encode(query)

}

func CreateSearchQuery(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var query searchQueryModel.SearchQuery
	_ = json.NewDecoder(r.Body).Decode(&query)
	searchQueryService.CreateSearchQuery(query)
	json.NewEncoder(w).Encode("SearchQuery created")

}

func DeleteSearchQuery(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	searchQueryService.DeleteSearchQuery(id)
	json.NewEncoder(w).Encode("SearchQuery deleted")
}
