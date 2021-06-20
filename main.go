package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return quote.Hello()
}
