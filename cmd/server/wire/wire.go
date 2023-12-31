//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go-gin-demo/internal/common/repository"
	"go-gin-demo/internal/common/routers"
	"go-gin-demo/internal/common/server"
	"go-gin-demo/internal/knowledge"
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

var KnowledgeSet = wire.NewSet(
	knowledge.NewKnowledgeClient,
	knowledge.NewCollectionHandler,
	knowledge.NewKnowledge,
)

func NewServer(viperViper *viper.Viper, _ *log.Logger) http.Handler {
	wire.Build(
		ServerSet,
		JwtSet,
		CommSet,
		SystemSet,
		KnowledgeSet,
		RouterSet,
	)
	return nil
}
