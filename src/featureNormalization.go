package main

import "math"

func normalizeFeatures(featuresMatrix [][][]float64, featuresLen int) [][][]float64 {
	if featuresLen == 0 {
		return featuresMatrix
	}

	featuresMin := make([]float64, featuresLen)
	featuresMax := make([]float64, featuresLen)
	for k := 0; k < featuresLen; k++ {
		featuresMin[k] = math.Inf(1)
		featuresMax[k] = math.Inf(-1)
	}

	for i := 0; i < featuresLen; i++ {

		for j := 0; j < m; j++ {

			for k := 0; k < m; k++ {
				z := featuresMatrix[j][k][i]

				if z < featuresMin[i] {
					featuresMin[i] = z
				}
				if z > featuresMax[i] {
					featuresMax[i] = z
				}
			}
		}
	}

	normalized := make([][][]float64, m)

	for i := 0; i < m; i++ {
		normalized[i] = make([][]float64, m)

		for j := 0; j < m; j++ {
			normalized[i][j] = make([]float64, featuresLen)

			for k := 0; k < featuresLen; k++ {
				z := featuresMatrix[i][j][k]

				if featuresMin[k] != featuresMax[k] {
					normalized[i][j][k] = (z - featuresMin[k]) / (featuresMax[k] - featuresMin[k])
				} else {
					normalized[i][j][k] = 0
				}
			}
		}
	}

	return normalized
}
