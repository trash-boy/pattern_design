package utlis

import (
	"fmt"
	"testing"
	"time"
)

func TestNewApiRequest(t *testing.T) {
	baseUrl := "liuyang"
	token := "123"
	appId := 1
	timeStamp := time.Now().Unix()

	apiRequest := NewApiRequest(baseUrl,token,appId,timeStamp)
	if apiRequest.GetBaseUrl() != baseUrl{
		t.Error("baseurl error")
	}
	if apiRequest.GetToken() != token{
		t.Error("token error")
	}
	if apiRequest.GetAppId() != appId{
		t.Error("appId error")
	}
	if apiRequest.GetTimeStamp() != timeStamp{
		t.Error("timeStamp error")
	}
}

func TestBuildFromFullUrl(t *testing.T) {
	baseUrl := "liuyang"
	token := "123"
	appId := 1
	timeStamp := time.Now().Unix()
	url := baseUrl + "+" + token + "+" + fmt.Sprintf("%d",appId)  + "+" + fmt.Sprintf("%d",timeStamp)

	apiRequest,err := BuildFromFullUrl(url)
	if err != nil{
		t.Error(err)
	}
	if apiRequest.GetBaseUrl() != baseUrl{
		t.Error("baseurl error")
	}
	if apiRequest.GetToken() != token{
		t.Error("token error")
	}
	if apiRequest.GetAppId() != appId{
		t.Error("appId error")
	}
	if apiRequest.GetTimeStamp() != timeStamp{
		t.Error("timeStamp error")
	}

}
