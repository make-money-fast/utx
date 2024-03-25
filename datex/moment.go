package datex

import (
	"fmt"
	"time"
)

func Moment(t time.Time) string {
	n := time.Now().Sub(t)
	switch {
	case n.Seconds() < 60:
		return `刚刚`
	case n.Minutes() < 60:
		return fmt.Sprintf("%d分钟前", int(n.Minutes()))
	case n.Hours() < 24:
		return fmt.Sprintf("%d小时前", int(n.Hours()))
	case n.Hours() > 24*1 && n.Hours() <= 24*31:
		return "1个月内"
	case n.Hours() > 31*24:
		return "很久以前"
	default:
		return "不久以后"
	}
}

// BeforeNow 是否在过去
func BeforeNow(t time.Time) bool {
	return time.Since(t) > 0
}

// AfterNow 是否在未来
func AfterNow(t time.Time) bool {
	return !BeforeNow(t)
}

// SinceNow 获取当前 - 过去时间.
func SinceNow(t time.Time) time.Duration {
	return time.Since(t)
}
