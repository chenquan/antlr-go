// Code generated from D:/chenquan/golang/antlr-go/csv\CSV.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // CSV

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCSVVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCSVVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCSVVisitor) VisitHdr(ctx *HdrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCSVVisitor) VisitRow(ctx *RowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCSVVisitor) VisitText(ctx *TextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCSVVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCSVVisitor) VisitEmpty(ctx *EmptyContext) interface{} {
	return v.VisitChildren(ctx)
}
