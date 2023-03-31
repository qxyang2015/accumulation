package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	LimitNum = 100000000
)

func main() {
	fmt.Println("start")
	t1 := time.Now()
	fmt.Println("计算结果为1：", Sum())
	fmt.Println("计算耗时1：", time.Since(t1).String())

	t2 := time.Now()
	fmt.Println("计算结果为2：", Sum2())
	fmt.Println("计算耗时2：", time.Since(t2).String())

	t3 := time.Now()
	fmt.Println("计算结果为3：", Sum3())
	fmt.Println("计算耗时3：", time.Since(t3).String())
	fmt.Println("done!")
}

func Sum() int {
	res := 0
	for i := 1; i < LimitNum; i++ {
		res += i
	}
	return res
}

func Sum2() int64 {
	var res int64
	n := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			start := (LimitNum / n) * idx
			end := start + (LimitNum / n)
			for j := start; j < end; j++ {
				atomic.AddInt64(&res, int64(j))
			}
		}(i)
	}
	wg.Wait()
	return res
}

func Sum3() int {
	res := 0
	n := runtime.GOMAXPROCS(0)
	s := make([]int, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			start := (LimitNum / n) * i
			end := start + (LimitNum / n)
			for j := start; j < end; j++ {
				s[idx] += j
			}
		}(i)
	}
	wg.Wait()
	for _, v := range s {
		res += v
	}
	return res
}

func Sum4() int {
	res := 0
	n := runtime.GOMAXPROCS(0)
	var lc sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			start := (LimitNum / n) * idx
			end := start + (LimitNum / n)
			for j := start; j < end; j++ {
				lc.Lock()
				res += j
				lc.Unlock()
			}
		}(i)
	}
	wg.Wait()
	return res
}
