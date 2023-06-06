package ye

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const uri = "https://api.kanye.rest"

func GetQuote() []byte {
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
