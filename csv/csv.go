package csv

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/antrl-go/csv/parser"
)

const Empty = ""

type Load struct {
	parser.BaseCSVListener
	rows                  [][]string
	header                []string
	currentRowFieldValues []string
}

// VisitTerminal is called when a terminal node is visited.
func (s *Load) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Load) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Load) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Load) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *Load) EnterFile(ctx *parser.FileContext) {}

// ExitFile is called when production file is exited.
func (s *Load) ExitFile(ctx *parser.FileContext) {}

// EnterHdr is called when production hdr is entered.
func (s *Load) EnterHdr(ctx *parser.HdrContext) {}

// ExitHdr is called when production hdr is exited.
func (s *Load) ExitHdr(ctx *parser.HdrContext) {
	s.header = append(s.header, s.currentRowFieldValues...)
}

// EnterRow is called when production row is entered.
func (s *Load) EnterRow(ctx *parser.RowContext) {
	s.currentRowFieldValues = make([]string, 0, len(s.header))
}

// ExitRow is called when production row is exited.
func (s *Load) ExitRow(ctx *parser.RowContext) {
	if hdrContext, ok := ctx.GetParent().(*parser.HdrContext); ok && hdrContext.GetRuleIndex() == parser.CSVParserRULE_hdr {
		fmt.Println(hdrContext.GetRuleIndex())
		return
	}
	//if ctx.GetRuleIndex() == parser.CSVParserRULE_hdr {
	//	return
	//}
	s.rows = append(s.rows, s.currentRowFieldValues)
}

// EnterText is called when production text is entered.
func (s *Load) EnterText(ctx *parser.TextContext) {}

// ExitText is called when production text is exited.
func (s *Load) ExitText(ctx *parser.TextContext) {
	s.currentRowFieldValues = append(s.currentRowFieldValues, ctx.TEXT().GetText())
}

// EnterString is called when production string is entered.
func (s *Load) EnterString(ctx *parser.StringContext) {}

// ExitString is called when production string is exited.
func (s *Load) ExitString(ctx *parser.StringContext) {
	s.currentRowFieldValues = append(s.currentRowFieldValues, ctx.STRING().GetText())
}

// EnterEmpty is called when production empty is entered.
func (s *Load) EnterEmpty(ctx *parser.EmptyContext) {}

// ExitEmpty is called when production empty is exited.
func (s *Load) ExitEmpty(ctx *parser.EmptyContext) {
	s.currentRowFieldValues = append(s.currentRowFieldValues, Empty)
}
