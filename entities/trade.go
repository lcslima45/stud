package entities

//	"../erros"

type Trade struct {
	Origem   string
	Ativo    string // Simbolo do ativo
	Data     string
	Hora     string  // Horário que o negócio ocorreu
	PrecoUlt float64 //  2 - Preço do último negócio Float
	QtdeNeg  int64   //  6 - Quantidade do negócio atual Inteiro (quantidade)
	//	VolAcum       int64   //  9 - Volume acumulado dos negócios Inteiro (quantidade)
	CdCorUltComp  int64  // 62 Código da corretora que realizou a última compra Inteiro
	CdCorUltVend  int64  // 63 Código da corretora que realizou a última venda Inteiro
	IdNeg         int64  // Identificador único do negócio
	IdReq         string // requisição> Identificador único de requisição
	Direto        string // 0 – Não Direto 1 – Direto
	Agressor_flag string // I – Indefinido A – Comprador V - Vendedor
	Side          string
	PrecoAsk      float64
	PrecoBid      float64
	Seriehist     string
}
