package calculate

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/antrl-go/calculate/parser"
	"strconv"
)

type (
	parseTreeVisit struct {
		parser.BaseLabeledExprVisitor
		memory map[string]interface{}
	}
)

func (v *parseTreeVisit) VisitProg(ctx *parser.ProgContext) interface{} {
	children := ctx.GetChildren()
	var val interface{}
	for _, child := range children {
		if tree, ok := child.(*antlr.TerminalNodeImpl); ok {
			if tree.GetSymbol().GetTokenType() == parser.LabeledExprLexerNEWLINE {
				continue
			}
		} else if _, ok := child.(*parser.BlankContext); ok {
			continue
		}

		val = v.Visit(child.(antlr.ParseTree))
	}
	return val
}
func (v *parseTreeVisit) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}
func (v *parseTreeVisit) VisitPrintExpr(ctx *parser.PrintExprContext) interface{} {
	value := v.Visit(ctx.Expr())
	return value
}

func (v *parseTreeVisit) VisitAssgin(ctx *parser.AssginContext) interface{} {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	v.memory[id] = value.(float64)
	return value
}

func (v *parseTreeVisit) VisitParens(ctx *parser.ParensContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *parseTreeVisit) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))

	if ctx.GetOp().GetTokenType() == parser.LabeledExprParserMUL {
		return left.(float64) * right.(float64)
	}
	return left.(float64) / right.(float64)
}

func (v *parseTreeVisit) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))

	if ctx.GetOp().GetTokenType() == parser.LabeledExprParserSUB {
		return left.(float64) - right.(float64)
	}
	return left.(float64) + right.(float64)
}

func (v *parseTreeVisit) VisitId(ctx *parser.IdContext) interface{} {
	id := ctx.ID().GetText()
	if val, ok := v.memory[id]; ok {
		return val
	}
	return 0
}

func (v *parseTreeVisit) VisitNumber(ctx *parser.NumberContext) interface{} {
	val, _ := strconv.ParseFloat(ctx.NUMBER().GetText(), 64)
	return val
}

var visit = &parseTreeVisit{}

// Calculate returns the computed value.
func Calculate(expr string) string {
	stream := antlr.NewInputStream(expr + "\n")
	lexer := parser.NewLabeledExprLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	exprParser := parser.NewLabeledExprParser(tokenStream)
	tree := exprParser.Prog()
	visit.memory = map[string]interface{}{}
	return fmt.Sprintf("%v", visit.Visit(tree))
}
