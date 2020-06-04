package firestore

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/pkg/errors"
	"github.com/reynaldoqs/x_resolver/internal/core/domain"
)

type farmersRepo struct {
	client *firestore.Client
}

func NewFarmersRepository(gapp *firebase.App) *farmersRepo {

	client, err := gapp.Firestore(context.Background())
	if err != nil {
		err = errors.Wrap(err, "farmersRepo.NewFarmersRepository")
		log.Fatalln(err)
	}

	crepo := farmersRepo{
		client: client,
	}
	return &crepo
}

func (cr *farmersRepo) GetAllFarmers() ([]*domain.Farmer, error) {
	collection := cr.client.Collection("farmerResolvers")

	var farmers []*domain.Farmer

	result, err := collection.Documents(context.TODO()).GetAll()
	if err != nil {
		err = errors.Wrap(err, "farmersRepo.NewFarmersRepository")
		return nil, err
	}

	for _, docSnap := range result {
		var myFarmer domain.Farmer

		err := docSnap.DataTo(&myFarmer)
		if err != nil {
			fmt.Println(err)
			err = errors.Wrap(err, "farmersRepo.NewFarmersRepository")
			return nil, err
		}

		farmers = append(farmers, &myFarmer)

	}
	return farmers, nil
}
