package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/pkg/errors"
	"github.com/reynaldoqs/x_resolver/internal/core/domain"
)

type communityRepo struct {
	client *firestore.Client
}

func NewCommunityRechargeRepository(gapp *firebase.App) *communityRepo {

	client, err := gapp.Firestore(context.Background())
	if err != nil {
		err = errors.Wrap(err, "communityrecharge.NewCommunityRechargeRepository")
		log.Fatalln(err)
	}

	crepo := communityRepo{
		client: client,
	}
	return &crepo
}

func (cr *communityRepo) Store(rechage *domain.CommunityRecharge) error {

	collection := cr.client.Collection("communityRecharge")
	_, _, err := collection.Add(context.TODO(), rechage)
	if err != nil {
		err = errors.Wrap(err, "communityrecharge.Store")
		return err
	}

	return err
}
