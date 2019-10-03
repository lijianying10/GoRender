// Code generated by compgen
package pages

import (
    "github.com/google/uuid"
    "github.com/lijianying10/GoRender/pkg/render"
)

// Render component Render
func (mpage *MainPage) Render(tid string) string {
    mpage.AppendRenderTraceID(tid)
    return mpage.render(tid)
}

// Destructor Comp DoDestructor
func (mpage *MainPage) Destructor() {
    for _, comp := range mpage.SubComponents {
        comp.Destructor()
    }
    mpage.destructor()
}

// Constructor Comp DoDestructor
func (mpage *MainPage) Constructor(grender *render.GoRender) {
    mpage.Runtime = grender
    mpage.UUID = uuid.New().String()
    mpage.constructor()
}