package main

import (
	"math/rand"
)

func generateRandomWB(featuresLen int) ([]float64, float64) {
	weights := make([]float64, featuresLen)

	for i := range weights {
		weights[i] = rand.Float64()*20 - 10
	}

	bias := rand.Float64()*20 - 10

	return weights, bias
}

func generateRandomNormalF_wb_xMatrix(f_wb_x [][]float64) [][]float64 {
	mu := 0.0
	standDevi := 0.7
	x := 0.0
	randomF_wb_x := make([][]float64, m)

	for i := 0; i < m; i++ {
		randomF_wb_x[i] = make([]float64, m)

		for j := 0; j < m; j++ {
			x = mu + float64(standDevi)*rand.NormFloat64()
			randomF_wb_x[i][j] = f_wb_x[i][j] - x
		}
	}

	return randomF_wb_x
}
