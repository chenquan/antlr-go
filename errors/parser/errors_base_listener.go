// Code generated from D:/chenquan/golang/antlr-go/errors\Errors.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Errors

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseErrorsListener is a complete listener for a parse tree produced by ErrorsParser.
type BaseErrorsListener struct{}

var _ ErrorsListener = &BaseErrorsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseErrorsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseErrorsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseErrorsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseErrorsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseErrorsListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseErrorsListener) ExitProg(ctx *ProgContext) {}

// EnterClassDef is called when production classDef is entered.
func (s *BaseErrorsListener) EnterClassDef(ctx *ClassDefContext) {}

// ExitClassDef is called when production classDef is exited.
func (s *BaseErrorsListener) ExitClassDef(ctx *ClassDefContext) {}

// EnterMember is called when production member is entered.
func (s *BaseErrorsListener) EnterMember(ctx *MemberContext) {}

// ExitMember is called when production member is exited.
func (s *BaseErrorsListener) ExitMember(ctx *MemberContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseErrorsListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseErrorsListener) ExitStat(ctx *StatContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseErrorsListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseErrorsListener) ExitExpr(ctx *ExprContext) {}
