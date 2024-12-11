package mathutils

func FindMin(a, b, c float64) float64 {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func CalculateAverage(a, b, c float64) float64 {
	return (a + b + c) / 3.0
}

func SolveLinearEquation(a, b float64) float64 {
	return -b / a
}
