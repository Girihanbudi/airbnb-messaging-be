package messagebird

import (
	"airbnb-messaging-be/internal/pkg/env"
	"airbnb-messaging-be/internal/pkg/log"
	"fmt"

	messagebird "github.com/messagebird/go-rest-api/v9"
)

const Instance string = "Messenger"

// global auth cache declaration
var MessengerClient *messagebird.DefaultClient

func InitAuthCache() {

	// Create a client.
	client := messagebird.New(env.CONFIG.Messenger.AccessKey)

	log.Event(Instance, fmt.Sprintf("connected to %s", env.CONFIG.Cache.Auth.Host))

	MessengerClient = client
}
