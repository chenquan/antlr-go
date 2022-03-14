// Code generated from D:/chenquan/golang/antlr-go/errors\Errors.g4 by ANTLR 4.9.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 61, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 6, 10,
	49, 10, 10, 13, 10, 14, 10, 50, 3, 11, 6, 11, 54, 10, 11, 13, 11, 14, 11,
	55, 3, 12, 3, 12, 3, 12, 3, 12, 2, 2, 13, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 3, 2, 5, 4, 2, 67, 92, 99,
	124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 62, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 25, 3, 2, 2, 2, 5, 31, 3,
	2, 2, 2, 7, 33, 3, 2, 2, 2, 9, 35, 3, 2, 2, 2, 11, 39, 3, 2, 2, 2, 13,
	41, 3, 2, 2, 2, 15, 43, 3, 2, 2, 2, 17, 45, 3, 2, 2, 2, 19, 48, 3, 2, 2,
	2, 21, 53, 3, 2, 2, 2, 23, 57, 3, 2, 2, 2, 25, 26, 7, 101, 2, 2, 26, 27,
	7, 110, 2, 2, 27, 28, 7, 99, 2, 2, 28, 29, 7, 117, 2, 2, 29, 30, 7, 117,
	2, 2, 30, 4, 3, 2, 2, 2, 31, 32, 7, 125, 2, 2, 32, 6, 3, 2, 2, 2, 33, 34,
	7, 127, 2, 2, 34, 8, 3, 2, 2, 2, 35, 36, 7, 107, 2, 2, 36, 37, 7, 112,
	2, 2, 37, 38, 7, 118, 2, 2, 38, 10, 3, 2, 2, 2, 39, 40, 7, 61, 2, 2, 40,
	12, 3, 2, 2, 2, 41, 42, 7, 42, 2, 2, 42, 14, 3, 2, 2, 2, 43, 44, 7, 43,
	2, 2, 44, 16, 3, 2, 2, 2, 45, 46, 7, 63, 2, 2, 46, 18, 3, 2, 2, 2, 47,
	49, 9, 2, 2, 2, 48, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 48, 3, 2, 2,
	2, 50, 51, 3, 2, 2, 2, 51, 20, 3, 2, 2, 2, 52, 54, 9, 3, 2, 2, 53, 52,
	3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2,
	56, 22, 3, 2, 2, 2, 57, 58, 9, 4, 2, 2, 58, 59, 3, 2, 2, 2, 59, 60, 8,
	12, 2, 2, 60, 24, 3, 2, 2, 2, 5, 2, 50, 55, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'class'", "'{'", "'}'", "'int'", "';'", "'('", "')'", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "ID", "INT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "ID", "INT",
	"WS",
}

type ErrorsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewErrorsLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ErrorsLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewErrorsLexer(input antlr.CharStream) *ErrorsLexer {
	l := new(ErrorsLexer)
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
	l.GrammarFileName = "Errors.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ErrorsLexer tokens.
const (
	ErrorsLexerT__0 = 1
	ErrorsLexerT__1 = 2
	ErrorsLexerT__2 = 3
	ErrorsLexerT__3 = 4
	ErrorsLexerT__4 = 5
	ErrorsLexerT__5 = 6
	ErrorsLexerT__6 = 7
	ErrorsLexerT__7 = 8
	ErrorsLexerID   = 9
	ErrorsLexerINT  = 10
	ErrorsLexerWS   = 11
)
