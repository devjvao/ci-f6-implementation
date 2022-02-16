package fitness

type Model struct {
	Fitness float64
}

type FunctionModel interface {
	Evaluate(x, y float64) float64
}
