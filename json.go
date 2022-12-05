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

func (o *Order) UnmarshalJSON(data []byte) error {
	type Alias Order

	aux := &struct {
		OpenedAt decimal.Decimal `json:"opentm"`
		StartAt  int64           `json:"starttm"`
		ExpireAt int64           `json:"expiretm"`
		Flags    string          `json:"oflags"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse opened_at
	o.OpenedAt = time.Unix(aux.OpenedAt.Floor().IntPart(), aux.OpenedAt.Sub(aux.OpenedAt.Floor()).Mul(decimal.NewFromInt(10000000)).IntPart())

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
	return nil
}
