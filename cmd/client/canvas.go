package main

import (
	"fmt"
	"syscall/js"
)

// Basic HTML5 Canvas Layer
type Canvas struct {
	element js.Value
	ctx js.Value

	width, height int
	charSize int
}

func NewCanvas(id string) *Canvas {
	element := getElement(id)
	ctx     := element.Call("getContext", "2d")

	ctx.Set("font", fmt.Sprintf("%dpx monospace", element.Get("width").Int() / CHAR_WIDTH))

	return &Canvas{
		element: element,
		ctx: ctx,
		width: element.Get("width").Int(),
		height: element.Get("height").Int(),
		charSize: element.Get("width").Int() / CHAR_WIDTH,
	}
}

func (can *Canvas) Clear() {
	can.ctx.Call("clearRect", 0, 0, can.width, can.height)
}

func (can *Canvas) DrawText(txt, color string, x, y int) {
	can.ctx.Set("fillStyle", color)
	can.ctx.Call("fillText", txt, x, y)
}