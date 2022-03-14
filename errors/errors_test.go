package errors

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/antrl-go/errors/parser"
)

func TestName(t *testing.T) {
	stream := antlr.NewInputStream(`
class T XX{ int i }
`)
	// 词法分析
	lexer := parser.NewErrorsLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	errorsParser := parser.NewErrorsParser(tokenStream)
	l := &VerboseListener{}
	errorsParser.RemoveErrorListeners()
	errorsParser.AddErrorListener(l)
	errorsParser.BuildParseTrees = true
	progTree := errorsParser.Prog()
	walker := antlr.NewParseTreeWalker()
	walker.Walk(l, progTree)
}
