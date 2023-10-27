package knowledge

import "go-gin-demo/internal/knowledge/handler"

var (
	NewCollectionHandler = handler.NewCollectionHandler
)

type Knowledge struct {
	CollectionHandler handler.CollectionHandler
}

func NewKnowledge(collectionHandler handler.CollectionHandler) *Knowledge {
	return &Knowledge{
		CollectionHandler: collectionHandler,
	}
}
