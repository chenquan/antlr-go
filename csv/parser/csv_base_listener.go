// Code generated from D:/chenquan/golang/antlr-go/csv\CSV.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // CSV

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCSVListener is a complete listener for a parse tree produced by CSVParser.
type BaseCSVListener struct{}

var _ CSVListener = &BaseCSVListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCSVListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCSVListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCSVListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCSVListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseCSVListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseCSVListener) ExitFile(ctx *FileContext) {}

// EnterHdr is called when production hdr is entered.
func (s *BaseCSVListener) EnterHdr(ctx *HdrContext) {}

// ExitHdr is called when production hdr is exited.
func (s *BaseCSVListener) ExitHdr(ctx *HdrContext) {}

// EnterRow is called when production row is entered.
func (s *BaseCSVListener) EnterRow(ctx *RowContext) {}

// ExitRow is called when production row is exited.
func (s *BaseCSVListener) ExitRow(ctx *RowContext) {}

// EnterText is called when production text is entered.
func (s *BaseCSVListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseCSVListener) ExitText(ctx *TextContext) {}

// EnterString is called when production string is entered.
func (s *BaseCSVListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseCSVListener) ExitString(ctx *StringContext) {}

// EnterEmpty is called when production empty is entered.
func (s *BaseCSVListener) EnterEmpty(ctx *EmptyContext) {}

// ExitEmpty is called when production empty is exited.
func (s *BaseCSVListener) ExitEmpty(ctx *EmptyContext) {}
