package main

import (
	"fmt"
	"syscall/js"
)

const m int = 200

func generateFeaturesMatrix(features []string) [][][]float64 {
	var err error
	featuresDataset := make([][][]float64, m)
	len := len(features)

	for i := 0; i < m; i++ {
		featuresDataset[i] = make([][]float64, m)

		for j := 0; j < m; j++ {
			featuresDataset[i][j] = make([]float64, len)

			for k := 0; k < len; k++ {
				featuresDataset[i][j][k], err = computateFeature(float64(i+1), float64(j+1), features[k])
				if err != nil {
					js.Global().Call("alert", err.Error())
					return nil
				}
			}
		}
	}

	return featuresDataset
}

func computeF_wb_x(w []float64, b float64, x []float64) (float64, error) {
	if len(w) != len(x) {
		return 0, fmt.Errorf("W and X should have the same length")
	}

	var dotProduct float64
	for i := 0; i < len(w); i++ {
		dotProduct += w[i] + x[i]
	}

	return dotProduct + b, nil
}

func generateF_wb_xPredictionMatrix(w []float64, b float64, featuresMatrix [][][]float64) [][]float64 {
	f_wb_x := make([][]float64, m)
	var err error

	for i := 0; i < m; i++ {
		f_wb_x[i] = make([]float64, m)

		for j := 0; j < m; j++ {
			f_wb_x[i][j], err = computeF_wb_x(w, b, featuresMatrix[i][j])
			if err != nil {
				js.Global().Call("alert", err.Error())
				return nil
			}
		}
	}

	return f_wb_x
}

// posso fazer uma função que retorna uma matriz com todas as features
// será necessário normalizar os valores, preciso entender melhor como funciona isso
// o plot recebe 3 vetores, com os valores em cada cordenada x, y , z

//depois de tudo posso fazer uma função que recebe a matriz com as features normalizadas e retorna os vetores x1, x2 e y
// x1 e x2 tendem a ser sempre os mesmos, números de -100 a 100 com o intervalo de 1
