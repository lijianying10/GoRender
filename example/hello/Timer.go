package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/lijianying10/GoRender/pkg/render"
)

// Timer a component
type Timer struct {
	render.Component
	Curr     int64
	StopFlag bool
}

func (t *Timer) constructor() {
	fmt.Println("Timer init")
	t.Curr = time.Now().Unix()
	t.StopFlag = false
	go func() {
		for {
			time.Sleep(time.Second)
			t.Curr = time.Now().Unix()
			t.Runtime.Render()
			//fmt.Println("debug 中文: timer ", time.Now().Unix())
			fmt.Println("debug: ", reflect.TypeOf(&t).String())
			if t.StopFlag {
				fmt.Println("stop signal received")
				break
			}
		}
	}()
}

func (t *Timer) destructor() {
	fmt.Println("Timer bye bye")
	t.StopFlag = true
}

func (t *Timer) render(tid string) string {
	return fmt.Sprintf("<p>Time Now is : %d </p>", t.Curr)
}
