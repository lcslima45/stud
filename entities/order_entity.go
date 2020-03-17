package entities

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type OrdType string

const (
	OrdType_MARKET                         OrdType = "1"
	OrdType_LIMIT                          OrdType = "2"
	OrdType_STOP                           OrdType = "3"
	OrdType_STOP_LIMIT                     OrdType = "4"
	OrdType_MARKET_ON_CLOSE                OrdType = "5"
	OrdType_WITH_OR_WITHOUT                OrdType = "6"
	OrdType_LIMIT_OR_BETTER                OrdType = "7"
	OrdType_LIMIT_WITH_OR_WITHOUT          OrdType = "8"
	OrdType_ON_BASIS                       OrdType = "9"
	OrdType_ON_CLOSE                       OrdType = "A"
	OrdType_LIMIT_ON_CLOSE                 OrdType = "B"
	OrdType_FOREX_MARKET                   OrdType = "C"
	OrdType_PREVIOUSLY_QUOTED              OrdType = "D"
	OrdType_PREVIOUSLY_INDICATED           OrdType = "E"
	OrdType_FOREX_LIMIT                    OrdType = "F"
	OrdType_FOREX_SWAP                     OrdType = "G"
	OrdType_FOREX_PREVIOUSLY_QUOTED        OrdType = "H"
	OrdType_FUNARI                         OrdType = "I"
	OrdType_MARKET_IF_TOUCHED              OrdType = "J"
	OrdType_MARKET_WITH_LEFT_OVER_AS_LIMIT OrdType = "K"
	OrdType_PREVIOUS_FUND_VALUATION_POINT  OrdType = "L"
	OrdType_NEXT_FUND_VALUATION_POINT      OrdType = "M"
	OrdType_PEGGED                         OrdType = "P"
	OrdType_COUNTER_ORDER_SELECTION        OrdType = "Q"
)

type OrderSide string // Direção da ordem. Valores válidos: 1 = Compra 2 = Venda

const (
	OrderSideCompra OrderSide = "1"
	OrderSideVenda  OrderSide = "2"
)

type Side string

const (
	Side_BUY                Side = "1"
	Side_SELL               Side = "2"
	Side_BUY_MINUS          Side = "3"
	Side_SELL_PLUS          Side = "4"
	Side_SELL_SHORT         Side = "5"
	Side_SELL_SHORT_EXEMPT  Side = "6"
	Side_UNDISCLOSED        Side = "7"
	Side_CROSS              Side = "8"
	Side_CROSS_SHORT        Side = "9"
	Side_CROSS_SHORT_EXEMPT Side = "A"
	Side_AS_DEFINED         Side = "B"
	Side_OPPOSITE           Side = "C"
	Side_SUBSCRIBE          Side = "D"
	Side_REDEEM             Side = "E"
	Side_LEND               Side = "F"
	Side_BORROW             Side = "G"
)

type AccAlloc struct {
	AllocAccount      string
	AllocAcctIDSource int // Fonte da conta. Valor aceito: 99 – Outro (custom or proprietary code)
}
type PartyOffer struct {
	PartyID       string // Identificador do Participante.
	PartyIDSource string // Identifica a origem do PartyID. O único valor acei-to é: D = Propietary/Custom code
	PartyRole     int    // Identifica o tipo do participante. Valores aceitos: 36 = Entering Trader
}

// Mensagem para envio de ordem simples para o mercado. As ordens devem conter identificador único (tag ClOr-deID será o identificador único) atribuído pelo sistema para o dia de pregão. As ordens com identificador único serão re-jeitadas.
type Order struct {
	ID                       bson.ObjectId `bson:"_id,omitempty"`
	ClOrdID                  string        `bson:"clordid,omitempty"` // 11 - ClOrdID   Identificador da ordem. Deve ser único por dia ou sessão, e definido pelo iniciador.
	ClOrdLinkID              string        `bson:"clordlinkid"`       // 583 - ClOrdLinkID  Identificador da ordem à qual essa ordem está as-sociada e possui dependência (só será enviada ao mercado caso a ordem dependente seja totalmente executada.
	Ativo                    string
	Lado                     string // 54 - Side
	Evento                   string
	IDdaordem                string
	Tipo                     OrdType
	PrecoStop                string
	Quantidade               string // 38 - OrderQty
	Tempo                    string
	Data                     string
	Validade                 string
	Login                    string
	Conta                    string
	TipodeAtivo              string
	Comprado                 string
	Vendido                  string
	PrecoStrike              string
	DataExp                  string
	OrdStatus                string
	OrdStatusStr             string
	OrdType                  OrdType //	Tipo da ordem. Pode ter um dos seguintes valores: A = OnClose 1 = Mercado 2 = Limite 4 = Stop Limit K = Mercado resto à limite S = Start Na versão MegaDirect_V3 foi retirado tipo 1.
	OrdTypeStr               string
	OrdRejReason             string
	OrdRejReasonStr          string
	OrdRejSource             string
	OrdQty                   float64
	LeavesQty                string
	CumQty                   string
	Price                    float64 //	Preço sugerido por contrato.
	AvgPx                    string
	LastQty                  string
	LastPx                   string
	Allocs                   []AccAlloc
	Parties                  []PartyOffer
	Side                     Side      // 	Direção da ordem. Valores válidos: 1 = Compra 2 = Venda
	MinQty                   int       // 	Quantidade mínima da ordem a ser executada.
	MaxFloor                 string    //	Número máximo de contratos da ordem a serem mostrados no livro de ofertas a cada momento.
	OrderRestrictions        string    // 	MultipleValueString  Restrições relacionadas à ordem. Se mais de uma restrição se aplicar à ordem, esse campo poderá conter instruções múltiplas separadas por espaço. Valores: 1 = Program trading 7 = Entidade estrangeira 8 = Participante de mercado estrangeiro
	StopPx                   string    // Preço de Stop para uma ordem Stop Limit. Obriga-tório apenas quando OrdType = 4 ou S.
	StopGainPx               float64   //	Preço de StopGain para uma ordem Stop Limit. Possível apenas em OrdType=4 ou S, mas não o-brigatório.
	Price2                   float64   //	Preço limite para StopGain.
	MovingStart              float64   // 	Preço de ínicio móvel, configurado para Ordens Start/Stop Móvel. Obrigatório quando InitialChan-ge é informado.
	InitialChangeType        int       //	Código que determina se valor em campo Initial-Change será representação percentual ou absoluta. Valores possíveis: 1 = Percentual
	InitialChange            float64   // Preço de alteração inicial do Price, quando Last-Trade ultrapassar MovingStart, configurado quan-do ordem Start/Stop Móvel. Obrigatório quando campos MovingStart e InitialChangeType forem informados. (Se percentual, 1.0 = 100%)
	PricePercentage          float64   //Percentual usado para calcular o preço limite em uma ordem Start/Stop Móvel.
	TimeInForce              string    // Especifica o tipo de duração da ordem. A ausência desse campo indica que seu valor é 0 (Dia). Pode ter um dos seguintes valores: 0 = Dia (ou sessão) 1 = Good Till Cancel (GTC) 2 = At The Opening (OPG) 3 = Executa (mesmo que parcial) imediatamente ou Cancela (IOC) 4 = Executa por completo ou Cancela (FOK) 6 = Good Till Date (GTD) 7 = At The Close (ATC) A = Good For Auction (GFA)
	ExpireTime               time.Time // Data e hora da expiração da ordem, este campo dever ser definido em UTC.Obrigatório quando TimeInForce = GTD (6).
	BrokerID                 string    //Código da corretora de destino.
	OrderStrategy            string    //Nome da estratégia.
	HandlInst                string    //Indica que é uma ordem administrada. Valor acei-to: 2 = Automática com possibilidade de intervenção do Broker 3 = Manual
	OrderTag                 string    //Tag (observação) informada para a ordem. Para ordens enviadas ao mercado através de uma sessão EntryPoint seu valor é repassado para o campo Memo (5149). Obs.: No EntryPoint são permitidos no máximo 50 caracteres para este campo.
	Currency                 float64   // Identifica a moeda corrente utilizada no preço.
	SettlCurrency            float64   //Código da moeda corrente de denominação do es-
	TargetStrategy           int       // A estratégia de destino da ordem.
	TargetStrategyParameters string    //Parametros da estratégia de destino definida entre as partes.
	SourceAddress            string    //Utilizado para armazenar a origem da ordem
	EventOrds                []EventOrd
	Mensagem                 string
}
