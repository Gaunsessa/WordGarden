package main

import (
	"syscall/js"
)

// Basic HTML5 Canvas Layer
type Canvas struct {
	element js.Value
	ctx js.Value

	width, height int
}

func NewCanvas(id string) *Canvas {
	element := getElement(id)
	ctx     := element.Call("getContext", "2d")

	ctx.Set("font", "30px monospace")

	return &Canvas{
		element: element,
		ctx: ctx,
		width: element.Get("width").Int(),
		height: element.Get("height").Int(),
	}
}

func (can *Canvas) Clear() {
	can.ctx.Set("fillStyle", "#000000");

	can.ctx.Call("fillRect", 0, 0, can.width, can.height)
}

func (can *Canvas) DrawText(txt, color string, x, y int) {
	can.ctx.Set("fillStyle", color)
	can.ctx.Call("fillText", txt, x, y)
}