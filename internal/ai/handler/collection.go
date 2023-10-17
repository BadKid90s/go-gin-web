package handler

import (
	"github.com/gin-gonic/gin"
	qdrant "github.com/qdrant/go-client/qdrant"
)

type CollectionHandler interface {
	Create(ctx *gin.Context)
}

func NewCollectionHandler(collectionsClient qdrant.CollectionsClient) CollectionHandler {
	return &collectionHandler{
		collectionsClient: collectionsClient,
	}
}

type collectionHandler struct {
	collectionsClient qdrant.CollectionsClient
}

func (c *collectionHandler) Create(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
