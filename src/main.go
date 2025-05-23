package main

import (
	"syscall/js"
)

func registerCallbacks() {
	js.Global().Set("generateRandomDots", js.FuncOf(generateRandomDots))
}

func main() {
	c := make(chan struct{})
	registerCallbacks()
	<-c
}
