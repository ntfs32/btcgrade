package btcgrade

import (
	"github.com/ntfs32/btcgrade/utils"
)

type signature string

const (
	API_VERSION = "2"
)

type Config struct {
	GOOGLE_AUTH_CODE string
	ACCESS_KEY       string
	SECRET_KEY       string
}

const (
	BTC  string = "btc"
	ETH  string = "eth"
	LTC  string = "ltc"
	DOGE string = "doge"
	YBC  string = "ybc"
)

const (
	API_URL string = "http://api.btctrade.com/api"
)

// 行情api返回
type QuotesResponse struct {
	High float64 `json:"hogh"`
	Low  float64 `json:"low"`
	Buy  float64 `json:"buy"`
	Sell float64 `json:"sell"`
	Last float64 `json:"last"`
	Vol  string  `json:"vol"`
	Time int64   `json:"time"`
}

// 公共订单数据
type CommonOrdersResponse struct {
	Date   string  `json:"date"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
	Tid    string  `json:"tid"`
	Type   string  `josn:"type"`
}

func (obj *Config) SetAccess_key(key string) Config {
	obj.ACCESS_KEY = key
	return *obj
}
func (obj *Config) SetSecret_key(key string) Config {
	obj.SECRET_KEY = key
	return *obj
}
func (obj *Config) SetGoogleAuthCode(key string) Config {
	obj.GOOGLE_AUTH_CODE = key
	return *obj
}

func Ticker(coin string) string {
	body, _ := utils.Get(API_URL + "/ticker?coin=" + coin)
	return string(body)
}

func DepthTicker(coin string) string {
	body, _ := utils.Get(API_URL + "/depth?coin=" + coin)
	return string(body)
}

func CommonOrders(coin string) string {
	body, _ := utils.Get(API_URL + "/trades?coin=" + coin)
	return string(body)
}

func AccountInfo(config Config) string {
	params := utils.Signature(config.ACCESS_KEY, config.SECRET_KEY, config.GOOGLE_AUTH_CODE, make(map[string]string, 0))
	body, _ := utils.Post(API_URL+"/balance", params)
	return string(body)
}

func AccountOrders(config Config, coin string, types string, since string, ob string) string {
	//	postParams := &map[string]string{"coin": coin, "type": types, "since": since, "ob": ob}
	postParams := make(map[string]string, 4)
	postParams["coin"] = coin
	postParams["type"] = types
	postParams["since"] = since
	postParams["ob"] = ob
	params := utils.Signature(config.ACCESS_KEY, config.SECRET_KEY, config.GOOGLE_AUTH_CODE, postParams)
	body, _ := utils.Post(API_URL+"/orders", params)
	return string(body)
}

func FetchOrder(config Config, orderId string) string {
	postParams := make(map[string]string, 1)
	postParams["id"] = orderId
	params := utils.Signature(config.ACCESS_KEY, config.SECRET_KEY, config.GOOGLE_AUTH_CODE, postParams)
	body, _ := utils.Post(API_URL+"/fetch_order", params)
	return string(body)
}

func CancelOrder(config Config, orderId string) string {
	postParams := make(map[string]string, 1)
	postParams["id"] = orderId
	params := utils.Signature(config.ACCESS_KEY, config.SECRET_KEY, config.GOOGLE_AUTH_CODE, postParams)
	body, _ := utils.Post(API_URL+"/cancel_order", params)
	return string(body)
}
