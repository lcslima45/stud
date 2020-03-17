package entities

import "time"

type OrderModal struct {
	Account      int       `json:"Account"`
	Id           string    `json:"Id"`
	IpAddress    string    `json:"IpAddress"`
	OnBehalf     string    `json:"OnBehalf"`
	OrderQty     int       `json:"OrderQty"`
	OrdType      byte      `json:"OrdType"`
	Price        float64   `json:"Price"`
	Verb         string    `json:"verb"`
	RId          uint64    `json:"RId"`
	Side         byte      `json:"Side"`
	OrderID      string    `json:"OrderID"`
	Source       string    `json:"Source"`
	Symbol       string    `json:"Symbol"`
	TimeInforce  byte      `json:"TimeInforce"`
	TransactTime time.Time `json:"TransactTime"`
}

const (
	OrdTypes_None                      = 0
	OrdTypes_Market                    = '1'
	OrdTypes_Limit                     = '2'
	OrdTypes_StopLoss                  = '3'
	OrdTypes_StopLimit                 = '4'
	OrdTypes_MarketWithLeftOverAsLimit = 'K'
	OrdTypes_Reject                    = 'R'
)

const (
	OfferSides_None = 0
	OfferSides_Buy  = '1'
	OfferSidesSell  = '2'

	TimeInforces_None              = 0
	TimeInforces_Day               = '0'
	TimeInforces_GoodTillCancel    = '1'
	TimeInforces_ImmediateOrCancel = '3'
	TimeInforces_FillOrKill        = '4'
	TimeInforces_GoodTillDate      = '6'
	TimeInforces_AtTheClose        = '7'
	TimeInforces_GoodForAuction    = 'A'
)
