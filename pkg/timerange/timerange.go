package timerange

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// TimeRange 时间范围 (以分钟为单位，0-1440)
type TimeRange struct {
	Start int // 开始时间 (分钟，0-1439)
	End   int // 结束时间 (分钟，1-1440)
}

// IsInRange 检查指定时间 (分钟) 是否在范围内
func (tr TimeRange) IsInRange(minute int) bool {
	// 处理跨天的情况
	if tr.Start <= tr.End {
		// 普通情况：start <= end
		return minute >= tr.Start && minute < tr.End
	} else {
		// 跨天情况：start > end (例如 23:00-01:00)
		return minute >= tr.Start || minute < tr.End
	}
}

// TimeRanges 时间范围切片
type TimeRanges []TimeRange

// IsInAnyRange 检查指定时间是否在任一范围内
func (trs TimeRanges) IsInAnyRange(minute int) bool {
	for _, tr := range trs {
		if tr.IsInRange(minute) {
			return true
		}
	}
	return false
}

// String 返回时间范围的字符串表示
func (trs TimeRanges) String() string {
	var ret []string
	for _, tr := range trs {
		ret = append(ret, fmt.Sprintf("%d-%d", tr.Start, tr.End))
	}
	return strings.Join(ret, ",")
}

// GetCurrentMinuteOfDay 获取当前时间在一天中的分钟数 (0-1439)
func GetCurrentMinuteOfDay() int {
	now := time.Now()
	return now.Hour()*60 + now.Minute()
}

// ParseTimeRanges 解析时间范围字符串
// 支持格式：
//   - 新格式：HH:MM-HH:MM,HH:MM-HH:MM (如 "06:00-08:00,12:00-14:00")
//   - 兼容格式：分钟数格式 (如 "360-480,720-840")
//   - 分隔符：只支持逗号(,)
//   - 全天：00:00-23:59 或 0-1440
func ParseTimeRanges(timedStr string) (TimeRanges, error) {
	if timedStr == "" {
		return TimeRanges{{Start: 0, End: 1440}}, nil // 默认全天
	}

	// 按逗号分割多个范围
	ranges := strings.Split(timedStr, ",")
	var timeRanges TimeRanges

	for _, rangeStr := range ranges {
		rangeStr = strings.TrimSpace(rangeStr)
		if rangeStr == "" {
			continue
		}

		// 按连字符分割开始和结束时间
		parts := strings.Split(rangeStr, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("时间范围格式错误，应为 'HH:MM-HH:MM': %s", rangeStr)
		}

		// 解析开始时间
		start, err := parseTimeToMinutes(parts[0])
		if err != nil {
			return nil, fmt.Errorf("无效的开始时间: %w", err)
		}

		// 解析结束时间
		end, err := parseTimeToMinutes(parts[1])
		if err != nil {
			return nil, fmt.Errorf("无效的结束时间: %w", err)
		}

		// 对于结束时间，如果是 23:59，转换为 1440 (第二天 00:00)
		if end == 23*60+59 {
			end = 1440
		}

		// 验证范围 (允许跨天，如 23:00-01:00)
		timeRange := TimeRange{Start: start, End: end}
		timeRanges = append(timeRanges, timeRange)
	}

	if len(timeRanges) == 0 {
		return TimeRanges{{Start: 0, End: 1440}}, nil // 如果没有有效范围，默认全天
	}

	return timeRanges, nil
}

// parseTimeToMinutes 将 HH:MM 格式的时间字符串转换为从0点开始的分钟数
func parseTimeToMinutes(timeStr string) (int, error) {
	// 先尝试解析 HH:MM 格式
	if strings.Contains(timeStr, ":") {
		parts := strings.Split(timeStr, ":")
		if len(parts) != 2 {
			return 0, fmt.Errorf("时间格式错误，应为 HH:MM 格式: %s", timeStr)
		}

		hour, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil || hour < 0 || hour > 23 {
			return 0, fmt.Errorf("无效的小时: %s (应为0-23)", parts[0])
		}

		minute, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil || minute < 0 || minute > 59 {
			return 0, fmt.Errorf("无效的分钟: %s (应为0-59)", parts[1])
		}

		return hour*60 + minute, nil
	}

	// 兼容旧格式：直接使用分钟数
	minutes, err := strconv.Atoi(strings.TrimSpace(timeStr))
	if err != nil || minutes < 0 || minutes > 1440 {
		return 0, fmt.Errorf("无效的分钟数: %s (应为0-1440或HH:MM格式)", timeStr)
	}

	return minutes, nil
}

// FindNextAllowedTime 找到下一个允许的时间分钟
func FindNextAllowedTime(currentMinute int, timeRanges TimeRanges) int {
	// 先检查当天剩余时间
	for minute := currentMinute + 1; minute < 1440; minute++ {
		if timeRanges.IsInAnyRange(minute) {
			return minute
		}
	}

	// 如果当天没有找到，检查第二天
	for minute := 0; minute < currentMinute; minute++ {
		if timeRanges.IsInAnyRange(minute) {
			return minute
		}
	}

	// 如果都没找到，返回当前时间 (理论上不应该发生)
	return currentMinute
}

// CalculateSleepDuration 计算需要休眠的时间
func CalculateSleepDuration(currentMinute, nextAllowedMinute int) time.Duration {
	var sleepMinutes int

	if nextAllowedMinute > currentMinute {
		// 同一天
		sleepMinutes = nextAllowedMinute - currentMinute
	} else {
		// 跨天
		sleepMinutes = (1440 - currentMinute) + nextAllowedMinute
	}

	return time.Duration(sleepMinutes) * time.Minute
}
