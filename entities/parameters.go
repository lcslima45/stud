package entities

// Parameters - Classe para tratar o Json entre os modulos da plataforma para inicializar a robo com os parametros configurados
type Parameters struct {
	Acc             Account            `json:"acc"`
	Instrument      Security           `json:"instrument"`
	Charts          ParameterClassEA   `json:"charts"`
	Backtest        ParameterClassEA   `json:"backtest"`
	Indicators      []Indicator        `json:"indicators"`
	Enters          []ParameterClassEA `json:"enters"`
	Exits           ParameterClassEA   `json:"exits"`
	DiaryExits      ParameterClassEA   `json:"diary_exits"`
	Timeconstraints ParameterClassEA   `json:"time_constraints"`
	Daytrade        ParameterClassEA   `json:"day_trade"`
	Range           ParameterClassEA   `json:"range"`
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

*/
