package entities

// ParametersEA  - Recebe os parametros auxiliares para configuração dos campos da estrategia
type ParametersEA struct {
	Id         string   `json:"id"`                   // Id
	Name       string   `json:"name"`                 // Nome do parametro para futura busca de valores armazenados
	Hidden     bool     `json:"hidden"`               // Define se o campo é oculto
	Comment    string   `json:"comment"`              // Nome a ser exibido na tela
	Type       string   `json:"type"`                 // Tipo : combo, checkbox, Numeric, Instrument, Account, String, DateTime, TimeSpan
	Value      string   `json:"value"`                // Valor armazenado
	Opcoes     []string `json:"opcoes,omitempty"`     // Opções permitidas ( aplicavel nos tipos combo)
	Associated bool     `json:"associated,omitempty"` // Deve ser setado para true caso o paramentro seja um "pai" dos demais ( Valido apenas para campos checkbox)
	Father     string   `json:"father,omitempty"`     // Se o "father" (pai) estiver definido, este paramentro só será exibido na tela se o valro do father for true
	Fields     []Field  `json:"fields,omitempty"`     // Aplicável apenas para campos range ( multivalorados )
}

/*
tipos de PAramentros

Combobox — 0
Checkbox — 1
Numeric — 2
Instrument — 4
Account — 5
String — 6
DateTime — 7
TimeSpan — 8
text.range - 9

Exemplos de Uso :

entities.ParametersEA{
				Id:      "1",
				Name:    "HistoryType",
				Hidden:  false,
				Comment: "Tipo de Historico",
				Type:    "combo",
				Value:   "Renko",
				Opcoes:  []string{"Segundo", "Minuto", "Hora", "Dia", "Renko", "Pontos", "Trades"}},

--  Seta o campo como "father"
			entities.ParametersEA{
				Id:         "4",
				Name:       "HabBacktest",
				Hidden:     false,
				Comment:    "Habilitar Backtest",
				Type:       "checkbox",
				Value:      "true",
				Opcoes:     []string{},
				Associated: true}

-- Definindo o campo dependente de um father:

			entities.ParametersEA{
				Id:      "1",
				Name:    "DataInicial",
				Hidden:  false,
				Comment: "Data inicial",
				Type:    "DateTime",
				Value:   "01/07/2018",
				Opcoes:  []string{},
				Father:  "HabBacktest"},

*/
