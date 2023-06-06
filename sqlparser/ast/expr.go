/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package ast

import (
	"github.com/guojia99/godb/sqlparser/keywords"
	tk "github.com/guojia99/godb/sqlparser/tokenizer"
)

// Expr expression.
type Expr interface {
	Node
	expr()
}

func (x *NumberExpr) node() {}
func (x *StringExpr) node() {}

func (x *NumberExpr) expr() {}
func (x *StringExpr) expr() {}

type BasicExpr struct {
	L Expr
	E tk.SqlToken
	R Expr
}

func (x *BasicExpr) String() string {
	switch x.E.Value {
	// punc
	case tk.PuncMinusSign, tk.PuncPlusSign, tk.PuncEqual, tk.PuncUnEqual, tk.PuncLess, tk.PuncEqualLess,
		tk.PuncGreater, tk.PuncEqualGreater, tk.PuncBar, tk.PuncForwardSlash, tk.PuncPercent,
		tk.PuncAmpersand, tk.PuncDoubleBar, tk.PuncDoubleLess, tk.PuncDoubleGreater,
		tk.PuncAsterisk:
		return x.L.String() + " " + x.E.Value + " " + x.R.String()
	}

	// key word
	switch x.E.KeyWord {
	case keywords.KwBetween, keywords.KwNotnull: // todo
		return x.L.String() + " " + x.E.Value + " " + x.R.String()
	}
	return ""
}

type NumberExpr struct {
	Value tk.SqlToken
}

func (x *NumberExpr) String() string { return x.Value.String() }

type StringExpr struct {
	Value tk.SqlToken
}

func (x *StringExpr) String() string { return x.Value.String() }
