package chromosome

import (
	"f6-implementation/internal/fitness"
	"f6-implementation/pkg/util"
	"math"
	"math/rand"
)

const binaryStringValues = "01"

// Model chromosome implementation
type Model struct {
	fitness.Model
	Bin   string
	XBin  string
	YBin  string
	XReal float64
	YReal float64
}

// NewChromosome creates a new random chromosome
func NewChromosome() Model {
	newBin := util.GenerateString(numberBits, binaryStringValues)
	return newChromosomeWithBin(newBin)
}

// NewPresetChromosome creates a chromosome with a binary specified
func NewPresetChromosome(bin string) Model {
	return newChromosomeWithBin(bin)
}

// newChromosomeWithBin creates a chromosome with a binary specified
func newChromosomeWithBin(bin string) Model {
	model := Model{
		Bin: bin,
	}
	model.ProcessBin()
	return model
}

// ProcessBin process the binary value
func (m *Model) ProcessBin() {
	xBin := m.Bin[0:numberBitsEach]
	yBin := m.Bin[numberBitsEach:numberBits]

	multiplier := (domainMax - domainMin) / (math.Pow(2, float64(numberBitsEach)) - 1)

	xReal := float64(util.BinaryToDecimal(xBin))*multiplier + domainMin
	yReal := float64(util.BinaryToDecimal(yBin))*multiplier + domainMin

	m.XBin = xBin
	m.YBin = yBin
	m.XReal = xReal
	m.YReal = yReal
	m.Fitness = functionModel.Evaluate(xReal, yReal)
}

// Mutate randomly mutates the binary value
func (m *Model) Mutate() {
	newBin := ""
	for i := range m.Bin {
		if rand.Float64() < mutationRate {
			newBin += flipBit(string(m.Bin[i]))
		} else {
			newBin += string(m.Bin[i])
		}
	}
	m.ProcessBin()
}

func (m Model) Clone() Model {
	return Model{
		Model: m.Model,
		Bin:   m.Bin,
		XBin:  m.XBin,
		YBin:  m.YBin,
		XReal: m.XReal,
		YReal: m.YReal,
	}
}

func flipBit(bit string) string {
	if bit == "1" {
		return "0"
	}
	return "1"
}
