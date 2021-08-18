package common

import (
	"math/rand"
	"time"
)

//生成[min, max]之间的随机数
func GetRandInt(min, max int) int {
	if min >= max {
		return min
	}
	rand.Seed(int64(time.Now().Nanosecond())) //种子
	var ret = rand.Intn(max + 1 - min)
	return ret + min
}

// z := x op y ? x:y
func GetCaseNum(op string, x int, y int) int {
	if op == "<" {
		if x > y {
			return y
		}
	} else if op == ">" {
		if x < y {
			return y
		}
	}
	return x
}
