package wire

import (
	"airbnb-messaging-be/internal/pkg/kafka"
	"airbnb-messaging-be/internal/pkg/kafka/router"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	wire.Struct(new(router.Options), "*"),
	router.NewRouter,

	wire.Struct(new(kafka.Options), "*"),
	kafka.NewEventListener,
)
