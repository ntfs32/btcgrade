package btcgrade

type Url string

type signature string

// 基本配置
//goole二次验证key（非动态密码，生成otp认证的密码）
//api名称
//私钥
type Config struct {
	GoogleAuthCode string
	ApiName        string
	PriviteKey     string
}

const (
	BTC  Url = "btc"
	ETH  Url = "eth"
	LTC  Url = "ltc"
	DOGE Url = "doge"
	YBC  Url = "ybc"
)

const (
	// ----------------行情API-------------------------
	//	获取行情 (GET)
	QuotesUrl Url = "http://api.btctrade.com/api/ticker?coin=%s"

	//	深度行情 (GET)
	DepthQuotesUrl Url = "http://api.btctrade.com/api/depth?coin=%s"

	//	成交记录 (GET)
	OrderUrl Url = "http://api.btctrade.com/api/trades?coin=%s"
)

const (
	//	------------------------交易API-----------------------
	//	获取账户信息 (POST)
	AccountInfoUrl Url = "http://api.btctrade.com/api/balance/"

	//	挂单查询 (POST): 指定时间后的挂单，可以根据类型查询，比如查看正在挂单和全部挂单

	AccoutPauseOrderUrl Url = "http://api.btctrade.com/api/orders/"

	//	查询订单信息 (POST)
	AccountOrderUrl Url = "http://api.btctrade.com/api/fetch_order/"
)
const (
	// -----------------	完整交易API------------------
	//取消挂单 (POST)
	CannelPauseOrderUrl Url = "http://api.btctrade.com/api/cancel_order/"

	//	挂买单 (POST)
	BuyOrderUrl Url = "http://api.btctrade.com/api/buy/"

	//	挂卖单 (POST)
	SellOrderUrl Url = "http://api.btctrade.com/api/sell/"
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
	High float64 `json:"hogh"`
	Low  float64 `json:"low"`
	Buy  float64 `json:"buy"`
	Sell float64 `json:"sell"`
	Last float64 `json:"last"`
	Vol  string  `json:"vol"`
	Time int64   `json:"time"`
}
