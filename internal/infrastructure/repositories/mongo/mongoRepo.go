package mongo

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoclient *mongo.Client

func NewMongoClient(mongoURL string, mongoTimeout int) *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()

	mongoclient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		err := errors.Wrap(err, "mongoRepo.NewMongoClient")
		log.Fatalln(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		err := errors.Wrap(err, "mongoRepo.NewMongoClient")
		log.Fatalln(err)
	}
	return mongoclient
}
