package trade

import (
	"fmt"
	"log"
	"stbot/poloniex"
	"time"
)

func Run(p *poloniex.Client) {
	// counters
	winCount := 0
	loseCount := 0

CalculateEntry:
	fmt.Print("\033[H\033[2J") // clear screen

	// Determine Sell Volume
	bal := p.GetBalances()

	// move only 98% of the total balance
	vol := bal.USDC * (0.98)
	fmt.Println("volume:      ", vol)

	// Determine Entry Price
	ob := p.GetOrderBook(poloniex.USDC_BTC)
	entryRate := ob.Asks[0].Rate + (ob.Asks[0].Rate * (0.002 / 100))
	fmt.Println("entry:       ", entryRate)

	// Determine Win Exit
	winExit := entryRate + (entryRate * (1.5 / 100))

	// Determine Lose Exit
	loseExit := entryRate - (entryRate * (0.4 / 100))

	// Calculate BTC amount
	amount := vol / entryRate
	fmt.Println("amount:      ", amount)

	// Buy
	err := p.Buy(poloniex.USDC_BTC, entryRate, amount)
	if err != nil {
		log.Println(err)
		time.Sleep(time.Second)
		goto CalculateEntry
	}

	bal = p.GetBalances()

	// Monitor
	for {
		fmt.Print("\033[H\033[2J")

		fmt.Println("Wins!: ", winCount)
		fmt.Println("Loses: ", loseCount)

		fmt.Println("time:        ", time.Now().UTC().Format(time.RFC3339))
		fmt.Println("entry:       ", entryRate)
		fmt.Println("volume:      ", vol)
		fmt.Println("win:         ", winExit)
		fmt.Println("lose:        ", loseExit)
		ob = p.GetOrderBook(poloniex.USDC_BTC)
		fmt.Println("current bid: ", ob.Bids[0].Rate)

	Sell:
		if ob.Bids[0].Rate >= winExit {
			err = p.Sell(poloniex.USDC_BTC, winExit-winExit*(0.005/100), bal.BTC)
			if err != nil {
				log.Println(err)
				goto Sell
			}
			winCount += 1
			goto CalculateEntry
		}

		if ob.Bids[0].Rate <= loseExit {
			err = p.Sell(poloniex.USDC_BTC, loseExit-loseExit*(0.005/100), bal.BTC)
			if err != nil {
				log.Println(err)
				goto Sell
			}
			loseCount += 1
			goto CalculateEntry
		}

		time.Sleep(time.Second)
	}

}
