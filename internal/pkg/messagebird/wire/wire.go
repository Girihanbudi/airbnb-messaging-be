package wire

import (
	"airbnb-messaging-be/internal/pkg/messagebird"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	wire.Struct(new(messagebird.Options), "*"),
	messagebird.InitMessenger,
)
