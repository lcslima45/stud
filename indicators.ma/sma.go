package indicators

import (
	entities "stud/entities"
	"errors"
)

// TEMA - Media movel Simples
type TSMA struct {
	valores     [][]float64 		// Buffer par armazenar os valroes calculados
	valLastTick float64     		// Variavel auxiliar par otimizacao de recalculo da media  - so calcular se for diferente do valor anterior
	bars        []entities.Bar 		// Armazena as barras  
	MaPeriod    int            		// Periodo da media movel
	SourcePrice entities.PriceType 	// Preço referencia para calulo da média movel
	Open 				float64
	High 				float64
	Low					float64
	Close				float64

}

//GetLenth - total de posicoes do buffer
//Parametro buffer - buffer de referencia
func (i *TSMA) GetLenth(buffer int) int {
	return len(i.valores[buffer])
}


// GetValue - Obter o valor de um buffer 
// Parametros buffer - Buffer do indicador
//            offset - Posicao no buffer (cresce da direita para esquerda, zero é a posicao mais atual do buffer )
func (i *TSMA) GetValue(buffer int, offset int) (float64, error) {

	// Verifica se o buffer requisitado existe
	if buffer > (len(i.valores) - 1) {

		// Retorna valor negativo e o erro
		return -1, errors.New("buffer is out of range")

	// Verifica se o offset é menor q o tamanho do slice	
	} else if offset > (len(i.valores[buffer]) - 1) {

		// Retorna valor negativo e o erro
		return -1, errors.New("offset is out of range")
	} else {

		// Retorna o valor requisitado
		return i.valores[buffer][offset], nil
	}

}

// SetValue - Atribuir um valor ao buffer 
// Parametros buffer - Buffer do indicador
//            offset - Posicao no buffer (cresce da direita para esquerda, zero é a posicao mais atual do buffer )
//            value  - Valor a ser armazenado
func (i *TSMA) SetValue(buffer int, offset int, Value float64) error {

	// Verifica se o buffer requisitado existe
	if buffer > (len(i.valores) - 1) {

		// Retorna erro
		return errors.New("buffer is out of range")
	}

	// Verifica se o offset é menor q o tamanho do slice	
	if offset > (len(i.valores[buffer]) - 1) {
		// Retorna erro
		return errors.New("offset is out of range")
	}

	// Armazena o valor na posição 
	i.valores[buffer][offset] = Value

	// Retorna erro nulo
	return nil

}

//Init - Inicializador do Indicador ( deve ser evocado no inicio para informe dos paramentros de inicializacao)
//Parametro paramenters - slice de interface com os parametros de inicializacao
func (i *TSMA) Init(maPeriod int, sourceprice entities.PriceType) error {

	// Periodo da media movel 
	i.MaPeriod = maPeriod

	// Preço referência da barra 
	i.SourcePrice = sourceprice

	// Inicializa Buffer
	i.valores = make([][]float64, 1)

	// inicializar a variavel auxiliar para reculculo da média
	i.valLastTick = -1

	return nil
}


func (i *TSMA) OnTick(trade entities.Trade) error {
	if i.valLastTick != trade.PrecoUlt{
		i.Close = trade.PrecoUlt
		if i.High < trade.PrecoUlt{
			i.High = trade.PrecoUlt
		}
		if i.Low > trade.PrecoUlt{
			i.Low = trade.PrecoUlt
		}

		if len(i.bars) > i.MaPeriod {
			var summa float64
			switch i.SourcePrice{
			case entities.Close:
				for j := 1; j<=i.MaPeriod; j++{
					close := i.bars[j].Close
					summa = summa + close

				}
			case entities.Low:
				for j := 1; j<=i.MaPeriod; j++{
					low := i.bars[j].Low
					summa = summa + low
				}
			case entities.High:
				for j := 1; j<= i.MaPeriod; j++{
					high := i.bars[j].High
					summa = summa + high
				}
			case entities.Open:
				for j := 1; j <= i.MaPeriod; j++{
					open := i.bars[j].Open
					summa = summa + open
				}
			}
			valor := summa/float64(i.MaPeriod)
			i.SetValue(0,0, valor)
		} else {
			i[0][0] = 0
		}
	}


	return nil
}

//OnNextBar - deve ser efocado a cada negociação do ativo target
//Parametro bar - barra atual
func (i *TSMA) OnNextBar(bar entities.Bar) error {
	// Iniciando da nova barra
	i.Open = bar.Close
	i.High = bar.Close
	i.Low = bar.Close
	i.Close = bar.Close
	// Adicionar um elemento no inicio do buffer
	if len(i.valores) == 0 {
		i.valores = make([][]float64, 1)
	}
	newBar := entities.Bar{}
	// Adicionar um elemento no inicio do buffer
	if len(i.valores[0]) > 0 {
		i.valores[0] = append([]float64{0}, i.valores[0]...)

	} else {
		i.valores[0] = append(i.valores[0], 0)
	}

	i.bars[0] = bar
	i.bars = append([]entities.Bar{newBar}, i.bars...)
	// Se o Valor do tick mudou
	i.valLastTick = bar.Close

	// Checking, if current amount of bars
	// more, than period of moving average. If it is
	// then the calculation is possible
	if len(i.bars) > i.MaPeriod {

		var summa float64 // Sum of prices
		summa = 0

		//var maint int
		//maint = int(i.MaPeriod)
		switch i.SourcePrice{
		case entities.Open:
			for j := 1; j<=i.MaPeriod; j++ {
				open := i.bars[j].Open
				summa = summa + open
			}

		case entities.Close:
			for j := 1; j <= i.MaPeriod; j++ {
				// Adding bar's price to the summa
				close := i.bars[j].Close

				summa = summa + close

			}
		case entities.High:
			for j := 1; j <= i.MaPeriod; j++ {
				high := i.bars[j].High
				summa = summa + high
			}

		case entities.Low:
			for j := 1; j <= i.MaPeriod; j++ {
				low := i.bars[j].Low
				summa = summa + low
			}

		}

		valor := summa / float64(i.MaPeriod)

		// Setting the current value of the indicator
		i.SetValue(0, 0, valor)

	} else {
		// Checking, if current amount of bars
		// more, than period of moving average. If it is
		// then the calculation is possible
		i.valores[0][0] = 0
		return nil
	}

	return nil

}
