// Code generated by compgen
package components

import (
    "github.com/google/uuid"
    "github.com/lijianying10/GoRender/pkg/render"
)

// Render component Render
func (input *Input) Render(tid string) string {
    input.AppendRenderTraceID(tid)
    return input.render(tid)
}

// Destructor Comp DoDestructor
func (input *Input) Destructor() {
    for _, comp := range input.SubComponents {
        comp.Destructor()
    }
    input.destructor()
}

// Constructor Comp DoDestructor
func (input *Input) Constructor(grender *render.GoRender) {
    input.Runtime = grender
    input.UUID = uuid.New().String()
    input.constructor()
}