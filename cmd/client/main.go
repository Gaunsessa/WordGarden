package main

import (
	"math/rand"
	"syscall/js"
	"time"
	"unicode/utf8"
)

type Player struct {
	sx int
	x, y int
	timeout float32
}

var PORT string
var IP string

var CANVAS = NewCanvas("garden")
var GARDEN = NewGarden(*CANVAS, "ws://" + IP + ":" + PORT + "/ws")
var PLAYER = Player{ 
	x: rand.Int() % CANVAS.width, 
	y: rand.Int() % CANVAS.height, 
	timeout: 3.0,
}

func update(this js.Value, inputs []js.Value) interface{} {
	dt := GetDeltaTime()

	CANVAS.Clear()

	PLAYER.timeout -= dt

	if (PLAYER.timeout <= 0) {
		PLAYER.timeout = 3

		PLAYER.x = rand.Int() % CANVAS.width
		PLAYER.y = rand.Int() % CANVAS.height

		PLAYER.sx = PLAYER.x
	}

	GARDEN.Update(dt)

	GARDEN.Render()

	js.Global().Call("requestAnimationFrame", js.FuncOf(update))

	return nil
}

func keyPressCb(this js.Value, inputs []js.Value) interface{} {
	key := inputs[0].Get("key").String()

	if key == "Enter" {
		PLAYER.y += CANVAS.charSize
		PLAYER.x = PLAYER.sx
	} else {
		GARDEN.PutText(key, PLAYER.x, PLAYER.y)
		PLAYER.x += CHAR_PIXEL_WIDTH
	}

	PLAYER.timeout = 3.0

	return nil
}

func pasteCb(this js.Value, inputs []js.Value) interface{} {
	var clip js.Value

	wData := js.Global().Get("clipboardData")
	if !wData.IsUndefined() { clip = wData }

	eData := inputs[0].Get("clipboardData")
	if !eData.IsUndefined() { clip = eData }

	data := clip.Call("getData", "text").String()

	GARDEN.PutText(data, PLAYER.x, PLAYER.y)

	PLAYER.timeout = 3.0
	PLAYER.x += CHAR_PIXEL_WIDTH * utf8.RuneCountInString(data)

	return nil
}

func main() {
	rand.Seed(time.Now().Unix())

	PLAYER.sx = PLAYER.x
	CHAR_PIXEL_WIDTH = int(float32(CANVAS.charSize) / 1.5)

	js.Global().Call("addEventListener", "keypress", js.FuncOf(keyPressCb))
	js.Global().Call("addEventListener", "paste", js.FuncOf(pasteCb))

	js.Global().Call("requestAnimationFrame", js.FuncOf(update))

	loop := make(chan bool)
	<-loop
}