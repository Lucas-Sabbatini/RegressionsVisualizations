package main

import (
	"fmt"
	"syscall/js"
)

func go3DToJS(matrix [][][]float64) js.Value {
	jsMatrix := js.Global().Get("Array").New()

	for i, row := range matrix {
		jsRow := js.Global().Get("Array").New()

		for j, col := range row {
			jsCol := js.Global().Get("Array").New()

			for k, val := range col {
				jsCol.SetIndex(k, val)
			}

			jsRow.SetIndex(j, jsCol)
		}

		jsMatrix.SetIndex(i, jsRow)
	}

	return jsMatrix
}

func jsTo3DFloat64(jsMatrix js.Value) [][][]float64 {
	rows := jsMatrix.Length()
	result := make([][][]float64, rows)

	for i := 0; i < rows; i++ {
		jsRow := jsMatrix.Index(i)
		cols := jsRow.Length()
		result[i] = make([][]float64, cols)

		for j := 0; j < cols; j++ {
			jsCell := jsRow.Index(j)
			depth := jsCell.Length()
			result[i][j] = make([]float64, depth)

			for k := 0; k < depth; k++ {
				result[i][j][k] = jsCell.Index(k).Float()
			}
		}
	}

	return result
}

func jsValueToFloat64Array(value js.Value) ([]float64, error) {
	if value.IsNull() || value.IsUndefined() {
		return nil, fmt.Errorf("o valor recebido é null ou undefined")
	}

	var length int
	isSupportedType := false

	lengthJS := value.Get("length")
	if lengthJS.Type() != js.TypeNumber {
		return nil, fmt.Errorf("o valor não possui uma propriedade 'length' numérica. Tipo: %s", value.Type().String())
	}
	length = lengthJS.Int()

	if value.Type() == js.TypeObject && value.Get("constructor").Equal(js.Global().Get("Array")) {
		isSupportedType = true
	} else if value.Type() == js.TypeObject {
		typedArrayConstructors := []string{"Float64Array", "Float32Array"}
		valueConstructor := value.Get("constructor")

		if valueConstructor.Type() == js.TypeFunction {
			for _, taName := range typedArrayConstructors {
				jsConstructorGlobal := js.Global().Get(taName)
				if !jsConstructorGlobal.IsUndefined() && !jsConstructorGlobal.IsNull() && valueConstructor.Equal(jsConstructorGlobal) {
					isSupportedType = true
					break
				}
			}
		}
	}

	if !isSupportedType {
		errMsg := fmt.Sprintf("o argumento não é um Array JavaScript padrão nem um TypedArray suportado. Tipo: %s", value.Type().String())
		if value.Type() == js.TypeObject {
			constructor := value.Get("constructor")
			constructorName := "N/A (não é função ou sem nome)"
			if constructor.Type() == js.TypeFunction {
				nameProp := constructor.Get("name")
				if nameProp.Type() == js.TypeString {
					constructorName = nameProp.String()
				}
			}
			errMsg = fmt.Sprintf("o argumento não é um Array JavaScript padrão nem um TypedArray suportado. Tipo: Object, Construtor: %s", constructorName)
		}
		return nil, fmt.Errorf(errMsg)
	}

	if length == 0 {
		return []float64{}, nil
	}

	slice := make([]float64, length)
	for i := 0; i < length; i++ {
		element := value.Index(i)
		if element.Type() != js.TypeNumber {
			return nil, fmt.Errorf("elemento do array no índice %d não é um número (tipo: %s)", i, element.Type().String())
		}
		slice[i] = element.Float()
	}
	return slice, nil
}

func float64MatrixToJsValue(matrix [][]float64) js.Value {
	if matrix == nil {
		return js.Null()
	}
	jsOuterArray := js.Global().Get("Array").New(len(matrix))
	for i, row := range matrix {
		jsInnerArray := js.Global().Get("Array").New(len(row))
		for j, val := range row {
			jsInnerArray.SetIndex(j, js.ValueOf(val))
		}
		jsOuterArray.SetIndex(i, jsInnerArray)
	}
	return jsOuterArray
}

func createGradientObject(newW []float64, newB float64, newJ float64, f_wb_xPlot [][]float64) interface{} {
	jsNewW := js.Global().Get("Array").New(len(newW))
	for i, val := range newW {
		jsNewW.SetIndex(i, js.ValueOf(val))
	}

	jsNewB := js.ValueOf(newB)

	jsNewJ := js.ValueOf(newJ)

	jsFwbXPlot := js.Global().Get("Array").New(len(f_wb_xPlot))
	for i, row := range f_wb_xPlot {
		jsRow := js.Global().Get("Array").New(len(row))
		for j, val := range row {
			jsRow.SetIndex(j, js.ValueOf(val))
		}
		jsFwbXPlot.SetIndex(i, jsRow)
	}

	resultObj := js.Global().Get("Object").New()
	resultObj.Set("w", jsNewW)
	resultObj.Set("b", jsNewB)
	resultObj.Set("j", jsNewJ)
	resultObj.Set("predictionPlot", jsFwbXPlot)

	return resultObj
}
