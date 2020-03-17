package entities

type Summary struct {
	RetornoLiquido  string `json:"retorno_liquido"`
	RetornoBruto    string `json:"retorno_bruto"`
	Patrimonio      string `json:"patrimonio"`
	DrawdownMax     string `json:"drawdown_max"`
	DrawdownMaxPerc string `json:"drawdown_max_perc"`
	TradesTotal     string `json:"trades_total"`
}
