package main

import (
	"ci-f6-implementation/internal/chromosome"
	"ci-f6-implementation/internal/crossover"
	"ci-f6-implementation/internal/f6"
	"ci-f6-implementation/internal/fitness"
	"ci-f6-implementation/internal/roulette"
	"ci-f6-implementation/pkg/log"
	"flag"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/sirupsen/logrus"
	"strings"
)

var logSpaces = strings.Repeat("=", 100)

var (
	bitsSize       int
	domainMin      float64
	domainMax      float64
	mutationRate   float64
	crossoverRate  float64
	populationSize int
	numGenerations int
)

func init() {
	log.InitLog(logrus.InfoLevel)

	flag.IntVar(&bitsSize, "b", 44, "The chromosome bits size")
	flag.IntVar(&populationSize, "p", 100, "The population size")
	flag.IntVar(&numGenerations, "g", 40, "The maximum number of generations")
	flag.Float64Var(&domainMin, "min", -100, "The domain minimum")
	flag.Float64Var(&domainMax, "max", 100, "The domain maximum")
	flag.Float64Var(&mutationRate, "m", 0.008, "The mutation rate")
	flag.Float64Var(&crossoverRate, "c", 0.65, "The crossover rate")

	flag.Parse()

	chromosome.Initialize(&f6.Impl{}, domainMin, domainMax, bitsSize, mutationRate)
	crossover.Initialize(crossoverRate)

	logrus.Info(logSpaces)
	logrus.Info("F6 implementation")
	logrus.Info(logSpaces)
	logrus.Info("Parameters: ")
	logrus.Info(fmt.Sprintf("\tChromosome bits size: %d", bitsSize))
	logrus.Info(fmt.Sprintf("\tDomain: { min: %.5f, max: %.5f }", domainMin, domainMax))
	logrus.Info(fmt.Sprintf("\tMutation rate: %.5f", mutationRate))
	logrus.Info(fmt.Sprintf("\tCrossover rate: %.5f", crossoverRate))
	logrus.Info(fmt.Sprintf("\tPopulation size: %d", populationSize))
	logrus.Info(fmt.Sprintf("\tMaximum generations: %d", numGenerations))
	logrus.Info(logSpaces)
}

func main() {
	var (
		bar                    = progressbar.Default(int64(numGenerations))
		bestIndividual         = chromosome.Model{}
		bestFitnessGenerations []float64
	)

	// Randomly generates the first population
	individuals := make([]chromosome.Model, populationSize)
	for i := range individuals {
		individuals[i] = chromosome.NewChromosome()
	}

	for g := 0; g < numGenerations; g++ {
		var (
			rouletteMax                 = float64(0)
			currentBestIndividual       = chromosome.Model{}
			currentWorstIndividualIndex = 0
		)

		// Calculate the maximum value for the roulette
		for i := range individuals {
			rouletteMax += individuals[i].Fitness
		}

		// Separate parents to generate new individuals
		var parentsList []crossover.Parents
		for i := 0; i < populationSize/2; i++ {
			parentsList = append(parentsList, crossover.Parents{
				ChromosomeA: roulette.SelectOne(individuals, rouletteMax),
				ChromosomeB: roulette.SelectOne(individuals, rouletteMax),
			})
		}

		// Apply crossover between parents
		var newIndividuals []chromosome.Model
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

		// Apply mutation on each chromosome
		for i, individual := range newIndividuals {
			individual.Mutate()
			if individual.Fitness < newIndividuals[currentWorstIndividualIndex].Fitness {
				currentWorstIndividualIndex = i
			}
		}

		// Replace the worst chromosome with the best one
		if bestIndividual.Bin != "" {
			newIndividuals[currentWorstIndividualIndex] = bestIndividual.Clone()
		}

		// Check the best chromosome of the current generation
		for _, individual := range newIndividuals {
			if currentBestIndividual.Fitness < individual.Fitness {
				currentBestIndividual = individual
			}
		}

		// Prepare the next generation
		individuals = newIndividuals

		// Add best fitness to list
		bestFitnessGenerations = append(bestFitnessGenerations, currentBestIndividual.Fitness)

		// Assign best individual
		if bestIndividual.Fitness < currentBestIndividual.Fitness {
			bestIndividual = currentBestIndividual
			_ = bar.Clear()
			logrus.Info(fmt.Sprintf("Chromosome improvement on generation %d: { XReal: %.5f, YReal: %.5f, Fitness: %.5f }",
				g, currentBestIndividual.XReal, currentBestIndividual.YReal, currentBestIndividual.Fitness))
		}

		_ = bar.Add(1)
	}

	// Calculate the fitness average
	fitnessSum := float64(0)
	for _, f := range bestFitnessGenerations {
		fitnessSum += f
	}
	fitnessAverage := fitnessSum / float64(len(bestFitnessGenerations))

	logrus.Info(logSpaces)
	logrus.Info(fmt.Sprintf("Fitness average: %.5f", fitnessAverage))
	logrus.Info(fmt.Sprintf("Best chromosome of all: { XReal: %.5f, YReal: %.5f, Fitness: %.5f }",
		bestIndividual.XReal, bestIndividual.YReal, bestIndividual.Fitness))
	logrus.Info(logSpaces)

	fitness.GenerateChart(bestFitnessGenerations)
}
