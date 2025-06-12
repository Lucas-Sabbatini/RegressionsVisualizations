package main

import (
	"math"
	"syscall/js"
)

// recieves -> y, array containig the respective y for each x1 and x2 | featuresMatrix -> 3d matrix with the value for each monomial assuming the values of x1 and x2
// recieves -> w0, array containig the original/target values from the vector w
// returns -> The gradient surface cost correctly reduced each W dimension using euclidian distance and direction.
func dimensionReduction(y []float64, featuresMatrix [][][]float64, w0 []float64) [][]float64 {
	numDimensions := len(w0)
	pace := 1
	n := 21
	move := 10
	costSurface := make([][]float64, n)

	for i := 0; i < n; i++ {

		costSurface[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			distanceVector := vectorScalarMult(vetDist(numDimensions), float64(i*pace-move))
			w := vectorPlusDot(distanceVector, w0)
			js.Global().Get("console").Call("log", scalerSignedDistance(w, w0))
			costSurface[i][j] = computeCost(w, float64(j*pace-move), y, featuresMatrix)
		}
	}

	return costSurface
}

func scalerSignedDistance(w []float64, w0 []float64) float64 {
	d := createD(len(w), math.Pow(1.0/float64(len(w)), 0.5))

	if len(w) != len(w0) {
		panic("signedDistance: w and w0 must have the same length")
	}
	var u float64
	for i := range w {
		u += (w[i] - w0[i]) * d[i]
	}
	return u
}

func vetDist(numDimensions int) []float64 {
	//quero q a distância inicial seja sempre 1

	catLen := math.Pow(1.0/float64(numDimensions), 0.5)
	return createD(numDimensions, catLen)
}

func createD(numDimensions int, catLen float64) []float64 {
	d := make([]float64, numDimensions)
	for i := range d {
		d[i] = catLen
	}

	return d
}

func vectorPlusDot(vector []float64, dot []float64) []float64 {
	result := make([]float64, len(vector))
	for i := range vector {
		result[i] = vector[i] + dot[i]
	}

	return result
}

func vectorScalarMult(vector []float64, scalar float64) []float64 {
	multD := make([]float64, len(vector))
	for i := range multD {
		multD[i] = vector[i] * scalar
	}

	return multD
}
