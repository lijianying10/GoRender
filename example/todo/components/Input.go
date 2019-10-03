package components

import (
	"fmt"

	"github.com/lijianying10/GoRender/example/todo/rt"
	"github.com/lijianying10/GoRender/pkg/render"
)

// Input a sample component
type Input struct {
	render.Component
	TodoRuntime *rt.TodoRuntime

	checkBox *InputCheckBox
	textBox  *InputTextBox
}

func (input *Input) constructor() {
	input.checkBox = &InputCheckBox{
		InputOtherHTMLTag: `aria-label="Checkbox for following text input"`,
		StateChangeCallBack: func(state bool) {
			if state {
				input.TodoRuntime.SetAllActive()
			} else {
				input.TodoRuntime.SetAllComplete()
			}
			input.Runtime.Render()
		},
	}
	input.RegisterSubComponent(input.checkBox)
	input.textBox = &InputTextBox{
		InputOtherHTMLTag: `class="form-control" aria-label="Text input with checkbox"`,
		KeyDownReturnCallBack: func() {
			todoContent := input.textBox.GetValue()
			input.TodoRuntime.Create(todoContent)
			input.textBox.SetValue("")
			input.Runtime.Render()
		},
	}
	input.RegisterSubComponent(input.textBox)
}

func (input *Input) destructor() {
}

func (input *Input) render(tid string) string {
	return fmt.Sprintf(`<div class="input-group">
              <div class="input-group-prepend">
                <div class="input-group-text">
				  %s
                </div>
              </div>
			  %s
            </div>`, input.checkBox.Render(tid), input.textBox.Render(tid))
}
