package main

import (
	"syscall/js"
)

func generateRandomDots(this js.Value, args []js.Value) interface{} {
	//recieves -> the normalized features matrix
	//returns -> a random set of dots tha follows the main function structure defined in the features matrix

	if len(args) < 1 {
		js.Global().Get("console").Call("error", "Go: Expected one matrix[][][] argument")
		return js.Null()
	}

	featuresMatrix := jsTo3DFloat64(args[0])
	featuresLen := len(featuresMatrix[0][0])

	weights, bias := generateRandomWB(featuresLen)

	f_wb_xMatrix := generateF_wb_xPredictionMatrix(weights, bias, featuresMatrix)

	f_wb_xNormalMatrix := generateRandomNormalF_wb_xMatrix(f_wb_xMatrix)

	result := mapPredictionMatrix(f_wb_xNormalMatrix)

	if _, ok := result.(js.Value); ok && result.(js.Value).IsNull() {
		js.Global().Get("console").Call("error", "Go: mapPredictionMatrix returned null, aborting.")
	}

	return result
}

func featuresMatrixToJs(this js.Value, args []js.Value) interface{} {
	//recieves -> a string containig the monomials that form the polynomial to be used as a model
	//returns -> the normalized set of features to be used in a 3d matrix

	if len(args) < 1 {
		js.Global().Get("console").Call("error", "Go: Expected one string argument")
		return js.Null()
	}

	input := args[0].String()
	features := parseFeatures(input)
	featuresLen := len(features)

	featuresMatrix := generateFeaturesMatrix(features)

	featuresMatrixNormalized := normalizeFeatures(featuresMatrix, featuresLen)

	return go3DToJS(featuresMatrixNormalized)
}

func costSurfaceToJs(this js.Value, args []js.Value) interface{} {
	//recieves -> the features matrix and the y axis mapped to a 1d array
	//returns -> a 2d array representing the dots in the cost plot
	if len(args) < 2 {
		js.Global().Get("console").Call("error", "Go: Expected two arguments in costSurface")
		return js.Undefined()
	}
	featuresMatrix := jsTo3DFloat64(args[0])
	yAxis, err := jsValueToFloat64Array(args[1])
	if err != nil {
		js.Global().Call("alert", err.Error())
		return js.Undefined()
	}

	costSurface := generateCostSurface(yAxis, featuresMatrix)
	return float64MatrixToJsValue(costSurface)
}

func registerCallbacks() {
	js.Global().Set("generateRandomDots", js.FuncOf(generateRandomDots))
	js.Global().Set("featuresMatrixToJs", js.FuncOf(featuresMatrixToJs))
	js.Global().Set("costSurfaceToJs", js.FuncOf(costSurfaceToJs))
}

func main() {
	c := make(chan struct{})
	registerCallbacks()
	<-c
}
