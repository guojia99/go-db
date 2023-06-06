/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package tokenizer

import (
	"fmt"
	"testing"

	"github.com/guojia99/go-tables/table"
)

func Test_parserSqlTokens(t *testing.T) {
	data := ` select 
                a.product_id,
                a.product_name,
                count(a.ins_id) as ins_num,
                -- 性别
                count(a.f) as f_num,
                count(a.m) as m_num,
                -- 成员数
                count(a.p_1) as p_1_num,
                count(a.p_2) as p_1_num,
                count(a.p_3) as p_1_num,
                count(a.gt3) as gt3_num,
                -- 年龄
                count(lt25) as lt25_num,
                count(gt25lt35) as gt25lt35_num,
                count(gt35lt45) as gt25lt35_num,
                count(gt45lt55) as gt25lt35_num,
                count(gt55) as gt55_num
        from(
                select 
                        a.ins_id,
                        b.product_id,
                        b.product_name,
                        c.cust_id,
                        c.cust_name,
                        c.cust_sex,
                        c.cust_age,
                        c.family_num,
                        -- 男
                        -- 这个地方根据数据库字段的不同，处理方式也不同
                        -- 如果数据库中cust_sex的数据类型本身就是0和1，那么就不需要转换
                        -- 只列出来即可
                        (case when c.cust_sex='男' then 1 else 0 end) as f,
                        -- 女
                        (case when c.cust_sex='女' then 1 else 0 end) as as m,
                        -- 其他的依次类推
                        -- 家庭成员数
                        (case when c.family_num=1 then 1 else 0 end) as p_1,
                        (case when c.family_num=2 then 1 else 0 end) as P_2,
                        (case when c.family_num=3 then 1 else 0 end) as p_3,
                        (case when c.family_num>3 then 1 else 0 end) as gt3,
                        -- 客户年龄
                        (case when c.cust_age<=25 then 1 else 0 end) as lt25,
                        (case when c.cust_age>25 and c.cust_age<=35 then 1 else 0 end) as gt25lt35,
                        (case when c.cust_age>35 and c.cust_age<=45 then 1 else 0 end) as gt35lt45,
                        (case when c.cust_age>45 and c.cust_age<=55 then 1 else 0 end) as gt45lt55,
                        (case when c.cust_age>55 then 1 else 0 end) as gt55
                from
                        insurance a,
                        product b,
                        customer c
                where 
                        a.product_id=b.product_id
                        and a.cust_id=c.cust_id
        ) a
        group by b.product_id, b.product_name`

	fmt.Println(data)
	tb := parserSqlTokens(data)
	fmt.Println(table.DefaultSimpleTable(tb))
}

func BenchmarkParserSqlTokens(b *testing.B) {
	data := `select * from table where a = b;`
	for i := 0; i < b.N; i++ {
		_ = parserSqlTokens(data)
	}
}

func BenchmarkParserSqlTokensByBigSQL(b *testing.B) {
	data := ` select 
                a.product_id,
                a.product_name,
                count(a.ins_id) as ins_num,
                -- 性别
                count(a.f) as f_num,
                count(a.m) as m_num,
                -- 成员数
                count(a.p_1) as p_1_num,
                count(a.p_2) as p_1_num,
                count(a.p_3) as p_1_num,
                count(a.gt3) as gt3_num,
                -- 年龄
                count(lt25) as lt25_num,
                count(gt25lt35) as gt25lt35_num,
                count(gt35lt45) as gt25lt35_num,
                count(gt45lt55) as gt25lt35_num,
                count(gt55) as gt55_num
        from(
                select 
                        a.ins_id,
                        b.product_id,
                        b.product_name,
                        c.cust_id,
                        c.cust_name,
                        c.cust_sex,
                        c.cust_age,
                        c.family_num,
                        -- 男
                        -- 这个地方根据数据库字段的不同，处理方式也不同
                        -- 如果数据库中cust_sex的数据类型本身就是0和1，那么就不需要转换
                        -- 只列出来即可
                        (case when c.cust_sex='男' then 1 else 0 end) as f,
                        -- 女
                        (case when c.cust_sex='女' then 1 else 0 end) as as m,
                        -- 其他的依次类推
                        -- 家庭成员数
                        (case when c.family_num=1 then 1 else 0 end) as p_1,
                        (case when c.family_num=2 then 1 else 0 end) as P_2,
                        (case when c.family_num=3 then 1 else 0 end) as p_3,
                        (case when c.family_num>3 then 1 else 0 end) as gt3,
                        -- 客户年龄
                        (case when c.cust_age<=25 then 1 else 0 end) as lt25,
                        (case when c.cust_age>25 and c.cust_age<=35 then 1 else 0 end) as gt25lt35,
                        (case when c.cust_age>35 and c.cust_age<=45 then 1 else 0 end) as gt35lt45,
                        (case when c.cust_age>45 and c.cust_age<=55 then 1 else 0 end) as gt45lt55,
                        (case when c.cust_age>55 then 1 else 0 end) as gt55
                from
                        insurance a,
                        product b,
                        customer c
                where 
                        a.product_id=b.product_id
                        and a.cust_id=c.cust_id
        ) a
        group by b.product_id, b.product_name`

	for i := 0; i < b.N; i++ {
		_ = parserSqlTokens(data)
	}
}
