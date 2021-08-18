package main

import (
	"context"
	"fmt"
	"time"
)

/*
处理日志模式
管控goroutine的生命周期
*/
type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch:   make(chan string, 10),
		stop: make(chan struct{}),
	}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
		fmt.Println("Shutdown: stop")
	case <-ctx.Done():
		fmt.Println("Shutdown: Done")
	}
}

func main() {
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.TODO(), "test")
	_ = tr.Event(context.TODO(), "test")
	_ = tr.Event(context.TODO(), "test")
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	tr.Shutdown(ctx)
	fmt.Println("done!")
}
