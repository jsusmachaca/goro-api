package repository

import (
	"context"

	"github.com/jsusmachaca/goroapi/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAll(client *mongo.Client) ([]model.ResultsResponse, error) {
	collection := client.Database("rnm").Collection("rnm_characters")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var results []model.ResultsResponse

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
