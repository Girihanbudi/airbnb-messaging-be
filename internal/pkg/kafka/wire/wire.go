package wire

import (
	"airbnb-messaging-be/internal/pkg/kafka"
	"airbnb-messaging-be/internal/pkg/kafka/router"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	router.NewRouter,

	wire.Struct(new(kafka.Options), "*"),
	kafka.NewEventListener,
)
