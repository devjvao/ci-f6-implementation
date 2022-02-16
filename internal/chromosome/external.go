package chromosome

import "ci-f6-implementation/internal/fitness"

var (
	functionModel  fitness.FunctionModel
	domainMin      float64
	domainMax      float64
	numberBits     int
	numberBitsEach int
	mutationRate   float64
)

// Initialize set the genetic parameters
func Initialize(funcModel fitness.FunctionModel, min, max float64, bits int, mutation float64) {
	functionModel = funcModel
	domainMin = min
	domainMax = max
	numberBits = bits
	numberBitsEach = bits / 2
	mutationRate = mutation
}
