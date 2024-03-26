package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Result struct {
	Docs          []bson.M
	TotalDocs     int64
	Limit         int64
	Page          int64
	TotalPages    int64
	HasNextPage   bool
	NextPage      int64
	HasPrevPage   bool
	PrevPage      *int64
	PagingCounter int64
}

func PaginateAggregate(collection *mongo.Collection, pipeline []interface{}, page, limit int64, collation *options.Collation) (result *Result, err error) {
	ctx := context.TODO()

	totalDocs, err := collection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 100
	}

	totalPages := (totalDocs + limit - 1) / limit
	offset := (page - 1) * limit
	options := options.Aggregate().SetMaxTime(2000)

	if collation != nil {
		options.SetCollation(collation)
	}

	pipeline = append(pipeline, bson.D{{Key: "$skip", Value: offset}}, bson.D{{Key: "$limit", Value: limit}})

	cursor, err := collection.Aggregate(ctx, pipeline, options)
	if err != nil {
		return nil, err
	}

	var docs []bson.M
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	result = &Result{
		Docs:          docs,
		TotalDocs:     totalDocs,
		Limit:         limit,
		Page:          page,
		TotalPages:    totalPages,
		HasNextPage:   page < totalPages,
		NextPage:      page + 1,
		HasPrevPage:   page > 1,
		PrevPage:      nil,
		PagingCounter: offset + 1,
	}

	if page > 1 {
		prevPage := page - 1
		result.PrevPage = &prevPage
	}

	return result, nil
}
