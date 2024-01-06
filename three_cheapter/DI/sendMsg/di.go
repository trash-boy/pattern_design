package sendMsg

import (
	"fmt"
)

type MessageSender interface {
	send(phone,msg string)error
}

type Notification struct {
	messageSender MessageSender
}

func NewNotification(messageSender MessageSender)*Notification{
	return &Notification{
		messageSender: messageSender,
	}
}
func (notify *Notification)Notify(msg,phone string)error{
	return notify.messageSender.send(phone,msg)
}


type SmsSender struct {

}

func (sms *SmsSender)send(phone,msg string)error{
	fmt.Println("sms", phone, msg)
	return nil
}

type InboxSender struct {

}
func(inbox *InboxSender)send(phone,msg string)error{
	fmt.Println("inbox ",phone, msg)
	return nil
}




