package core

import (
	"fmt"
	"os/exec"
	"strings"
)

// getWindowIDs 获取活跃的窗口id
func getWindowIDs() []string {
	cmd := exec.Command("wmctrl", "-l")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("获取窗口列表失败:", err)
		return nil
	}

	var windowIDs []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) > 0 {
			windowIDs = append(windowIDs, parts[0])
		}
	}
	return windowIDs
}

// getProcessIDs 获取对应进程id
func getProcessIDs(windowIDs []string) []string {
	var pids []string
	for _, wid := range windowIDs {
		cmd := exec.Command("xprop", "-id", wid)
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("获取窗口属性失败（窗口ID %s）: %v\n", wid, err)
			pids = append(pids, "")
			continue
		}

		pid := ""
		for _, line := range strings.Split(string(output), "\n") {
			if strings.Contains(line, "_NET_WM_PID(CARDINAL) =") {
				parts := strings.Split(line, "=")
				if len(parts) == 2 {
					pid = strings.TrimSpace(parts[1])
				}
			}
		}
		pids = append(pids, pid)
	}
	return pids
}

// getProcessNames 获取对应进程名字
func getProcessNames(processIDs []string) []string {
	var names []string
	for _, pid := range processIDs {
		if pid == "" {
			names = append(names, "")
			continue
		}
		cmd := exec.Command("ps", "-p", pid, "-o", "comm=")
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("获取进程名失败（PID %s）: %v\n", pid, err)
			names = append(names, "")
			continue
		}
		names = append(names, strings.TrimSpace(string(output)))
	}
	return names
}
