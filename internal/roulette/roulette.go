package roulette

import (
	"errors"
	"f6-implementation/internal/chromosome"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

func SelectOne(individuals []*chromosome.Model, max float64) *chromosome.Model {
	rouletteRandom := rand.Float64() * max
	rouletteSum := float64(0)

	for _, i := range individuals {
		rouletteSum += i.Fitness
		if rouletteSum > rouletteRandom {
			return i
		}
	}

	log.Fatal(errors.New("internal failure"))
	return nil
}
