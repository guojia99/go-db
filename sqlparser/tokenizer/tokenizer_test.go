package tokenizer

import (
	"fmt"
	"strings"
	"testing"
)

//var dd = "select * from ( select rownum as rn,tb1.stuid,tb1.summary from ( select stuid,sum(score) as summary from gk_score group by stuid order by summary desc ) tb1 order by tb1.summary desc ) tb2 where rn<11"

func TestParserTokens(t *testing.T) {
	data := `select * from (select row_num as rn,tb1.stu_id,tb1.summary from tb1 order by tb1.summary desc ) tb2 where rn<11 and tt>=-1.1`
	data = strings.Trim(data, "\n")
	tks := ParserTokens(data)
	for _, tk := range tks {
		if tk.Type == TokenTypeSpace {
			continue
		}
		fmt.Printf("[%s]-[%s]\n", tk.Type, tk.Value)
	}
}
