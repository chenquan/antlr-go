// Code generated from D:/chenquan/golang/antlr-go/json\JSON.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // JSON

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JSONListener is a complete listener for a parse tree produced by JSONParser.
type JSONListener interface {
	antlr.ParseTreeListener

	// EnterJson is called when entering the json production.
	EnterJson(c *JsonContext)

	// EnterAnObject is called when entering the AnObject production.
	EnterAnObject(c *AnObjectContext)

	// EnterAEmptyObject is called when entering the AEmptyObject production.
	EnterAEmptyObject(c *AEmptyObjectContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterAnArray is called when entering the AnArray production.
	EnterAnArray(c *AnArrayContext)

	// EnterAEmptyArray is called when entering the AEmptyArray production.
	EnterAEmptyArray(c *AEmptyArrayContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterAtom is called when entering the Atom production.
	EnterAtom(c *AtomContext)

	// EnterObjectValue is called when entering the ObjectValue production.
	EnterObjectValue(c *ObjectValueContext)

	// EnterArrayValue is called when entering the ArrayValue production.
	EnterArrayValue(c *ArrayValueContext)

	// ExitJson is called when exiting the json production.
	ExitJson(c *JsonContext)

	// ExitAnObject is called when exiting the AnObject production.
	ExitAnObject(c *AnObjectContext)

	// ExitAEmptyObject is called when exiting the AEmptyObject production.
	ExitAEmptyObject(c *AEmptyObjectContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitAnArray is called when exiting the AnArray production.
	ExitAnArray(c *AnArrayContext)

	// ExitAEmptyArray is called when exiting the AEmptyArray production.
	ExitAEmptyArray(c *AEmptyArrayContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitAtom is called when exiting the Atom production.
	ExitAtom(c *AtomContext)

	// ExitObjectValue is called when exiting the ObjectValue production.
	ExitObjectValue(c *ObjectValueContext)

	// ExitArrayValue is called when exiting the ArrayValue production.
	ExitArrayValue(c *ArrayValueContext)
}
