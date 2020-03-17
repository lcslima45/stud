package indicators.core

import (
	entities "github.com/autotrend/entities"
)

type Indicador interface {

	// GetValue - Obter o valor de um buffer 
	// Parametros buffer - Buffer do indicador
	//            offset - Posicao no buffer (cresce da direita para esquerda, zero é a posicao mais atual do buffer )
	GetValue(buffer int, offset int) (float64, error)

	// SetValue - Atribuir um valor ao buffer 
	// Parametros buffer - Buffer do indicador
	//            offset - Posicao no buffer (cresce da direita para esquerda, zero é a posicao mais atual do buffer )
	//            value  - Valor a ser armazenado
	SetValue(buffer int, offset int, Value float64) error

	//Ontick - deve ser efocado a cada negociação do ativo target
	// Parametro trade - ultimo negocio recebido
	OnTick(trade entities.Trade) error 

	//OnNextBar - deve ser efocado a cada negociação do ativo target
	//Parametro bar - barra atual
	OnNextBar(bar entities.Bar) error

	//GetLenth - total de posicoes do buffer
	//Parametro buffer - buffer de referencia
	GetLenth(buffer int) int

	//Init - Inicializador do Indicador ( deve ser evocado no inicio para informe dos paramentros de inicializacao)
	////Parametro paramenters - slice de interface com os parametros de inicializacao
	Init(parameters ...interface{}) error

}

