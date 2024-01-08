package observer

import "fmt"

//   a ------------>[b,c,d,e]
// 被观察者          观察者
//  subject         observer
// publisher        subscriber
// producer         consumer

type Observer interface {
	update(msg string)
}

type Subject interface {
	registerObserver(o Observer)
	removeObserver(o Observer)
	notifyObserver(msg string)
}

type Teacher struct {
	students []Observer
}

func (t *Teacher)registerObserver(o Observer){
	t.students = append(t.students, o)
}

func(t *Teacher)removeObserver(o Observer){
	for index,val := range t.students{
		if val == o{
			t.students = append(t.students[:index],t.students[index+1:]...)
			break
		}
	}
}

func(t *Teacher)notifyObserver(msg string){
	for _,o := range t.students{
		o.update(msg)
	}
}

type Student struct {
	name string
}
func(s *Student)update(msg string){
	fmt.Println("s.name finish ",msg)
}


