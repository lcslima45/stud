package entities

type OrdemWeb struct {
	OrdemID       string `json:"ordem_id"`
	DataHora      string `json:"data_hora"`
	Ativo         string `json:"ativo"`
	CV            string `json:"cv"`
	Quantidade    string `json:"quantidade"`
	PrecoLimite   string `json:"preco_limite"`
	Status        string `json:"status"`
	Tipo          string `json:"tipo"`
	QtdeExec      string `json:"qtde_executada"`
	PrecoMedio    string `json:"preco_medio"`
	ResAbsoluto   string `json:"resultado_abs"`
	ResRelativo   string `json:"resultado_rel"`
	ResFinanceiro string `json:"resultado_financ"`
}
