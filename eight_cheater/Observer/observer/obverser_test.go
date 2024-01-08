package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObserver(t *testing.T){
	teacher := Teacher{
		students: make([]Observer,0),
	}

	stu1 := &Student{
		"Huawei",
	}
	stu2 := &Student{
		"Xiaomi",
	}
	teacher.registerObserver(stu1)
	teacher.registerObserver(stu2)

	teacher.notifyObserver("我从来不卖车")

	teacher.removeObserver(stu1)
	assert.Equal(t, len(teacher.students), 1)
}
