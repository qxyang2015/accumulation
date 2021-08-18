package common

import (
	"fmt"
	"strings"
)

const (
	SEP_T = "\t"
)

//新增获取slice内获取值位置函数
func GetIndexFromStr(s string, tag string) int {
	ss := strings.Split(s, SEP_T)
	for i, v := range ss {
		if v == tag {
			return i
		}
	}
	return -1
}

func GetIndexsFromStr(s string, tags []string) ([]int, error) {
	ss := strings.Split(strings.TrimSpace(s), SEP_T)
	ssMap := make(map[string]int)
	for idx, t := range ss {
		ssMap[t] = idx
	}
	var ret []int
	for i, v := range tags {
		if idx, ok := ssMap[v]; ok {
			ret = append(ret, idx)
		} else {
			return ret, fmt.Errorf("发现未找到的列[%v]列名[%v]", i, v)
		}
	}
	return ret, nil
}
