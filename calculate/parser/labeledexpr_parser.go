// Code generated from D:/chenquan/golang/antlr-go/calculate\LabeledExpr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // LabeledExpr

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 45, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 6, 2, 10, 10, 2, 13, 2, 14, 2, 11,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 23, 10, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 32, 10, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 40, 10, 4, 12, 4, 14, 4, 43, 11, 4, 3, 4,
	2, 3, 6, 5, 2, 4, 6, 2, 4, 3, 2, 10, 11, 3, 2, 12, 13, 2, 48, 2, 9, 3,
	2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 31, 3, 2, 2, 2, 8, 10, 5, 4, 3, 2, 9, 8,
	3, 2, 2, 2, 10, 11, 3, 2, 2, 2, 11, 9, 3, 2, 2, 2, 11, 12, 3, 2, 2, 2,
	12, 3, 3, 2, 2, 2, 13, 14, 5, 6, 4, 2, 14, 15, 7, 8, 2, 2, 15, 23, 3, 2,
	2, 2, 16, 17, 7, 6, 2, 2, 17, 18, 7, 3, 2, 2, 18, 19, 5, 6, 4, 2, 19, 20,
	7, 8, 2, 2, 20, 23, 3, 2, 2, 2, 21, 23, 7, 8, 2, 2, 22, 13, 3, 2, 2, 2,
	22, 16, 3, 2, 2, 2, 22, 21, 3, 2, 2, 2, 23, 5, 3, 2, 2, 2, 24, 25, 8, 4,
	1, 2, 25, 32, 7, 7, 2, 2, 26, 32, 7, 6, 2, 2, 27, 28, 7, 4, 2, 2, 28, 29,
	5, 6, 4, 2, 29, 30, 7, 5, 2, 2, 30, 32, 3, 2, 2, 2, 31, 24, 3, 2, 2, 2,
	31, 26, 3, 2, 2, 2, 31, 27, 3, 2, 2, 2, 32, 41, 3, 2, 2, 2, 33, 34, 12,
	7, 2, 2, 34, 35, 9, 2, 2, 2, 35, 40, 5, 6, 4, 8, 36, 37, 12, 6, 2, 2, 37,
	38, 9, 3, 2, 2, 38, 40, 5, 6, 4, 7, 39, 33, 3, 2, 2, 2, 39, 36, 3, 2, 2,
	2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 7, 3,
	2, 2, 2, 43, 41, 3, 2, 2, 2, 7, 11, 22, 31, 39, 41,
}
var literalNames = []string{
	"", "'='", "'('", "')'", "", "", "", "", "'*'", "'/'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "ID", "NUMBER", "NEWLINE", "WS", "MUL", "DIV", "ADD", "SUB",
}

var ruleNames = []string{
	"prog", "stat", "expr",
}

type LabeledExprParser struct {
	*antlr.BaseParser
}

// NewLabeledExprParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *LabeledExprParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLabeledExprParser(input antlr.TokenStream) *LabeledExprParser {
	this := new(LabeledExprParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "LabeledExpr.g4"

	return this
}

// LabeledExprParser tokens.
const (
	LabeledExprParserEOF     = antlr.TokenEOF
	LabeledExprParserT__0    = 1
	LabeledExprParserT__1    = 2
	LabeledExprParserT__2    = 3
	LabeledExprParserID      = 4
	LabeledExprParserNUMBER  = 5
	LabeledExprParserNEWLINE = 6
	LabeledExprParserWS      = 7
	LabeledExprParserMUL     = 8
	LabeledExprParserDIV     = 9
	LabeledExprParserADD     = 10
	LabeledExprParserSUB     = 11
)

// LabeledExprParser rules.
const (
	LabeledExprParserRULE_prog = 0
	LabeledExprParserRULE_stat = 1
	LabeledExprParserRULE_expr = 2
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LabeledExprParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LabeledExprParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *ProgContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LabeledExprVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LabeledExprParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LabeledExprParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LabeledExprParserT__1)|(1<<LabeledExprParserID)|(1<<LabeledExprParserNUMBER)|(1<<LabeledExprParserNEWLINE))) != 0) {
		{
			p.SetState(6)
			p.Stat()
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LabeledExprParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LabeledExprParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyFrom(ctx *StatContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlankContext struct {
	*StatContext
}

func NewBlankContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlankContext {
	var p = new(BlankContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *BlankContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(LabeledExprParserNEWLINE, 0)
}

func (s *BlankContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.EnterBlank(s)
	}
}

func (s *BlankContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.ExitBlank(s)
	}
}

func (s *BlankContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LabeledExprVisitor:
		return t.VisitBlank(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssginContext struct {
	*StatContext
}

func NewAssginContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssginContext {
	var p = new(AssginContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *AssginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssginContext) ID() antlr.TerminalNode {
	return s.GetToken(LabeledExprParserID, 0)
}

func (s *AssginContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssginContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(LabeledExprParserNEWLINE, 0)
}

func (s *AssginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.EnterAssgin(s)
	}
}

func (s *AssginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.ExitAssgin(s)
	}
}

func (s *AssginContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LabeledExprVisitor:
		return t.VisitAssgin(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintExprContext struct {
	*StatContext
}

func NewPrintExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintExprContext {
	var p = new(PrintExprContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *PrintExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintExprContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(LabeledExprParserNEWLINE, 0)
}

func (s *PrintExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.EnterPrintExpr(s)
	}
}

func (s *PrintExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.ExitPrintExpr(s)
	}
}

func (s *PrintExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LabeledExprVisitor:
		return t.VisitPrintExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LabeledExprParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LabeledExprParserRULE_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPrintExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(11)
			p.expr(0)
		}
		{
			p.SetState(12)
			p.Match(LabeledExprParserNEWLINE)
		}

	case 2:
		localctx = NewAssginContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(14)
			p.Match(LabeledExprParserID)
		}
		{
			p.SetState(15)
			p.Match(LabeledExprParserT__0)
		}
		{
			p.SetState(16)
			p.expr(0)
		}
		{
			p.SetState(17)
			p.Match(LabeledExprParserNEWLINE)
		}

	case 3:
		localctx = NewBlankContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(19)
			p.Match(LabeledExprParserNEWLINE)
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LabeledExprParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LabeledExprParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumberContext struct {
	*ExprContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(LabeledExprParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LabeledExprVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensContext struct {
	*ExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LabeledExprVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivContext struct {
	*ExprContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MulDivContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(LabeledExprParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(LabeledExprParserDIV, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LabeledExprVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	*ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(LabeledExprParserSUB, 0)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(LabeledExprParserADD, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LabeledExprVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	*ExprContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(LabeledExprParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LabeledExprListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LabeledExprVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LabeledExprParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *LabeledExprParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, LabeledExprParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LabeledExprParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(23)
			p.Match(LabeledExprParserNUMBER)
		}

	case LabeledExprParserID:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(24)
			p.Match(LabeledExprParserID)
		}

	case LabeledExprParserT__1:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(25)
			p.Match(LabeledExprParserT__1)
		}
		{
			p.SetState(26)
			p.expr(0)
		}
		{
			p.SetState(27)
			p.Match(LabeledExprParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(37)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LabeledExprParserRULE_expr)
				p.SetState(31)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(32)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LabeledExprParserMUL || _la == LabeledExprParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(33)
					p.expr(6)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LabeledExprParserRULE_expr)
				p.SetState(34)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(35)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LabeledExprParserADD || _la == LabeledExprParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(36)
					p.expr(5)
				}

			}

		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

func (p *LabeledExprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LabeledExprParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
