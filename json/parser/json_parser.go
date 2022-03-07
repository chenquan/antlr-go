// Code generated from D:/chenquan/golang/antlr-go/json\JSON.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // JSON

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 60, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 5,
	2, 15, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 21, 10, 3, 12, 3, 14, 3, 24,
	11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 40, 10, 5, 12, 5, 14, 5, 43, 11, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 5, 5, 49, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 5, 6, 58, 10, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 2, 2, 65, 2, 14,
	3, 2, 2, 2, 4, 29, 3, 2, 2, 2, 6, 31, 3, 2, 2, 2, 8, 48, 3, 2, 2, 2, 10,
	57, 3, 2, 2, 2, 12, 15, 5, 4, 3, 2, 13, 15, 5, 8, 5, 2, 14, 12, 3, 2, 2,
	2, 14, 13, 3, 2, 2, 2, 15, 3, 3, 2, 2, 2, 16, 17, 7, 8, 2, 2, 17, 22, 5,
	6, 4, 2, 18, 19, 7, 12, 2, 2, 19, 21, 5, 6, 4, 2, 20, 18, 3, 2, 2, 2, 21,
	24, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 25, 3, 2, 2,
	2, 24, 22, 3, 2, 2, 2, 25, 26, 7, 9, 2, 2, 26, 30, 3, 2, 2, 2, 27, 28,
	7, 8, 2, 2, 28, 30, 7, 9, 2, 2, 29, 16, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2,
	30, 5, 3, 2, 2, 2, 31, 32, 7, 3, 2, 2, 32, 33, 7, 13, 2, 2, 33, 34, 5,
	10, 6, 2, 34, 7, 3, 2, 2, 2, 35, 36, 7, 10, 2, 2, 36, 41, 5, 10, 6, 2,
	37, 38, 7, 12, 2, 2, 38, 40, 5, 10, 6, 2, 39, 37, 3, 2, 2, 2, 40, 43, 3,
	2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 44, 3, 2, 2, 2, 43,
	41, 3, 2, 2, 2, 44, 45, 7, 11, 2, 2, 45, 49, 3, 2, 2, 2, 46, 47, 7, 10,
	2, 2, 47, 49, 7, 11, 2, 2, 48, 35, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49,
	9, 3, 2, 2, 2, 50, 58, 7, 3, 2, 2, 51, 58, 7, 4, 2, 2, 52, 58, 5, 4, 3,
	2, 53, 58, 5, 8, 5, 2, 54, 58, 7, 5, 2, 2, 55, 58, 7, 6, 2, 2, 56, 58,
	7, 7, 2, 2, 57, 50, 3, 2, 2, 2, 57, 51, 3, 2, 2, 2, 57, 52, 3, 2, 2, 2,
	57, 53, 3, 2, 2, 2, 57, 54, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 56, 3,
	2, 2, 2, 58, 11, 3, 2, 2, 2, 8, 14, 22, 29, 41, 48, 57,
}
var literalNames = []string{
	"", "", "", "'true'", "'false'", "'null'", "'{'", "'}'", "'['", "']'",
	"','", "':'",
}
var symbolicNames = []string{
	"", "STRING", "NUMBER", "TRUE", "FALSE", "NULL", "LEFT_BRACE", "RIGHT_BRACE",
	"OPEN_BRACKET", "CLOSE_BRACKET", "COMMA", "COLON", "WS",
}

var ruleNames = []string{
	"json", "object", "pair", "array", "value",
}

type JSONParser struct {
	*antlr.BaseParser
}

// NewJSONParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *JSONParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewJSONParser(input antlr.TokenStream) *JSONParser {
	this := new(JSONParser)
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
	this.GrammarFileName = "JSON.g4"

	return this
}

// JSONParser tokens.
const (
	JSONParserEOF           = antlr.TokenEOF
	JSONParserSTRING        = 1
	JSONParserNUMBER        = 2
	JSONParserTRUE          = 3
	JSONParserFALSE         = 4
	JSONParserNULL          = 5
	JSONParserLEFT_BRACE    = 6
	JSONParserRIGHT_BRACE   = 7
	JSONParserOPEN_BRACKET  = 8
	JSONParserCLOSE_BRACKET = 9
	JSONParserCOMMA         = 10
	JSONParserCOLON         = 11
	JSONParserWS            = 12
)

// JSONParser rules.
const (
	JSONParserRULE_json   = 0
	JSONParserRULE_object = 1
	JSONParserRULE_pair   = 2
	JSONParserRULE_array  = 3
	JSONParserRULE_value  = 4
)

// IJsonContext is an interface to support dynamic dispatch.
type IJsonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonContext differentiates from other interfaces.
	IsJsonContext()
}

type JsonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonContext() *JsonContext {
	var p = new(JsonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_json
	return p
}

func (*JsonContext) IsJsonContext() {}

func NewJsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContext {
	var p = new(JsonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_json

	return p
}

func (s *JsonContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *JsonContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *JsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterJson(s)
	}
}

func (s *JsonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitJson(s)
	}
}

func (s *JsonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitJson(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Json() (localctx IJsonContext) {
	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JSONParserRULE_json)

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

	p.SetState(12)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JSONParserLEFT_BRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			p.Object()
		}

	case JSONParserOPEN_BRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(11)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) CopyFrom(ctx *ObjectContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AnObjectContext struct {
	*ObjectContext
}

func NewAnObjectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AnObjectContext {
	var p = new(AnObjectContext)

	p.ObjectContext = NewEmptyObjectContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObjectContext))

	return p
}

func (s *AnObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnObjectContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(JSONParserLEFT_BRACE, 0)
}

func (s *AnObjectContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *AnObjectContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *AnObjectContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(JSONParserRIGHT_BRACE, 0)
}

func (s *AnObjectContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JSONParserCOMMA)
}

func (s *AnObjectContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JSONParserCOMMA, i)
}

func (s *AnObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterAnObject(s)
	}
}

func (s *AnObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitAnObject(s)
	}
}

func (s *AnObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitAnObject(s)

	default:
		return t.VisitChildren(s)
	}
}

type AEmptyObjectContext struct {
	*ObjectContext
}

func NewAEmptyObjectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AEmptyObjectContext {
	var p = new(AEmptyObjectContext)

	p.ObjectContext = NewEmptyObjectContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObjectContext))

	return p
}

func (s *AEmptyObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AEmptyObjectContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(JSONParserLEFT_BRACE, 0)
}

func (s *AEmptyObjectContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(JSONParserRIGHT_BRACE, 0)
}

func (s *AEmptyObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterAEmptyObject(s)
	}
}

func (s *AEmptyObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitAEmptyObject(s)
	}
}

func (s *AEmptyObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitAEmptyObject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Object() (localctx IObjectContext) {
	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JSONParserRULE_object)
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

	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAnObjectContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)
			p.Match(JSONParserLEFT_BRACE)
		}
		{
			p.SetState(15)
			p.Pair()
		}
		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JSONParserCOMMA {
			{
				p.SetState(16)
				p.Match(JSONParserCOMMA)
			}
			{
				p.SetState(17)
				p.Pair()
			}

			p.SetState(22)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(23)
			p.Match(JSONParserRIGHT_BRACE)
		}

	case 2:
		localctx = NewAEmptyObjectContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Match(JSONParserLEFT_BRACE)
		}
		{
			p.SetState(26)
			p.Match(JSONParserRIGHT_BRACE)
		}

	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(JSONParserSTRING, 0)
}

func (s *PairContext) COLON() antlr.TerminalNode {
	return s.GetToken(JSONParserCOLON, 0)
}

func (s *PairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitPair(s)
	}
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JSONParserRULE_pair)

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
	{
		p.SetState(29)
		p.Match(JSONParserSTRING)
	}
	{
		p.SetState(30)
		p.Match(JSONParserCOLON)
	}
	{
		p.SetState(31)
		p.Value()
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) CopyFrom(ctx *ArrayContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AEmptyArrayContext struct {
	*ArrayContext
}

func NewAEmptyArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AEmptyArrayContext {
	var p = new(AEmptyArrayContext)

	p.ArrayContext = NewEmptyArrayContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayContext))

	return p
}

func (s *AEmptyArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AEmptyArrayContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(JSONParserOPEN_BRACKET, 0)
}

func (s *AEmptyArrayContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(JSONParserCLOSE_BRACKET, 0)
}

func (s *AEmptyArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterAEmptyArray(s)
	}
}

func (s *AEmptyArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitAEmptyArray(s)
	}
}

func (s *AEmptyArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitAEmptyArray(s)

	default:
		return t.VisitChildren(s)
	}
}

type AnArrayContext struct {
	*ArrayContext
}

func NewAnArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AnArrayContext {
	var p = new(AnArrayContext)

	p.ArrayContext = NewEmptyArrayContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayContext))

	return p
}

func (s *AnArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnArrayContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(JSONParserOPEN_BRACKET, 0)
}

func (s *AnArrayContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *AnArrayContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AnArrayContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(JSONParserCLOSE_BRACKET, 0)
}

func (s *AnArrayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JSONParserCOMMA)
}

func (s *AnArrayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JSONParserCOMMA, i)
}

func (s *AnArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterAnArray(s)
	}
}

func (s *AnArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitAnArray(s)
	}
}

func (s *AnArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitAnArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JSONParserRULE_array)
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

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAnArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Match(JSONParserOPEN_BRACKET)
		}
		{
			p.SetState(34)
			p.Value()
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JSONParserCOMMA {
			{
				p.SetState(35)
				p.Match(JSONParserCOMMA)
			}
			{
				p.SetState(36)
				p.Value()
			}

			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(42)
			p.Match(JSONParserCLOSE_BRACKET)
		}

	case 2:
		localctx = NewAEmptyArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Match(JSONParserOPEN_BRACKET)
		}
		{
			p.SetState(45)
			p.Match(JSONParserCLOSE_BRACKET)
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ObjectValueContext struct {
	*ValueContext
}

func NewObjectValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectValueContext {
	var p = new(ObjectValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ObjectValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectValueContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *ObjectValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterObjectValue(s)
	}
}

func (s *ObjectValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitObjectValue(s)
	}
}

func (s *ObjectValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitObjectValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	*ValueContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(JSONParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayValueContext struct {
	*ValueContext
}

func NewArrayValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayValueContext {
	var p = new(ArrayValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterArrayValue(s)
	}
}

func (s *ArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitArrayValue(s)
	}
}

func (s *ArrayValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitArrayValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtomContext struct {
	*ValueContext
}

func NewAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomContext {
	var p = new(AtomContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(JSONParserNUMBER, 0)
}

func (s *AtomContext) TRUE() antlr.TerminalNode {
	return s.GetToken(JSONParserTRUE, 0)
}

func (s *AtomContext) FALSE() antlr.TerminalNode {
	return s.GetToken(JSONParserFALSE, 0)
}

func (s *AtomContext) NULL() antlr.TerminalNode {
	return s.GetToken(JSONParserNULL, 0)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JSONVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JSONParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JSONParserRULE_value)

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

	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JSONParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Match(JSONParserSTRING)
		}

	case JSONParserNUMBER:
		localctx = NewAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Match(JSONParserNUMBER)
		}

	case JSONParserLEFT_BRACE:
		localctx = NewObjectValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Object()
		}

	case JSONParserOPEN_BRACKET:
		localctx = NewArrayValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(51)
			p.Array()
		}

	case JSONParserTRUE:
		localctx = NewAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(52)
			p.Match(JSONParserTRUE)
		}

	case JSONParserFALSE:
		localctx = NewAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(53)
			p.Match(JSONParserFALSE)
		}

	case JSONParserNULL:
		localctx = NewAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(54)
			p.Match(JSONParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
