package main

import (
	"fmt"
	"time"

	"github.com/HCH1212/gofish/core"
)

func main() {
	for {
		time.Sleep(2 * time.Second)
		name := core.GetActiveWindow()
		fmt.Println("当前活动窗口:", name)
	}
}
