package Car

//将一个物品从多个维度分解，利用依赖注入的方式，构建出一个物品
type Car struct {
	color IColor
	hub IHub
}

func (c Car)GetConfig()string{
	return c.color.GetColor() + c.hub.GetHub()
}

type IColor interface {
	GetColor()string
}

type IHub interface {
	GetHub()string
}

type Red struct {
}

func (r *Red)GetColor()string{
	return "red"
}

type Blue struct {
}

func (r *Blue)GetColor()string{
	return "blue"
}


type HuaWei struct {
}

func (huawei *HuaWei)GetHub()string{
	return "Huawei"
}


type Xiaomi struct {
}

func (x *Xiaomi)GetHub()string{
	return "Xiaomi"
}

