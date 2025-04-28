package core

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetActiveWindow 获取当前活动窗口的标题
func GetActiveWindow() string {
	// 获取当前活动窗口的 ID
	cmd := exec.Command("xprop", "-root", "_NET_ACTIVE_WINDOW")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("无法获取活动窗口 ID:", err)
		return ""
	}

	// 提取窗口 ID
	windowID := extractWindowID(string(output))
	if windowID == "" {
		fmt.Println("未找到活动窗口 ID")
		return ""
	}

	// 获取窗口标题
	cmd = exec.Command("xprop", "-id", windowID, "_NET_WM_NAME")
	output, err = cmd.Output()
	if err != nil {
		fmt.Println("无法获取窗口标题:", err)
		return ""
	}

	// 提取窗口标题
	title := extractWindowTitle(string(output))
	return title
}

// extractWindowID 从 xprop 输出中提取窗口 ID
func extractWindowID(output string) string {
	parts := strings.Fields(output)
	if len(parts) > 0 {
		return parts[len(parts)-1] // 窗口 ID 通常是最后一个字段
	}
	return ""
}

// extractWindowTitle 从 xprop 输出中提取窗口标题
func extractWindowTitle(output string) string {
	parts := strings.Split(output, "=")
	if len(parts) > 1 {
		return strings.Trim(parts[1], " \"\n")
	}
	return ""
}
