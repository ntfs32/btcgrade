package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
)

func init() {
	defer func() {
		if err := recover(); error != nil {
			fmt.Println("Error painc:", err.Error())
			os.Exit(1)
		}
	}()
}

func Md5(sourceString string) string {
	cryptObj := md5.New()
	nx, err := cryptObj.Write([]byte(sourceString))
	if nx != len(sourceString) {
		panic("write crypt md5 failed")
	}

	return fmt.Sprintf("%x", cryptObj.Sum([]byte{}))
}

func Sha256(sourceString string) string {
	cryptObj := sha256.New()
	nx, err := cryptObj.Write([]byte(sourceString))
	if nx != len(sourceString) {
		panic("write crypt md5 failed")
	}

	return fmt.Sprintf("%x", cryptObj.Sum([]byte{}))

}

func GetGoogleAuthCode(secret string) string {

	code, err := totp.GenerateCode(secret, time.Now().Second())
	if(err != nil){
		panic('Generategoogle Code Failed',err.Error())
	}
	return code
}

func Signature(postParams map[string]interface{}) map[string]interface{}{
	var beforeSignStrig string
	for ix,item := range postParams{
		if beforeSignStrig ==""{
			beforeSignStrig =string(ix)+"="+string(item)
		}else{
			beforeSignStrig = beforeSignStrig+"&"+string(ix)+"="+string(item)
		}
	}
	postParams["signature"] = Md5(beforeSignStrig)
	return postParams
	
}
