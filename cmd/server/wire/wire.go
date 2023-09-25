//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go-gin-demo/internal/common"
	"go-gin-demo/internal/routers"
	"go-gin-demo/internal/server"
	"go-gin-demo/internal/system/handler"
	"go-gin-demo/internal/system/repository"
	"go-gin-demo/internal/system/service"
	"go-gin-demo/pkg/jwt"
	"go-gin-demo/pkg/log"
	"net/http"
)

var ServerSet = wire.NewSet(
	server.NewServer,
)
var JwtSet = wire.NewSet(
	jwt.NewJwt,
)

var CommSet = wire.NewSet(
	common.NewHandler,
	common.NewDB,
	common.NewRepository,
)

var SystemSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	handler.NewUserHandler,
	handler.NewSystemHandler,
)

var RouterSet = wire.NewSet(
	routers.NewRouter,
)

func NewServer(viperViper *viper.Viper, _ *log.Logger) http.Handler {
	wire.Build(
		ServerSet,
		JwtSet,
		CommSet,
		SystemSet,
		RouterSet,
	)
	return nil
}
