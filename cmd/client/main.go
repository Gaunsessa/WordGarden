package main

import (
	"math/rand"
	"syscall/js"
	"time"
)

type Player struct {
	x, y int
	timeout float32
}

var CANVAS = NewCanvas("garden")
var GARDEN = NewGarden(*CANVAS, "ws://localhost:8000/ws")
var PLAYER = Player{ rand.Int() % CANVAS.width, rand.Int() % CANVAS.height, 3.0 }

func update(this js.Value, inputs []js.Value) interface{} {
	dt := GetDeltaTime()

	CANVAS.Clear()

	PLAYER.timeout -= dt

	if (PLAYER.timeout <= 0) {
		PLAYER.timeout = 3

		PLAYER.x = rand.Int() % CANVAS.width
		PLAYER.y = rand.Int() % CANVAS.height
	}

	GARDEN.Update(dt)

	GARDEN.Render()

	js.Global().Call("requestAnimationFrame", js.FuncOf(update))

	return nil
}

func keyPressCb(this js.Value, inputs []js.Value) interface{} {
	key := inputs[0].Get("key").String()

	GARDEN.PutText(key, PLAYER.x, PLAYER.y)

	PLAYER.timeout = 3.0
	PLAYER.x += 20

	return nil
}

func main() {
	rand.Seed(time.Now().Unix())

	js.Global().Call("addEventListener", "keypress", js.FuncOf(keyPressCb))

	js.Global().Call("requestAnimationFrame", js.FuncOf(update))

	loop := make(chan bool)
	<-loop

}