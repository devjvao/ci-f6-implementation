package roulette

import (
	"ci-f6-implementation/internal/chromosome"
	"errors"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

// SelectOne selects one random chromosome
func SelectOne(individuals []chromosome.Model, max float64) chromosome.Model {
	rouletteRandom := rand.Float64() * max
	rouletteSum := float64(0)

	for _, i := range individuals {
		rouletteSum += i.Fitness
		if rouletteSum > rouletteRandom {
			return i
		}
	}

	log.Fatal(errors.New("internal failure"))
	return chromosome.Model{}
}
