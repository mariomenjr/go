package models

import (
	"fmt"
	"log"
	"os"

	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() {
	user := os.Getenv("MONGO_USER")
	pass := os.Getenv("MONGO_PASS")
	host := os.Getenv("MONGO_HOST")
	db := os.Getenv("MONGO_DB")

	url := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", user, pass, host, db)

	err := mgm.SetDefaultConfig(nil, db, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("links/models: Could'nt connect to database server")
	}
}
