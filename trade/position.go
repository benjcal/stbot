package trade

import (
	"fmt"
	"stbot/poloniex"
)

const LONG = "LONG"
const SHORT = "SHORT"

type Position struct {
	Type         string
	CurrencyPair string
	Rate         float64
	Amount       float64
	Win          float64
	Lose         float64
}

func NewShort(cp string, rate, amount, margin, bail float64) *Position {
	win := rate - rate*(margin/100)
	lose := rate + rate*(bail/100)

	return &Position{
		Type:         SHORT,
		CurrencyPair: cp,
		Rate:         rate,
		Amount:       amount,
		Win:          win,
		Lose:         lose,
	}

}

func NewLong(cp string, rate, amount, margin, bail float64) *Position {
	win := rate + rate*(margin/100)
	lose := rate - rate*(bail/100)

	return &Position{
		Type:         LONG,
		CurrencyPair: cp,
		Rate:         rate,
		Amount:       amount,
		Win:          win,
		Lose:         lose,
	}

}

func (pos *Position) Excecute(p *poloniex.Client) {
	if pos.Type == LONG {
		// LONG
		// buy at 1 & sell at 1.5 for 0.5 profit

		fmt.Printf("at %v buy %v BTC for a total of %v USDC\n", pos.Rate, pos.Amount, pos.Rate*pos.Amount)

		fmt.Printf("at %v sell %v BTC for a total of %v USDC\n", pos.Win, pos.Amount, pos.Win*pos.Amount)
		fmt.Printf("at %v sell %v BTC for a total of %v USDC\n", pos.Lose, pos.Amount, pos.Lose*pos.Amount)

		//err := p.Buy(pos.CurrencyPair, pos.Rate, pos.Amount)
		//if err != nil {
		//	// do something
		//}

		// monitor bids
		//...

		//err = p.Sell(pos.CurrencyPair, pos.Rate, )
	}

	if pos.Type == SHORT {
		fmt.Printf("at %v sell %v BTC for a total of %v USDC\n", pos.Rate, pos.Amount, pos.Rate*pos.Amount)

		fmt.Printf("at %v buy %v BTC for a total of %v USDC\n", pos.Win, pos.Amount, pos.Win*pos.Amount)
		fmt.Printf("at %v buy %v BTC for a total of %v USDC\n", pos.Lose, pos.Amount, pos.Lose*pos.Amount)
	}

	// SHORT
	// sell at 1.5 & buy same amount at 1 for 0.5 profit

}
