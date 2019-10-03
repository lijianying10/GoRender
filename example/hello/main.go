package main

import (
	"fmt"

	"github.com/lijianying10/GoRender/pkg/render"
)

func main() {
	fmt.Println("hello here is main function")
	app := render.NewGoRender("#app")
	app.MountRootComponent(&Timer{})
}
