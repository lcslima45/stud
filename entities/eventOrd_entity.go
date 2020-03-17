package entities

import (
	"gopkg.in/mgo.v2/bson"
)

// Eventos da ordem
type EventOrd struct {
	ID               bson.ObjectId `bson:"_id,omitempty"`
	ClOrdID          string        `bson:"clordid,omitempty"` // 11 - ClOrdID   Identificador da ordem. Deve ser único por dia ou sessão, e definido pelo iniciador.
	ClOrdLinkID      string        `bson:"clordlinkid"`       // 583 - ClOrdLinkID  Identificador da ordem à qual essa ordem está as-sociada e possui dependência (só será enviada ao mercado caso a ordem dependente seja totalmente executada.
	DataHora         string     
	Evento           string
	Descricao        string
	OrderID          string
	OrigClOrdID      string
	OrdStatusReqID   string
	SecondaryOrderID string
	ExecID           string
	ExecType         string
	ExecTypeStr      string
	OrdStatus        string
	OrdStatusStr     string
	OrdType          string
	OrdTypeStr       string
	OrdRejReason     string
	OrdRejReasonStr  string
	Side             string
	OrdQty           string
	LeavesQty        string
	CumQty           string
	Price            string //	Preço sugerido por contrato.
	AvgPx            string
	LastQty          string
	LastPx           string
	MinQty           int     // 	Quantidade mínima da ordem a ser executada.
	MaxFloor         string  //	Número máximo de contratos da ordem a serem mostrados no livro de ofertas a cada momento.
	StopPx           string  // Preço de Stop para uma ordem Stop Limit. Obriga-tório apenas quando OrdType = 4 ou S.
	StopGainPx       float64 //	Preço de StopGain para uma ordem Stop Limit. Possível apenas em OrdType=4 ou S, mas não o-brigatório.

	TransactTime string

	Mensagem string
}
