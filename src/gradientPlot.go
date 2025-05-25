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

	return costSurface
}
