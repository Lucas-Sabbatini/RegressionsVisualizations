package main

import (
	"fmt"
	"math"
	"syscall/js"
)

const m int = 200

func generateFeaturesMatrix(features []string) [][][]float64 {
	var err error
	lenF := len(features)

	featuresDataset := make([][][]float64, m)
	for i := 0; i < m; i++ {
		featuresDataset[i] = make([][]float64, m)
		for j := 0; j < m; j++ {
			featuresDataset[i][j] = make([]float64, lenF)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < lenF; k++ {
				featuresDataset[j][i][k], err = computateFeature(float64(i+1), float64(j+1), features[k])
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
		dotProduct += w[i] * x[i]
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

func mapPredictionMatrix(f_wb_xMatrix [][]float64) interface{} {
	if f_wb_xMatrix == nil {
		js.Global().Get("console").Call("error", "mapPredictionMatrix received nil f_wb_xMatrix")
		return js.Null()
	}

	xGoSlice := make([]float64, 0, m*m)
	yGoSlice := make([]float64, 0, m*m)
	zGoSlice := make([]float64, 0, m*m)

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			xGoSlice = append(xGoSlice, float64(i+1))
			yGoSlice = append(yGoSlice, float64(j+1))
			if i < len(f_wb_xMatrix) && j < len(f_wb_xMatrix[j]) {
				zGoSlice = append(zGoSlice, f_wb_xMatrix[j][i])
			} else {
				js.Global().Get("console").Call("error", fmt.Sprintf("Index out of bounds accessing f_wb_xMatrix[%d][%d] during mapping", i, j))
				zGoSlice = append(zGoSlice, math.NaN())
			}
		}
	}

	xJSArray := sliceToFloat64Array(xGoSlice)
	yJSArray := sliceToFloat64Array(yGoSlice)
	zJSArray := sliceToFloat64Array(zGoSlice)

	return map[string]interface{}{
		"x": xJSArray,
		"y": yJSArray,
		"z": zJSArray,
	}
}
