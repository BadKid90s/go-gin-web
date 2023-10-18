package handler

import (
	"github.com/gin-gonic/gin"
	qdrant "github.com/qdrant/go-client/qdrant"
	"go-gin-demo/internal/common/resp"
	"go-gin-demo/internal/knowledgebase/pkg/xfyun"
)

type CollectionHandler interface {
	Create(ctx *gin.Context)
}

func NewCollectionHandler(collectionsClient qdrant.CollectionsClient, embedding *xfyun.Embedding) CollectionHandler {
	return &collectionHandler{
		collectionsClient: collectionsClient,
		embedding:         embedding,
	}
}

type collectionHandler struct {
	collectionsClient qdrant.CollectionsClient
	embedding         *xfyun.Embedding
}

func (c *collectionHandler) Create(ctx *gin.Context) {
	embedding, err := c.embedding.GetEmbedding("你好啊")
	if err != nil {
		resp.HandleError(ctx, err, nil)
		return
	}
	resp.HandleSuccess(ctx, embedding)
}
