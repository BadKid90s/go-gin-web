package client

import (
	"github.com/spf13/viper"
	"log"
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
	log.Panicln("embedding config is not found")
	return nil
}
