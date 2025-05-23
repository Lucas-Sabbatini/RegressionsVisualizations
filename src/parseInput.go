package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"syscall/js"
	"unicode"
)

func isValidMonomial(s string) bool {
	pattern := `^x[12](\^[0-9]+)?(x[12](\^[0-9]+)?)*$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}

func parseNumber(input string, start int) (float64, int, error) {
	end := start
	for end < len(input) && (unicode.IsDigit(rune(input[end])) || input[end] == '.') {
		end++
	}
	numStr := input[start:end]
	val, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, end, fmt.Errorf("'%s' is not a valid float", numStr)
	}
	return val, end, nil
}

func computateFeature(x1 float64, x2 float64, mon string) (float64, error) {
	if mon == "" {
		return 0, fmt.Errorf("empty monomial")
	}

	var result float64 = 1
	for i := 0; i < len(mon); {
		if mon[i] != 'x' {
			return 0, fmt.Errorf("expecting 'x' in position %d", i)
		}
		i++

		if i >= len(mon) || (mon[i] != '1' && mon[i] != '2') {
			return 0, fmt.Errorf("'1' or '2' expected after 'x' in %d", i)
		}
		varVal := x1
		if mon[i] == '2' {
			varVal = x2
		}
		i++

		exp := 1.0
		if i < len(mon) && mon[i] == '^' {
			num, next, err := parseNumber(mon, i+1)
			if err != nil {
				return 0, err
			}
			exp = num
			i = next
		}

		result *= math.Pow(varVal, exp)
	}

	return result, nil
}

func parseFeatures(featuresString string) []string {
	if len(featuresString) < 2 || featuresString[0] != '{' || featuresString[len(featuresString)-1] != '}' {
		return []string{}
	}

	trimmedString := featuresString[1 : len(featuresString)-1]

	if trimmedString == "" {
		return []string{}
	}

	parts := strings.Split(trimmedString, ",")
	return parts
}

func sliceToFloat64Array(slice []float64) js.Value {
	uint8ArrayConstructor := js.Global().Get("Float64Array")
	buffer := js.Global().Get("ArrayBuffer").New(len(slice) * 8)
	view := uint8ArrayConstructor.New(buffer)

	for i, v := range slice {
		view.SetIndex(i, v)
	}
	return view
}
