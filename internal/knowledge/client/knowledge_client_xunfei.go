package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

type KnowledgeClientXunfei struct {
	appid     string
	apiKey    string
	apiSecret string
	url       string
}

func newKnowledgeClientXunfei(conf *viper.Viper) KnowledgeClient {
	appid := conf.GetString("embedding.xunfei.appid")
	apiKey := conf.GetString("embedding.xunfei.apiKey")
	apiSecret := conf.GetString("embedding.xunfei.apiSecret")
	url := conf.GetString("embedding.xunfei.url")

	return &KnowledgeClientXunfei{
		appid:     appid,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		url:       url,
	}
}

func (c *KnowledgeClientXunfei) GetEmbedding(text string) ([]float32, error) {
	authorization, err := c.getAuthorization()
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient

	param := map[string]string{
		"prompt": text,
	}

	jsonData, _ := json.Marshal(param)

	request, err := http.NewRequest("POST", c.url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", authorization)
	if err != nil {
		return nil, err

	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err

	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type JsonResult struct {
		Success bool
		Msg     string
		Data    struct {
			Embedding []float32
		}
	}
	var result JsonResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.Success {
		return result.Data.Embedding, nil
	}
	return nil, errors.New(result.Msg)

}

// getAuthorization 获取用户鉴权
func (c *KnowledgeClientXunfei) getAuthorization() (string, error) {

	return "", nil
}
