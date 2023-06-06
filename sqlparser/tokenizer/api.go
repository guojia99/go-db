/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package tokenizer

func ParserTokens(str string) (tokens Tokens) {
	parserTokenIterator(str, func(typ TokenType, value string, curIdx int) error {
		tokens = append(tokens, Token{
			Value: value,
			Type:  typ,
		})
		return nil
	})
	return
}

func ParserSqlTokens(str string) []SqlToken { return parserSqlTokens(str) }
