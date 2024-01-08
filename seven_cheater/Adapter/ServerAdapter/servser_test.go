package ServerAdapter

import "testing"

func TestAdapter(t *testing.T){
	aa := AliyunAdapter{aliyun: new(Aliyun)}
	ba := BaiduYunAdapter{baiduYun: new(BaiduYun)}

	aa.CreateServer(100,100)
	ba.CreateServer(80,90)
}
