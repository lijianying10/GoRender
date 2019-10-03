package main

import (
	"fmt"

	"github.com/lijianying10/GoRender/example/todo/pages"
	"github.com/lijianying10/GoRender/example/todo/rt"
	"github.com/lijianying10/GoRender/pkg/render"
)

func main() {
	fmt.Println("app todo start!")
	app := render.NewGoRender("#app")
	app.MountRootComponent(&pages.MainPage{
		TodoRuntime: rt.NewTodoRuntime(),
	})
}
