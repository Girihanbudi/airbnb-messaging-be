//go:build wireinject
// +build wireinject

package app

import (
	env "airbnb-messaging-be/internal/pkg/env/wire"
	gorm "airbnb-messaging-be/internal/pkg/gorm/wire"
	http "airbnb-messaging-be/internal/pkg/http/server/wire"
	kafka "airbnb-messaging-be/internal/pkg/kafka/wire"
	messagebird "airbnb-messaging-be/internal/pkg/messagebird/wire"

	sms "airbnb-messaging-be/internal/app/sms/wire"

	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	wire.Struct(new(Options), "*"),
	wire.Struct(new(App), "*"),
)

func NewApp() (*App, error) {
	panic(
		wire.Build(
			env.PackageSet,
			gorm.PackageSet,
			messagebird.PackageSet,
			http.PackageSet,
			kafka.PackageSet,

			AppSet,

			sms.ModuleSet,
		),
	)
}
