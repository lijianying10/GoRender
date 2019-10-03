package render

import "encoding/json"

// IComponent COMP IFACE
type IComponent interface {
	Constructor(*GoRender)
	Destructor()
	Render(string) string
	GetInfo() string
}

// Component Component 基础算法
type Component struct {
	Runtime       *GoRender
	SubComponents []IComponent
	UUID          string
	RenderTraceID []string
}

// RegisterSubComponent 创建子组件维护树形结构
func (comp *Component) RegisterSubComponent(c IComponent) {
	comp.Runtime.CreateComponent(c)
	comp.SubComponents = append(comp.SubComponents, c)
}

// AppendRenderTraceID append trace id
func (comp *Component) AppendRenderTraceID(tid string) {
	comp.RenderTraceID = append(comp.RenderTraceID, tid)
}

// GetInfo get component info
func (comp *Component) GetInfo() string {
	res := []string{}
	lena := len(comp.RenderTraceID)
	if lena > 15 {
		res = comp.RenderTraceID[lena-15:]
	} else {
		res = comp.RenderTraceID
	}
	renderIDsBody, _ := json.Marshal(res)
	return comp.UUID + " " + string(renderIDsBody)
}
