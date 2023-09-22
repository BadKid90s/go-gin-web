package pkg

import "golang.org/x/crypto/bcrypt"

// HashPassword 哈希密码
func HashPassword(password string) (string, error) {
	// 生成密码的哈希值，使用默认的成本设置（10）
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// CheckPassword 验证密码
func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
