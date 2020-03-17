package entities

type RiscoEA struct {
	FatorDeLucro        string `json:"fator_de_lucro"`
	DrawdownAbsoluto    string `json:"drawdown_abs"`
	DrawdownRelativo    string `json:"drawdown_rel"`
	DrawdownMaxAbsoluto string `json:"drawdown_max_abs"`
	DrawdownMaxRElativo string `json:"drawdown_max_rel"`
}
