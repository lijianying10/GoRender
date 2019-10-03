// Code generated by compgen
package main 

import (
    "github.com/google/uuid"
    "github.com/lijianying10/GoRender/pkg/render"
)

// Render component Render
func (timer *Timer) Render(tid string) string {
    timer.AppendRenderTraceID(tid)
    return timer.render(tid)
}

// Destructor Comp DoDestructor
func (timer *Timer) Destructor() {
    for _, comp := range timer.SubComponents {
        comp.Destructor()
    }
    timer.destructor()
}

// Constructor Comp DoDestructor
func (timer *Timer) Constructor(grender *render.GoRender) {
    timer.Runtime = grender
    timer.UUID = uuid.New().String()
    timer.constructor()
}
