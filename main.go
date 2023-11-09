package main

import (
	"fmt"

	ye "github.com/i8abyte/ye/client"
)

func main() {
	censor := true
	fmt.Println(ye.QuoteYe(censor))
}
