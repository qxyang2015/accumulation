package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

func main() {
	v3, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("v3", err)
		return
	}
	fmt.Println("cpu percent：", v3[0])
	vm, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("获取内存信息错误", err)
		return
	}
	fmt.Println("memeory used percent", vm.UsedPercent)
	fmt.Println("done")
}
