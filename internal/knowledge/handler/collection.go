package handler

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/common/resp"
)

type CollectionHandler interface {
	Create(ctx *gin.Context)
}

func NewCollectionHandler() CollectionHandler {
	return &collectionHandler{}
}

type collectionHandler struct {
}

func (c *collectionHandler) Create(ctx *gin.Context) {
	//
	//collectionName := "collectionName"
	//
	//var defaultSegmentNumber uint64 = 2
	//_, err := c.qdrantClient.CollectionsClient.Create(ctx, &qdrant.CreateCollection{
	//	CollectionName: collectionName,
	//	VectorsConfig: &qdrant.VectorsConfig{Config: &qdrant.VectorsConfig_Params{
	//		Params: &qdrant.VectorParams{
	//			Size:     1024,
	//			Distance: qdrant.Distance_Dot,
	//		},
	//	}},
	//	OptimizersConfig: &qdrant.OptimizersConfigDiff{
	//		DefaultSegmentNumber: &defaultSegmentNumber,
	//	},
	//})
	//
	//if err != nil {
	//	resp.HandleError(ctx, err, nil)
	//	return
	//}
	//
	//strs := "WRYKotlin是一种静态类型的编程语言，由JetBrains开发。它可以编译成Java字节码，也可以编译成JavaScript代码。 WRYKotlin是一种跨平台语言，可以用于Android、Java、JavaScript和Native等领域。WRYKotlin具有许多现代编程语言的特性，例如函数式编程和空安全检查等。WRYKotlin的语法简单易懂，非常适合初学者学习。"
	//
	//var upsertPoints []*qdrant.PointStruct
	//
	//for i, str := range strings.Split(strs, "。") {
	//
	//	embedding, _ := c.embedding.GetEmbedding(str)
	//
	//	point := qdrant.PointStruct{
	//		// Point Id is number or UUID
	//		Id: &qdrant.PointId{
	//			PointIdOptions: &qdrant.PointId_Num{Num: uint64(i)},
	//		},
	//		Vectors: &qdrant.Vectors{VectorsOptions: &qdrant.Vectors_Vector{Vector: &qdrant.Vector{Data: embedding}}},
	//		Payload: map[string]*qdrant.Value{
	//			"message": {
	//				Kind: &qdrant.Value_StringValue{StringValue: str},
	//			},
	//		},
	//	}
	//
	//	upsertPoints = append(upsertPoints, &point)
	//}
	//waitUpsert := true
	//_, err = c.qdrantClient.PointsClient.Upsert(ctx, &qdrant.UpsertPoints{
	//	CollectionName: collectionName,
	//	Wait:           &waitUpsert,
	//	Points:         upsertPoints,
	//})
	resp.HandleSuccess(ctx, "embedding")
}
