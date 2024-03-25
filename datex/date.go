package datex

import "time"

const (
	LayoutYMDHIS = "2006-01-02 15:04:05"
	LayoutYMD    = "2006-01-02"
	LayoutHIS    = "15:04:05"
	LayoutHI     = "15:04"
	LayoutMD     = "01-12"
)

const (
	Sec   = time.Second
	Min   = 60 * time.Second
	Hour  = 60 * Min
	Day   = 24 * Hour
	Month = 31 * Day
	Year  = 365 * Day
)

// Timestamp 秒
func Timestamp() int64 {
	return time.Now().Unix()
}

// TimestampMS 毫秒
func TimestampMS() int64 {
	return time.Now().UnixMilli()
}

func FromTimestamp(n int64) time.Time {
	return time.Unix(n, 0)
}

func FromTimestampMS(n int64) time.Time {
	seconds := n / 1000                 // 转换为秒
	nanoseconds := (n % 1000) * 1000000 // 转换为纳秒
	// 使用time.Unix()函数将秒和纳秒转换为time.Time类型
	return time.Unix(seconds, nanoseconds)
}

// Format 格式化时间.
func Format(t time.Time, layouts ...string) string {
	var layout = LayoutYMDHIS
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	return t.Format(layout)
}
