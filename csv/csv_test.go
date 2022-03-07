package csv

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/antrl-go/csv/parser"
	"testing"
)

func TestName(t *testing.T) {
	stream := antlr.NewInputStream(
		`a,b
1,3
1,2
`)
	lexer := parser.NewCSVLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	csvParser := parser.NewCSVParser(tokenStream)
	tree := csvParser.File()
	l := &Load{}
	walker := antlr.NewParseTreeWalker()
	walker.Walk(l, tree)
	fmt.Println(l)
}
