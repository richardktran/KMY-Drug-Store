package utils

import "math"

func CalculatePercentageChange(previous float64, current float64) float64 {
	percentage := (current - previous) / previous * 100

	return math.Round(percentage*100) / 100
}
