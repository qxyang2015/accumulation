package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"time"
	// "github.com/shirou/gopsutil/mem"  // to use v2
)

func main() {
	v1, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Println("v1", err)
		return
	}
	v2, err := cpu.Percent(0, true)
	if err != nil {
		fmt.Println("v2", err)
		return
	}
	sum := 0.0
	for _, v := range v2 {
		sum += v
	}
	fmt.Println("不间隔时间:", v1[0], sum/float64(len(v2)))
	for i := 0; i < 10; i++ {
		v3, err := cpu.Percent(time.Second, false)
		if err != nil {
			fmt.Println("v3", err)
			return
		}
		v4, err := cpu.Percent(time.Second, true)
		if err != nil {
			fmt.Println("v4", err)
			return
		}
		sum := 0.0
		for _, v := range v4 {
			sum += v
		}
		fmt.Println("间隔1s时间:", v3[0], sum/float64(len(v4)))
	}
	fmt.Println("done")
}
