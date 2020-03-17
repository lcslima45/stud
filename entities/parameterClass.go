package entities

// ParameterClassEA - Classe para aramazenar os parametros de um determinado campo
type ParameterClassEA struct {
	Hidden bool           `json:"hiden"`  // Indica se o campo deve ser oculto
	Titulo string         `json:"titulo"` // Titulo a ser mostrado na tela de input
	Param  []ParametersEA `json:"param"`  // Paramentros auxiliares do campo
}

// ObterParametro - Retorna o Valor de um parametro especifico
func (p ParameterClassEA) ObterParametro(key string) (valor string) {

	//Percorre a lista de parametros procuranod por uma chave especifica
	for _, prm := range p.Param {

		// Compara o nome do parametro com a chave especifica
		if prm.Name == key {

			// Se encontrou a chave, obter o valor do parametro
			valor = prm.Value

			// Encerrar o loop
			break
		}
	}

	// Retornar o valor padrão caso não encontre a chave
	return valor

}
