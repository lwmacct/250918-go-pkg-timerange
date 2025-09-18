# TimeRange - Go 时间范围管理包

[![Go Version](https://img.shields.io/github/go-mod/go-version/lwmacct/250918-go-pkg-timerange)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/lwmacct/250918-go-pkg-timerange)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/lwmacct/250918-go-pkg-timerange)](https://goreportcard.com/report/github.com/lwmacct/250918-go-pkg-timerange)

一个功能强大的 Go 时间范围管理包，专为需要在特定时间段内控制任务执行而设计。支持跨天时间范围、多时间段组合以及灵活的时间格式。

## ✨ 功能特性

- 🌅 **跨天支持**: 完美处理跨天时间范围（如 23:00-01:00）
- 🔄 **多范围组合**: 支持多个时间段的组合使用
- 📅 **灵活格式**: 支持 HH:MM 格式和分钟数格式
- ⏰ **实时检查**: 提供当前时间是否在允许范围内的快速检查
- 🔍 **智能查找**: 自动查找下一个允许的执行时间
- ⏱️ **等待计算**: 精确计算需要等待的时间长度

## 🚀 快速开始

### 安装

```bash
go get github.com/lwmacct/250918-go-pkg-timerange
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lwmacct/250918-go-pkg-timerange/pkg/timerange"
)

func main() {
    // 解析时间范围字符串
    ranges, err := timerange.ParseTimeRanges("06:00-08:00,12:00-14:00")
    if err != nil {
        panic(err)
    }

    // 获取当前时间（分钟数）
    currentMinute := timerange.GetCurrentMinuteOfDay()
    
    // 检查当前时间是否在任一范围内
    if ranges.IsInAnyRange(currentMinute) {
        fmt.Println("✅ 当前时间在允许范围内，可以执行任务")
    } else {
        fmt.Println("❌ 当前时间不在允许范围内，需要等待")
        
        // 查找下一个允许时间
        nextMinute := timerange.FindNextAllowedTime(currentMinute, ranges)
        sleepDuration := timerange.CalculateSleepDuration(currentMinute, nextMinute)
        fmt.Printf("⏳ 需要等待 %v 直到下一个允许时间\n", sleepDuration)
    }
}
```

### 高级用法

```go
// 跨天时间范围示例
ranges, _ := timerange.ParseTimeRanges("23:00-01:00,06:00-08:00")

// 创建单个时间范围
tr := timerange.TimeRange{Start: 360, End: 480} // 6:00-8:00
if tr.IsInRange(420) { // 检查 7:00 是否在范围内
    fmt.Println("7:00 在允许范围内")
}
```

## 📖 详细文档

更多使用方法和 API 详细说明，请查看：

- [📋 Go Doc API 文档](https://pkg.go.dev/github.com/lwmacct/250918-go-pkg-timerange/pkg/timerange)
- [📝 更新日志](CHANGELOG.md)
- [🔒 安全政策](SECURITY.md)

## 🛠️ 开发环境

本项目使用现代化的开发工具链：

- **任务管理**: [Taskfile](https://taskfile.dev) - 查看所有可用任务
- **开发环境**: [Dev Container](https://code.visualstudio.com/docs/devcontainers/containers) - 一键搭建开发环境
- **Go 版本**: 1.25.1+

### 开发命令

```bash
# 查看所有可用任务
task -a

# 运行测试
task test

# 构建项目
task build
```

### 开发环境设置

1. 使用 Dev Container（推荐）：
   - 在 VS Code 中打开项目
   - 选择 "Reopen in Container"
   - 环境将自动配置完成

2. 手动设置：
   - 确保安装 Go 1.25.1+
   - 安装 [Taskfile](https://taskfile.dev/installation/)
   - 运行 `go mod download`

## 🤝 贡献指南

欢迎贡献代码！请遵循以下步骤：

1. Fork 这个仓库
2. 创建你的功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的修改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启一个 Pull Request

## 📄 许可证

本项目基于 [LICENSE](LICENSE) 许可证开源。

## 🔗 相关链接

- [GitHub 仓库](https://github.com/lwmacct/250918-go-pkg-timerange)
- [作者主页](https://github.com/lwmacct)
- [Dev Container 文档](https://www.yuque.com/lwmacct/vscode/dev-containers)

---

<div align="center">
Made with ❤️ by <a href="https://github.com/lwmacct">lwmacct</a>
</div>
