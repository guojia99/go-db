package tokenizer

import (
	"fmt"
	"testing"

	"github.com/guojia99/go-tables/table"
)

//var dd = "select * from ( select rownum as rn,tb1.stuid,tb1.summary from ( select stuid,sum(score) as summary from gk_score group by stuid order by summary desc ) tb1 order by tb1.summary desc ) tb2 where rn<11"

func TestParserTokens(t *testing.T) {

	tests := []struct {
		name string
		data string
	}{
		{
			name: "sql",
			data: "select * from \n (select row_num as rn,tb1.stu_id,tb1.summary from tb1 order by tb1.summary desc ) tb2 where rn<11 and tt>=-1.1",
		},
		{
			name: "numbers",
			data: "-1.111,-222.2312,1+1",
		},
		{
			name: "words",
			data: "abc,def fig",
		},
		{
			name: "hex_number",
			data: "0x01 0x02 0x1234 0xffff 0xer",
		},
		{
			name: "calc_number",
			data: "1+2+3.444-4.5+1>=1",
		},
		{
			name: "calc_number_2",
			data: "1+2+3.444-4.5+1>=1 $1, 2+3*%",
		},
		{
			name: "string",
			data: "`TEST`'TEST'\"TEST\"",
		},
		{
			name: "comment",
			data: "abc /*abc*/ abc",
		},
		{
			name: "comment2",
			data: "abc /*abc -0-0- * *",
		},
		{
			name: "comment3",
			data: "abc //abc \n abc",
		},
		{
			name: "comment3",
			data: "abc ---adb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.data)
			var newTks Tokens
			for _, tk := range ParserTokens(tt.data) {
				if tk.Type == TTSpace {
					continue
				}
				newTks = append(newTks, tk)
			}
			fmt.Println(table.DefaultSimpleTable(newTks))
		})
	}
}
