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

func (msgn *messageNotifier) FarmerNotify(farmer *domain.Farmer, data map[string]string) error {
	//get from resolvers
	fmt.Println("Ya empezamos 1")

	notification := messaging.Notification{
		Title: farmer.DeviceID,
		Body:  fmt.Sprintf("el numero %v necesita recarga", farmer.PhoneNumber),
	}
	fmt.Println("Ya empezamos 2")
	data = map[string]string{
		"execCode": "*#62#",
	}
	fmt.Println("Ya empezamos 3")
	message := messaging.Message{
		Token:        farmer.MsgToken,
		Data:         data,
		Notification: &notification,
	}
	fmt.Println("Ya empezamos 4")
	result, err := msgn.client.Send(context.TODO(), &message)
	if err != nil {
		fmt.Println(err)
		err = errors.Wrap(err, "cloudmessaging.RechargeNotify")
		return err
	}
	fmt.Println("Ya mandamos")
	fmt.Println(result)
	return err
}
