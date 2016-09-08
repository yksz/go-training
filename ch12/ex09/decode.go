// Package sexpr provides a means for converting Go objects to and
// from S-expressions.
package sexpr

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

type Token interface{}

type Symbol string
type String string
type Int int
type StartList string
type EndList string

type Decoder struct {
	lex *lexer
}

func NewDecoder(r io.Reader) *Decoder {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(r)
	return &Decoder{lex}
}

func (dec *Decoder) Token() (tok Token, err error) {
	dec.lex.next() // get the first token
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			if x == io.EOF {
				err = io.EOF
			} else {
				err = fmt.Errorf("error at %s: %v", dec.lex.scan.Position, x)
			}
		}
	}()
	tok = read(dec.lex)
	return
}

type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

func read(lex *lexer) interface{} {
	switch lex.token {
	case scanner.EOF:
		panic(io.EOF)
	case scanner.Ident:
		return Symbol(lex.text())
	case scanner.String:
		return String(lex.text())
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
		return Int(i)
	case '(':
		return StartList(lex.text())
	case ')':
		return EndList(lex.text())
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}
