package pages

import (
	"fmt"

	"github.com/lijianying10/GoRender/example/todo/components"
	"github.com/lijianying10/GoRender/example/todo/rt"
	"github.com/lijianying10/GoRender/pkg/render"
)

// MainPage a sample component
type MainPage struct {
	render.Component
	TodoRuntime *rt.TodoRuntime

	todoInputComponent *components.Input
	listTODOItems      *components.ListTODOItems

	allBtnRef      *components.Button
	actionBtnRef   *components.Button
	completeBtnRef *components.Button
}

func (mpage *MainPage) constructor() {
	mpage.todoInputComponent = &components.Input{
		TodoRuntime: mpage.TodoRuntime,
	}
	mpage.RegisterSubComponent(mpage.todoInputComponent)

	mpage.listTODOItems = &components.ListTODOItems{
		TodoRuntime: mpage.TodoRuntime,
	}
	mpage.RegisterSubComponent(mpage.listTODOItems) // TODO: 这个过程能否省略?

	mpage.allBtnRef = &components.Button{
		OtherHTMLTag: `type="button" class="btn btn-secondary"`,
		Text:         "all",
		CallBack: func() {
			mpage.TodoRuntime.SetShowStatusAll()
			mpage.Runtime.Render()
		},
	}
	mpage.RegisterSubComponent(mpage.allBtnRef)

	mpage.actionBtnRef = &components.Button{
		OtherHTMLTag: `type="button" class="btn btn-secondary"`,
		Text:         "active",
		CallBack: func() {
			mpage.TodoRuntime.SetShowStatusActive()
			mpage.Runtime.Render()
		},
	}
	mpage.RegisterSubComponent(mpage.actionBtnRef)

	mpage.completeBtnRef = &components.Button{
		OtherHTMLTag: `type="button" class="btn btn-secondary"`,
		Text:         "complete",
		CallBack: func() {
			mpage.TodoRuntime.SetShowStatusComplete()
			mpage.Runtime.Render()
		},
	}
	mpage.RegisterSubComponent(mpage.completeBtnRef)
}

func (mpage *MainPage) destructor() {
}

func (mpage *MainPage) renderBtn(tid string) string {
	return mpage.allBtnRef.Render(tid) + mpage.actionBtnRef.Render(tid) + mpage.completeBtnRef.Render(tid)
}

func (mpage *MainPage) render(tid string) string {
	return fmt.Sprintf(`
    <div class="container">
        <div class="pageframe">
            <div class="row">
                <h1 class="apptitle">TODO</h1>
            </div>
            <div class="row appbox">
                %s
            </div>
            <div class="row appbox">
			    %s
            </div>
            <div class="row appbox">
                <div class="btn-toolbar" role="toolbar" aria-label="Toolbar with button groups">
                  <div class="input-group">
                    <div class="input-group-prepend">
                      <div class="input-group-text" id="btnGroupAddon">%d items left</div>
                    </div>
                  </div>
                  <div class="btn-group" role="group" aria-label="First group">
				  %s
                  </div>
                </div>
            </div>
        </div>
    </div>
	`, mpage.todoInputComponent.Render(tid), mpage.listTODOItems.Render(tid), mpage.TodoRuntime.TodoActiveCount(), mpage.renderBtn(tid))
}
