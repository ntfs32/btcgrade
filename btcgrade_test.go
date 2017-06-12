package btcgrade

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	ACCESS_KEY = "3499-3493-3dccfcvdc-f44s-3dff"
	SECRET_KEY = "T3dOsq()*:LKKJIY$@@#$%TGHH"
)

func TestTicker(t *testing.T) {
	ticker := Ticker("doge")
	fmt.Println(ticker)
	var testString string
	require.IsTypef(t, ticker, testString, "获取行情失败")
}

func TestDepthTicker(t *testing.T) {
	ticker := DepthTicker("doge")
	fmt.Println(ticker)
	var testString string
	require.IsType(t, ticker, testString, "获取深度行情失败")
}

func TestCommonOrders(t *testing.T) {
	ticker := CommonOrders("doge")
	fmt.Println(ticker)
	var testString string
	require.IsType(t, ticker, testString, "获取公共交易数据失败")
}

func TestAccountInfo(t *testing.T) {
	var config Config
	config.SetAccess_key(ACCESS_KEY)
	config.SetSecret_key(SECRET_KEY)
	info := AccountInfo(config)
	fmt.Println(info)
	var testString bool
	require.IsType(t, info, testString, "获取账户的余额信息失败")
}
