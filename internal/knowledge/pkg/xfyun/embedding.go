package xfyun

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

type Embedding struct {
	apiKey string
	url    string
}

func NewEmbedding(conf *viper.Viper) *Embedding {
	apiKey := conf.GetString("embedding.apiKey")
	url := conf.GetString("embedding.url")

	return &Embedding{
		apiKey: apiKey,
		url:    url,
	}
}

func (e *Embedding) GetEmbedding(text string) ([]float32, error) {
	authorization, err := getAuthorization(e.apiKey)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient

	param := map[string]string{
		"prompt": text,
	}

	jsonData, _ := json.Marshal(param)

	request, err := http.NewRequest("POST", e.url, bytes.NewBuffer(jsonData))
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
