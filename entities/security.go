package entities

type Security struct {
	Exchange ParametersEA `json:"exchange"`
	Symbol   string `json:"simbol"`
}
