package firebasemsg

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

var firebaseApp *firebase.App

func NewFirebaseApp(configPath string) *firebase.App {
	cOptions := option.WithCredentialsFile(configPath)

	firebaseApp, err := firebase.NewApp(context.Background(), nil, cOptions)
	if err != nil {
		err = errors.Wrap(err, "firebase.NewFirebaseApp")
		log.Fatalln(err)
	}
	return firebaseApp
}
