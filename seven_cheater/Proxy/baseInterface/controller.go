package baseInterface

import (
	"fmt"
	"time"
)
//基于接口实现，并且添加 与原始类不相关的功能
type IUserController interface {
	Login(Name string, Password string)bool
	Register(Name string, Password string)bool
}

type IStorage interface {
	Add(string, string)bool
	Find(string , string)bool
}

type UserController struct {
	storage IStorage
}

func (UserController *UserController)Register(name string, password string)bool{
	time.Sleep(time.Second)
	return UserController.storage.Add(name,password)
}

func (UserController *UserController)Login(name string, password string)bool{
	time.Sleep(time.Second * 3)
	return UserController.storage.Find(name, password)
}

type ProxyUserController struct {
	uc *UserController
}

func (puc *ProxyUserController)Login(name string, password string)bool{
	startTime := time.Now()
	ans :=  puc.uc.Login(name,password)
	fmt.Println("login cost time ", time.Now().Sub(startTime))
	return ans
}

func(puc *ProxyUserController)Register(name string, password string)bool{
	startTime := time.Now()
	ans :=  puc.uc.Register(name,password)
	fmt.Println("Register cost time ", time.Now().Sub(startTime))
	return ans

}