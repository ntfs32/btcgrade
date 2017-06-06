package main

import (
	"fmt"

	"./quotes"
)

func main() {
	res := quotes.GetQuotes("doge")
	fmt.Println(res)
}
