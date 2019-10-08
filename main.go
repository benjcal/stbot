package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
)

type config struct {
	Key    string
	Secret string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Pass the config file as first value")
		os.Exit(0)
	}

	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var c config
	_, err = toml.Decode(string(f), &c)
	if err != nil {
		log.Fatal(err)
	}

	//p := poloniex.NewClient(c.Key, c.Secret)

	//trade.Run(p)
}
