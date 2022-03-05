// Code generated from D:/chenquan/golang/antrl-go/calculate\LabeledExpr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // LabeledExpr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLabeledExprListener is a complete listener for a parse tree produced by LabeledExprParser.
type BaseLabeledExprListener struct{}

var _ LabeledExprListener = &BaseLabeledExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLabeledExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLabeledExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLabeledExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLabeledExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseLabeledExprListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseLabeledExprListener) ExitProg(ctx *ProgContext) {}

// EnterPrintExpr is called when production printExpr is entered.
func (s *BaseLabeledExprListener) EnterPrintExpr(ctx *PrintExprContext) {}

// ExitPrintExpr is called when production printExpr is exited.
func (s *BaseLabeledExprListener) ExitPrintExpr(ctx *PrintExprContext) {}

// EnterAssgin is called when production assgin is entered.
func (s *BaseLabeledExprListener) EnterAssgin(ctx *AssginContext) {}

// ExitAssgin is called when production assgin is exited.
func (s *BaseLabeledExprListener) ExitAssgin(ctx *AssginContext) {}

// EnterBlank is called when production blank is entered.
func (s *BaseLabeledExprListener) EnterBlank(ctx *BlankContext) {}

// ExitBlank is called when production blank is exited.
func (s *BaseLabeledExprListener) ExitBlank(ctx *BlankContext) {}

// EnterParens is called when production parens is entered.
func (s *BaseLabeledExprListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production parens is exited.
func (s *BaseLabeledExprListener) ExitParens(ctx *ParensContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseLabeledExprListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseLabeledExprListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseLabeledExprListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseLabeledExprListener) ExitAddSub(ctx *AddSubContext) {}

// EnterId is called when production id is entered.
func (s *BaseLabeledExprListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseLabeledExprListener) ExitId(ctx *IdContext) {}

// EnterInt is called when production int is entered.
func (s *BaseLabeledExprListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BaseLabeledExprListener) ExitInt(ctx *IntContext) {}
