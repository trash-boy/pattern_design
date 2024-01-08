package ServerAdapter

import "fmt"

//提供一个同义的接口，让所有不同的实现，实现为同样的接口
type ICreateServer interface {
	CreateServer(mem, cpu int)error
}

type Aliyun struct {
}

func (aliyun *Aliyun)RunServer(mem, cpu int)error{
	fmt.Println("阿里云服务器启动",mem,cpu)
	return nil
}

type BaiduYun struct {
}

func (baiduyun *BaiduYun)RunInstance(mem, cpu int)error{
	fmt.Println("百度云服务器启动",mem, cpu)
	return nil
}

type AliyunAdapter struct {
	aliyun *Aliyun
}
func(ap *AliyunAdapter)CreateServer(mem, cpu int)error{
	return ap.aliyun.RunServer(mem,cpu)
}

type BaiduYunAdapter struct {
	baiduYun *BaiduYun
}
func(ba *BaiduYunAdapter)CreateServer(mem,cpu int)error{
	return ba.baiduYun.RunInstance(mem,cpu)
}
