package baseInterface

import (
	"fmt"
	"testing"
)

type MockStorage struct {

}

func (ms *MockStorage)Add(name, password string)bool{
	fmt.Println("添加", name , password)
	return true
}
func (ms *MockStorage)Find(name, password string)bool{
	fmt.Println("查找", name , password)
	return true
}


func TestProxyUserController_Login(t *testing.T) {
	uc := &UserController{
		storage: new(MockStorage),
	}

	puc := ProxyUserController{
		uc: uc,
	}
	puc.Login("liuyang", "123456")
	puc.Register("liuyang", "123")
}
