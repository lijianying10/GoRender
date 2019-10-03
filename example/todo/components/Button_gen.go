// Code generated by compgen
package components

import (
    "github.com/google/uuid"
    "github.com/lijianying10/GoRender/pkg/render"
)

// Render component Render
func (btn *Button) Render(tid string) string {
    btn.AppendRenderTraceID(tid)
    return btn.render(tid)
}

// Destructor Comp DoDestructor
func (btn *Button) Destructor() {
    for _, comp := range btn.SubComponents {
        comp.Destructor()
    }
    btn.destructor()
}

// Constructor Comp DoDestructor
func (btn *Button) Constructor(grender *render.GoRender) {
    btn.Runtime = grender
    btn.UUID = uuid.New().String()
    btn.constructor()
}