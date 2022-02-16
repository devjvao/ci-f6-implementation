package chromosome

import (
	"f6-implementation/internal/binary"
	"f6-implementation/internal/fitness"
	"f6-implementation/internal/rand"
	"math"
	rand2 "math/rand"
)

const binaryStringValues = "01"

type Model struct {
	fitness.Model
	Bin   string
	XBin  string
	YBin  string
	XReal float64
	YReal float64
}

func NewChromosome() *Model {
	bin := rand.StringBytes(int(numberBits), binaryStringValues)
	return newChromosomeWithBin(bin)
}

func NewPresetChromosome(bin string) *Model {
	return newChromosomeWithBin(bin)
}

func newChromosomeWithBin(bin string) *Model {
	model := &Model{
		Bin: bin,
	}
	model.ProcessBin()
	return model
}

func (m *Model) ProcessBin() {
	xBin := m.Bin[0:numberBitsEach]
	yBin := m.Bin[numberBitsEach:numberBits]

	multiplier := (domainMax - domainMin) / (math.Pow(2, float64(numberBitsEach)) - 1)

	xReal := float64(binary.ToDecimal(xBin))*multiplier + domainMin
	yReal := float64(binary.ToDecimal(yBin))*multiplier + domainMin

	m.XBin = xBin
	m.YBin = yBin
	m.XReal = xReal
	m.YReal = yReal
	m.Fitness = functionModel.Evaluate(xReal, yReal)
}

func (m *Model) Mutate() {
	newBin := ""
	for i := range m.Bin {
		if rand2.Float64() < mutationRate {
			newBin += flipBit(string(m.Bin[i]))
		} else {
			newBin += string(m.Bin[i])
		}
	}
	m.ProcessBin()
}

func flipBit(bit string) string {
	if bit == "1" {
		return "0"
	}
	return "1"
}
