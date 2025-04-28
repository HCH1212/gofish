package core

import (
    "fmt"
    "os"
    "time"
)

// LogFishingTime 记录摸鱼时间
func LogFishingTime(duration time.Duration) {
    f, err := os.OpenFile("fishing.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("无法打开日志文件:", err)
        return
    }
    defer f.Close()

    log := fmt.Sprintf("%s - 摸鱼时间: %v\n", time.Now().Format("2006-01-02 15:04:05"), duration)
    f.WriteString(log)
}