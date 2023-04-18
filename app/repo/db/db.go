package db

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/lib/pq"
)

// declare a Db object, where we can use throughout the model package
// so in blog.go, we have access to this object
var db *mongo.Client

// a struct to hold all the Db connection information
type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetDB() *mongo.Client {
	return db
}

func Init() {
	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err.Error())
		return
	}

	connInfo := connection{
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		User:     os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PWD"),
		DBName:   os.Getenv("MONGO_DB"),
	}

	// try to open our postgresql connection with our connection info
	clientOptions := options.Client().ApplyURI(connToString(connInfo))
	clientOptions.SetAuth(options.Credential{AuthSource: "admin", Username: connInfo.User, Password: connInfo.Password})

	db, err = mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		fmt.Printf("Error connecting to the DB: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB is open\n")
	}

	// check if we can ping our DB
	err = db.Ping(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error could not ping database: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB pinged successfully\n")
	}

}

// Take our connection struct and convert to a string for our Db connection info
func connToString(info connection) string {

	return fmt.Sprintf("mongodb://%v:%v@%v:%v/fbMarketplaceScrapper?authSource=admin",
		info.User, info.Password, info.Host, info.Port)

}
