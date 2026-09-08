// Package format 提供字节大小等数值的人性化格式化工具。
// 本文件将字节数转换为人类可读的字符串(如 "1.5 GB"),用于 CLI 输出与日志展示。
package format

import (
	"fmt"
	"math"
)

// 字节单位常量:SI 十进制单位(1000 进制)与 IEC 二进制单位(1024 进制)。
const (
	Byte = 1

	// 十进制单位(KB/MB/GB/TB),按 1000 进位,常用于面向用户的展示。
	KiloByte = Byte * 1000
	MegaByte = KiloByte * 1000
	GigaByte = MegaByte * 1000
	TeraByte = GigaByte * 1000

	// 二进制单位(KiB/MiB/GiB),按 1024 进位,常用于精确内存/显存统计。
	KibiByte = Byte * 1024
	MebiByte = KibiByte * 1024
	GibiByte = MebiByte * 1024
)

// HumanBytes 把 int64 字节数格式化为十进制单位的人性化字符串。
// 规则:从 TB 到 KB 逐级匹配最大可用单位;
//   - 数值 >= 10 时取整显示(如 "42 GB");
//   - 数值带小数时保留 1 位小数(如 "1.5 GB");
//   - 不足 1KB 时直接返回 "N B"。
func HumanBytes(b int64) string {
	var value float64
	var unit string

	switch {
	case b >= TeraByte:
		value = float64(b) / TeraByte
		unit = "TB"
	case b >= GigaByte:
		value = float64(b) / GigaByte
		unit = "GB"
	case b >= MegaByte:
		value = float64(b) / MegaByte
		unit = "MB"
	case b >= KiloByte:
		value = float64(b) / KiloByte
		unit = "KB"
	default:
		return fmt.Sprintf("%d B", b)
	}

	switch {
	case value >= 10:
		return fmt.Sprintf("%d %s", int(value), unit)
	case value != math.Trunc(value):
		return fmt.Sprintf("%.1f %s", value, unit)
	default:
		return fmt.Sprintf("%d %s", int(value), unit)
	}
}

// HumanBytes2 把 uint64 字节数格式化为二进制单位(KiB/MiB/GiB)的人性化字符串。
// 与 HumanBytes 的区别:使用 1024 进制单位,且始终保留 1 位小数。
// 典型用途:显存统计等需要二进制语义的场景。
func HumanBytes2(b uint64) string {
	switch {
	case b >= GibiByte:
		return fmt.Sprintf("%.1f GiB", float64(b)/GibiByte)
	case b >= MebiByte:
		return fmt.Sprintf("%.1f MiB", float64(b)/MebiByte)
	case b >= KibiByte:
		return fmt.Sprintf("%.1f KiB", float64(b)/KibiByte)
	default:
		return fmt.Sprintf("%d B", b)
	}
}
