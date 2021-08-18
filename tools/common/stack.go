package common

import (
	"container/list"
	"sync"
)

type Stack struct {
	list  *list.List
	Mutex sync.RWMutex
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list, sync.RWMutex{}}
}

func (stack *Stack) Push(value interface{}) {
	stack.Mutex.Lock()
	stack.list.PushBack(value)
	stack.Mutex.Unlock()
}

func (stack *Stack) Pop() interface{} {
	stack.Mutex.Lock()
	defer stack.Mutex.Unlock()

	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	stack.Mutex.RLock()
	defer stack.Mutex.RUnlock()

	e := stack.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (stack *Stack) Len() int {
	stack.Mutex.RLock()
	defer stack.Mutex.RUnlock()

	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	stack.Mutex.RLock()
	defer stack.Mutex.RUnlock()

	return stack.list.Len() == 0
}
