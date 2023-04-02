package wire

import (
	"airbnb-messaging-be/internal/pkg/gorm"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	gorm.NewORM,
)
