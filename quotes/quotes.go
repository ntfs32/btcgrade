package quotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	btcgrade "github.com/ntfs32/btcgrade"
)

func init() {
}

func GetQuotes(types btcgrade.Url) btcgrade.QuotesResponse {
	response, resErr := http.Get(fmt.Sprintf(string(btcgrade.QuotesUrl), string(types)))
	check(resErr)
	defer response.Body.Close()
	body, errReadBody := ioutil.ReadAll(response.Body)
	check(errReadBody)
	var resObj btcgrade.QuotesResponse
	jsonEncodeErr := json.Unmarshal(body, &resObj)
	check(jsonEncodeErr)
	return resObj
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
