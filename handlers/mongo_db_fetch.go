package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/simple-go-server/config"
	"github.com/simple-go-server/db"
	"github.com/simple-go-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleMongoFetch(w http.ResponseWriter, r *http.Request) error {
	cfg, ok := r.Context().Value(config.ConfigKey).(*config.Config)
	if !ok {
		return fmt.Errorf("config not found in context")
	}

	var filter struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		MinCount  int    `json:"minCount"`
		MaxCount  int    `json:"maxCount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		return fmt.Errorf("error decoding request body: %v", err)
	}

	// Parse dates
	startDate, err := time.Parse("2006-01-02", filter.StartDate)
	if err != nil {
		return fmt.Errorf("invalid start date: %v", err)
	}
	endDate, err := time.Parse("2006-01-02", filter.EndDate)
	if err != nil {
		return fmt.Errorf("invalid end date: %v", err)
	}
	if filter.MinCount < 0 {
		return fmt.Errorf("invalid min count should not be less than 0")
	}
	if filter.MaxCount < 0 {
		return fmt.Errorf("invalid max count should not be less than 0")
	}
	collection := cfg.DB.Database(db.DB_NAME).Collection(db.TABLE_NAME)

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "createdAt", Value: bson.D{
				{Key: "$gte", Value: startDate},
				{Key: "$lte", Value: endDate},
			}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "createdAt", Value: 1},
			{Key: "key", Value: 1},
			{Key: "totalCount", Value: 1},
		}}},
		{{Key: "$match", Value: bson.D{
			{Key: "totalCount", Value: bson.D{
				{Key: "$gte", Value: filter.MinCount},
				{Key: "$lte", Value: filter.MaxCount},
			}},
		}}},
	}

	// Execute the aggregation
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return fmt.Errorf("error executing MongoDB aggregation: %v", err)
	}
	defer cursor.Close(context.TODO())

	var results []models.Record
	if err = cursor.All(context.TODO(), &results); err != nil {
		return fmt.Errorf("error decoding MongoDB results: %v", err)
	}

	resp := &models.MongoFetchRespnse{
		Code:    0,
		Status:  "Success",
		Records: results,
	}
	return writeJSON(w, http.StatusOK, resp)
}
