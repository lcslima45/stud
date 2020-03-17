package entities

type RelatorioEA struct {
	Financeiro   FinanceiroEA   `json:"financeiro"`
	EstatTrades  EstatTradesEA  `json:"estatistica_trades"`
	Risco        RiscoEA        `json:"risco"`
	Consolidacao ConsolidacaoEA `json:"consolidacao"`
}
