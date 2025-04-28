package core

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"time"
)

var lastActivityTime time.Time

// inputEvent 定义 Linux 输入事件结构
type inputEvent struct {
	Time  [2]uint64 // 时间戳（秒和微秒）
	Type  uint16    // 事件类型
	Code  uint16    // 事件代码
	Value int32     // 事件值
}

func GetLastActivityTime() time.Time {
	return lastActivityTime
}

// UpdateActivity 更新最后活动时间
func UpdateActivity() {
	lastActivityTime = time.Now()
}

// IsIdle 检查是否超过指定时间无活动
func IsIdle(idleDuration time.Duration) bool {
	return time.Since(lastActivityTime) > idleDuration
}

// MonitorActivity 根据文件名持续监测鼠标和键盘活动
func MonitorActivity(devicePath string) {
	// 打开输入设备文件
	file, err := os.Open(devicePath)
	if err != nil {
		panic(fmt.Sprintf("无法打开设备文件 %s: %v", devicePath, err))
	}
	defer file.Close()

	for {
		var event inputEvent
		// 从设备文件中读取事件
		err := binary.Read(file, binary.LittleEndian, &event)
		if err != nil {
			continue
		}

		// 检测键盘或鼠标活动
		if event.Type == 1 && event.Value == 1 { // 键盘按键按下
			UpdateActivity()
		} else if event.Type == 2 { // 鼠标移动
			UpdateActivity()
		}
	}
}

// FindInputDevices 获取所有与input设备有关的文件名
func FindInputDevices() []string {
	file, err := os.Open("/proc/bus/input/devices")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var devices []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "H: Handlers=") {
			// 检查是否包含键盘或鼠标的标识符
			if strings.Contains(line, "kbd") || strings.Contains(line, "mouse") {
				parts := strings.Fields(line)
				for _, part := range parts {
					if strings.HasPrefix(part, "event") {
						devices = append(devices, "/dev/input/"+part)
					}
				}
			}
		}
	}
	return devices
}
