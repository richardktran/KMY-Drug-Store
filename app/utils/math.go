package utils

func CalculatePercentageChange(previous float64, current float64) float64 {
	return (current - previous) / previous * 100
}
