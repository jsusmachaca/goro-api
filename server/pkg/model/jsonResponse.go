package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ResultsResponse struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name    string             `bson:"name" json:"name"`
	Status  string             `bson:"status" json:"status"`
	Species string             `bson:"species" json:"species"`
	Type    string             `bson:"type" json:"type"`
	Gender  string             `bson:"gender" json:"gender"`
	Origin  struct {
		Name string `bson:"name" json:"name"`
		URL  string `bson:"url" json:"url"`
	} `bson:"origin" json:"origin"`
	Location struct {
		Name string `bson:"name" json:"name"`
		URL  string `bson:"url" json:"url"`
	} `bson:"location" json:"location"`
	Image   string    `bson:"image" json:"image"`
	Episode []string  `bson:"episode" json:"episode"`
	URL     string    `bson:"url" json:"url"`
	Created time.Time `bson:"created" json:"created"`
}

type JsonResponse struct {
	Info struct {
		Count int         `json:"count"`
		Pages int         `json:"pages"`
		Next  string      `json:"next"`
		Prev  interface{} `json:"prev"`
	} `json:"info"`
	Results []ResultsResponse `json:"results"`
}
