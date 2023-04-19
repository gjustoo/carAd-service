package queryService

import (
	searchQueryModel "fm-scrapper-go/app/model/query"
	searchQueryRepo "fm-scrapper-go/app/repo/query"
	"log"
)

func GetAllSearchQueries() []searchQueryModel.SearchQuery {

	queries, err := searchQueryRepo.GetAllSearchQueries()
	if err != nil {
		log.Println(err)
		return nil
	}
	return queries
}

func GetSearchQuery(id uint64) searchQueryModel.SearchQuery {

	query, err := searchQueryRepo.GetSearchQuery(id)
	if err != nil {
		log.Println("There was an error retrieving query : ", id, err)
		return searchQueryModel.SearchQuery{}
	}
	return query

}

func CreateSearchQuery(post searchQueryModel.SearchQuery) {

	err := searchQueryRepo.CreateSearchQuery(post)
	if err != nil {
		log.Println("There was an error creating query : ", post, err)
		return
	}

	log.Print("Query created : ", post)
}

func DeleteSearchQuery(id uint64) {

	err := searchQueryRepo.DeleteSearchQuery(id)
	if err != nil {
		log.Println("There was an error deleting query : ", id, err)
		return
	}

	log.Print("Query deleted : ", id)
}
