package main

import (
	"f6-implementation/internal/chromosome"
	"f6-implementation/internal/crossover"
	"f6-implementation/internal/f6"
	"f6-implementation/internal/roulette"
	"f6-implementation/pkg/log"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	populationSize = 100
	numGenerations = 40
)

func main() {
	log.InitLog(logrus.InfoLevel)

	fitnessModel := &f6.Impl{}
	chromosome.Initialize(fitnessModel, -100, 100, 44, 0.1)
	crossover.Initialize(0.6)

	var bestIndividual *chromosome.Model

	for g := 0; g < numGenerations; g++ {
		var currentBestIndividual *chromosome.Model

		rouletteMax := float64(0)
		individuals := make([]*chromosome.Model, populationSize)

		for i := range individuals {
			individuals[i] = chromosome.NewChromosome()
			rouletteMax += individuals[i].Fitness
		}

		var parentsList []*crossover.Parents
		for i := 0; i < populationSize/2; i++ {
			parentsList = append(parentsList, &crossover.Parents{
				ChromosomeA: roulette.SelectOne(individuals, rouletteMax),
				ChromosomeB: roulette.SelectOne(individuals, rouletteMax),
			})
		}

		var newIndividuals []*chromosome.Model
		for _, parents := range parentsList {
			if parents.ShouldDoCrossover() {
				newA, newB := parents.DoExchangeOnCutPoint()
				newIndividuals = append(newIndividuals, newA)
				newIndividuals = append(newIndividuals, newB)
			} else {
				newIndividuals = append(newIndividuals, parents.ChromosomeA)
				newIndividuals = append(newIndividuals, parents.ChromosomeB)
			}
		}

		for _, individual := range newIndividuals {
			individual.Mutate()
			if currentBestIndividual == nil || currentBestIndividual.Fitness < individual.Fitness {
				currentBestIndividual = individual
			}
		}

		logrus.Info(fmt.Sprintf("[%d] Best individual: {XReal: %f, YReal: %f, Fitness: %f}", g, currentBestIndividual.XReal, currentBestIndividual.YReal, currentBestIndividual.Fitness))

		if bestIndividual == nil || bestIndividual.Fitness < currentBestIndividual.Fitness {
			bestIndividual = currentBestIndividual
		}
	}

	logrus.Info(fmt.Sprintf("[All] Best individual: {XReal: %f, YReal: %f, Fitness: %f}", bestIndividual.XReal, bestIndividual.YReal, bestIndividual.Fitness))
}
