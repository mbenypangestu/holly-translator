package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mygetzu/holly-translator/models"
)

const DB_NAME = "holly_dev"
const DB_COLLECTION = "location"
const DB_CONNECTIONSTRING = "mongodb://localhost:27017"

var currentTime = time.Now()

func init() {
	fmt.Println("[ ", currentTime.Format("2001-01-02 15:04:05 Monday"), "] Initializing ...")

	var locations []models.Location

	client, err := options.Client().ApplyURI(DB_CONNECTIONSTRING)
	if err != nil {
		log.Fatal(err)
	}

	// err = client.Connect(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db := client.Database(DB_NAME)
}

func main() {
	fmt.Println("[ ", currentTime.Format("2001-01-02 15:04:05 Monday"), "] Start Translating review ...")
}
