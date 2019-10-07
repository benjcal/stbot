package poloniex

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Order struct {
	Rate   float64
	Volume float64
}

type OrderBook struct {
	Asks []Order
	Bids []Order
}

func (c *Client) GetOrderBook(currencyPair string) *OrderBook {
	res, _ := http.Get(fmt.Sprintf("https://poloniex.com/public?command=returnOrderBook&currencyPair=%s&depth=50", currencyPair))
	body, _ := ioutil.ReadAll(res.Body)

	bodyMap := gjson.ParseBytes(body)

	asks := parseOrders(bodyMap.Value().(map[string]interface{})["asks"])
	bids := parseOrders(bodyMap.Value().(map[string]interface{})["bids"])

	return &OrderBook{
		Asks: asks,
		Bids: bids,
	}
}

func parseOrders(orders interface{}) []Order {
	var ordersArray []Order

	for _, v := range orders.([]interface{}) {
		r, _ := strconv.ParseFloat(v.([]interface{})[0].(string), 64)

		o := Order{
			Rate:   r,
			Volume: v.([]interface{})[1].(float64),
		}

		ordersArray = append(ordersArray, o)
	}

	return ordersArray
}
