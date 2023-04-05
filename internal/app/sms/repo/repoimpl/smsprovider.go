package repoimpl

import (
	"airbnb-messaging-be/internal/pkg/gorm"
)

type Options struct {
	Gorm *gorm.Engine
}

type Repo struct {
	Options
}

func NewSmsRepo(options Options) *Repo {
	return &Repo{options}
}
