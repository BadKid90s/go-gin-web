//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go-gin-demo/internal/ai"
	"go-gin-demo/internal/common/repository"
	"go-gin-demo/internal/routers"
	"go-gin-demo/internal/server"
	"go-gin-demo/internal/system"
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
	repository.NewDB,
	repository.NewRepository,
)

var RouterSet = wire.NewSet(
	routers.NewRouter,
)

var SystemSet = wire.NewSet(
	system.NewUserHandler,
	system.NewUserService,
	system.NewUserRepository,
	system.NewSystem,
)

var AiSet = wire.NewSet(
	ai.NewQdrant,
	ai.NewQdrantClient,
	ai.NewCollectionsClient,
	ai.NewPointsClient,
	ai.NewSnapshotsClient,
	ai.NewCollectionHandler,
	ai.NewAi,
)

func NewServer(viperViper *viper.Viper, _ *log.Logger) http.Handler {
	wire.Build(
		ServerSet,
		JwtSet,
		CommSet,
		SystemSet,
		AiSet,
		RouterSet,
	)
	return nil
}
