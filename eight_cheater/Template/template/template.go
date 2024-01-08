package template

import "fmt"

//回调函数的定义
//A的F函数注册到B类中，然后A类在调用B类方法p的时候运行F函数，F函数就叫做回调函数
//将B类和CallBack固定，那么A类就是可以扩展的
type CallBack interface {
	callback()
}

type A struct {
	b *B
}

type B struct {
}

func (b *B)process(cb CallBack){
	cb.callback()
}

func (a *A)callback(){
	fmt.Println("这是a函数")
}

type Game interface {
	Initialize()
	StartGame()
	EndPlay()
}
type BaseGame struct {

}

func (b *BaseGame)Initialize(){
	fmt.Println("base init")
}
func (b *BaseGame)StartGame(){
	fmt.Println("base start game")
}
func (b *BaseGame)EndPlay(){
	fmt.Println("base end play")
}
//func (b *BaseGame)Play(){
//	b.Initialize()
//	b.Play()
//	b.EndPlay()
//}

type Cricket struct {

}

func (b *Cricket)Initialize(){
	fmt.Println("Cricket init")
}
func (b *Cricket)StartGame(){
	fmt.Println("Cricket start game")
}
func (b *Cricket)EndPlay(){
	fmt.Println("Cricket end play")
}
//func (b *Cricket)Play(){
//	b.Initialize()
//	b.Play()
//	b.EndPlay()
//}



func Play(g Game){
	g.Initialize()
	g.StartGame()
	g.EndPlay()
}







