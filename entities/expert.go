package entities

import "gopkg.in/mgo.v2/bson"

type Expert struct {
	ID                 bson.ObjectId `json:"_id,omitempty"`
	Nome               string
	Robo               string
	Resumo             Summary
	Relatorio          RelatorioEA
	DrawdownMaximo     string
	GrafLucroid        string
	GrafLucroX         []string
	GrafLucroY         []float64
	GrafPriceid        string
	GrafPriceX         []string
	GrafPriceY         [][4]float64
	Parametros         Parameters
	SaldoInicial       string
	Resultado          string
	Taxas              string
	TradesLucro        string
	TradesPreju        string
	FatorLucro         string
	DrawdownAbs        string
	DrawdownRelat      string
	DrawdownMaxRelat   string
	LucroBruto         string
	PrejuBruto         string
	ResultadoBruto     string
	Ordens             []OrdemWeb
	GraficoPerformance GraficoPerformanceEA
}
