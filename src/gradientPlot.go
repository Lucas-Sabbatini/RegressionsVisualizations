package main

import (
	"math"
	"syscall/js"
)

func computeCost(w []float64, b float64, y []float64, featuresMatrix [][][]float64) float64 {
	if 1 > len(y) {
		return 0
	}
	var acumulatedErr float64

	for k, val := range y {
		i := k % m
		j := k / m

		f_wb_x, computeError := computeF_wb_x(w, b, featuresMatrix[i][j])
		if computeError != nil {
			js.Global().Call("alert", computeError.Error())
			return 0
		}
		err := f_wb_x - val
		acumulatedErr += math.Pow(err, 2.0)
	}

	return acumulatedErr / (2 * float64(m))
}

func generateCostSurface(y []float64, featuresMatrix [][][]float64) [][]float64 {

	return costSurface
}
