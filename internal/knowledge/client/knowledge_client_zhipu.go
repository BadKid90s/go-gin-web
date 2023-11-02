package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type knowledgeClientZhipu struct {
	apiKey       string
	embeddingUrl string
	modelUrl     string
}

func newKnowledgeClientZhipu(conf *viper.Viper) KnowledgeClient {
	apiKey := conf.GetString("embedding.zhipu.apiKey")
	embeddingUrl := conf.GetString("embedding.zhipu.embeddingUrl")
	modelUrl := conf.GetString("embedding.zhipu.modelUrl")

	return &knowledgeClientZhipu{
		apiKey:       apiKey,
		embeddingUrl: embeddingUrl,
		modelUrl:     modelUrl,
	}
}
func (c *knowledgeClientZhipu) ChatMessage(text string) (chan string, error) {
	authorization, err := c.getAuthorization()
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient

	param := map[string]any{
		"prompt": []map[string]any{
			{
				"role":    "user",
				"content": text,
			},
		},
	}

	jsonData, _ := json.Marshal(param)

	request, err := http.NewRequest("POST", c.modelUrl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", authorization)
	if err != nil {
		return nil, err

	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err

	}

	eventChan := make(chan string) // 创建一个用于存储SSE事件的通道

	go func() {
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
		decoder := NewDecoder(resp.Body) // 自定义SSE解码器
		for {
			event, err := decoder.Decode()
			if err != nil {
				if err == io.EOF {
					close(eventChan) // 关闭通道
					return
				}
				log.Println(err)
			}
			eventChan <- event.Data // 将解码后的SSE事件放入通道中
		}
	}()
	return eventChan, nil
}

func (c *knowledgeClientZhipu) GetEmbedding(text string) ([]float32, error) {
	authorization, err := c.getAuthorization()
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient

	param := map[string]string{
		"prompt": text,
	}

	jsonData, _ := json.Marshal(param)

	request, err := http.NewRequest("POST", c.modelUrl, bytes.NewBuffer(jsonData))
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
func (c *knowledgeClientZhipu) getAuthorization() (string, error) {
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

// SSEEvent 包含SSE事件的各个字段
type SSEEvent struct {
	Data  string
	Event string
	ID    string
	Retry int
}

// SSEDecoder SSE解码器
type SSEDecoder struct {
	reader *bufio.Reader
}

// NewDecoder 创建一个新的SSE解码器
func NewDecoder(r io.Reader) *SSEDecoder {
	return &SSEDecoder{
		reader: bufio.NewReader(r),
	}
}

// Decode 解码SSE消息为事件对象
func (d *SSEDecoder) Decode() (*SSEEvent, error) {
	event := &SSEEvent{}

	for {
		line, err := d.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.TrimRight(line, "\r\n")

		if line == "" {
			// 空行表示SSE消息的结束，返回解码后的事件对象
			return event, nil
		}

		field := strings.SplitN(line, ":", 2)
		//if len(field) != 2 {
		//	return nil, errors.New("invalid field line: " + line)
		//}

		switch field[0] {
		case "data":
			event.Data = field[1]
		case "event":
			event.Event = field[1]
		case "id":
			event.ID = field[1]
		case "retry":
			// 将retry字段解析为整数
			retry, err := strconv.Atoi(field[1])
			if err != nil {
				return nil, errors.New("invalid retry value: " + field[1])
			}
			event.Retry = retry
		}
	}
}
