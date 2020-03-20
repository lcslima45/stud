package indicators

import (
	entities "stud/entities"
	"errors"
)

// TEMA - Media movel Exponencial
type TEMA struct {
	valores     [][]float64     	// Buffer par armazenar os valroes calculados
	bars        []entities.Bar  	// Armazena as barras
	valLastTick float64         	// Variavel auxiliar par otimizacao de recalculo da media  - so calcular se for diferente do valor anterior
	MaPeriod    int             	// Periodo da media movel exponencial 
	SourcePrice entities.PriceType 	// Preço referencia para calulo da média movel exponencial
	Open        float64
	Close		float64
	High		float64
	Low			float64

}

//GetLenth - total de posicoes do buffer
//Parametro buffer - buffer de referencia
func (i *TEMA) GetLenth(buffer int) int {

	// Retorna o tamanho do slice do buffer em questão
	return len(i.valores[buffer])
}


// GetValue - Obter o valor de um buffer 
// Parametros buffer - Buffer do indicador
//            offset - Posicao no buffer (cresce da direita para esquerda, zero é a posicao mais atual do buffer )
func (i *TEMA) GetValue(buffer int, offset int) (float64, error) {

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
func (i *TEMA) SetValue(buffer int, offset int, Value float64) error {

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
func (i *TEMA) Init(parameters ...interface{}) error {

	// Periodo da media movel 
	i.MaPeriod = parameters[0].(int)

	// Preço referência da barra 
	i.SourcePrice = parameters[1].(entities.PriceType)

	// Inicializa Buffer
	i.valores = make([][]float64, 1)

	//  Inicializa o buffer de medias
	i.valores[0] = make([]float64, 1)

	// inicializar a variavel auxiliar para reculculo da média
	i.valLastTick = -1

	return nil
}

//OnNextBar - deve ser efocado a cada negociação do ativo target
//Parametro bar - barra atual
func (i *TEMA) OnNextBar(bar entities.Bar) error {
	//No começo a nova barra tem os preços semelhantes ao fechamento da barra anterior.
	i.Open  = bar.Close
	i.Low   = bar.Close
	i.High  = bar.Close
	i.Close = bar.Close
	i.valLastTick = bar.Close
	newBar := entities.Bar{} //verificando se o vetor de valores de i é vazio
	if len(i.valores) == 0 {
		i.valores = make([][]float64, 1)
	}
	// Adicionar um elemento no inicio do buffer
	if len(i.valores) > 0 {
		i.valores[0] = append([]float64{0}, i.valores[0]...)
	} else {
		i.valores[0] = append(i.valores[0], 0)
	}

	i.bars[0] = bar
	i.bars = append([]entities.Bar{newBar}, i.bars...)
	//se houverem barras suficientes realizar o cálculo da média exponencial
	var mediaExp float64
	if len(i.bars) > i.MaPeriod {
		var novoPreco float64
		k := 2/float64(i.MaPeriod + 1)
		prevMediaEx, _ := i.GetValue(0, 1)

		switch i.SourcePrice {

		case entities.Open:
			novoPreco = bar.Open
			mediaExp = prevMediaEx + k*(novoPreco-prevMediaEx)

		case entities.Low:
			novoPreco = bar.Low
			mediaExp = prevMediaEx + k*(novoPreco-prevMediaEx)

		case entities.Close:
			novoPreco = bar.Close
			mediaExp = prevMediaEx + k*(novoPreco-prevMediaEx)

		}
		i.SetValue(0, 0, mediaExp)
	} else if (len(i.bars)-1) == i.MaPeriod {
		// calcular Média Simples como base para a Média exponencial base
		var summa float64
		switch i.SourcePrice {
		case entities.Open:
			for j := 1; j<=i.MaPeriod; j++ {
				summa = summa + i.bars[j].Open
			}
		case entities.Low:
			for j := 1; j<=i.MaPeriod; j++ {
				summa = summa + i.bars[j].Low
			}
		case entities.Close:
			for j := 1; j<= i.MaPeriod; j++ {
				summa = summa + i.bars[j].Close
			}
		case entities.High:
			for j := 1; j <= i.MaPeriod; j++ {
				summa = summa + i.bars[j].High
			}
		}
		mediaSimples := summa/float64(i.MaPeriod)
		i.SetValue(0, 0, mediaSimples)
	} else {
		i.SetValue(0,0,0)
	}
	//Retorna erro nulo
	return nil

}

//Ontick - deve ser efocado a cada negociação do ativo target
// Parametro trade - ultimo negocio recebido
func (i *TEMA) OnTick(trade entities.Trade) error {

	// Se o Valor do tick mudou
	if i.valLastTick != trade.PrecoUlt {

		// Armazena novo valor 
		i.valLastTick = trade.PrecoUlt

		// obtem preço corrente 
		price := trade.PrecoUlt

		// Verificação, se a quantidade atual de barras for maior, do que o período da média móvel. 
		//Se for, então o cálculo é possível
		if len(i.valores[0]) > i.MaPeriod {

			// Calculando o coeficiente
			var k float64
			k = 2.0 / float64(i.MaPeriod+1)

			// Obtendo a media anterior
			prevEMA, _ := i.GetValue(0, 1)

			// Calcular o novo valor da media
			ema := prevEMA + k*(price-prevEMA)

			// Atualizar o valor da media
			i.SetValue(0, 0, ema)

		} else {

			// Inicializar o valor da media com o primeiro valor
			i.SetValue(0, 0, price)
		}

	}

	return nil
}


