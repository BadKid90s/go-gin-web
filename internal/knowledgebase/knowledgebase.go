package knowledgebase

import (
	qdrant "github.com/qdrant/go-client/qdrant"
	"github.com/spf13/viper"
	hd "go-gin-demo/internal/knowledgebase/handler"
	"go-gin-demo/internal/knowledgebase/pkg/xfyun"
	"go-gin-demo/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	NewCollectionHandler = hd.NewCollectionHandler
	NewEmbedding         = xfyun.NewEmbedding
)

func NewQdrant(conf *viper.Viper, logger *log.Logger) *grpc.ClientConn {

	url := conf.GetString("qdrant.url")

	clientConn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("connect qdrant failed")
	}
	defer clientConn.Close()

	return clientConn
}

func NewQdrantClient(conn *grpc.ClientConn) qdrant.QdrantClient {
	return qdrant.NewQdrantClient(conn)
}
func NewCollectionsClient(conn *grpc.ClientConn) qdrant.CollectionsClient {
	return qdrant.NewCollectionsClient(conn)
}
func NewPointsClient(conn *grpc.ClientConn) qdrant.PointsClient {
	return qdrant.NewPointsClient(conn)
}
func NewSnapshotsClient(conn *grpc.ClientConn) qdrant.SnapshotsClient {
	return qdrant.NewSnapshotsClient(conn)
}

type Knowledgebase struct {
	Collection hd.CollectionHandler
}

func NewKnowledgebase(collection hd.CollectionHandler) *Knowledgebase {
	return &Knowledgebase{
		Collection: collection,
	}
}
