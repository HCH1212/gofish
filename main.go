package main

import (
	"fmt"
	"time"

	"github.com/HCH1212/gofish/core"
)

func main() {
	// 初始化
	core.UpdateActivity()
	idleDuration := 15 * time.Minute // input设备检查
	checkInterval := 5 * time.Minute // 页面检查

	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for {
		<-ticker.C

		// 检查鼠标和键盘活动
		if core.IsIdle(idleDuration) {
			fmt.Println("用户长时间无活动，视为摸鱼！")
			core.LogFishingTime(checkInterval)
			continue
		}

		// 检查顶级窗口
		windowTitle := core.GetActiveWindow()
		if core.IsFishingPage(windowTitle, fishingKeywords) {
			fmt.Printf("当前窗口标题: %s，视为摸鱼！\n", windowTitle)
			core.LogFishingTime(checkInterval)
		} else {
			fmt.Println("未检测到摸鱼行为，继续努力！")
		}
	}
}
