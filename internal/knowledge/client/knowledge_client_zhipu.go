package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"strings"
	"time"
)

type KnowledgeClientZhipu struct {
	apiKey string
	url    string
}

func newKnowledgeClientZhipu(conf *viper.Viper) KnowledgeClient {
	apiKey := conf.GetString("embedding.zhipu.apiKey")
	url := conf.GetString("embedding.zhipu.url")

	return &KnowledgeClientZhipu{
		apiKey: apiKey,
		url:    url,
	}
}

func (c *KnowledgeClientZhipu) GetEmbedding(text string) ([]float32, error) {
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
func (c *KnowledgeClientZhipu) getAuthorization() (string, error) {
	keys := strings.Split(c.apiKey, ".")

	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	timestamp := time.Now().Unix()

	mapClaims := jwt.MapClaims{
		"api_key":   keys[0],
		"exp":       expirationTime,
		"timestamp": timestamp,
	}

	// 创建一个新的token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	// 设置header
	token.Header["alg"] = "HS256"
	token.Header["sign_type"] = "SIGN"

	// 生成签名密钥
	signingKey := []byte(keys[1])

	// 生成JWT字符串
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
