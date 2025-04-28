package core

import "strings"

// IsFishingPage 判断当前页面是否为摸鱼页面
func IsFishingPage(windowTitle string, fishingKeywords []string) bool {
    for _, keyword := range fishingKeywords {
        if strings.Contains(strings.ToLower(windowTitle), strings.ToLower(keyword)) {
            return true
        }
    }
    return false
}