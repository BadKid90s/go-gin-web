//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go-gin-demo/internal/routers"
	"go-gin-demo/internal/server"
	"go-gin-demo/internal/system/handler"
	"net/http"
)

var ServerSet = wire.NewSet(
	server.NewServer,
)

var RouterSet = wire.NewSet(
	routers.NewRouter,
)

var HandlerSet = wire.NewSet(
	handler.NewUserHandler,
	handler.NewSystemHandler,
)

func NewServer(_ *viper.Viper) http.Handler {
	wire.Build(
		ServerSet,
		HandlerSet,
		RouterSet,
	)
	return nil
}
