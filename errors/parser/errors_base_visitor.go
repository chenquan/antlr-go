// Code generated from D:/chenquan/golang/antlr-go/errors\Errors.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Errors

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseErrorsVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseErrorsVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseErrorsVisitor) VisitClassDef(ctx *ClassDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseErrorsVisitor) VisitMember(ctx *MemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseErrorsVisitor) VisitStat(ctx *StatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseErrorsVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}
