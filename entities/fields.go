package entities

type Field struct {
	Comment string   `json:"comment"`
	Hidden  bool     `json:"hidden"`
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Opcoes  []string `json:"opcoes"`
	Type    string   `json:"type"`
	Value   string   `json:"value"`
}
