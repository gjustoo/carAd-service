package carAdRepo

import (
	"context"
	carAdModel "fm-scrapper-go/app/model/carAd"
	"fm-scrapper-go/app/repo/db"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCarAds() ([]carAdModel.CarAd, error) {

	var posts []carAdModel.CarAd

	session, err := db.GetDB().StartSession()
	if err != nil {
		return posts, err
	}

	defer session.EndSession(context.Background())

	cursor, err := session.Client().Database("db_name").Collection("fbMarketplace").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var post carAdModel.CarAd
		if err := cursor.Decode(&post); err != nil {
			fmt.Print("ERROR", err.Error())
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return posts, nil

}

func GetCarAd(id uint64) (carAdModel.CarAd, error) {

	var post carAdModel.CarAd

	session, err := db.GetDB().StartSession()
	if err != nil {
		return post, err
	}

	defer session.EndSession(context.Background())

	err = session.Client().Database("db_name").Collection("fbMarketplace").FindOne(context.Background(), bson.M{"id": id}).Decode(&post)
	if err != nil {
		return post, err
	}

	return post, nil
}

func CreateCarAd(post carAdModel.CarAd) error {

	session, err := db.GetDB().StartSession()
	if err != nil {
		return err
	}

	defer session.EndSession(context.Background())

	_, err = session.Client().Database("db_name").Collection("fbMarketplace").InsertOne(context.Background(), post)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCarAd(id uint64) error {

	session, err := db.GetDB().StartSession()
	if err != nil {
		return err
	}

	defer session.EndSession(context.Background())

	_, err = session.Client().Database("db_name").Collection("fbMarketplace").DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return err
	}

	return nil
}
