package utlis

import (
	"fmt"
	"testing"
	"time"
)

func TestAuthenticator_Auth(t *testing.T)  {
	s := &MyStorage{
		count: make(map[int]string),
	}
	s.count[1] = "123"
	auth := MyAuth{storage: s}

	baseUrl := "liuyang"
	password := "123"
	appId := 1
	timeStamp := time.Now().Unix()
	content := fmt.Sprintf("%s+%d+%s+%d",baseUrl,appId,password,timeStamp)
	token := Sha256(content)
	url := generateUrl(baseUrl,token,appId,timeStamp)
	t.Logf(token)
	isvalid,err := auth.auth(url)
	if err != nil{
		t.Error(err)
	}
	if !isvalid{
		t.Error(err)
	}



}
