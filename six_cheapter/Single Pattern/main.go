package main

import (
	"fmt"
	"sync"
	"time"
)

type Singleton struct {
	Name string
}

var (
	instance *Singleton
	lock sync.Mutex
)

func GetInstance()*Singleton{
	//一些线程读一些线程写 数据竞争
	if instance == nil{
		lock.Lock()
		defer lock.Unlock()
		if instance == nil{
			instance = &Singleton{"liuyang"}
		}
	}
	return instance
}

func main(){
	for i := 0;i < 10;i++{
		go func() {
			instance = GetInstance()
			fmt.Println(instance.Name)
		}()
	}

	time.Sleep(time.Second)
}


