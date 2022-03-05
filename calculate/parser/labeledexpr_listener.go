// Code generated from D:/chenquan/golang/antlr-go/calculate\LabeledExpr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // LabeledExpr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LabeledExprListener is a complete listener for a parse tree produced by LabeledExprParser.
type LabeledExprListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterPrintExpr is called when entering the printExpr production.
	EnterPrintExpr(c *PrintExprContext)

	// EnterAssgin is called when entering the assgin production.
	EnterAssgin(c *AssginContext)

	// EnterBlank is called when entering the blank production.
	EnterBlank(c *BlankContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterParens is called when entering the parens production.
	EnterParens(c *ParensContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitPrintExpr is called when exiting the printExpr production.
	ExitPrintExpr(c *PrintExprContext)

	// ExitAssgin is called when exiting the assgin production.
	ExitAssgin(c *AssginContext)

	// ExitBlank is called when exiting the blank production.
	ExitBlank(c *BlankContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitParens is called when exiting the parens production.
	ExitParens(c *ParensContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)
}
