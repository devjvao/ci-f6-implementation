package chromosome

import "f6-implementation/internal/fitness"

var (
	functionModel  fitness.FunctionModel
	domainMin      float64
	domainMax      float64
	numberBits     int64
	numberBitsEach int64
	mutationRate   float64
)

func Initialize(funcModel fitness.FunctionModel, min, max float64, bits int64, mutation float64) {
	functionModel = funcModel
	domainMin = min
	domainMax = max
	numberBits = bits
	numberBitsEach = bits / 2
	mutationRate = mutation
}
