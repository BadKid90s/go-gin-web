package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type knowledgeClientXunfei struct {
	appid        string
	apiKey       string
	apiSecret    string
	embeddingUrl string
	modelUrl     string
}

func newKnowledgeClientXunfei(conf *viper.Viper) KnowledgeClient {
	appid := conf.GetString("embedding.xunfei.appid")
	apiKey := conf.GetString("embedding.xunfei.apiKey")
	apiSecret := conf.GetString("embedding.xunfei.apiSecret")
	embeddingUrl := conf.GetString("embedding.xunfei.embeddingUrl")
	modelUrl := conf.GetString("embedding.xunfei.modelUrl")

	return &knowledgeClientXunfei{
		appid:        appid,
		apiKey:       apiKey,
		apiSecret:    apiSecret,
		embeddingUrl: embeddingUrl,
		modelUrl:     modelUrl,
	}
}

func (c *knowledgeClientXunfei) ChatMessage(text string) (chan string, error) {
	targetUrl, err := c.getUrl(c.modelUrl, "GET")
	conn, _, err := websocket.DefaultDialer.Dial(targetUrl, nil)
	if err != nil {
		log.Println("dial:", err)
	}

	params := c.genParams(c.appid, text)
	err = conn.WriteJSON(params)
	if err != nil {
		log.Println("WriteMessage:", err)
	}

	eventChan := make(chan string) // 创建一个用于存储SSE事件的通道

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("read message error:", err)
				break
			}

			var data map[string]interface{}
			err1 := json.Unmarshal(msg, &data)
			if err1 != nil {
				log.Println("Error parsing JSON:", err)
				close(eventChan)
				return
			}

			header := data["header"].(map[string]interface{})
			code := header["code"].(float64)
			if code != 0 {
				message := header["message"].(string)
				log.Println("Xfyun err:", errors.New(fmt.Sprintf("code:%f,message:%s", code, message)))
				close(eventChan)
				return
			}
			//解析数据
			payload := data["payload"].(map[string]interface{})
			choices := payload["choices"].(map[string]interface{})
			status := choices["status"].(float64)
			text := choices["text"].([]interface{})
			content := text[0].(map[string]interface{})["content"].(string)
			if status != 2 {
				eventChan <- content
			} else {
				log.Println("收到最终结果")
				eventChan <- content
				usage := payload["usage"].(map[string]interface{})
				temp := usage["text"].(map[string]interface{})
				totalTokens := temp["total_tokens"].(float64)
				log.Println("total_tokens:", totalTokens)
				_ = conn.Close()
				close(eventChan)
				break
			}
		}
	}()
	return eventChan, nil
}

func (c *knowledgeClientXunfei) GetEmbedding(text string) ([]float32, error) {
	targetUrl, err := c.getUrl(c.embeddingUrl, "POST")
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
		Header struct {
			Code    int
			Message string
			Sid     string
		}
		Payload struct {
			Text struct {
				Vector string
			}
		}
	}
	var result JsonResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return c.strToFloatArray(result.Payload.Text.Vector)
}

// getAuthorization 获取用户鉴权
func (c *knowledgeClientXunfei) getUrl(URL string, method string) (string, error) {
	parse, err := url.Parse(URL)
	if err != nil {
		return "", err
	}
	host := parse.Host
	path := parse.Path

	now := time.Now().UTC()
	date := now.Format(time.RFC1123)

	signatureOrigin := fmt.Sprintf("host: %s\ndate: %s\n%s %s HTTP/1.1", host, date, method, path)
	h := hmac.New(sha256.New, []byte(c.apiSecret))
	h.Write([]byte(signatureOrigin))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	authorizationOrigin := fmt.Sprintf(`api_key="%s", algorithm="hmac-sha256", headers="host date request-line", signature="%s"`, c.apiKey, signature)
	authorization := base64.StdEncoding.EncodeToString([]byte(authorizationOrigin))

	v := url.Values{}
	v.Set("authorization", authorization)
	v.Set("date", date)
	v.Set("host", host)

	return fmt.Sprintf("%s?%s", URL, v.Encode()), nil
}

func (c *knowledgeClientXunfei) strToFloatArray(str string) ([]float32, error) {
	// 删除字符串中的 "[" 和 "]" 字符
	str = str[1 : len(str)-1]

	// 按 "," 分割字符串并存储为 slice
	strArr := strings.Split(str, ",")

	// 创建长度为 slice 长度的 float32 数组
	floatArr := make([]float32, len(strArr))

	// 转换每个字符串元素为 float32 类型，并存储到数组中
	for i, s := range strArr {
		f, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return nil, err
		}
		floatArr[i] = float32(f)
	}
	return floatArr, nil
}

// 生成参数
func (c *knowledgeClientXunfei) genParams(appid, question string) map[string]interface{} { // 根据实际情况修改返回的数据结构和字段名

	messages := []Message{
		{Role: "user", Content: question},
	}

	data := map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
		"header": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"app_id": appid, // 根据实际情况修改返回的数据结构和字段名
		},
		"parameter": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"chat": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"domain":      "generalv2",  // 根据实际情况修改返回的数据结构和字段名
				"temperature": float64(0.8), // 根据实际情况修改返回的数据结构和字段名
				"top_k":       int64(6),     // 根据实际情况修改返回的数据结构和字段名
				"max_tokens":  int64(2048),  // 根据实际情况修改返回的数据结构和字段名
				"auditing":    "default",    // 根据实际情况修改返回的数据结构和字段名
			},
		},
		"payload": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"message": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"text": messages, // 根据实际情况修改返回的数据结构和字段名
			},
		},
	}
	return data // 根据实际情况修改返回的数据结构和字段名
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
