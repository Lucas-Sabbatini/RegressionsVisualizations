package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"syscall/js"
	"unicode"
)

const m int16 = 100
const n int16 = 100

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

func generateRandomDots(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		js.Global().Get("console").Call("error", "Expected one string argument")
		return nil
	}

	input := args[0].String()

	result, err := computateFeature(2.0, 2.0, input)
	if err != nil {
		js.Global().Call("alert", err.Error())
		return nil
	}
	js.Global().Get("console").Call("log", strconv.FormatFloat(result, 'f', 2, 64))

	return 0
}

func registerCallbacks() {
	js.Global().Set("generateRandomDots", js.FuncOf(generateRandomDots))
}

func main() {
	c := make(chan struct{})
	registerCallbacks()
	<-c
}
