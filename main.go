package mymath

import "math"

// Sqrt - функция для вычисления квадратного корня из числа
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	return x / y
}

func floor(x float64) float64 {
	return math.Floor(x)
}

func ceil(x float64) float64 {
	return math.Ceil(x)
}

func round(x float64) float64 {
	return math.Round(x)
}

func trunc(x float64) float64 {
	return math.Trunc(x)
}

func maximum(x, y float64) float64 {
	return math.Max(x, y)
}

func minimum(x, y float64) float64 {
	return math.Min(x, y)
}

func power(x, y float64) float64 {
	return math.Pow(x, y)
}

func squareRoot(x float64) float64 {
	return math.Sqrt(x)
}

func sine(x float64) float64 {
	return math.Sin(x)
}

func cosine(x float64) float64 {
	return math.Cos(x)
}

func tangent(x float64) float64 {
	return math.Tan(x)
}
