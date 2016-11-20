package number

import (
	"math/rand"
	"time"
)

// GenerateInt : 随机生成 n 个 最大为 max 的 Int 切片
func GenerateInt(n, max int) (data []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		data = append(data, rand.Intn(max))
	}
	return
}
