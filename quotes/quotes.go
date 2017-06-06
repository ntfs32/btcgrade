package quotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Url string

const (
	QuotesUrl Url = "http://api.btctrade.com/api/ticker?coin=%s"
)

var QuotesResponseName map[string]string = map[string]string{
	"high": "最高价",
	"low":  "最低价",
	"buy":  "买一价",
	"sell": "卖一价",
	"last": "最新成交价",
	"vol":  "成交量(最近的24小时)",
	"time": "返回数据时服务器时间",
}

type QuotesResponse struct {
	high float64
	low  float64
	buy  float64
	sell float64
	last float64
	vol  float64
	time int64
}

func init() {
}

func GetQuotes(types string) QuotesResponse {
	response, resErr := http.Get(fmt.Sprintf(string(QuotesUrl), string(types)))
	check(resErr)
	defer response.Body.Close()
	body, errReadBody := ioutil.ReadAll(response.Body)
	check(errReadBody)
	resObj := &QuotesResponse{}
	jsonEncodeErr := json.Unmarshal(body, resObj)
	check(jsonEncodeErr)
	fmt.Println(resObj.last)
	return *resObj
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
