package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	cn := runtime.NumCPU()
	betch := 10000
	sumList := make([]int, cn)
	wg := sync.WaitGroup{}
	for i := 0; i < cn; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			sum := 0
			for j := idx * betch; j < 1000000000000; j++ {
				sum += j
			}
			sumList[idx] = sum
		}(i)
	}
	wg.Wait()
	total := 0
	for _, v := range sumList {
		total += v
	}
	fmt.Println(total)
	fmt.Println("耗时：", time.Since(now))
	fmt.Println("done")
}
