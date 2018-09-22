// parser project parser.go
package parser

import (
	"webreader"
)

type Parser struct {
	Options *webreader.RequestOptions
}

var ParserObj = new(Parser)

func GetParser() *Parser {
	if ParserObj.Options == nil {
		ParserObj.Options = webreader.GetOptions()
	}
	return ParserObj
}
