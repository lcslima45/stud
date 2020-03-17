package entities

type Indicator struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Comment     string         `json:"comment"`
	Hidden      bool           `json:"hidden"`
	Parameteres []ParametersEA `json:"parameteres"`
}
