package fitness

import (
	"ci-f6-implementation/pkg/util"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/wcharczuk/go-chart/v2"
	"os"
)

// Model abstract struct that contains the Fitness value
type Model struct {
	Fitness float64
}

// FunctionModel interface that has an Evaluate expression
type FunctionModel interface {
	Evaluate(x, y float64) float64
}

// GenerateChart generates the fitness chart
func GenerateChart(bestFitnessGenerations []float64) {
	xTicks, xValues := util.MakeXAxisTicks(0, len(bestFitnessGenerations))

	graph := chart.Chart{
		Title:      "Fitness Evolution by Generation",
		TitleStyle: chart.StyleTextDefaults(),
		XAxis: chart.XAxis{
			Name:  "Generation",
			Ticks: xTicks,
		},
		YAxis: chart.YAxis{
			Name: "Fitness",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: bestFitnessGenerations,
			},
		},
	}

	f, err := os.Create("fitness-chart.png")
	if err != nil {
		logrus.Error(errors.New("failed to create/open fitness chart file"))
		return
	}
	defer util.Ignore(f.Close)

	err = graph.Render(chart.PNG, f)
	if err != nil {
		logrus.Error(errors.New("failed to generate fitness chart"))
		return
	}
}
