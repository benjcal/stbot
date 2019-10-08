package trade

type Position struct {
	CurrencyPair string
	Rate         float64
	Amount       float64
	Win          float64
	Lose         float64
}
