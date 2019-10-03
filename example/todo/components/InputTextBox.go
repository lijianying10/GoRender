package components

import (
	"fmt"

	"github.com/lijianying10/GoRender/pkg/render"
)

// InputTextBox a component
type InputTextBox struct {
	render.Component

	listenDoen        chan struct{}
	listenKeyDownDoen chan struct{}

	value        string
	InitialValue string

	InputOtherHTMLTag string

	KeyDownReturnCallBack func()
}

func (input *InputTextBox) constructor() {
	input.value = input.InitialValue

	input.listenDoen = make(chan struct{}, 1)
	input.listenKeyDownDoen = make(chan struct{}, 1)

	go func() {
		channel := input.Runtime.EventsUtil.Bind("input", input.UUID)
		for {
			select {
			case event := <-channel:
				input.value = event.Get("target").Get("value").String()
				input.Runtime.Render()
			case <-input.listenDoen:
				fmt.Println("InputCheckBox listen done routine exit")
				return
			}
		}
	}()

	if input.KeyDownReturnCallBack != nil {
		go func() {
			channel := input.Runtime.EventsUtil.Bind("keydown", input.UUID)
			for {
				select {
				case event := <-channel:
					inputKey := event.Get("key").String()
					if inputKey == "Enter" {
					} else {
						continue
					}
					input.KeyDownReturnCallBack()
				case <-input.listenKeyDownDoen:
					return
				}
			}
		}()

	}
}

// SetValue set a new value to component
func (input *InputTextBox) SetValue(val string) {
	input.value = val
	input.Runtime.Render()
}

// GetValue get component value
func (input *InputTextBox) GetValue() string {
	return input.value
}

func (input *InputTextBox) destructor() {
	input.listenDoen <- struct{}{}
	input.Runtime.EventsUtil.Unbind("keypress", input.UUID)

	if input.KeyDownReturnCallBack != nil {
		input.listenKeyDownDoen <- struct{}{}
		input.Runtime.EventsUtil.Unbind("keydown", input.UUID)
	}
}

func (input *InputTextBox) render(tid string) string {
	return fmt.Sprintf(`<input grid="%s" type="text" value="%s" %s/>`, input.UUID, input.value, input.InputOtherHTMLTag)
}
