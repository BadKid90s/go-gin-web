package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
	"time"
)

type knowledgeClientXunfei struct {
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

	return &knowledgeClientXunfei{
		appid:     appid,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		url:       url,
	}
}

func (c *knowledgeClientXunfei) GetEmbedding(text string) ([]float32, error) {
	targetUrl, err := c.getUrl()
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient

	param := map[string]any{
		"header": map[string]string{
			"app_id": c.appid,
		},
		"payload": map[string]string{
			"text": text,
		},
	}

	jsonData, _ := json.Marshal(param)

	request, err := http.NewRequest("POST", targetUrl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
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
		Payload struct {
			Text struct {
				vector []float32
			}
		}
	}
	var result JsonResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result.Payload.Text.vector, nil
}

// getAuthorization 获取用户鉴权
func (c *knowledgeClientXunfei) getUrl() (string, error) {
	parse, err := url.Parse(c.url)
	if err != nil {
		return "", err
	}
	host := parse.Host
	path := parse.Path

	now := time.Now().UTC()
	date := now.Format(time.RFC1123)

	signatureOrigin := fmt.Sprintf("host: %s\ndate: %s\nGET %s HTTP/1.1", host, date, path)
	h := hmac.New(sha256.New, []byte(c.apiSecret))
	h.Write([]byte(signatureOrigin))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	authorizationOrigin := fmt.Sprintf(`api_key="%s", algorithm="hmac-sha256", headers="host date request-line", signature="%s"`, c.apiKey, signature)
	authorization := base64.StdEncoding.EncodeToString([]byte(authorizationOrigin))

	v := url.Values{}
	v.Set("authorization", authorization)
	v.Set("date", date)
	v.Set("host", host)

	return fmt.Sprintf("%s?%s", c.url, v.Encode()), nil
}
