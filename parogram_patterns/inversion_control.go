package parogram_patterns

import (
	"github.com/pkg/errors"
)

type IntSet0 struct {
	data map[int]bool
}

//其中实现了 Add() 、Delete() 和 Contains() 三个操作，前两个是写操作，后一个是读操作。
func NewIntSet0() IntSet0 {
	return IntSet0{make(map[int]bool)}
}
func (set *IntSet0) Add(x int) {
	set.data[x] = true
}
func (set *IntSet0) Delete(x int) {
	delete(set.data, x)
}
func (set *IntSet0) Contains(x int) bool {
	return set.data[x]
}

type UndoableIntSet struct { // Poor style
	IntSet    // Embedding (delegation)
	functions []func()
}

//现在，我们想实现一个 Undo 的功能。我们可以再包装一下 IntSet ，变成 UndoableIntSet ，代码如下所示：
func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}

func (set *UndoableIntSet) Add(x int) { // Override
	if !set.Contains(x) {
		set.data[x] = true
		set.functions = append(set.functions, func() { set.Delete(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Delete(x int) { // Override
	if set.Contains(x) {
		delete(set.data, x)
		set.functions = append(set.functions, func() { set.Add(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(set.functions) - 1
	if function := set.functions[index]; function != nil {
		function()
		set.functions[index] = nil // For garbage collection
	}
	set.functions = set.functions[:index]
	return nil
}

/*
我来解释下这段代码。
	我们在 UndoableIntSet 中嵌入了IntSet ，然后 Override 了 它的 Add()和 Delete() 方法；
	Contains() 方法没有 Override，所以，就被带到 UndoableInSet 中来了。
	在 Override 的 Add()中，记录 Delete 操作；在 Override 的 Delete() 中，记录 Add 操作；
	在新加入的 Undo() 中进行 Undo 操作。
*/
/*
用这样的方式为已有的代码扩展新的功能是一个很好的选择。这样，就可以在重用原有代码功能和新的功能中达到一个平衡。
但是，这种方式最大的问题是，Undo 操作其实是一种控制逻辑，并不是业务逻辑，
所以，在复用 Undo 这个功能时，是有问题的，因为其中加入了大量跟 IntSet 相关的业务逻辑。
*/

type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // For garbage collection
	}
	*undo = functions[:index]
	return nil
}

type IntSet struct {
	data map[int]bool
	undo Undo
}

func NewIntSet() IntSet {
	return IntSet{data: make(map[int]bool)}
}

func (set *IntSet) Undo() error {
	return set.undo.Undo()
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}

/*
这个就是控制反转，不是由控制逻辑 Undo  来依赖业务逻辑 IntSet，而是由业务逻辑 IntSet 依赖 Undo 。
这里依赖的是其实是一个协议，这个协议是一个没有参数的函数数组。可以看到，这样一来，我们 Undo 的代码就可以复用了。
*/
