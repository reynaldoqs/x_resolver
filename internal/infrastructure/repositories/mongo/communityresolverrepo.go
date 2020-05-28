package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/reynaldoqs/x_resolver/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type communityRepo struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func NewCommunityResolverRepository(mongoClient *mongo.Client, mongoDB string) *communityRepo {
	repo := &communityRepo{
		client:   mongoClient,
		database: mongoDB,
		timeout:  time.Duration(30) * time.Second,
	}
	return repo

}

func (r *communityRepo) GetAllC() ([]*domain.CommunityResolver, error) {

	collection := r.client.Database(r.database).Collection("communityResolvers")
	findOptions := options.Find()

	var results []*domain.CommunityResolver

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		err := errors.Wrap(err, "comminityresolverrepo.GetAllC")
		return nil, err
	}

	for cur.Next(context.TODO()) {

		var elem domain.CommunityResolver
		err := cur.Decode(&elem)
		if err != nil {
			err := errors.Wrap(err, "comminityresolverrepo.GetAllC")
			return nil, err
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		err := errors.Wrap(err, "comminityresolverrepo.GetAllC")
		return nil, err
	}

	cur.Close(context.TODO())
	return results, nil

}
func (r *communityRepo) SaveC(rslr *domain.CommunityResolver) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection("communityResolvers")
	fmt.Println(rslr)
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"ID":                rslr.ID,
			"msgToken":          rslr.MsgToken,
			"notify":            rslr.Notify,
			"resolverRecharges": rslr.CRecharges,
			"resolvers":         rslr.Resolvers,
		},
	)

	if err != nil {
		err := errors.Wrap(err, "comminityresolverrepo.SaveC")
		return err
	}

	return nil

}
