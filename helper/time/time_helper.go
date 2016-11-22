package time

import "time"

// CurrentMillis : 获取当前时间，毫秒
func CurrentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
