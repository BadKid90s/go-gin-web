//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go-gin-demo/internal/common"
	"go-gin-demo/internal/routers"
	"go-gin-demo/internal/server"
	"go-gin-demo/internal/system/handler"
	"go-gin-demo/pkg/log"
	"net/http"
)

var ServerSet = wire.NewSet(
	server.NewServer,
)

var CommSet = wire.NewSet(
	common.NewHandler,
	common.NewDB,
	common.NewRepository,
)

var SystemHandlerSet = wire.NewSet(
	handler.NewUserHandler,
	handler.NewSystemHandler,
)

var RouterSet = wire.NewSet(
	routers.NewRouter,
)

func NewServer(viperViper *viper.Viper, logger *log.Logger) http.Handler {
	wire.Build(
		ServerSet,
		CommSet,
		SystemHandlerSet,
		RouterSet,
	)
	return nil
}