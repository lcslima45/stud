package entities

import (
	"time"
)

type TickType uint

const (
	NewTick TickType = iota
	NewBar
)

type BarsAggregations int

const (
	None      BarsAggregations = 1 + iota //0 This value mean no, or classic aggregations type.
	ByWicks                               // 1 This value will result in creating a renko data with wicks, which will provide more information to the trader than the standard renko bricks do.
	ByHighLow                             //  2 Use high/low method, for Renko and PnF.
)

type HistoryType int

const (
	Simple    HistoryType = 0 + iota // 0 Simple bars objects represent the Time, Tick, Price Range bars of data.
	Renko                            // 1 Represent access to Renko bars objects.
	Kagi                             // 2 Represent access to Kagi bars objects.
	LineBreak                        // 3 Represent access to Line Break bars objects.
	PnF                              // 4 Represent access to PnF bars objects.
	Point
	HistoryType_Segundo HistoryType = 0 + iota
	HistoryType_Minuto
	HistoryType_Hora
	HistoryType_Dia
	HistoryType_Semana
	HistoryType_Mes
	HistoryType_Renko
	HistoryType_Pontos
	HistoryType_Trade
)

type AppliedExtrem int

const (
	HIGHLOW AppliedExtrem = 0 + iota
	HIGHLOWOPEN
	HIGHLOWCLOSE
	OPENHIGHLOW
	CLOSEHIGHLOW
)

type PriceType int

const (
	Open     PriceType = 0 + iota //0 Use open price.
	Close                         // 1 Use close price.
	High                          // 2 Use high price.
	Low                           // 3 Use low price.
	Medium                        // 4 Use medium price.
	Typical                       // 5 Use typical price.
	Weighted                      // 6 Use weihgted price.

)

type MAMode int

const (
	MASimple MAMode = 0 + iota
	MAExponential
	MAModified
	MALinearWeighted
)

type Period time.Duration

const (
	Day      = time.Hour * 24   //  86400 // 1 Day Period.
	Hour     = time.Hour        // 1 Hour Period.
	Hour4    = time.Hour * 4    //4 Hour Period.
	Min      = time.Minute      // 1 Minute Period.
	Min15    = time.Minute * 15 // 15 Minute Period.
	Min30    = time.Minute * 30 // 30 Minute Period.
	Min5     = time.Minute * 5  // 5 Minute Period.
	Second30 = time.Second * 30 // Second Period.
	Second15 = time.Second * 15 // Second Period.
	Second5  = time.Second * 5  // Second Period.
	Second   = time.Second      // Second Period.
	Tick     = 0                // Tick Period.
)

type PeriodPT int

const (
	TickPT   = 0      // Tick Period.
	MinPT    = 1      // 1 Minute Period.
	Min5PT   = 5      // 5 Minute Period.
	Min15PT  = 15     // 15 Minute Period.
	Min30PT  = 30     // 30 Minute Period.
	HourPT   = 60     // 1 Hour Period.
	Hour4PT  = 240    // 4 Hour Period.
	DayPT    = 1440   // 1 Day Period.
	WeekPT   = 10080  // 1 Week Pertiod.
	MonthPT  = 43200  // 1 Monthl Period.
	YearPT   = 518400 // 1 Year Period.
	SecondPT = 10093  // 1 Second Period.
)
