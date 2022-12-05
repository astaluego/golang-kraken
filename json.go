package kraken

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func (a *AssetPairsInfo) UnmarshalJSON(data []byte) error {
	type Alias AssetPairsInfo

	aux := &struct {
		Fees      [][]float64 `json:"fees"`
		FeesMaker [][]float64 `json:"fees_maker"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse fees
	for _, fee := range aux.Fees {
		if len(fee) != 2 {
			continue
		}

		f := Fee{
			Volume:  int64(fee[0]),
			Percent: decimal.NewFromFloat(fee[1]),
		}
		a.Fees = append(a.Fees, f)
	}

	// Parse fees_maker
	for _, feeMaker := range aux.FeesMaker {
		if len(feeMaker) != 2 {
			continue
		}

		f := Fee{
			Volume:  int64(feeMaker[0]),
			Percent: decimal.NewFromFloat(feeMaker[1]),
		}
		a.FeesMaker = append(a.FeesMaker, f)
	}

	return nil
}

func (t *AssetTickerInfo) UnmarshalJSON(data []byte) error {
	type Alias AssetTickerInfo

	aux := &struct {
		Ask             []decimal.Decimal `json:"a"`
		Bid             []decimal.Decimal `json:"b"`
		LastTradeClosed []decimal.Decimal `json:"c"`
		Volume          []decimal.Decimal `json:"v"`
		VWAP            []decimal.Decimal `json:"p"`
		CountTrades     []int64           `json:"t"`
		Low             []decimal.Decimal `json:"l"`
		High            []decimal.Decimal `json:"h"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse a
	t.Ask.Price = aux.Ask[0]
	t.Ask.WholeLotVolume = aux.Ask[1]
	t.Ask.Volume = aux.Ask[2]

	// Parse b
	t.Bid.Price = aux.Bid[0]
	t.Bid.WholeLotVolume = aux.Bid[1]
	t.Bid.Volume = aux.Bid[2]

	// Parse c
	t.LastTradeClosed.Price = aux.LastTradeClosed[0]
	t.LastTradeClosed.Volume = aux.LastTradeClosed[1]

	// Parse v
	t.Volume.Today = aux.Volume[0]
	t.Volume.Last24hours = aux.Volume[1]

	// Parse p
	t.VWAP.Today = aux.VWAP[0]
	t.VWAP.Last24hours = aux.VWAP[1]

	// Parse t
	t.CountTrades.Today = aux.CountTrades[0]
	t.CountTrades.Last24hours = aux.CountTrades[1]

	// Parse l
	t.Low.Today = aux.Low[0]
	t.Low.Last24hours = aux.Low[1]

	// Parse h
	t.High.Today = aux.High[0]
	t.High.Last24hours = aux.High[1]

	return nil
}

func (o *Order) UnmarshalJSON(data []byte) error {
	type Alias Order

	aux := &struct {
		OpenedAt float64 `json:"opentm"`
		StartAt  int64   `json:"starttm"`
		ExpireAt int64   `json:"expiretm"`
		Flags    string  `json:"oflags"`
		ClosedAt float64 `json:"closetm"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse opened_at
	o.OpenedAt = time.UnixMicro(int64(aux.OpenedAt * 1000000))

	// Parse start_at
	o.StartAt = time.Unix(aux.StartAt, 0)

	// Parse expire_at
	o.ExpireAt = time.Unix(aux.ExpireAt, 0)

	// Parse oflags
	flags := strings.Split(aux.Flags, ",")
	for _, flag := range flags {
		switch flag {
		case "post":
			o.Flags = append(o.Flags, Post)
		case "fcib":
			o.Flags = append(o.Flags, Fcib)
		case "fciq":
			o.Flags = append(o.Flags, Fciq)
		case "nompp":
			o.Flags = append(o.Flags, Nompp)
		case "viqc":
			o.Flags = append(o.Flags, Viqc)
		}
	}

	// Parse closed_at
	o.ClosedAt = time.UnixMicro(int64(aux.ClosedAt * 1000000))
	return nil
}
