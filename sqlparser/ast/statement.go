/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package ast

import "github.com/guojia99/godb/sqlparser/tokenizer"

type Statement interface {
	Node
	stmt()
	Clone() Statement
}

type CreateTableStatement struct {
	Create tokenizer.SqlToken
	Table  tokenizer.SqlToken
	Name   tokenizer.SqlToken

	Columns []Column
}
