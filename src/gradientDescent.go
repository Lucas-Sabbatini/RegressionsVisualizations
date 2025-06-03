package main

import (
	"syscall/js"
)

func dj_dwi(w []float64, b float64, y []float64, featuresMatrix [][][]float64, ki int) float64 {
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
		acumulatedErr += err * featuresMatrix[i][j][ki]
	}

	return acumulatedErr / float64(m)
}

func dj_db(w []float64, b float64, y []float64, featuresMatrix [][][]float64) float64 {
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
		acumulatedErr += err
	}

	return acumulatedErr / float64(m)
}

func computeNewW(w []float64, learningRate float64, b float64, y []float64, featuresMatrix [][][]float64) []float64 {
	wLen := len(w)
	newW := make([]float64, wLen)

	for i := 0; i < wLen; i++ {
		newW[i] = w[i] - learningRate*dj_dwi(w, b, y, featuresMatrix, i)
	}

	return newW
}

func gradientDescent(w []float64, b float64, y []float64, featuresMatrix [][][]float64) ([]float64, float64, float64, [][]float64) {
	const learningRate float64 = 0.006
	newW := computeNewW(w, learningRate, b, y, featuresMatrix)
	newB := b - learningRate*dj_db(w, b, y, featuresMatrix)

	newJ := computeCost(newW, newB, y, featuresMatrix)
	f_wb_xPlot := generateF_wb_xPredictionMatrix(newW, newB, featuresMatrix)

	return newW, newB, newJ, f_wb_xPlot
}
