package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/reynaldoqs/x_resolver/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type rechargesRepo struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func NewRechargesRepository(mongoClient *mongo.Client, mongoDB string) *rechargesRepo {
	repo := &rechargesRepo{
		client:   mongoClient,
		database: mongoDB,
		timeout:  time.Duration(30) * time.Second,
	}
	return repo

}

func (r *rechargesRepo) GetAllR() ([]*domain.Recharge, error) {

	collection := r.client.Database(r.database).Collection("recharges")
	findOptions := options.Find()

	var results []*domain.Recharge

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		err := errors.Wrap(err, "rechargesrepo.GetAll")
		return nil, err
	}

	for cur.Next(context.TODO()) {

		var elem domain.Recharge
		err := cur.Decode(&elem)
		if err != nil {
			err := errors.Wrap(err, "rechargesrepo.GetAll")
			return nil, err
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		err := errors.Wrap(err, "rechargesrepo.GetAll")
		return nil, err
	}

	cur.Close(context.TODO())
	return results, nil

}
func (r *rechargesRepo) SaveR(recharge *domain.Recharge) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection("recharges")

	result, err := collection.InsertOne(
		ctx,
		bson.M{
			"ID":          recharge.ID,
			"phoneNumber": recharge.PhoneNumber,
			"company":     recharge.Company,
			"cardNumber":  recharge.CardNumber,
			"status":      recharge.Status,
			"mount":       recharge.Mount,
			"idResolver":  recharge.IDResolver,
			"createdAt":   recharge.CreatedAt,
			"resolvedAt":  recharge.ResolvedAt,
		},
	)

	if err != nil {
		err := errors.Wrap(err, "rechargesrepo.Save")
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		recharge.ID = oid.Hex()
	}

	return nil

}
