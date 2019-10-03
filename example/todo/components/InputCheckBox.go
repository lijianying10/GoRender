package components

import (
	"fmt"

	"github.com/lijianying10/GoRender/pkg/render"
)

// InputCheckBox a sample component
type InputCheckBox struct {
	render.Component

	listenDoen chan struct{}

	checked             bool
	CheckInitialState   bool
	InputOtherHTMLTag   string
	StateChangeCallBack func(bool)
}

func (input *InputCheckBox) constructor() {
	if input.CheckInitialState {
		input.checked = true
	}

	input.listenDoen = make(chan struct{}, 1)

	go func() {
		channel := input.Runtime.EventsUtil.Bind("click", input.UUID)
		for {
			select {
			case _ = <-channel:
				if input.checked {
					input.checked = false
				} else {
					input.checked = true
				}
				if input.StateChangeCallBack != nil {
					input.StateChangeCallBack(input.checked)
				}
				input.Runtime.Render()
			case <-input.listenDoen:
				if input.Runtime.DebugMod {
					fmt.Println("InputCheckBox listen done routine exit")
				}
				return
			}
		}
	}()
}

func (input *InputCheckBox) destructor() {
	input.listenDoen <- struct{}{}
	input.Runtime.EventsUtil.Unbind("click", input.UUID)
}

func (input *InputCheckBox) getHTMLCheckedTag() string {
	if input.checked {
		return "checked"
	}
	return ""
}

func (input *InputCheckBox) render(tid string) string {
	return fmt.Sprintf(`<input grid="%s" type="checkbox" %s %s/>`, input.UUID, input.getHTMLCheckedTag(), input.InputOtherHTMLTag)
}
