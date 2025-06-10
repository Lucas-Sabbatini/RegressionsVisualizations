package main

import (
	"math"
)

func dimensionReduction(y []float64, featuresMatrix [][][]float64, w0 []float64) [][]float64 {
	//recieves -> y, array containig the respective y for each x1 and x2 | featuresMatrix -> 3d matrix with the value for each monomial assuming the values of x1 and x2
	//recieves -> w0, array containig the original/target values from the vector w
	//returns -> The gradient surface cost correctly reduced each W dimension using euclidian distance and direction.

	pace := 1
	n := 20
	move := 10
	costSurface := make([][]float64, n)

	for i := 0; i < n; i++ {

		costSurface[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			w := float64(i*pace - move)
			costSurface[i][j] = computeCost(generateW(w, len(featuresMatrix[0][0])), float64(j*pace-move), y, featuresMatrix)
		}
	}

	for i := 0; i < n; i++ {

		costSurface[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			w := float64(i*pace - move)
			costSurface[i][j] = computeCost(generateW(w, len(featuresMatrix[0][0])), float64(j*pace-move), y, featuresMatrix)
		}
	}

	return costSurface
}

func generateCostTable(y []float64, featuresMatrix [][][]float64, w0 []float64) [][]float64 {
	//recieves -> y, array containig the respective y for each x1 and x2 | featuresMatrix -> 3d matrix with the value for each monomial assuming the values of x1 and x2
	//recieves -> w0, array containig the original/target values from the vector w
	//returns -> a table with columns for: values B, W1, W2... distance from W and W0 and the respective cost these values generate.
	numW := len(featuresMatrix[0][0])

	w := make([]float64, numW)
	i := 0
	costTable := make([][]float64, int(math.Pow(21, float64(numW+1))))

	for b := -10; b <= 10; b++ {
		iterateW(w, 0, func(comb []float64) {
			costTable[i] = make([]float64, numW+3)

			costTable[i][0] = float64(b)
			for j := 1; i < numW+1; j++ {
				costTable[i][j] = comb[j-1]
			}
			costTable[i][numW+1] = scalerSignedDistance(comb, w0)
			costTable[i][numW+2] = computeCost(comb, float64(b), y, featuresMatrix)
			i++
		})
	}
	return costTable
}

func iterateW(w []float64, idx int, callback func([]float64)) {
	if idx == len(w) {
		comb := make([]float64, len(w))
		copy(comb, w)
		callback(comb)
		return
	}

	for val := -10; val <= 10; val++ {
		w[idx] = float64(val)
		iterateW(w, idx+1, callback)
	}
}

func scalerSignedDistance(w []float64, w0 []float64) float64 {
	d := createD(len(w))

	if len(w) != len(w0) {
		panic("signedDistance: w and w0 must have the same length")
	}
	var u float64
	for i := range w {
		u += (w[i] - w0[i]) * d[i]
	}
	return u
}

func createD(n int) []float64 {
	d := make([]float64, n)
	for i := range d {
		d[i] = 1.0
	}

	return d
}
