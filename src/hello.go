package main

import (
	"syscall/js"
)

func generateRandomDots(this js.Value, args []js.Value) interface{} {
	js.Global().Get("console").Call("log", "Hello from Go!")
	return nil
}

func registerCallbacks() {
	js.Global().Set("generateRandomDots", js.FuncOf(generateRandomDots))
}

func main() {
	c := make(chan struct{})
	registerCallbacks()
	<-c // impede que o programa Go termine
}
