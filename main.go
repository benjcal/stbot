package main

import (
	"flag"
	tm "github.com/buger/goterm"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func main() {
	var winMargin = flag.Float64("win margin in percent", 1.5, "enter margin to call a win")
	flag.Parse()
	tm.Clear()
	entry := 0.0
	winExit := 0.0
	loseExit := 0.0
	wins := 0
	loses := 0

StartTrade:
	entry = getEntry()
	entry = entry + (entry * percent(0.01))

	winExit = entry + (entry * percent(*winMargin))
	loseExit = entry - (entry * percent(0.4))

	for {
		tm.Printf("------------------------------------\n\n")
		tm.Printf("Margin: %.3f\n", *winMargin)
		tm.Printf("Wins: %v\n", wins)
		tm.Printf("Loses: %v\n\n", loses)
		tm.Printf("Entry:       %f\n", entry)
		tm.Printf("Win Exit:    %f\n", winExit)
		tm.Printf("Lose Exit:   %f\n", loseExit)

		b := trade(entry, winExit, loseExit)
		tm.Printf("current bid: %f\n", b)

		if b >= winExit {
			//if true {
			wins += 1
			tm.Flush()
			goto StartTrade

		}
		if b <= loseExit {
			loses += 1
			tm.Flush()
			goto StartTrade
		}
		tm.Printf("\n------------------------------------\n")
		tm.Flush()
		time.Sleep(time.Second)
	}
}

func trade(e, w, l float64) float64 {
	res, _ := http.Get("https://poloniex.com/public?command=returnOrderBook&currencyPair=USDC_BTC&depth=1")
	body, _ := ioutil.ReadAll(res.Body)

	j := gjson.GetBytes(body, "bids")

	b := j.Value()

	var bids float64

	for _, v := range b.([]interface{}) {
		f, _ := strconv.ParseFloat(v.([]interface{})[0].(string), 64)

		bids += f
	}

	return bids / 1

}

func getEntry() float64 {
	res, _ := http.Get("https://poloniex.com/public?command=returnOrderBook&currencyPair=USDC_BTC&depth=1")
	body, _ := ioutil.ReadAll(res.Body)

	j := gjson.GetBytes(body, "asks")

	a := j.Value()

	var asks float64

	for _, v := range a.([]interface{}) {
		f, _ := strconv.ParseFloat(v.([]interface{})[0].(string), 64)

		asks += f
	}

	return asks / 1
}

func percent(n float64) float64 {
	return n / 100
}
