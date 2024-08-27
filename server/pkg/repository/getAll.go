package repository

import (
	"context"

	"github.com/jsusmachaca/goroapi/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchData(client *mongo.Client) ([]models.ResultsResponse, error) {
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("rnm").Collection("rnm_characters")

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var results []models.ResultsResponse

	err = cursor.All(context.TODO(), &results)

	if err != nil {
		return nil, err
	}
	return results, nil
}
