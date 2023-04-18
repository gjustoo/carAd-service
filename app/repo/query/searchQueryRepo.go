package searchQueryRepo

import (
	"context"
	searchQueryModel "fm-scrapper-go/app/model/query"
	"fm-scrapper-go/app/repo/db"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllQueries() ([]searchQueryModel.SearchQuery, error) {

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
