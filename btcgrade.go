package btcgrade

import (
	"strconv"

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

//	执行post请求数据
func PostAction(urls string, config Config, params map[string]string) string {
	sign := utils.Signature(config.ACCESS_KEY, config.SECRET_KEY, config.GOOGLE_AUTH_CODE, params)
	body, _ := utils.Post(API_URL+urls, sign)
	return string(body)
}

//	设置ACCESS_KEY
func (obj *Config) SetAccess_key(key string) Config {
	obj.ACCESS_KEY = key
	return *obj
}

//	设置SECRET_KEY
func (obj *Config) SetSecret_key(key string) Config {
	obj.SECRET_KEY = key
	return *obj
}

//	设置google auth code
func (obj *Config) SetGoogleAuthCode(key string) Config {
	obj.GOOGLE_AUTH_CODE = key
	return *obj
}

//	换取行情信息
//coin: 交易币种（btc,eth,ltc,doge,ybc）
func Ticker(coin string) string {
	body, _ := utils.Get(API_URL + "/ticker?coin=" + coin)
	return string(body)
}

//	深度行情信息
//coin: 交易币种（btc,eth,ltc,doge,ybc）
func DepthTicker(coin string) string {
	body, _ := utils.Get(API_URL + "/depth?coin=" + coin)
	return string(body)
}

//	公共成交记录
//coin: 交易币种（btc,eth,ltc,doge,ybc）
func CommonOrders(coin string) string {
	body, _ := utils.Get(API_URL + "/trades?coin=" + coin)
	return string(body)
}

//	当前帐号的账户余额信息
func AccountInfo(config Config) string {
	return PostAction("/balance", config, make(map[string]string, 0))
}

//	当前账户的挂单历史列表
//type: 挂单类型[open:正在挂单, all:所有挂单]
//since: 时间戳, 查询某个时间戳之后的挂单
//ob: 排序, ASC,DESC
func AccountOrders(config Config, coin string, types string, since string, ob string) string {
	postParams := make(map[string]string, 4)
	postParams["coin"] = coin
	postParams["type"] = types
	postParams["since"] = since
	postParams["ob"] = ob
	return PostAction("/orders", config, postParams)
}

//	当前账户的挂单详情
//id: 挂单ID
func FetchOrder(config Config, id string) string {
	postParams := make(map[string]string, 1)
	postParams["id"] = id
	return PostAction("/fetch_order", config, postParams)
}

//	取消挂单
//id: 挂单ID
func CancelOrder(config Config, id string) string {
	postParams := make(map[string]string, 1)
	postParams["id"] = id
	return PostAction("/cancel_order", config, postParams)
}

//	挂买单
//coin: 交易币种（btc,eth,ltc,doge,ybc）
//amount: 购买数量
//price: 购买价格
func Buy(config Config, coin string, amount int, price float64) string {
	postParams := make(map[string]string, 3)
	postParams["coin"] = coin
	postParams["amount"] = strconv.Itoa(amount)
	postParams["price"] = strconv.FormatFloat(price, 'f', 9, 64)
	return PostAction("/buy", config, postParams)
}

//	挂卖单
//coin: 交易币种（btc,eth,ltc,doge,ybc）
//amount: 购买数量
//price: 购买价格
func Sell(config Config, coin string, amount int, price float64) string {
	postParams := make(map[string]string, 3)
	postParams["coin"] = coin
	postParams["amount"] = strconv.Itoa(amount)
	postParams["price"] = strconv.FormatFloat(price, 'f', 9, 64)
	return PostAction("/sell", config, postParams)
}
