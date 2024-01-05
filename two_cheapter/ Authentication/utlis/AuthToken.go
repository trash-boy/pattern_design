package utlis

import (
	"fmt"
	"time"
)

var EXPIREDTIMEINTERVAL int64 = 60 * 1000
type AuthToken struct {
	token string
	createTime int64
	expiredTimeInterval int64
}

func WithExpiredTimeInterval(expiredTimeInterval int64)func(*AuthToken){
	return func(token *AuthToken) {
		token.expiredTimeInterval = expiredTimeInterval
	}
}


func NewAuthToken(token string, createTime int64, opts ...func(*AuthToken))*AuthToken{
	authToken := AuthToken{
		token: token,
		createTime: createTime,
		expiredTimeInterval: EXPIREDTIMEINTERVAL,
	}
	for _,opt := range opts{
		opt(&authToken)
	}
	return &authToken
}

func (auth *AuthToken)GetToken()string{
	return auth.token
}
func (auth *AuthToken)GetCreateTime()int64{
	return auth.createTime
}
func (auth *AuthToken)GetExpiredTimeInterval()int64{
	return auth.expiredTimeInterval
}

func (auth *AuthToken)generateToken(url string, appId int, password string, timeStamp int64)string{
	content := fmt.Sprintf("%s+%d+%s+%d",url,appId,password,timeStamp)
	return Sha256(content)
}

func (auth *AuthToken)Generate(url string, appId int, password string, timeStamp int64)*AuthToken{
	//生成token,组成成为authToken类型
	token := auth.generateToken(url, appId,password,timeStamp)
	return NewAuthToken(token, timeStamp)
}

func (auth *AuthToken)IsExpired()bool{
	timeStamp := auth.GetCreateTime()
	nowStamp := time.Now().Unix()
	fmt.Println(timeStamp,nowStamp)
	if nowStamp - timeStamp <= auth.GetExpiredTimeInterval(){
		return true
	}
	return false
}

func (auth *AuthToken)Match(serverAuth *AuthToken)bool{
	return auth.GetToken() == serverAuth.GetToken()
}
