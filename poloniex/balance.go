package poloniex

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Balances struct {
	USDC float64 `json:"USDC,string"`
	BTC  float64 `json:"BTC,string"`
}

func (c *Client) GetBalances() *Balances {
	nonce := time.Now().UnixNano()
	data := fmt.Sprintf("command=returnBalances&nonce=%v", nonce)

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

	var b Balances
	err = json.Unmarshal(body, &b)
	if err != nil {
		log.Fatal(err)
	}

	return &b
}
