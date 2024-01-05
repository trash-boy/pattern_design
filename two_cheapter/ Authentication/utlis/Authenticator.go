package utlis

import (
	"errors"
)

type Authenticator interface {
	auth(url string)bool
}


type MyAuth struct {
	storage CredentialStorage
}

func (auth *MyAuth)auth(url string)(bool,error){
	//1.根据url构造apirequest对象
	//2.根据apirequest对象构造authtoken对象
	//3.查询本地模型验证authtoken对象是否正确
	apiRequest,err := BuildFromFullUrl(url)
	if err != nil{
		return false,err
	}
	url = apiRequest.GetBaseUrl()
	token := apiRequest.GetToken()
	timeStamp := apiRequest.GetTimeStamp()
	appId := apiRequest.GetAppId()
	clientAuthToken := NewAuthToken(token,timeStamp)
	if !clientAuthToken.IsExpired(){
		return false,errors.New("token 过期")
	}
	password := auth.storage.getPasswordByAppId(appId)
	serverAuthToken := clientAuthToken.Generate(url,appId,password,timeStamp)
	if !clientAuthToken.Match(serverAuthToken){
		return false,errors.New("与本地不匹配")
	}
	return true,nil

}
