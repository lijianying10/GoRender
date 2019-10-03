package components

import (
	"fmt"

	"github.com/lijianying10/GoRender/example/todo/rt"
	"github.com/lijianying10/GoRender/pkg/render"
)

// ListTODOItems a component
type ListTODOItems struct {
	render.Component
	items []*ListTODOItem

	TodoRuntime *rt.TodoRuntime
}

func (list *ListTODOItems) constructor() {
	go func() {
		channel := list.TodoRuntime.Bind(list.UUID)
		for {
			_, ok := <-channel
			if !ok {
				fmt.Println("already closed watch runtime data change")
				return
			}

			for _, item := range list.items {
				item.Destructor()
			}
			list.items = make([]*ListTODOItem, 0)

			for _, todoDataItem := range list.TodoRuntime.GetList() {
				item := &ListTODOItem{
					TodoData:    todoDataItem,
					TodoRuntime: list.TodoRuntime,
				}
				list.items = append(list.items, item)
				list.RegisterSubComponent(item)
			}
			list.Runtime.Render()
		}
	}()
}

func (list *ListTODOItems) destructor() {
	list.TodoRuntime.Unbind(list.UUID)
}

func (list *ListTODOItems) render(tid string) string {
	res := ""
	for _, item := range list.items {
		res += item.Render(tid)
	}
	return res
}
