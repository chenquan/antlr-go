// Code generated from D:/chenquan/golang/antlr-go/csv\CSV.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // CSV

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CSVParser.
type CSVVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CSVParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by CSVParser#hdr.
	VisitHdr(ctx *HdrContext) interface{}

	// Visit a parse tree produced by CSVParser#row.
	VisitRow(ctx *RowContext) interface{}

	// Visit a parse tree produced by CSVParser#text.
	VisitText(ctx *TextContext) interface{}

	// Visit a parse tree produced by CSVParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by CSVParser#empty.
	VisitEmpty(ctx *EmptyContext) interface{}
}
