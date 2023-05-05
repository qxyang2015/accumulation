package expr

// 表达式计算
type Expr interface {
	Eval(env Env) float64
	Check(vars map[string]bool) error
}

// 变量
type Var string

// 常量
type Const float64

// 一元运算符
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

// 二元运算符
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

// 函数
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

type Node struct {
	Type   string
	Value  interface{}
	CNodes []Node
}

func Parse(node Node) {

}
