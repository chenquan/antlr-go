// Code generated from D:/chenquan/golang/antlr-go/calculate\LabeledExpr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 68, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 6, 5, 33, 10, 5, 13, 5, 14, 5, 34,
	3, 6, 5, 6, 38, 10, 6, 3, 6, 3, 6, 3, 6, 6, 6, 43, 10, 6, 13, 6, 14, 6,
	44, 5, 6, 47, 10, 6, 3, 7, 5, 7, 50, 10, 7, 3, 7, 3, 7, 3, 8, 6, 8, 55,
	10, 8, 13, 8, 14, 8, 56, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 2, 2, 13, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	17, 10, 19, 11, 21, 12, 23, 13, 3, 2, 6, 4, 2, 67, 92, 99, 124, 3, 2, 47,
	47, 3, 2, 50, 59, 3, 2, 11, 11, 2, 73, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 25, 3, 2, 2, 2, 5, 27, 3, 2, 2, 2, 7, 29,
	3, 2, 2, 2, 9, 32, 3, 2, 2, 2, 11, 37, 3, 2, 2, 2, 13, 49, 3, 2, 2, 2,
	15, 54, 3, 2, 2, 2, 17, 60, 3, 2, 2, 2, 19, 62, 3, 2, 2, 2, 21, 64, 3,
	2, 2, 2, 23, 66, 3, 2, 2, 2, 25, 26, 7, 63, 2, 2, 26, 4, 3, 2, 2, 2, 27,
	28, 7, 42, 2, 2, 28, 6, 3, 2, 2, 2, 29, 30, 7, 43, 2, 2, 30, 8, 3, 2, 2,
	2, 31, 33, 9, 2, 2, 2, 32, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 32,
	3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 10, 3, 2, 2, 2, 36, 38, 9, 3, 2, 2,
	37, 36, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 46, 9,
	4, 2, 2, 40, 42, 7, 48, 2, 2, 41, 43, 9, 4, 2, 2, 42, 41, 3, 2, 2, 2, 43,
	44, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 47, 3, 2, 2,
	2, 46, 40, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 12, 3, 2, 2, 2, 48, 50,
	7, 15, 2, 2, 49, 48, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2,
	51, 52, 7, 12, 2, 2, 52, 14, 3, 2, 2, 2, 53, 55, 9, 5, 2, 2, 54, 53, 3,
	2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57,
	58, 3, 2, 2, 2, 58, 59, 8, 8, 2, 2, 59, 16, 3, 2, 2, 2, 60, 61, 7, 44,
	2, 2, 61, 18, 3, 2, 2, 2, 62, 63, 7, 49, 2, 2, 63, 20, 3, 2, 2, 2, 64,
	65, 7, 45, 2, 2, 65, 22, 3, 2, 2, 2, 66, 67, 7, 47, 2, 2, 67, 24, 3, 2,
	2, 2, 9, 2, 34, 37, 44, 46, 49, 56, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'('", "')'", "", "", "", "", "'*'", "'/'", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "ID", "NUMBER", "NEWLINE", "WS", "MUL", "DIV", "ADD", "SUB",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "ID", "NUMBER", "NEWLINE", "WS", "MUL", "DIV",
	"ADD", "SUB",
}

type LabeledExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewLabeledExprLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *LabeledExprLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLabeledExprLexer(input antlr.CharStream) *LabeledExprLexer {
	l := new(LabeledExprLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "LabeledExpr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LabeledExprLexer tokens.
const (
	LabeledExprLexerT__0    = 1
	LabeledExprLexerT__1    = 2
	LabeledExprLexerT__2    = 3
	LabeledExprLexerID      = 4
	LabeledExprLexerNUMBER  = 5
	LabeledExprLexerNEWLINE = 6
	LabeledExprLexerWS      = 7
	LabeledExprLexerMUL     = 8
	LabeledExprLexerDIV     = 9
	LabeledExprLexerADD     = 10
	LabeledExprLexerSUB     = 11
)
