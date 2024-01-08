package Department

//树状结构，就是利用组合模式
type IOrigination interface {
	Count()int
}

type Employee struct {
	name string
}
func (e Employee)Count()int{
	return 1
}

type Department struct {
	name string
	o []IOrigination
}
func (d Department)Count()int{
	ans := 0
	for _,e := range d.o{
		ans += e.Count()
	}
	return ans
}

func (d *Department)Add(o IOrigination){
	d.o = append(d.o, o)
}

func NewOrganization() IOrigination {
	root := &Department{name: "root"}
	for i := 0; i < 10; i++ {
		root.Add(&Employee{})
		root.Add(&Department{name: "sub", o: []IOrigination{&Employee{}}})
	}
	return root
}
