package mymath

import "math"

// Sqrt - функция для вычисления квадратного корня из числа
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
func Add(x, y float64) float64 {
	return x + y
}

func Subtract(x, y float64) float64 {
	return x - y
}

func Multiply(x, y float64) float64 {
	return x * y
}

func Divide(x, y float64) float64 {
	return x / y
}

func Floor(x float64) float64 {
	return math.Floor(x)
}

func Ceil(x float64) float64 {
	return math.Ceil(x)
}

func Round(x float64) float64 {
	return math.Round(x)
}

func Trunc(x float64) float64 {
	return math.Trunc(x)
}

func Maximum(x, y float64) float64 {
	return math.Max(x, y)
}

func Minimum(x, y float64) float64 {
	return math.Min(x, y)
}

func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

func SquareRoot(x float64) float64 {
	return math.Sqrt(x)
}

func Sine(x float64) float64 {
	return math.Sin(x)
}

func Cosine(x float64) float64 {
	return math.Cos(x)
}

func Tangent(x float64) float64 {
	return math.Tan(x)
}
func Abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}
func Max(x, y float64) float64 {
    if x > y {
        return x
    }
    return y
}
func Sin(x float64) float64 {
    return math.Sin(x)
}
func Cos(x float64) float64 {
    return math.Cos(x)
}
func Yn(n int, x float64) float64 {
    if n == 0 {
        return (math.Sin(x) / x)
    } else if n == 1 {
        return ((math.Sin(x) / x) - (math.Cos(x) / x))
    } else {
        return (((2 * float64(n) - 1) / x) * Yn(n-1, x) - Yn(n-2, x))
    }
}
