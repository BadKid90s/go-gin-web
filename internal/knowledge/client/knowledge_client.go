package client

import (
	"fmt"
	"github.com/spf13/viper"
)

type KnowledgeClient interface {
	GetEmbedding(text string) ([]float32, error)
}

func NewKnowledgeClient(conf *viper.Viper) KnowledgeClient {
	if conf.InConfig("embedding.xunfei") {
		return newKnowledgeClientXunfei(conf)
	}
	if conf.InConfig("embedding.zhipu") {
		return newKnowledgeClientZhipu(conf)
	}
	panic(fmt.Sprintf("embedding config is not found"))
	return nil
}
