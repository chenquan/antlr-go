package errors

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/antrl-go/errors/parser"
)

type VerboseListener struct {
	parser.BaseErrorsListener
	antlr.ConsoleErrorListener
}

func (l VerboseListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.ConsoleErrorListener.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
	underLineError(recognizer, offendingSymbol.(antlr.Token), line, column)
}

func underLineError(recognizer antlr.Recognizer, offendingToken antlr.Token, line, column int) {
	inputStream := recognizer.(*antlr.BaseParser).GetTokenStream().(*antlr.CommonTokenStream).GetTokenSource().GetInputStream()
	input := inputStream.(*antlr.InputStream).String()
	lines := strings.Split(input, "\n")
	errorLine := lines[line-1]
	fmt.Println(errorLine)
	for i := 0; i < column; i++ {
		fmt.Printf(" ")
	}
	start := offendingToken.GetStart()
	stop := offendingToken.GetStop()
	if start >= 0 && stop >= 0 {
		for i := start; i <= stop; i++ {
			fmt.Printf("^")
		}
	}
	fmt.Println()
}
