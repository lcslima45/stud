package entities

// Account - Dados da Conta que irá rodar a estrategia
type Account struct {
	Id         int          `json:"id"`         // Id no Bco de dados
	AccNum     string       `json:"acc_num"`    // Numero da conta na corretora
	Exchange   Exchange     `json:"corretora"`  // Dados da Corretora
	Enviroment ParametersEA `json:"enviroment"` // Ambiente que irá rodar {"Backtest", "Simulador", "Real"}
	Status     string       `json:"status"`     // Status  ( Activated / Deactivated )
}

/*

Exemplo de Uso

	Acc = entities.Account{
		Id:       1,
		AccNum:   "0001",
		Exchange: entities.Exchange{1, "Orama", "Activated"},
		Enviroment: entities.ParametersEA{
			Id:      "1",
			Name:    "Enviroment",
			Hidden:  false,
			Comment: "Ambiente",
			Type:    "combo",
			Value:   "Simulador",
			Opcoes:  []string{"Backtest", "Simulador", "Real"},
		},
		Status: "Ativated",
	}

*/
