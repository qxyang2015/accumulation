package common

import (
	"time"
)

var Loc, _ = time.LoadLocation("Asia/Shanghai")

const (
	Layoutday         = "20060102"
	LayoutSecond      = "20060102150405"
	LayoutMilliSecond = "2006-01-02 15:04:05.000"
	LayoutMinuteHor   = "2006-01-02 15:04"
	LayoutSecondHor   = "2006-01-02 15:04:05"
)

func TimeNow() time.Time {
	return time.Now().In(Loc)
}
