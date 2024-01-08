package Car

import "testing"

func TestBridge(t *testing.T){
	car1 := Car{
		color: new(Red),
		hub: new(HuaWei),
	}
	if car1.GetConfig() != "redHuawei"{
		t.Error(car1.GetConfig(),"car1, bridge错误")
	}
	car2 := Car{
		color: new(Blue),
		hub: new(Xiaomi),
	}
	if car2.GetConfig() != "blueXiaomi"{
		t.Error("car2, bridge错误")
	}
}
