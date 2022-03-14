// Code generated from D:/chenquan/golang/antlr-go/errors\Errors.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Errors

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ErrorsParser.
type ErrorsVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ErrorsParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by ErrorsParser#classDef.
	VisitClassDef(ctx *ClassDefContext) interface{}

	// Visit a parse tree produced by ErrorsParser#member.
	VisitMember(ctx *MemberContext) interface{}

	// Visit a parse tree produced by ErrorsParser#stat.
	VisitStat(ctx *StatContext) interface{}

	// Visit a parse tree produced by ErrorsParser#expr.
	VisitExpr(ctx *ExprContext) interface{}
}
