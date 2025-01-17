package uuid

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func GenUUID() string {
	return uuid.NewString()
}

// GenerateSMSCode 生成 6 位短信验证码
func GenerateSMSCode() string {
	// 使用 New() 创建新的随机数生成器，避免使用 global rand
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := r.Intn(1000000)          // 生成一个 0 到 999999 的随机数
	return fmt.Sprintf("%06d", code) // 返回 6 位的验证码，前面补零
}
