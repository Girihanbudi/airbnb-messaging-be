package wire

import (
	"airbnb-messaging-be/internal/pkg/credential"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	wire.Struct(new(credential.Options), "*"),
	credential.NewTLSCredentials,
)
