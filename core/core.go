package core

import (
	"fmt"
	"strings"
	"time"
)

// FishTime 累积摸鱼时间
var FishTime int

// IsActivity 判断程序是否活跃
func IsActivity(processNames []string) bool {
	wids := getWindowIDs()
	pids := getProcessIDs(wids)
	names := getProcessNames(pids)

	for _, name := range names {
		for _, pname := range processNames {
			if pname == "" {
				continue
			}
			if strings.Contains(strings.ToLower(name), strings.ToLower(pname)) {
				return true
			}
		}
	}
	return false
}

// RunMonitor 持续监控程序是否活跃，活跃则累计摸鱼时间
func RunMonitor(processNames []string) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		if IsActivity(processNames) {
			FishTime += 5
			fmt.Printf("又在摸鱼了宝宝，累计摸鱼时间：%d 秒\n", FishTime)
		} else {
			fmt.Println("太认真啦乖宝宝~")
		}
	}
}
