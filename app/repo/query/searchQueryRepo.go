package searchQueryRepo

import (
	"context"
	searchQueryModel "fm-scrapper-go/app/model/query"
	"fm-scrapper-go/app/repo/db"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllSearchQueries() ([]searchQueryModel.SearchQuery, error) {

	var queries []searchQueryModel.SearchQuery

	session, err := db.GetDB().StartSession()
	if err != nil {
		return queries, err
	}

	defer session.EndSession(context.Background())

	cursor, err := session.Client().Database("db_name").Collection("searchQueries").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var query searchQueryModel.SearchQuery
		if err := cursor.Decode(&query); err != nil {
			fmt.Print("ERROR", err.Error())
			return nil, err
		}
		queries = append(queries, query)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return queries, nil

}

func GetSearchQuery(id uint64) (searchQueryModel.SearchQuery, error) {

	var query searchQueryModel.SearchQuery

	session, err := db.GetDB().StartSession()
	if err != nil {
		return searchQueryModel.SearchQuery{}, err
	}

	defer session.EndSession(context.Background())

	err = session.Client().Database("db_name").Collection("searchQueries").FindOne(context.Background(), bson.M{"id": id}).Decode(&query)
	if err != nil {
		return searchQueryModel.SearchQuery{}, err
	}

	return query, nil
}

func CreateSearchQuery(post searchQueryModel.SearchQuery) error {
	session, err := db.GetDB().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	_, err = session.Client().Database("db_name").Collection("searchQueries").InsertOne(context.Background(), post)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSearchQuery(id uint64) error {
	session, err := db.GetDB().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	_, err = session.Client().Database("db_name").Collection("searchQueries").DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return err
	}

	return nil
}
