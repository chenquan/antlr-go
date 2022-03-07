// Code generated from D:/chenquan/golang/antlr-go/json\JSON.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // JSON

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseJSONVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJSONVisitor) VisitJson(ctx *JsonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitAnObject(ctx *AnObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitAEmptyObject(ctx *AEmptyObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitAnArray(ctx *AnArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitAEmptyArray(ctx *AEmptyArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitObjectValue(ctx *ObjectValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitArrayValue(ctx *ArrayValueContext) interface{} {
	return v.VisitChildren(ctx)
}
