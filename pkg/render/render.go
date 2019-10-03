package render

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/gopherjs/gopherjs/js"
)

// GoRender 这是一个Runtime
type GoRender struct {
	Components   []IComponent
	Root         IComponent
	RootSelector string // Mount 组件的Selector
	EventsUtil   *Listener
	DebugMod     bool
}

// NewGoRender Return a new GoRender
func NewGoRender(sel string) *GoRender {
	runtime := &GoRender{
		RootSelector: sel,
		EventsUtil:   NewListener(sel),
	}

	js.Global.Set("GoRender_DebugON", func() {
		fmt.Println("GoRender debuging mod is on")
		runtime.DebugMod = true
	})

	js.Global.Set("GoRender_DebugOff", func() {
		fmt.Println("GoRender debuging mod is off")
		runtime.DebugMod = false
	})

	js.Global.Set("GoRender_ComponentsCount", func() {
		runtime.componentCount()
	})

	return runtime
}

func (gr *GoRender) componentCount() {
	fmt.Println("count:", len(gr.Components))
	for _, comp := range gr.Components {
		fmt.Println(reflect.TypeOf(&comp).String(), comp.GetInfo())
	}
	fmt.Println("count end")
}

// Render render all pages
func (gr *GoRender) Render() {
	uuid := uuid.New().String()
	if gr.DebugMod {
		fmt.Println("debug render :", uuid)
	}
	js.Global.Call("gorender", gr.Root.Render(uuid), gr.RootSelector)
}

// CreateComponent 创建Comp
func (gr *GoRender) CreateComponent(comp IComponent) {
	gr.Components = append(gr.Components, comp)
	comp.Constructor(gr)
}

// MountRootComponent Runtime 绑定根组件
func (gr *GoRender) MountRootComponent(comp IComponent) {
	gr.CreateComponent(comp)
	gr.Root = comp
	gr.Render()
}
