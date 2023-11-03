package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/common/resp"
	"go-gin-demo/internal/knowledge/client"
	"time"
)

type CollectionHandler interface {
	Create(ctx *gin.Context)
}

func NewCollectionHandler(client client.KnowledgeClient) CollectionHandler {
	return &collectionHandler{
		client,
	}
}

type collectionHandler struct {
	knowledgeClient client.KnowledgeClient
}

func (c *collectionHandler) Create(ctx *gin.Context) {
	message := "你好啊"
	text := ctx.Query("text")
	if text != "" {
		message = text
	}
	chatChan, err := c.knowledgeClient.ChatMessage(message)
	if err != nil {
		resp.HandleError(ctx, err, nil)
		return
	}
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	for {
		select {
		case message, ok := <-chatChan:
			if !ok {
				return // 当通道关闭时结束流式连接
			}
			_, err := fmt.Fprintf(ctx.Writer, "%s", message)
			if err != nil {
				return
			}
			ctx.Writer.Flush()
			time.Sleep(100 * time.Millisecond) // 添加延迟以刷新输出缓冲区
		case <-ctx.Writer.CloseNotify():
			// Client closed connection, stop sending events
			return
		}

	}
}
