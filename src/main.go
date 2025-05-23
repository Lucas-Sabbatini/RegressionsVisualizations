package main

import (
	"fmt"
	"syscall/js"
)

func generateRandomDots(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		js.Global().Get("console").Call("error", "Go: Expected one string argument")
		return js.Null() // Usa js.Null() para erros
	}
	input := args[0].String()
	js.Global().Get("console").Call("log", fmt.Sprintf("Go: Input string: '%s'", input))

	features := parseFeatures(input)
	js.Global().Get("console").Call("log", fmt.Sprintf("Go: Parsed features: %v (count: %d)", features, len(features)))

	featuresLen := len(features)
	weights, bias := generateRandomWB(featuresLen)
	js.Global().Get("console").Call("log", fmt.Sprintf("Go: Generated weights: %.2f, bias: %.2f", weights, bias))

	featuresMatrix := generateFeaturesMatrix(features)
	js.Global().Get("console").Call("log", "Go: Features Matrix generation complete.")

	featuresMatrixNormalized := normalizeFeatures(featuresMatrix, featuresLen)

	f_wb_xMatrix := generateF_wb_xPredictionMatrix(weights, bias, featuresMatrixNormalized)
	js.Global().Get("console").Call("log", "Go: Predictions Matrix generation complete.")

	f_wb_xNormalMatrix := generateRandomNormalF_wb_xMatrix(f_wb_xMatrix)

	result := mapPredictionMatrix(f_wb_xNormalMatrix)

	if _, ok := result.(js.Value); ok && result.(js.Value).IsNull() {
		js.Global().Get("console").Call("error", "Go: mapPredictionMatrix returned null, aborting.")
	} else {
		js.Global().Get("console").Call("log", "Go: Matrix mapping complete. Returning data to JS.")
	}

	return result
}

func registerCallbacks() {
	js.Global().Set("generateRandomDots", js.FuncOf(generateRandomDots))
}

func main() {
	c := make(chan struct{})
	registerCallbacks()
	<-c
}
