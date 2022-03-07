package json

import (
	"bytes"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chenquan/antrl-go/json/parser"
	"strings"
)

type XMLEmitter struct {
	parser.BaseJSONListener
	xml   map[antlr.Tree]string
	depth int
}

// ExitJson is called when production json is exited.
func (s *XMLEmitter) ExitJson(ctx *parser.JsonContext) {
	s.setXML(ctx, s.getXML(ctx.GetChild(0)))
}

// EnterAnObject is called when production AnObject is entered.
func (s *XMLEmitter) EnterAnObject(ctx *parser.AnObjectContext) {}

// ExitAnObject is called when production AnObject is exited.
func (s *XMLEmitter) ExitAnObject(ctx *parser.AnObjectContext) {
	buffer := &bytes.Buffer{}
	buffer.WriteByte('\n')
	for _, pairContext := range ctx.AllPair() {
		buffer.WriteString(s.getXML(pairContext))
	}
	s.setXML(ctx, buffer.String())
}

// EnterAEmptyObject is called when production EmptyObject is entered.
func (s *XMLEmitter) EnterAEmptyObject(ctx *parser.AEmptyObjectContext) {}

// ExitAEmptyObject is called when production EmptyObject is exited.
func (s *XMLEmitter) ExitAEmptyObject(ctx *parser.AEmptyObjectContext) {
	s.setXML(ctx, "")
}

// EnterPair is called when production pair is entered.
func (s *XMLEmitter) EnterPair(ctx *parser.PairContext) {}

// ExitPair is called when production pair is exited.
func (s *XMLEmitter) ExitPair(ctx *parser.PairContext) {
	tag := stripQuotes(ctx.STRING().GetText())
	s.setXML(ctx, fmt.Sprintf("<%s>%s</%s>\n", tag, s.getXML(ctx.Value()), tag))
}

// EnterAnArray is called when production AnArray is entered.
func (s *XMLEmitter) EnterAnArray(ctx *parser.AnArrayContext) {}

// ExitAnArray is called when production AnArray is exited.
func (s *XMLEmitter) ExitAnArray(ctx *parser.AnArrayContext) {
	buffer := &bytes.Buffer{}
	buffer.WriteByte('\n')
	for _, valueContext := range ctx.AllValue() {
		buffer.WriteString("<element>")
		buffer.WriteString(s.getXML(valueContext))
		buffer.WriteString("</element>")
		buffer.WriteByte('\n')
	}
	s.setXML(ctx, buffer.String())
}

// EnterAEmptyArray is called when production EmptyArray is entered.
func (s *XMLEmitter) EnterAEmptyArray(ctx *parser.AEmptyArrayContext) {}

// ExitAEmptyArray is called when production EmptyArray is exited.
func (s *XMLEmitter) ExitAEmptyArray(ctx *parser.AEmptyArrayContext) {
	s.setXML(ctx, "")
}

// EnterString is called when production String is entered.
func (s *XMLEmitter) EnterString(ctx *parser.StringContext) {}

// ExitString is called when production String is exited.
func (s *XMLEmitter) ExitString(ctx *parser.StringContext) {
	s.setXML(ctx, stripQuotes(ctx.GetText()))
}

// EnterAtom is called when production Atom is entered.
func (s *XMLEmitter) EnterAtom(ctx *parser.AtomContext) {}

// ExitAtom is called when production Atom is exited.
func (s *XMLEmitter) ExitAtom(ctx *parser.AtomContext) {
	s.setXML(ctx, ctx.GetText())
}

// EnterObjectValue is called when production ObjectValue is entered.
func (s *XMLEmitter) EnterObjectValue(ctx *parser.ObjectValueContext) {}

// ExitObjectValue is called when production ObjectValue is exited.
func (s *XMLEmitter) ExitObjectValue(ctx *parser.ObjectValueContext) {
	s.setXML(ctx, s.getXML(ctx.Object()))
}

// EnterArrayValue is called when production ArrayValue is entered.
func (s *XMLEmitter) EnterArrayValue(ctx *parser.ArrayValueContext) {}

// ExitArrayValue is called when production ArrayValue is exited.
func (s *XMLEmitter) ExitArrayValue(ctx *parser.ArrayValueContext) {
	s.setXML(ctx, s.getXML(ctx.Array()))
}

func (s *XMLEmitter) getXML(ctx antlr.Tree) string {
	if text, ok := s.xml[ctx]; ok {
		return text
	}

	return ""
}

func (s *XMLEmitter) setXML(ctx antlr.Tree, text string) {
	s.xml[ctx] = text
}

func stripQuotes(text string) string {
	return strings.Trim(text, `"`)
}

func (s *XMLEmitter) ToXML(ctx antlr.Tree) string {
	if text, ok := s.xml[ctx]; ok {
		return text
	}

	return ""
}
