package json

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/antrl-go/json/parser"
	"testing"
)

func TestJson(t *testing.T) {
	stream := antlr.NewInputStream(
		`[1,{"a":[3,4]}]`)
	lexer := parser.NewJSONLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	jsonParser := parser.NewJSONParser(tokenStream)
	jsonTree := jsonParser.Json()
	walker := antlr.NewParseTreeWalker()
	x := &XMLEmitter{xml: map[antlr.Tree]string{}}
	walker.Walk(x, jsonTree)
	fmt.Println(x.ToXML(jsonTree))
}
