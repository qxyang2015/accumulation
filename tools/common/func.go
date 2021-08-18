package common

import (
	"flag"
	"math"
)

//获取命令行启动参数　无默认值
func GetCommandLineParam(aSlice []string) map[string]*string {

	ret := make(map[string]*string)
	for _, k := range aSlice {
		ret[k] = new(string)
		flag.StringVar(ret[k], k, "", "usage")
	}
	flag.Parse()
	for k, v := range ret {
		if len(*v) == 0 {
			errmsg := k + "信息缺失"
			panic(errmsg)
		}
	}
	return ret
}

func Margin2Probs(vals []float64) []float64 {
	probs := make([]float64, len(vals))
	for idx, val := range vals {
		probs[idx] = 1.0 / (1.0 + math.Exp(-val))
	}
	return probs
}
