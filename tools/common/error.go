package common

import (
	"fmt"
	"runtime"
)

func Assert(err error) {
	if err != nil {
		panic(err)
	}
}

//捕获panic转为log输出， 建议加入有可能panic的函数中，用defer方式调用
func CatchPanicLog(rcv interface{}) string {
	errmsg := fmt.Sprintf("panic[%v]\n", rcv)
	errmsg += "/---------------------------------------/\n"
	//获取上层调用者信息
	for skip := 0; ; skip++ {
		if pc, file, line, ok := runtime.Caller(skip); ok {
			funcinfo := runtime.FuncForPC(pc)
			var funcname string = "unknown"
			if funcinfo != nil {
				funcname = funcinfo.Name()
			}
			errmsg += fmt.Sprintf("file[%s],line[%d],funcname[%s]\n", file, line, funcname)
		} else {
			break
		}
	}
	errmsg += "/---------------------------------------/"
	return errmsg
}
