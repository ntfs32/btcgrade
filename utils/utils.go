package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pquerna/otp/totp"
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error painc:", err)
			os.Exit(1)
		}
	}()
}

func Md5(sourceString string) string {
	cryptObj := md5.New()
	nx, err := cryptObj.Write([]byte(sourceString))
	if nx != len(sourceString) || err != nil {
		panic("write crypt md5 failed")
	}

	return fmt.Sprintf("%x", cryptObj.Sum([]byte{}))
}

func Sha256(sourceString string, key string) string {
	cryptObj := hmac.New(sha256.New, []byte(key))
	nx, err := cryptObj.Write([]byte(sourceString))
	if nx != len(sourceString) || err != nil {
		panic("write crypt md5 failed")
	}

	return fmt.Sprintf("%x", cryptObj.Sum([]byte{}))

}

func GetGoogleAuthCode(secret string) (string, error) {

	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return code, nil
}

func Signature(access_key string, secret_key string, google_auth_code string, postParams map[string]string) string {
	urls := &url.Values{}
	for key, item := range postParams {
		urls.Add(key, item)
	}
	urls.Add("key", access_key)
	urls.Add("version", "2")
	urls.Add("nonce", strconv.Itoa(int(time.Now().Unix())+10))
	postParams["signature"] = Sha256(urls.Encode(), Md5(secret_key))
	urls.Add("signature", postParams["signature"])
	return urls.Encode()

}

func Get(url string) (body []byte, errBody error) {
	response, resErr := http.Get(url)
	if resErr != nil {
		return nil, resErr
	}
	defer response.Body.Close()
	body, errBody = ioutil.ReadAll(response.Body)
	return
}

func Post(urlStr string, params string) (string, error) {
	resp, errPost := http.Post(urlStr, "application/x-www-form-urlencoded", strings.NewReader(params))
	defer resp.Body.Close()
	if errPost != nil {
		return "", errPost
	}
	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return "", errRead
	}
	return string(body), nil
}
