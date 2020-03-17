package entities

import (
	"time"
)

type Bar struct {
	Ativo        string
	Tipo         string
	Tamanho      string
	Time         time.Time
	Open         float64
	High         float64
	Low          float64
	Close        float64
	Volume       int64
	Ticks        int
	Agressor     float64
	VolCompra    int64
	VolVenda     int64
	VolDireto    int64
	Complementar bool
}
