package components

import (
	"fmt"

	"github.com/lijianying10/GoRender/pkg/render"
)

// Button a component
type Button struct {
	render.Component
	listenDoen chan struct{}

	OtherHTMLTag string
	Text         string
	CallBack     func()
}

func (btn *Button) constructor() {
	if btn.CallBack == nil {
		panic("Callback must not be nil")
	}
	btn.listenDoen = make(chan struct{}, 1)
	go func() {
		channel := btn.Runtime.EventsUtil.Bind("click", btn.UUID)
		for {
			select {
			case _ = <-channel:
				btn.CallBack()
			case <-btn.listenDoen:
				return
			}
		}
	}()
}

func (btn *Button) destructor() {
	btn.listenDoen <- struct{}{}
	btn.Runtime.EventsUtil.Unbind("click", btn.UUID)
}

func (btn *Button) render(tid string) string {
	return fmt.Sprintf(`<button grid="%s" %s>%s</button>`, btn.UUID, btn.OtherHTMLTag, btn.Text)
}
