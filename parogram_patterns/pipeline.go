package parogram_patterns

import (
	"fmt"
	"math"
	"sync"
	"testing"
)

func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func odd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 != 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

type EchoFunc func([]int) <-chan int
type PipeFunc func(<-chan int) <-chan int

func pipeline(nums []int, echo EchoFunc, pipeFns ...PipeFunc) <-chan int {
	ch := echo(nums)
	for i := range pipeFns {
		ch = pipeFns[i](ch)
	}
	return ch
}

func pipelineTest1(t testing.T) {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for n := range pipeline(nums, echo, odd, sq, sum) {
		fmt.Println(n)
	}
}

/*
动用 Go 语言的 Go Routine 和 Channel 还有一个好处，就是可以写出 1 对多，或多对 1 的 Pipeline，也就是 Fan In/ Fan Out。下面，我们来看一个 Fan in 的示例。
假设我们要通过并发的方式对一个很长的数组中的质数进行求和运算，我们想先把数组分段求和，然后再把它们集中起来。
下面是我们的主函数：
*/

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func pipelineTest2(t testing.T) {
	nums := makeRange(1, 10000)
	in := echo(nums)

	const nProcess = 5
	var chans [nProcess]<-chan int
	for i := range chans {
		chans[i] = sum(prime(in))
	}

	for n := range sum(merge(chans[:])) {
		fmt.Println(n)
	}
}

func is_prime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func prime(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if is_prime(n) {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

/*
我来简单解释下这段代码。首先，我们制造了从 1 到 10000 的数组；然后，把这堆数组全部 echo到一个 Channel 里—— in；
此时，生成 5 个 Channel，接着都调用 sum(prime(in)) ，
于是，每个 Sum 的 Go Routine 都会开始计算和；最后，再把所有的结果再求和拼起来，得到最终的结果。
*/
func merge(cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
