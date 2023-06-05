package main

import (
	"fmt"

	api "github.com/imran-salim/ye/api"
)

func main() {
	censor := true
	fmt.Println(api.QuoteKanyeWest(censor))
}
