// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Generic expression parser/evaluator
// See https://github.com/golang/talks/blob/master/2010/io/talk.pdf

type Value interface {
	String() string
	BinaryOp(op string, y Value) Value
}




// Command-line expression evaluator

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		//fmt.Printf("%s\n", Eval(precTab, trace(newVal), line))
		fmt.Printf("%s\n", Eval(precTab, newVal, line))
	}
}

// Eval expression in argument string using an operator precedence map and a func
// to turn strings into Values.
// Use the newVal for expression evaluation.
// Use the trace for expression evaluation with tracing.
func Eval(precTab map[string]int, newVal func(string) Value, src string) Value {
	var p Parser
	p.precTab = precTab
	p.newVal = newVal
	p.src = src
	p.next()
	return p.binaryExpr(1) // start with lowest precedence
}

// Custom grammar and values

// PrecTab is a map were key is an operator like "!=" and value is its precedence.
var precTab = map[string]int{
	"&&": 1,
	"||": 2,
	"==": 3,
	"!=": 3,
	"<":  3,
	"<=": 3,
	">":  3,
	">=": 3,
	"+":  4,
	"-":  4,
	"*":  5,
	"/":  5,
	"%":  5,
}

// NewValue returns a Value most appropriate for argument string.
func newVal(lit string) Value {
	x, err := strconv.Atoi(lit)
	if err == nil {
		return Int(x)
	}
	b, err := strconv.ParseBool(lit)
	if err == nil {
		return Bool(b)
	}
	s, err := strconv.Unquote(lit)
	if err == nil {
		return String(s)
	}
	return Error(fmt.Sprintf("illegal literal '%s'", lit))
}

type Error string

func (e Error) BinaryOp(op string, y Value) Value {
	return e
}

func (e Error) String() string {
	return string(e)
}

type Int int

func (x Int) BinaryOp(op string, y Value) Value {
	switch y := y.(type) {
	case Error:
		return y
	case String:
		switch op {
		case "*":
			return String(strings.Repeat(string(y), int(x)))
		}
	case Int:
		switch op {
		case "+":
			return x + y
		case "-":
			return x - y
		case "*":
			return x * y
		case "/":
			return x / y
		case "%":
			return x % y
		case "==":
			return Bool(x == y)
		case "!=":
			return Bool(x != y)
		case "<":
			return Bool(x < y)
		case "<=":
			return Bool(x <= y)
		case ">":
			return Bool(x > y)
		case ">=":
			return Bool(x >= y)
		}
	}
	return Error(fmt.Sprintf("illegal operation: '%v %s %v'", x, op, y))
}

func (x Int) String() string {
	return strconv.Itoa(int(x))
}

type Bool bool

func (x Bool) BinaryOp(op string, y Value) Value {
	switch y := y.(type) {
	case Error:
		return y
	case Bool:
		switch op {
		case "&&":
			return Bool(x && y)
		case "||":
			return Bool(x || y)
		case "==":
			return Bool(x == y)
		case "!=":
			return Bool(x != y)
		}
	}
	return Error(fmt.Sprintf("illegal operation: '%v %s %v'", x, op, y))
}

func (x Bool) String() string {
	return strconv.FormatBool(bool(x))
}


type String string

func (x String) BinaryOp(op string, y Value) Value {
	switch y := y.(type) {
	case Error:
		return y
	case Int:
		switch op {
		case "*":
			return String(strings.Repeat(string(x), int(y)))
		}
	case String:
		switch op {
		case "+":
			return x + y
		case "<":
			return Bool(x < y)
		}
	}
	return Error(fmt.Sprintf("illegal operation: '%v %s %v'", x, op, y))
}

func (x String) String() string {
	return strconv.Quote(string(x))
}


// Trace wraps the argument function and prints tracing info.
func trace(newVal func(string) Value) func(string) Value {
	return func(s string) Value {
		v := newVal(s)
		fmt.Printf("\tnewVal(%q) = %s\n", s, fmtv(v))
		return &traceValue{v}
	}
}

type traceValue struct {
	Value
}

func (x *traceValue) BinaryOp(op string, y Value) Value {
	z := x.Value.BinaryOp(op, y.(*traceValue).Value)
	fmt.Printf("\t%s.BinaryOp(%q, %s) = %s\n", fmtv(x.Value), op, fmtv(y.(*traceValue).Value), fmtv(z))
	return &traceValue{z}
}

func (x *traceValue) String() string {
	s := x.Value.String()
	fmt.Printf("\t%s.String() = %#v\n", fmtv(x.Value), s)
	return s
}

// Fmtv returns a string with argument v type and value.
func fmtv(v Value) string {
	t := fmt.Sprintf("%T", v)
	if i := strings.LastIndex(t, "."); i >= 0 { // strip package
		t = t[i+1:]
	}
	return fmt.Sprintf("%s(%#v)", t, v)
}

/*** Parser *******************************************************************/

type Parser struct {
	precTab map[string]int
	newVal  func(string) Value
	src     string
	pos     int
	tok     string
}

const alphanum = "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (p *Parser) stop(c uint8) bool {
	switch {
	case p.pos >= len(p.src):
		return true
	case c == '"':
		if p.src[p.pos] == '"' {
			p.pos++
			return true
		}
		return false
	case strings.IndexRune(alphanum, rune(c)) >= 0:
		return strings.IndexRune(alphanum, rune(p.src[p.pos])) < 0
	}
	return true
}

func (p *Parser) next() {
	// skip blanks
	for ; p.pos < len(p.src) && p.src[p.pos] <= ' '; p.pos++ {
	}
	if p.pos >= len(p.src) {
		p.tok = ""
		return
	}
	start := p.pos
	c := p.src[p.pos]
	for p.pos < len(p.src) {
		p.pos++
		if p.stop(c) {
			break
		}
	}
	p.tok = p.src[start:p.pos]
}

// BinaryExpr evaluates binary expression of equal or higher precedence then arg.
// On entry p.tok is a value and p.next() tok is an operator.
// If the operator precedence >= then what is on the stack then add it to the stack else unwind stack.
// see http://www.engr.mun.ca/~theo/Misc/exp_parsing.htm#climbing
func (p *Parser) binaryExpr(prec1 int) Value {
	fmt.Println("prec1", prec1)

	x := p.newVal(p.tok)
	p.next()
	for prec := p.precTab[p.tok]; prec >= prec1; prec-- {
		fmt.Println("prec", prec, "tok", p.tok)
		for p.precTab[p.tok] == prec {
			op := p.tok
			p.next()
			y := p.binaryExpr(prec + 1)
			fmt.Println(x, op, y)
			x = x.BinaryOp(op, y)
		}
	}
	return x
}
