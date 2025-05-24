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

	// Tenta obter a propriedade 'length'. Se não existir ou não for um número, não é array-like.
	lengthJS := value.Get("length")
	if lengthJS.Type() != js.TypeNumber {
		return nil, fmt.Errorf("o valor não possui uma propriedade 'length' numérica. Tipo: %s", value.Type().String())
	}
	length = lengthJS.Int()

	// Verifica se é um Array padrão
	if value.Type() == js.TypeObject && value.Get("constructor").Equal(js.Global().Get("Array")) {
		isSupportedType = true
	} else if value.Type() == js.TypeObject {
		// Verifica se é um TypedArray comum.
		// Adicione outros construtores de TypedArray se necessário (ex: "Int32Array", "Uint8Array").
		typedArrayConstructors := []string{"Float64Array", "Float32Array"}
		valueConstructor := value.Get("constructor") // O construtor do objeto JS

		// Verifica se valueConstructor é uma função (construtores são funções)
		if valueConstructor.Type() == js.TypeFunction {
			for _, taName := range typedArrayConstructors {
				jsConstructorGlobal := js.Global().Get(taName) // O construtor global (ex: window.Float64Array)
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
		return []float64{}, nil // Retorna slice vazio se o array JS estiver vazio
	}

	slice := make([]float64, length)
	for i := 0; i < length; i++ {
		element := value.Index(i) // Funciona para Array e TypedArray
		if element.Type() != js.TypeNumber {
			return nil, fmt.Errorf("elemento do array no índice %d não é um número (tipo: %s)", i, element.Type().String())
		}
		slice[i] = element.Float()
	}
	return slice, nil
}

// float64_2D_ToJsValue converte [][]float64 para um js.Value (array JS de arrays).
func float64MatrixToJsValue(matrix [][]float64) js.Value {
	if matrix == nil {
		return js.Null() // Ou js.Undefined() dependendo da preferência
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
