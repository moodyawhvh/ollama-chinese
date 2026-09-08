// Package format 提供时间等数值的人性化格式化工具。
// 本文件把 time.Time 转换为 "4 hours ago"、"About a minute" 这类人类可读的相对时间描述。
package format

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// humanDuration 返回时长的人性化近似描述(如 "About a minute"、"4 hours ago" 等)。
// 分级规则:秒 -> 分钟 -> 小时 -> 天 -> 周 -> 月 -> 年,
// 每级都留出约两倍的缓冲区间,避免边界值在相邻单位间来回跳动。
// 注意:输出为英文文案,因该结果直接展示给用户且历史行为如此,保持原样。
func humanDuration(d time.Duration) string {
	seconds := int(d.Seconds())

	switch {
	case seconds < 1:
		return "Less than a second"
	case seconds == 1:
		return "1 second"
	case seconds < 60:
		return fmt.Sprintf("%d seconds", seconds)
	}

	minutes := int(d.Minutes())
	switch {
	case minutes == 1:
		return "About a minute"
	case minutes < 60:
		return fmt.Sprintf("%d minutes", minutes)
	}

	hours := int(math.Round(d.Hours()))
	switch {
	case hours == 1:
		return "About an hour"
	case hours < 48:
		return fmt.Sprintf("%d hours", hours)
	case hours < 24*7*2:
		// 2 天到 14 天之间:以"天"为单位
		return fmt.Sprintf("%d days", hours/24)
	case hours < 24*30*2:
		// 2 周到 2 个月之间:以"周"为单位
		return fmt.Sprintf("%d weeks", hours/24/7)
	case hours < 24*365*2:
		// 2 个月到 2 年之间:以"月"为单位
		return fmt.Sprintf("%d months", hours/24/30)
	}

	// 超过两年:以"年"为单位
	return fmt.Sprintf("%d years", int(d.Hours())/24/365)
}

// HumanTime 返回给定时间相对于当前时刻的人性化描述。
// 若时间为零值,则返回调用方提供的 zeroValue(例如 "never")。
func HumanTime(t time.Time, zeroValue string) string {
	return humanTime(t, zeroValue)
}

// HumanTimeLower 与 HumanTime 相同,但把结果整体转为小写,
// 用于对大小写敏感的展示场景(如表格统一风格)。
func HumanTimeLower(t time.Time, zeroValue string) string {
	return strings.ToLower(humanTime(t, zeroValue))
}

// humanTime 是 HumanTime/HumanTimeLower 的内部实现:
//   - 零值时间直接返回 zeroValue;
//   - 目标时间早于 20 年以上(例如 Unix 纪元附近的占位时间)视为 "Forever";
//   - 未来时间返回 "xxx from now",过去时间返回 "xxx ago"。
func humanTime(t time.Time, zeroValue string) string {
	if t.IsZero() {
		return zeroValue
	}

	delta := time.Since(t)
	if int(delta.Hours())/24/365 < -20 {
		return "Forever"
	} else if delta < 0 {
		return humanDuration(-delta) + " from now"
	}

	return humanDuration(delta) + " ago"
}
