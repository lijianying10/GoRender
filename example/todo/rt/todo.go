package rt

import (
	"github.com/google/uuid"
)

// This file is todo's

// Todo kernel data structure
type Todo struct {
	UUID    string
	Content string
	Status  string // Active and Complete
}

// IsComplete define is complete
func (t *Todo) IsComplete() bool {
	return t.Status == "Complete"
}

// TodoRuntime Runtime stuct
type TodoRuntime struct {
	todoList        []*Todo
	showStatus      string
	dataChangeEvent *DataChangeEvent
}

// NewTodoRuntime new rt
func NewTodoRuntime() *TodoRuntime {
	return &TodoRuntime{
		showStatus:      "All",
		dataChangeEvent: NewDataChangeEvent(),
	}
}

// Create Create a new todo
func (rt *TodoRuntime) Create(Content string) {
	rt.todoList = append(rt.todoList, &Todo{
		UUID:    uuid.New().String(),
		Content: Content,
		Status:  "Active",
	})
	rt.bcustChangeEvent()
}

// TodoActiveCount count items
func (rt *TodoRuntime) TodoActiveCount() int {
	return len(rt.getTodoListByStatus("Active"))
}

// SetAllComplete set all status to done
func (rt *TodoRuntime) SetAllComplete() {
	for idx := range rt.todoList {
		rt.todoList[idx].Status = "Complete"
	}
	rt.bcustChangeEvent()
}

// SetAllActive set all status to done
func (rt *TodoRuntime) SetAllActive() {
	for idx := range rt.todoList {
		rt.todoList[idx].Status = "Active"
	}
	rt.bcustChangeEvent()
}

// SetShowStatusAll set show status all
func (rt *TodoRuntime) SetShowStatusAll() {
	rt.showStatus = "All"
	rt.bcustChangeEvent()
}

// SetShowStatusActive set show status active
func (rt *TodoRuntime) SetShowStatusActive() {
	rt.showStatus = "Active"
	rt.bcustChangeEvent()
}

// SetShowStatusComplete set show status complete
func (rt *TodoRuntime) SetShowStatusComplete() {
	rt.showStatus = "Complete"
	rt.bcustChangeEvent()
}

// GetList get todo list by todo data status
func (rt *TodoRuntime) GetList() []*Todo {
	return rt.getTodoListByStatus(rt.showStatus)
}

func (rt *TodoRuntime) getTodoListByStatus(status string) []*Todo {
	if status == "All" {
		return rt.todoList
	}
	var res []*Todo
	for _, todo := range rt.todoList {
		if todo.Status == status {
			res = append(res, todo)
		}
	}
	return res
}

func (rt *TodoRuntime) setStatus(uuid, status string) {
	for _, todo := range rt.todoList {
		if todo.UUID == uuid {
			todo.Status = status
			return
		}
	}

	rt.bcustChangeEvent()

	// Ignore error handle for sample app
}

// ActiveTodo active a todo item
func (rt *TodoRuntime) ActiveTodo(uuid string) {
	rt.setStatus(uuid, "Active")
}

// CompleteTodo complete a todo item
func (rt *TodoRuntime) CompleteTodo(uuid string) {
	rt.setStatus(uuid, "Complete")
}

// Bind change event
func (rt *TodoRuntime) Bind(id string) chan struct{} {
	return rt.dataChangeEvent.Bind(id)
}

// Unbind change event
func (rt *TodoRuntime) Unbind(id string) {
	rt.dataChangeEvent.Unbind(id)
}

func (rt *TodoRuntime) bcustChangeEvent() {
	rt.dataChangeEvent.BCust(struct{}{})
}
