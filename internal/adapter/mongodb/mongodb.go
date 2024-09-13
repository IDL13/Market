package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/IDL13/Market/internal/entity"
	"github.com/IDL13/Market/pkg/time_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB interface {
	AddStatistics(json entity.SaveRequest) int8
	ResetStatistics() int8
	GetStatistics(json entity.ShowRequest) []entity.Statistics
}

type mongoDB struct {
	port       string
	datatbase  string
	collection string
}

func New(port, datatbase, collection string) MongoDB {
	return mongoDB{
		port:       port,
		datatbase:  datatbase,
		collection: collection,
	}
}

func connect(port string) (*mongo.Client, context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	url := fmt.Sprintf("mongodb://localhost:%s", port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	return client, ctx
}

func disconect(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func (m mongoDB) AddStatistics(json entity.SaveRequest) int8 {
	client, ctx := connect(m.port)
	defer disconect(client, ctx)

	collection := client.Database(m.datatbase).Collection(m.collection)

	insert_ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	statistics := entity.New(
		json.Date,
		json.Views,
		json.Clicks,
		float64(json.Cost),
	)

	_, err := collection.InsertOne(insert_ctx, statistics)
	if err != nil {
		e := time_errors.New(err)
		fmt.Println(e.Error())

		return 0
	}

	return 1
}
func (m mongoDB) ResetStatistics() int8 {
	client, ctx := connect(m.port)
	defer disconect(client, ctx)

	collection := client.Database(m.datatbase).Collection(m.collection)

	insert_ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.Drop(insert_ctx)
	if err != nil {
		e := time_errors.New(err)
		fmt.Println(e.Error())

		return 0
	}

	return 1
}

func (m mongoDB) GetStatistics(json entity.ShowRequest) []entity.Statistics {
	var result []entity.Statistics

	client, ctx := connect(m.port)
	defer disconect(client, ctx)

	collection := client.Database(m.datatbase).Collection(m.collection)

	insert_ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"date", bson.D{{"$gte", json.From}}}},
				bson.D{{"date", bson.D{{"$lte", json.To}}}},
			}},
	}

	cur, err := collection.Find(insert_ctx, filter)
	if err != nil {
		e := time_errors.New(err)
		fmt.Println(e.Error())
	}
	defer cur.Close(insert_ctx)

	for cur.Next(insert_ctx) {
		var statistics entity.Statistics

		if err = cur.Decode(&statistics); err != nil {
			e := time_errors.New(err)
			fmt.Println(e.Error())
		}

		result = append(result, statistics)
	}

	if err = cur.Err(); err != nil {
		e := time_errors.New(err)
		fmt.Println(e.Error())
	}

	return result
}
