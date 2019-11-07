package utils

import "strconv"

func ToFloat(value string) float64 {
	if value == "" {
		return 0.00
	}
	value2, err := strconv.ParseFloat(value, 64);
	if err != nil {

		return 0
	} else {
		return value2
	}
}
