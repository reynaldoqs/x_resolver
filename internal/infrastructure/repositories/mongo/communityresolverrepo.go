package mongo

import (
	"context"
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
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"id":                rslr.ID,
			"msgToken":          rslr.MsgToken,
			"notify":            rslr.Notify,
			"resolverRecharges": rslr.CRecharges,
			"resolvers":         rslr.Resolvers,
		},
	)

	if err != nil {
		err := errors.Wrap(err, "communityresolverrepo.SaveC")
		return err
	}

	return err

}

// UpdateC updates all fields without recharges
func (r *communityRepo) UpdateC(id string, resolver *domain.CommunityResolver) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection("communityResolvers")

	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"id":                resolver.ID,
			"msgToken":          resolver.MsgToken,
			"notify":            resolver.Notify,
			"resolverRecharges": resolver.CRecharges,
			"resolvers":         resolver.Resolvers,
		},
	}
	upset := true
	after := options.After

	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upset,
	}

	result := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	if result.Err() != nil {
		err := errors.Wrap(result.Err(), "communityresolverrepo.GetOneC")
		return err
	}

	return nil
}
func (r *communityRepo) GetOneC(id string) (*domain.CommunityResolver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection("communityResolvers")

	result := collection.FindOne(ctx, bson.M{"id": id})

	var resolver domain.CommunityResolver

	err := result.Decode(&resolver)
	if err != nil {
		err = errors.Wrap(err, "communityresolverrepo.GetOneC")
		return nil, err
	}

	return &resolver, nil
}

func (r *communityRepo) RemoveC(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection("communityResolvers")

	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		err = errors.Wrap(err, "comminityresolverrepo.RemoveC")
		return err
	}

	return err
}
