/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package ast

import (
	tk "github.com/guojia99/godb/sqlparser/tokenizer"
)

// https://www.w3schools.com/sql/sql_default.asp

type Constraint interface {
	Node
	constraint()
}

func (x *PrimaryKeyConstraint) node() {}
func (x *NotNullConstraint) node()    {}
func (x *UniqueConstraint) node()     {}
func (x *ForeignKeyConstraint) node() {}
func (x *CheckConstraint) node()      {}

func (x *PrimaryKeyConstraint) constraint() {}
func (x *NotNullConstraint) constraint()    {}
func (x *UniqueConstraint) constraint()     {}
func (x *ForeignKeyConstraint) constraint() {}
func (x *CheckConstraint) constraint()      {}

func getColumnsString(columns []tk.SqlToken) string {
	var out string
	if len(columns) > 0 {
		out += " ("
		for i := range columns {
			if i != 0 {
				out += ","
			}
			out += columns[i].String()
		}
		out += ")"
	}
	return out
}

// PrimaryKeyConstraint the complete form is as follows:
// constraint pk_pk primary key (id, o_id)
type PrimaryKeyConstraint struct {
	ConstraintName tk.SqlToken
	PrimaryKey     tk.SqlToken
	Columns        []tk.SqlToken
	Autoincrement  tk.SqlToken
}

func (x *PrimaryKeyConstraint) String() string {
	var out string
	if !x.ConstraintName.IsEmpty() {
		out += "constraint " + x.ConstraintName.String() + " "
	}
	if !x.PrimaryKey.IsEmpty() {
		out += "primary key"
	}
	out += getColumnsString(x.Columns)
	if !x.Autoincrement.IsEmpty() {
		out += " auto_increment"
	}
	return out
}

type NotNullConstraint struct {
	NotNull tk.SqlToken
}

func (x *NotNullConstraint) String() string {
	if x.NotNull.IsEmpty() {
		return ""
	}
	return "not null"
}

type UniqueConstraint struct {
	ConstraintName tk.SqlToken
	Unique         tk.SqlToken
	Columns        []tk.SqlToken
}

func (x *UniqueConstraint) String() string {
	var out string
	if !x.ConstraintName.IsEmpty() {
		out += "constraint " + x.ConstraintName.String() + " "
	}
	if !x.Unique.IsEmpty() {
		out += "unique"
	}
	out += getColumnsString(x.Columns)
	return out
}

type ForeignKeyConstraint struct {
	ConstraintName    tk.SqlToken
	ForeignKey        tk.SqlToken
	ForeignKeyColumns []tk.SqlToken
	References        tk.SqlToken
	ReferencesTable   tk.SqlToken
	ReferencesColumns []tk.SqlToken

	// args
	OnUpdate   tk.SqlToken
	OnDelete   tk.SqlToken
	SetNull    tk.SqlToken
	SetDefault tk.SqlToken
	Cascade    tk.SqlToken
	Restrict   tk.SqlToken
	NoAction   tk.SqlToken
}

func (x *ForeignKeyConstraint) String() string {
	var out string
	if !x.ConstraintName.IsEmpty() {
		out += "constraint " + x.ConstraintName.String() + " "
	}
	if !x.ForeignKey.IsEmpty() {
		out += "foreign key" + getColumnsString(x.ForeignKeyColumns)
	}
	if !x.References.IsEmpty() && !x.ReferencesTable.IsEmpty() {
		out += " references " + x.ReferencesTable.String() + getColumnsString(x.ReferencesColumns)
	}
	if !x.OnUpdate.IsEmpty() {
		out += " on update"
	}
	if !x.OnDelete.IsEmpty() {
		out += " on delete"
	}
	if !x.SetNull.IsEmpty() {
		out += " set null"
	}
	if !x.SetDefault.IsEmpty() {
		out += " set default"
	}
	if !x.Cascade.IsEmpty() {
		out += " cascade"
	}
	if !x.Restrict.IsEmpty() {
		out += " restrict"
	}
	if !x.NoAction.IsEmpty() {
		out += " no action"
	}
	return out
}

type CheckConstraint struct {
	ConstraintName tk.SqlToken
	CheckExpr      Expr
}

func (x *CheckConstraint) String() string {
	var out string
	if !x.ConstraintName.IsEmpty() {
		out += "constraint " + x.ConstraintName.String() + " "
	}
	out += "check (" + x.CheckExpr.String() + ")"
	return out
}

type DefaultConstraint struct {
	ConstraintName tk.SqlToken
	DefaultExpr    Expr
}

func (x *DefaultConstraint) String() string {
	var out string
	if !x.ConstraintName.IsEmpty() {
		out += "constraint " + x.ConstraintName.String() + " "
	}
	out += "default " + x.DefaultExpr.String()
	return out
}
