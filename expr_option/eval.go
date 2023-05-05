package expr

import (
	"fmt"
	"github.com/spf13/cast"
)

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (c Const) Eval(_ Env) float64 {
	return float64(c)
}

func (u unary) Eval(env Env) float64 {
	v1 := cast.ToFloat64(u.x.Eval(env))
	switch u.op {
	case '+':
		v1 = +v1
	case '-':
		v1 = -v1
	default:
	}
	return v1
}

func (b binary) Eval(env Env) float64 {
	v1, v2 := cast.ToFloat64(b.x.Eval(env)), cast.ToFloat64(b.y.Eval(env))
	switch b.op {
	case '+':
		return v1 + v2
	case '-':
		return v1 - v2
	case '*':
		return v1 * v2
	case '/':
		return v1 / v2
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "sum":
		//TODO
	case "count":
		//TODO
	case "max":
		//TODO
	case "min":
		//TODO
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
