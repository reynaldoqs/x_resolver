package firebasemsg

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/pkg/errors"
	"github.com/reynaldoqs/x_resolver/internal/core/domain"
)

type messageNotifier struct {
	client *messaging.Client
}

func NewCloudMessaging(gapp *firebase.App) *messageNotifier {

	client, err := gapp.Messaging(context.Background())
	if err != nil {
		err = errors.Wrap(err, "cloudmessaging.NewCloudMessaging")
		log.Fatalln(err)
	}

	msgn := messageNotifier{
		client: client,
	}
	return &msgn
}

func (msgn *messageNotifier) RechargeNotify(recharge *domain.Recharge, resolvers []*domain.CommunityResolver) error {
	//get from resolvers
	var messages []*messaging.Message

	notification := messaging.Notification{
		Title: recharge.Company,
		Body:  fmt.Sprintf("el numero %v necesita recarga", recharge.PhoneNumber),
	}

	data := map[string]string{
		"donde": "audiman",
	}

	for _, v := range resolvers {
		msg := messaging.Message{Token: v.MsgToken, Data: data, Notification: &notification}
		messages = append(messages, &msg)
	}

	_, err := msgn.client.SendAll(context.TODO(), messages)
	if err != nil {
		err = errors.Wrap(err, "cloudmessaging.RechargeNotify")
		return err
	}
	return err
}
