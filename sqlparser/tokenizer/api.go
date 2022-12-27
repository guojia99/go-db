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
