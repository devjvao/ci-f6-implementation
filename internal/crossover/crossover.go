package crossover

import (
	"ci-f6-implementation/internal/chromosome"
	"math"
	"math/rand"
)

// Parents struct that contains two chromosomes to do the crossover
type Parents struct {
	ChromosomeA chromosome.Model
	ChromosomeB chromosome.Model
}

// ShouldDoCrossover randomly check if the crossover will be done
func (c Parents) ShouldDoCrossover() bool {
	return rand.Float64() < crossoverRate
}

// DoExchangeOnCutPoint swap chromosome tails at the cut-off point
func (c Parents) DoExchangeOnCutPoint() (chromosome.Model, chromosome.Model) {
	newBinA := ""
	newBinB := ""

	cutPoint := int(math.Floor(rand.Float64()*float64(len(c.ChromosomeA.Bin))-1)) + 1

	for i := range c.ChromosomeA.Bin {
		if i < cutPoint {
			newBinA += string(c.ChromosomeA.Bin[i])
			newBinB += string(c.ChromosomeB.Bin[i])
		} else {
			newBinA += string(c.ChromosomeB.Bin[i])
			newBinB += string(c.ChromosomeA.Bin[i])
		}
	}

	newChromosomeA := chromosome.NewPresetChromosome(newBinA)
	newChromosomeB := chromosome.NewPresetChromosome(newBinA)

	return newChromosomeA, newChromosomeB
}
