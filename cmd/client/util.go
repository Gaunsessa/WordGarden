package main

import (
	"syscall/js"
	"time"
)

var DOCUMENT = js.Global().Get("document")

var PT int64 = time.Now().UnixMilli()

func GetDeltaTime() float32 {
	ct := time.Now().UnixMilli()
	dt := float32(ct - PT)
	PT = ct

	return dt / 1000
}

func getElement(id string) js.Value {
	return DOCUMENT.Call("getElementById", id)
}