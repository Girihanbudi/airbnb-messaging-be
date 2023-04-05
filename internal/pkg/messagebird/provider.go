package messagebird

import (
	"airbnb-messaging-be/internal/pkg/messagebird/config"

	messagebird "github.com/messagebird/go-rest-api/v9"
)

const Instance string = "Messenger"

type Options struct {
	config.Config
}

type Messenger struct {
	Client *messagebird.DefaultClient
	Options
}

func InitMessenger(options Options) *Messenger {
	// Create a client.
	client := messagebird.New(options.AccessKey)

	return &Messenger{
		Client:  client,
		Options: options,
	}
}
