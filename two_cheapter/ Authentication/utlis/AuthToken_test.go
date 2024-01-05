package utlis

import (
	"testing"
	"time"
)

func TestNewAuthToken(t *testing.T){
	authtoken := NewAuthToken("123",int64(12312), WithExpiredTimeInterval(72))

	token := "123"
	var createTime int64 = 12312
	var expiredTimeInterval int64 = 72
	if authtoken.GetToken() != token{
		t.Error("token not equal")
	}
	if authtoken.GetCreateTime() != createTime{
		t.Error("createTime not equal")
	}
	if authtoken.GetExpiredTimeInterval() != expiredTimeInterval{
		t.Error("expiredTimeInterval not equal")
	}
}

func TestAuthToken_Generate(t *testing.T) {
	url := "liuyang"
	appID := 1
	password := "123"
	timestamp := int64(1313)

	clientAuthToken := NewAuthToken("123",int64(12312), WithExpiredTimeInterval(72))
	serverAuthToken := clientAuthToken.Generate(url,appID,password,timestamp)
	if serverAuthToken.GetToken() != clientAuthToken.generateToken(url, appID,password,timestamp){
		t.Error("token 错误")
	}
}

func TestAuthToken_IsExpired(t *testing.T) {
	AuthToken := NewAuthToken("123",time.Now().Unix(), WithExpiredTimeInterval(2))

	expect := true
	if AuthToken.IsExpired() != expect{
		t.Error("错误")
	}
	time.Sleep(3 * time.Second)
	expect = false
	if AuthToken.IsExpired() != expect{
		t.Error("错误", AuthToken.GetExpiredTimeInterval())
	}

}
