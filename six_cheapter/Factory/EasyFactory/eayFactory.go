package EasyFactory

import "fmt"

type Parser interface {
	parse(string)
}

type ParserFactory struct {
}

func (pf *ParserFactory)CreateParser(rule string) Parser {
	var parser Parser
	switch rule {
	case "json":
		parser = new(JsonParser)
	case "xml":
		parser = new(XmlParser)
	case "yaml":
		parser = new(YamlParser)
	default:
		parser = new(WrongParser)
	}
	return parser

}


type JsonParser struct {

}
func (jp *JsonParser)parse(text string){
	fmt.Println("json解析结果")
}


type XmlParser struct {

}
func (xp *XmlParser)parse(text string){
	fmt.Println("xml解析结果")
}

type WrongParser struct {

}
func (wp *WrongParser)parse(text string){
	fmt.Println("json解析结果")
}
type YamlParser struct {

}
func (yp *YamlParser )parse(text string){
	fmt.Println("Yaml解析结果")
}