package Draw

//添加原始模型的增强的相关功能，例如以前画图，我现在变成画带颜色的图
type Draw interface {
	draw()string
}

type Square struct {

}
func (s Square)draw()string{
	return "square"
}

type ColorSquare struct {
	s Square
	color string
}

func (cs ColorSquare)draw()string{
	return cs.color + cs.s.draw()
}
