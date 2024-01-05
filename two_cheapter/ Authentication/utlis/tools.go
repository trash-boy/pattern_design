package utlis

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Sha256(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

func generateUrl(baseUrl, token string, appId int, timeStamp int64)string{

	return  baseUrl + "+" + token + "+" + fmt.Sprintf("%d",appId)  + "+" + fmt.Sprintf("%d",timeStamp)
}
