package core

import (
	"fmt"
	"log"
	"sns/utils"
	"strings"
)

type Stock struct {
	Name  string
	Code  string
	Price float64
	Chg   float64
	ChgR  float64
}

func (s Stock) String() string {
	return fmt.Sprintf("name=%s, code=%s, price=%.3f, chgr=%.2f%%, chg=%.3f", s.Name, s.Code, s.Price, s.ChgR, s.Chg)
}

func ParseStocks(rowData string, codes []string) []Stock {
	var stocks []Stock
	lines := strings.Split(rowData, ";")
	for i := 0; i < len(lines)-1; i++ {
		cols := strings.Split(lines[i], "\"")
		if len(cols) != 3 {
			log.Printf("stock cols parse error, line: %s, i: %d", lines[i], i)
			continue
		}
		values := strings.Split(cols[1], ",")
		stocks = append(stocks, newStock(codes[i], values))
	}
	return stocks
}

func newStock(code string, values []string) Stock {
	priceY, price := utils.ParseFloat64(values[2]), utils.ParseFloat64(values[3])
	chg := utils.ConvertToFixDecimal(price - priceY)
	chgr := utils.ConvertToFixDecimal((price - priceY) / price * 100)
	return Stock{
		Name:  values[0],
		Code:  code,
		Price: price,
		Chg:   chg,
		ChgR:  chgr,
	}
}
