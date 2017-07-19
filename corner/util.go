package corner

import "math"

func minFloatSlice(slice []float64) float64 {
	m := slice[0]
	for _, x := range slice {
		m = math.Min(m, x)
	}
	return m
}

func maxFloatSlice(slice []float64) float64 {
	m := slice[0]
	for _, x := range slice {
		m = math.Max(m, x)
	}
	return m
}
