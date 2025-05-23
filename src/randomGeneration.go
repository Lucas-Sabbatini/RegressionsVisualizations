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
