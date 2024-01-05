package utlis

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ApiRequest struct {
	baseUrl string
	token string
	appId int
	timeStamp int64
}

func NewApiRequest(baseUrl, token string, appId int, timeStamp int64)*ApiRequest{
	return &ApiRequest{
		baseUrl: baseUrl,
		token: token,
		appId: appId,
		timeStamp: timeStamp,
	}
}

func (ar *ApiRequest)GetBaseUrl()string{
	return ar.baseUrl
}
func (ar *ApiRequest)GetToken()string{
	return ar.token
}
func (ar *ApiRequest)GetAppId()int{
	return ar.appId
}
func(ar *ApiRequest)GetTimeStamp()int64{
	return ar.timeStamp
}

func BuildFromFullUrl(url string)(*ApiRequest,error){
	words := strings.Split(url, "+")
	fmt.Println("words", words)
	if len(words) != 4{
		return nil,errors.New("url 错误")
	}
	appId,_ := strconv.Atoi(words[2])
	timeStamp,_ := strconv.ParseInt(words[3], 10, 64)
	return NewApiRequest(words[0],words[1],appId,timeStamp),nil
}

