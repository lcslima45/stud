package entities

type ConsolidacaoEA struct {
	NumTrades      string `json:"num_trades"`
	LucroBruto     string `json:"lucro_bruto"`
	PrejuizoBruto  string `json:"prejuizo_bruto"`
	ResultadoBruto string `json:"resultado_bruto"`
}
