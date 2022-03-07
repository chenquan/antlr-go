// Code generated from D:/chenquan/golang/antlr-go/json\JSON.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // JSON

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by JSONParser.
type JSONVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JSONParser#json.
	VisitJson(ctx *JsonContext) interface{}

	// Visit a parse tree produced by JSONParser#AnObject.
	VisitAnObject(ctx *AnObjectContext) interface{}

	// Visit a parse tree produced by JSONParser#AEmptyObject.
	VisitAEmptyObject(ctx *AEmptyObjectContext) interface{}

	// Visit a parse tree produced by JSONParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by JSONParser#AnArray.
	VisitAnArray(ctx *AnArrayContext) interface{}

	// Visit a parse tree produced by JSONParser#AEmptyArray.
	VisitAEmptyArray(ctx *AEmptyArrayContext) interface{}

	// Visit a parse tree produced by JSONParser#String.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by JSONParser#Atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by JSONParser#ObjectValue.
	VisitObjectValue(ctx *ObjectValueContext) interface{}

	// Visit a parse tree produced by JSONParser#ArrayValue.
	VisitArrayValue(ctx *ArrayValueContext) interface{}
}
