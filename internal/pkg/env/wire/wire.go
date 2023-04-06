package wire

import (
	"airbnb-messaging-be/internal/pkg/env"
	"airbnb-messaging-be/internal/pkg/env/tool"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	env.ProvideEnv,
	tool.ExtractServerConfig,
	tool.ExtractDBConfig,
	tool.ExtractMessengerConfig,
	tool.ExtractKafkaConfig,
	tool.ExtractKafkaRouterConfig,
)
