// Code generated from D:/chenquan/golang/antlr-go/csv\CSV.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // CSV

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 37, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 6, 2, 13, 10,
	2, 13, 2, 14, 2, 14, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 22, 10, 4, 12,
	4, 14, 4, 25, 11, 4, 3, 4, 5, 4, 28, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 5, 5, 35, 10, 5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 2, 2, 37, 2, 10, 3, 2,
	2, 2, 4, 16, 3, 2, 2, 2, 6, 18, 3, 2, 2, 2, 8, 34, 3, 2, 2, 2, 10, 12,
	5, 4, 3, 2, 11, 13, 5, 6, 4, 2, 12, 11, 3, 2, 2, 2, 13, 14, 3, 2, 2, 2,
	14, 12, 3, 2, 2, 2, 14, 15, 3, 2, 2, 2, 15, 3, 3, 2, 2, 2, 16, 17, 5, 6,
	4, 2, 17, 5, 3, 2, 2, 2, 18, 23, 5, 8, 5, 2, 19, 20, 7, 3, 2, 2, 20, 22,
	5, 8, 5, 2, 21, 19, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2,
	23, 24, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 26, 28, 7,
	4, 2, 2, 27, 26, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29,
	30, 7, 5, 2, 2, 30, 7, 3, 2, 2, 2, 31, 35, 7, 6, 2, 2, 32, 35, 7, 7, 2,
	2, 33, 35, 3, 2, 2, 2, 34, 31, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 33,
	3, 2, 2, 2, 35, 9, 3, 2, 2, 2, 6, 14, 23, 27, 34,
}
var literalNames = []string{
	"", "','", "'\r'", "'\n'",
}
var symbolicNames = []string{
	"", "", "", "", "TEXT", "STRING",
}

var ruleNames = []string{
	"file", "hdr", "row", "field",
}

type CSVParser struct {
	*antlr.BaseParser
}

// NewCSVParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *CSVParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCSVParser(input antlr.TokenStream) *CSVParser {
	this := new(CSVParser)
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
	this.GrammarFileName = "CSV.g4"

	return this
}

// CSVParser tokens.
const (
	CSVParserEOF    = antlr.TokenEOF
	CSVParserT__0   = 1
	CSVParserT__1   = 2
	CSVParserT__2   = 3
	CSVParserTEXT   = 4
	CSVParserSTRING = 5
)

// CSVParser rules.
const (
	CSVParserRULE_file  = 0
	CSVParserRULE_hdr   = 1
	CSVParserRULE_row   = 2
	CSVParserRULE_field = 3
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CSVParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CSVParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) Hdr() IHdrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHdrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHdrContext)
}

func (s *FileContext) AllRow() []IRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRowContext)(nil)).Elem())
	var tst = make([]IRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRowContext)
		}
	}

	return tst
}

func (s *FileContext) Row(i int) IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitFile(s)
	}
}

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CSVVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CSVParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CSVParserRULE_file)
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
	{
		p.SetState(8)
		p.Hdr()
	}
	p.SetState(10)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CSVParserT__0)|(1<<CSVParserT__1)|(1<<CSVParserT__2)|(1<<CSVParserTEXT)|(1<<CSVParserSTRING))) != 0) {
		{
			p.SetState(9)
			p.Row()
		}

		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHdrContext is an interface to support dynamic dispatch.
type IHdrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHdrContext differentiates from other interfaces.
	IsHdrContext()
}

type HdrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHdrContext() *HdrContext {
	var p = new(HdrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CSVParserRULE_hdr
	return p
}

func (*HdrContext) IsHdrContext() {}

func NewHdrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HdrContext {
	var p = new(HdrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CSVParserRULE_hdr

	return p
}

func (s *HdrContext) GetParser() antlr.Parser { return s.parser }

func (s *HdrContext) Row() IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
}

func (s *HdrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HdrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HdrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterHdr(s)
	}
}

func (s *HdrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitHdr(s)
	}
}

func (s *HdrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CSVVisitor:
		return t.VisitHdr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CSVParser) Hdr() (localctx IHdrContext) {
	localctx = NewHdrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CSVParserRULE_hdr)

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
		p.SetState(14)
		p.Row()
	}

	return localctx
}

// IRowContext is an interface to support dynamic dispatch.
type IRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRowContext differentiates from other interfaces.
	IsRowContext()
}

type RowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowContext() *RowContext {
	var p = new(RowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CSVParserRULE_row
	return p
}

func (*RowContext) IsRowContext() {}

func NewRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowContext {
	var p = new(RowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CSVParserRULE_row

	return p
}

func (s *RowContext) GetParser() antlr.Parser { return s.parser }

func (s *RowContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *RowContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *RowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterRow(s)
	}
}

func (s *RowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitRow(s)
	}
}

func (s *RowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CSVVisitor:
		return t.VisitRow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CSVParser) Row() (localctx IRowContext) {
	localctx = NewRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CSVParserRULE_row)
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
	{
		p.SetState(16)
		p.Field()
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CSVParserT__0 {
		{
			p.SetState(17)
			p.Match(CSVParserT__0)
		}
		{
			p.SetState(18)
			p.Field()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CSVParserT__1 {
		{
			p.SetState(24)
			p.Match(CSVParserT__1)
		}

	}
	{
		p.SetState(27)
		p.Match(CSVParserT__2)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CSVParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CSVParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) CopyFrom(ctx *FieldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringContext struct {
	*FieldContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(CSVParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CSVVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type TextContext struct {
	*FieldContext
}

func NewTextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TextContext {
	var p = new(TextContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) TEXT() antlr.TerminalNode {
	return s.GetToken(CSVParserTEXT, 0)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitText(s)
	}
}

func (s *TextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CSVVisitor:
		return t.VisitText(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmptyContext struct {
	*FieldContext
}

func NewEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyContext {
	var p = new(EmptyContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *EmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterEmpty(s)
	}
}

func (s *EmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitEmpty(s)
	}
}

func (s *EmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CSVVisitor:
		return t.VisitEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CSVParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CSVParserRULE_field)

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

	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CSVParserTEXT:
		localctx = NewTextContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Match(CSVParserTEXT)
		}

	case CSVParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Match(CSVParserSTRING)
		}

	case CSVParserT__0, CSVParserT__1, CSVParserT__2:
		localctx = NewEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
