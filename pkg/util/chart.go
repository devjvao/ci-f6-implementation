package util

import (
	"github.com/wcharczuk/go-chart/v2"
	"strconv"
)

// MakeXAxisTicks generates the chart.Ticks and the x values
func MakeXAxisTicks(min, max int) (chart.Ticks, []float64) {
	ticks := chart.Ticks{}
	var values []float64

	for i := min; i < max; i++ {
		values = append(values, float64(i))
		label := ""
		if i%2 == 0 {
			label = strconv.Itoa(i)
		}
		ticks = append(ticks, chart.Tick{
			Value: float64(i),
			Label: label,
		})
	}

	return ticks, values
}
