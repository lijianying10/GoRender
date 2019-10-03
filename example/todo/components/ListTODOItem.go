package components

import (
	"fmt"

	"github.com/lijianying10/GoRender/example/todo/rt"
	"github.com/lijianying10/GoRender/pkg/render"
)

// ListTODOItem a component
type ListTODOItem struct {
	render.Component
	checkBox *InputCheckBox

	TodoData    *rt.Todo
	TodoRuntime *rt.TodoRuntime
}

func (item *ListTODOItem) constructor() {
	item.checkBox = &InputCheckBox{
		InputOtherHTMLTag: `aria-label="Checkbox for following text input"`,
		CheckInitialState: item.TodoData.IsComplete(),
		StateChangeCallBack: func(state bool) {
			if state {
				item.TodoRuntime.CompleteTodo(item.TodoData.UUID)
			} else {
				item.TodoRuntime.ActiveTodo(item.TodoData.UUID)
			}
			item.Runtime.Render()
		},
	}
	item.RegisterSubComponent(item.checkBox)
}

func (item *ListTODOItem) destructor() {
}

func (item *ListTODOItem) renderContent() string {
	if item.TodoData.IsComplete() {
		return fmt.Sprintf(`<strike>%s</strike>`, item.TodoData.Content)
	}
	return item.TodoData.Content
}

func (item *ListTODOItem) render(tid string) string {
	return fmt.Sprintf(`
    <div class="input-group">
        <div class="input-group-prepend"><div class="input-group-text">
            %s
        </div></div>
		<div class="form-control">
		    %s
        </div>
    </div>
	`, item.checkBox.Render(tid), item.renderContent())
}
