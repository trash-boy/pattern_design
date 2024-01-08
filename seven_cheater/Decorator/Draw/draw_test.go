package Draw

import "testing"

func TestColorDraw(t *testing.T){
	s := Square{}
	cs := ColorSquare{
		s: s,
		color: "red",
	}

	if cs.draw() != "redsquare"{
		t.Error("draw 方法错误")
	}
}
