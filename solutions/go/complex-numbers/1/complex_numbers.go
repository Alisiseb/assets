package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	real      float64
	imaginary float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imaginary
}

func (n1 Number) Add(n2 Number) Number {
	new := Number{}
	new.real = n1.real + n2.real
	new.imaginary = n1.imaginary + n2.imaginary
	return new
}

func (n1 Number) Subtract(n2 Number) Number {
	new := Number{}
	new.real = n1.real - n2.real
	new.imaginary = n1.imaginary - n2.imaginary
	return new
}

func (n1 Number) Multiply(n2 Number) Number {
	new := Number{}
	new.real = n1.real*n2.real - n1.imaginary*n2.imaginary
	new.imaginary = n1.real*n2.imaginary + n1.imaginary*n2.real
	return new
}

func (n Number) Times(factor float64) Number {
	new := Number{}
	new.real = n.real * factor
	new.imaginary = n.imaginary * factor
	return new
}

func (n1 Number) Divide(n2 Number) Number {
	new := Number{}
	denominator := n2.real*n2.real + n2.imaginary*n2.imaginary
	new.real = (n1.real*n2.real + n1.imaginary*n2.imaginary) / denominator
	new.imaginary = (n1.imaginary*n2.real - n1.real*n2.imaginary) / denominator
	return new
}

func (n Number) Conjugate() Number {
	new := Number{}
	new.real = n.real
	new.imaginary = -n.imaginary
	return new
}
func (n Number) Abs() float64 {
	return math.Sqrt(n.real*n.real + n.imaginary*n.imaginary)
}

func (n Number) Exp() Number {
	new := Number{}
	eToTheX := math.Exp(n.real)
	new.real = eToTheX * math.Cos(n.imaginary)
	new.imaginary = eToTheX * math.Sin(n.imaginary)
	return new
}
