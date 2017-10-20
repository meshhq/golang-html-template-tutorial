package main

import "html/template"
import "math"

// ThirdViewFormattingFuncMap returns a map of functions
// that will be used by the Third View template
func ThirdViewFormattingFuncMap() template.FuncMap {
	return template.FuncMap{
		"formatOddOrEven": formatOddOrEven,
	}
}

func formatOddOrEven(number int) string {
	remainder := int(math.Abs(float64(number))) % 2
	if remainder == 1 {
		return "Odd"
	}
	return "Even"
}
