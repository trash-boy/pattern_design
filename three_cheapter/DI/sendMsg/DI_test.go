package sendMsg

import "testing"

func TestNotification_Notify(t *testing.T) {
	smsSender := new(SmsSender)

	notifycation := NewNotification(smsSender)
	if notifycation.Notify("liuyang", "1524844412454") != nil{
		t.Error("发送错误")
	}

}
