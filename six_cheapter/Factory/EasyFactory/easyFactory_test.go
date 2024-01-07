package EasyFactory

import "testing"

func TestParserFactory_CreateParser(t *testing.T) {
	fp := new(ParserFactory)

	xp :=  fp.CreateParser("xml")
	xp.parse("123")
}
