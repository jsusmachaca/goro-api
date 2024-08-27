package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/jsusmachaca/goroapi/api/handler"
	"github.com/jsusmachaca/goroapi/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

var (
	PORT string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	URI := os.Getenv("MONGO_URI")
	if URI == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	PORT = fmt.Sprintf(":%s", os.Getenv("PORT"))
	if PORT == ":" {
		log.Fatal("Set your 'PORT' environment variable.")
	}

	var err error
	client, err = database.MongoConection(URI)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	http.HandleFunc("/api/rnm", func(w http.ResponseWriter, r *http.Request) {
		handler.SendData(w, r, client)
	})

	http.ListenAndServe(PORT, nil)
}
