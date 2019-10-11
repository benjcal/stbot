package poloniex

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func (c *Client) Sell(currencyPair string, rate, amount float64) error {
	nonce := time.Now().UnixNano()
	data := fmt.Sprintf("command=sell&currencyPair=%v&rate=%f&amount=%f&fillOrKill=1&immediateOrCancel=1&nonce=%v", currencyPair, rate, amount, nonce)

	req, err := http.NewRequest(http.MethodPost, "https://poloniex.com/tradingApi", strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	h := hmac.New(sha512.New, []byte(c.secret))
	h.Write([]byte(data))

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Key", c.key)
	req.Header.Add("Sign", hex.EncodeToString(h.Sum(nil)))

	res, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	resErr := gjson.GetBytes(body, "error")
	if resErr.Exists() {
		return errors.New(resErr.Value().(string))
	}

	return nil

}
