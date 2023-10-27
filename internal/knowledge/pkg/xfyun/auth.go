package xfyun

import (
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

// getAuthorizationL 获取用户鉴权
func getAuthorization(apiKey string) (string, error) {
	keys := strings.Split(apiKey, ".")

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
