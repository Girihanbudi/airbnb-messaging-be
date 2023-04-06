package wire

import (
	"airbnb-messaging-be/internal/app/sms/api/event"
	"airbnb-messaging-be/internal/app/sms/repo"
	"airbnb-messaging-be/internal/app/sms/repo/repoimpl"
	"airbnb-messaging-be/internal/app/sms/usecase"
	"airbnb-messaging-be/internal/app/sms/usecase/usecaseimpl"

	"github.com/google/wire"
)

var ModuleSet = wire.NewSet(
	repoSet,
	usecaseSet,
	apiSet,
)

var repoSet = wire.NewSet(
	wire.Struct(new(repoimpl.Options), "*"),
	repoimpl.NewSmsRepo,
	wire.Bind(new(repo.ISms), new(*repoimpl.Repo)),
)

var usecaseSet = wire.NewSet(
	wire.Struct(new(usecaseimpl.Options), "*"),
	usecaseimpl.NewSmsUsecase,
	wire.Bind(new(usecase.ISms), new(*usecaseimpl.Usecase)),
)

var apiSet = wire.NewSet(
	wire.Struct(new(event.Options), "*"),
	event.NewSmsHandler,
)
