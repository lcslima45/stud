package entities

type GraficoPerformanceEA struct {
	Titulo  string    `json:"titulo"`
	Labels  []string  `json:"labels"`
	Valores []float64 `json:"values"`
}
