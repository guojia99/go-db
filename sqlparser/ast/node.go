/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package ast

import (
	"fmt"

	"github.com/guojia99/godb/sqlparser/tokenizer"
)

type Node interface {
	node()
	fmt.Stringer
}

func (x *ColumnType) node() {}

type Column struct {
	Name        tokenizer.Token
	Type        ColumnType
	Constraints []Constraint
}

type ColumnType struct {
	Name      tokenizer.SqlToken
	Precision *NumberExpr
	Scale     *NumberExpr
}

func (x *ColumnType) String() string {
	if x.Precision != nil && x.Scale != nil {
		return fmt.Sprintf("%s(%s,%s)", x.Name, x.Precision, x.Scale)
	}
	if x.Precision != nil {
		return fmt.Sprintf("%s(%s)", x.Name, x.Precision)
	}
	return x.Name.String()
}
