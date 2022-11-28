package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
)

type Character struct {
	char rune
	alpha float32

	x, y int
}

type Garden struct {
	canvas Canvas
	chars []Character

	ws js.Value
}

func NewGarden(canvas Canvas, ws string) *Garden {
	g := Garden{
		canvas: canvas,
		ws: js.Global().Get("WebSocket").New(ws),
	}

	msgCb := func(this js.Value, inputs []js.Value) interface{} {
		msg := strings.Split(inputs[0].Get("data").String(), "|")
		if len(msg) != 3 { return nil }

		x, err := strconv.ParseFloat(msg[0], 32)
		if err != nil { return nil }

		y, err := strconv.ParseFloat(msg[1], 32)
		if err != nil { return nil }

		for i, c := range msg[2] {
			g.chars = append(g.chars, Character{
				char: c,
				alpha: 0xFF,
				x: int(float64(g.canvas.width) * x) + i * CHAR_PIXEL_WIDTH,
				y: int(float64(g.canvas.height) * y),
			})
		}

		return nil
	}

	g.ws.Call("addEventListener", "message", js.FuncOf(msgCb))

	return &g
}

func (g *Garden) PutText(txt string, x, y int) {
	g.ws.Call("send", fmt.Sprintf("%f|%f|%s", float32(x) / float32(g.canvas.width), float32(y) / float32(g.canvas.height), txt))

	for i, c := range txt {
		g.chars = append(g.chars, Character{
			char: c,
			alpha: 0xFF,
			x: x + i * CHAR_PIXEL_WIDTH,
			y: y,
		})
	}
}

func (g *Garden) Update(dt float32) {
	for i := 0; i < len(g.chars); i++ {
		g.chars[i].alpha -= 0x88 * dt

		if g.chars[i].alpha <= 0x11 {
			g.chars = append(g.chars[:i], g.chars[i + 1:]...)

			i--
		}
	}
}

func (g *Garden) Render() {
	for _, c := range g.chars {
		CANVAS.DrawText(string(c.char), "#e1e1e6" + fmt.Sprintf("%X", int(c.alpha)), c.x, c.y)
	}
}