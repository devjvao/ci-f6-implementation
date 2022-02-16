package f6

import (
	"ci-f6-implementation/internal/fitness"
	"math"
)

// Impl struct that implements the F6 function
type Impl struct {
	fitness.FunctionModel
}

// Evaluate implementation of the F6 function
func (f Impl) Evaluate(x, y float64) float64 {
	squareSum := math.Pow(x, 2) + math.Pow(y, 2)
	dividend := math.Pow(math.Sin(math.Sqrt(squareSum)), 2) - 0.5
	divisor := math.Pow(1+0.001*squareSum, 2)
	return 0.5 - (dividend / divisor)
}
