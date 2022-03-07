// Code generated from D:/chenquan/golang/antlr-go/csv\CSV.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // CSV

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CSVListener is a complete listener for a parse tree produced by CSVParser.
type CSVListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterHdr is called when entering the hdr production.
	EnterHdr(c *HdrContext)

	// EnterRow is called when entering the row production.
	EnterRow(c *RowContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterEmpty is called when entering the empty production.
	EnterEmpty(c *EmptyContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitHdr is called when exiting the hdr production.
	ExitHdr(c *HdrContext)

	// ExitRow is called when exiting the row production.
	ExitRow(c *RowContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitEmpty is called when exiting the empty production.
	ExitEmpty(c *EmptyContext)
}
