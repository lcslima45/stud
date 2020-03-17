package entities

// Exchange - Corretora que irá rodar a estratégia
type Exchange struct {
	Id     int    `json:"id"`     // Id da Corretora no Bco de Dados
	Name   string `json:"name"`   // Nome da Corretora
	Status string `json:"status"` // Status  ( Activated / Deactivated )
}

/*

Exemplo de Uso

	Exchange: entities.Exchange{
		Id: 1,
		Name: "Orama",
		Status: "Activated"},

*/
