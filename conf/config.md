>	"GoogleAuthCode": Google Authenticator Token
>	"ApiName":	API Name default account email 
>	"PriviteKey": privite key

package main

import (
	"fmt"

	"github.com/ntfs32/btcgrade"
	"github.com/ntfs32/btcgrade/utils"
)

func main() {
	quotesData := btcgrade.Ticker(btcgrade.DOGE)
	fmt.Println(quotesData)
	fmt.Println(btcgrade.DepthTicker(btcgrade.DOGE))
	fmt.Println(btcgrade.CommonOrders(btcgrade.DOGE))

	fmt.Println(utils.GetGoogleAuthCode("code"))
	var config btcgrade.Config
	config.SetAccess_key("code")
	config.SetSecret_key("code")
	fmt.Println(btcgrade.AccountInfo(config))
	fmt.Println(btcgrade.AccountOrders(config, btcgrade.DOGE, "all", "1493500731", "ASC"))
	fmt.Println(btcgrade.FetchOrder(config, "360964718"))
}
