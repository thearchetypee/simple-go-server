package models

import "time"

type Record struct {
	Key        string    `json:"key" bson:"key"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	TotalCount int       `json:"totalCount" bson:"totalCount"`
}

type MongoFetchRespnse struct {
	Code    int      `json:"code"`
	Status  string   `json:"status"`
	Records []Record `json:"records"`
}
