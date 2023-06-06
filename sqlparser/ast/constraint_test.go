/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package ast

import (
	"reflect"
	"testing"

	"github.com/guojia99/godb/sqlparser/tokenizer"
)

func checkString(t *testing.T, name string, want string, node Node) {
	t.Run(name, func(t *testing.T) {

	})
}

func TestPrimaryKeyConstraint_String(t *testing.T) {
	// create table tb_name(
	//      o_id int,
	//    	id int primary key,
	//      id int auto_increment,
	//    	primary key (id),
	//    	primary key (id, o_id),
	//    	constraint pk_pk primary key (id, o_id)
	//	)

	tests := []struct {
		name string
		want string
		node Node
	}{
		{
			name: "default",
			want: "primary key",
			node: &PrimaryKeyConstraint{
				ConstraintName: tokenizer.SqlToken{},
				PrimaryKey:     tokenizer.SqlToken{Value: "primary key"},
				Columns:        []tokenizer.SqlToken{},
				Autoincrement:  tokenizer.SqlToken{},
			},
		},
		{
			name: "has_columns",
			want: "primary key (id)",
			node: &PrimaryKeyConstraint{
				ConstraintName: tokenizer.SqlToken{},
				PrimaryKey:     tokenizer.SqlToken{Value: "primary key"},
				Columns: []tokenizer.SqlToken{
					{Value: "id"},
				},
				Autoincrement: tokenizer.SqlToken{},
			},
		},
		{
			name: "has_columns_many",
			want: "primary key (id,o_id)",
			node: &PrimaryKeyConstraint{
				ConstraintName: tokenizer.SqlToken{},
				PrimaryKey:     tokenizer.SqlToken{Value: "primary key"},
				Columns: []tokenizer.SqlToken{
					{Value: "id"},
					{Value: "o_id"},
				},
				Autoincrement: tokenizer.SqlToken{},
			},
		},
		{
			name: "has_constraint",
			want: "constraint tk primary key",
			node: &PrimaryKeyConstraint{
				ConstraintName: tokenizer.SqlToken{Value: "tk"},
				PrimaryKey:     tokenizer.SqlToken{Value: "primary key"},
				Columns:        []tokenizer.SqlToken{},
				Autoincrement:  tokenizer.SqlToken{},
			},
		},
		{
			name: "has_constraint_columns_many",
			want: "constraint tk primary key (id,o_id)",
			node: &PrimaryKeyConstraint{
				ConstraintName: tokenizer.SqlToken{Value: "tk"},
				PrimaryKey:     tokenizer.SqlToken{Value: "primary key"},
				Columns: []tokenizer.SqlToken{
					{Value: "id"},
					{Value: "o_id"},
				},
				Autoincrement: tokenizer.SqlToken{},
			},
		}, {
			name: "auto_increment",
			want: "constraint tk primary key (id,o_id) auto_increment",
			node: &PrimaryKeyConstraint{
				ConstraintName: tokenizer.SqlToken{Value: "tk"},
				PrimaryKey:     tokenizer.SqlToken{Value: "primary key"},
				Columns: []tokenizer.SqlToken{
					{Value: "id"},
					{Value: "o_id"},
				},
				Autoincrement: tokenizer.SqlToken{Value: "auto_increment"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := tt.node.String()
			if str != tt.want {
				t.Errorf("[%s] check string want `%s` got `%s`", reflect.TypeOf(tt.node).Name(), tt.want, str)
			}
		})
	}
}

func TestNotNullConstraint_String(t *testing.T) {
	tests := []struct {
		name string
		want string
		node Node
	}{
		{
			name: "default",
			want: "not null",
			node: &NotNullConstraint{
				NotNull: tokenizer.SqlToken{Value: "not null"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := tt.node.String()
			if str != tt.want {
				t.Errorf("[%s] check string want `%s` got `%s`", reflect.TypeOf(tt.node).Name(), tt.want, str)
			}
		})
	}
}

func TestUniqueConstraint_String(t *testing.T) {
	tests := []struct {
		name string
		want string
		node Node
	}{
		{
			name: "default",
			want: "unique",
			node: &UniqueConstraint{
				ConstraintName: tokenizer.SqlToken{},
				Unique:         tokenizer.SqlToken{Value: "unique"},
				Columns:        []tokenizer.SqlToken{},
			},
		},
		{
			name: "has_constraint",
			want: "constraint tk unique (id)",
			node: &UniqueConstraint{
				ConstraintName: tokenizer.SqlToken{Value: "tk"},
				Unique:         tokenizer.SqlToken{Value: "unique"},
				Columns: []tokenizer.SqlToken{
					{Value: "id"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := tt.node.String()
			if str != tt.want {
				t.Errorf("[%s] check string want `%s` got `%s`", reflect.TypeOf(tt.node).Name(), tt.want, str)
			}
		})
	}
}

func TestForeignKeyConstraint_String(t *testing.T) {
	// create table tb_name(
	//      id int foreign key references tb (tb_id),
	//    	id int foreign key (id) references tb_1 (tb_1_id),
	//      constraint tk foreign key (id),
	//      foreign key (id) references tb_2 (tb_2_id) on delete cascade,
	//	)

	tests := []struct {
		name string
		want string
		node Node
	}{
		{
			name: "default",
			want: "foreign key references tb (tb_id)",
			node: &ForeignKeyConstraint{
				ConstraintName:    tokenizer.SqlToken{},
				ForeignKey:        tokenizer.SqlToken{Value: "foreign key"},
				ForeignKeyColumns: []tokenizer.SqlToken{},
				References:        tokenizer.SqlToken{Value: "references"},
				ReferencesTable:   tokenizer.SqlToken{Value: "tb"},
				ReferencesColumns: []tokenizer.SqlToken{{Value: "tb_id"}},
			},
		},
		{
			name: "has foreign key columns",
			want: "foreign key (id) references tb_1 (tb_1_id)",
			node: &ForeignKeyConstraint{
				ConstraintName:    tokenizer.SqlToken{},
				ForeignKey:        tokenizer.SqlToken{Value: "foreign key"},
				ForeignKeyColumns: []tokenizer.SqlToken{{Value: "id"}},
				References:        tokenizer.SqlToken{Value: "references"},
				ReferencesTable:   tokenizer.SqlToken{Value: "tb_1"},
				ReferencesColumns: []tokenizer.SqlToken{{Value: "tb_1_id"}},
			},
		},
		{
			name: "has constraint",
			want: "constraint tk foreign key (id)",
			node: &ForeignKeyConstraint{
				ConstraintName:    tokenizer.SqlToken{Value: "tk"},
				ForeignKey:        tokenizer.SqlToken{Value: "foreign key"},
				ForeignKeyColumns: []tokenizer.SqlToken{{Value: "id"}},
			},
		},
		{
			name: "has on delete",
			want: "foreign key (id) references tb_2 (tb_2_id) on delete cascade",
			node: &ForeignKeyConstraint{
				ConstraintName:    tokenizer.SqlToken{},
				ForeignKey:        tokenizer.SqlToken{Value: "foreign key"},
				ForeignKeyColumns: []tokenizer.SqlToken{{Value: "id"}},
				References:        tokenizer.SqlToken{Value: "references"},
				ReferencesTable:   tokenizer.SqlToken{Value: "tb_2"},
				ReferencesColumns: []tokenizer.SqlToken{{Value: "tb_2_id"}},
				OnDelete:          tokenizer.SqlToken{Value: "on delete"},
				Cascade:           tokenizer.SqlToken{Value: "cascade"},
			},
		},
		{
			name: "has on update",
			want: "foreign key (id) references tb_2 (tb_2_id) on update cascade",
			node: &ForeignKeyConstraint{
				ConstraintName:    tokenizer.SqlToken{},
				ForeignKey:        tokenizer.SqlToken{Value: "foreign key"},
				ForeignKeyColumns: []tokenizer.SqlToken{{Value: "id"}},
				References:        tokenizer.SqlToken{Value: "references"},
				ReferencesTable:   tokenizer.SqlToken{Value: "tb_2"},
				ReferencesColumns: []tokenizer.SqlToken{{Value: "tb_2_id"}},
				OnUpdate:          tokenizer.SqlToken{Value: "on update"},
				Cascade:           tokenizer.SqlToken{Value: "cascade"},
			},
		},
		{
			name: "has set null",
			want: "foreign key (id) references tb_2 (tb_2_id) on update set null cascade",
			node: &ForeignKeyConstraint{
				ConstraintName:    tokenizer.SqlToken{},
				ForeignKey:        tokenizer.SqlToken{Value: "foreign key"},
				ForeignKeyColumns: []tokenizer.SqlToken{{Value: "id"}},
				References:        tokenizer.SqlToken{Value: "references"},
				ReferencesTable:   tokenizer.SqlToken{Value: "tb_2"},
				ReferencesColumns: []tokenizer.SqlToken{{Value: "tb_2_id"}},
				OnUpdate:          tokenizer.SqlToken{Value: "on update"},
				SetNull:           tokenizer.SqlToken{Value: "set null"},
				Cascade:           tokenizer.SqlToken{Value: "cascade"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := tt.node.String()
			if str != tt.want {
				t.Errorf("[%s] check string want `%s` got `%s`", reflect.TypeOf(tt.node).Name(), tt.want, str)
			}
		})
	}
}

func TestCheckConstraint_String(t *testing.T) {
	// todo
	// create table tb (
	// 		id int check (age >= 10),
	//      constraint check_tk check (age >= 10)
	// )

	tests := []struct {
		name string
		want string
		node Node
	}{
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := tt.node.String()
			if str != tt.want {
				t.Errorf("[%s] check string want `%s` got `%s`", reflect.TypeOf(tt.node).Name(), tt.want, str)
			}
		})
	}
}
