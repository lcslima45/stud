package entities

type Charts struct {
	Hidden bool           `json:"hiden"`
	Titulo string         `json:"titulo"`
	Param  []ParametersEA `json:"param"`
}
