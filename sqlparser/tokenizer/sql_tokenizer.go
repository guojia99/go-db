/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package tokenizer

import (
	"strings"

	"github.com/guojia99/godb/sqlparser/keywords"
)

type SqlTokenType int

func (s SqlTokenType) String() string {
	switch s {
	case SttComment:
		return "comment"
	case SttSpace:
		return "space"
	case SttString:
		return "string"
	case SttInteger:
		return "integer"
	case SttFloat:
		return "float"
	case SttPunc:
		return "punc"
	case SttKeyword:
		return "keyword"
	case SttValue:
		return "value"
	}
	return "none"
}

const (
	SttEOF SqlTokenType = iota
	SttComment
	SttSpace

	SttString  // like TTString by 'abc' -> abc
	SttInteger // like 123
	SttFloat   // like 123.456
	SttPunc

	// SttKeyword is keywords.KeyWord
	SttKeyword
	SttValue
)

type SqlToken struct {
	Type    SqlTokenType
	Value   string
	KeyWord keywords.KeyWord
	Offset  int
	EndSet  int
	Column  int
}

func (s SqlToken) String() string {
	switch s.Type {
	case SttString:
		return `'` + strings.Replace(s.Value, `"`, `""`, -1) + `'`
	case SttValue:
		return `"` + strings.Replace(s.Value, `"`, `""`, -1) + `"`
	}
	return s.Value
}

func (s SqlToken) IsEmpty() bool { return len(s.Value) == 0 }

func (s SqlToken) IsKeyWord() bool { return s.KeyWord != keywords.KwEOF }

func parserSqlTokens(str string) (out []SqlToken) {
	idx, lastOffset := 0, 0
	newToken := func(typ SqlTokenType, value string, kw keywords.KeyWord, offset int) SqlToken {
		idx++
		tk := SqlToken{Type: typ, Value: value, Column: idx, KeyWord: kw, Offset: lastOffset, EndSet: offset}
		lastOffset = offset
		return tk
	}

	//lastKeyWordToken := SqlToken{}
	parserTokenIterator(str, func(typ TokenType, val string, curIdx int) error {
		switch typ {
		case TTNumber:
			if strings.Contains(val, ".") {
				out = append(out, newToken(SttFloat, val, keywords.KwEOF, curIdx))
				return nil
			}
			out = append(out, newToken(SttInteger, val, keywords.KwEOF, curIdx))
		case TTString:
			out = append(out, newToken(SttString, val[1:len(val)-1], keywords.KwEOF, curIdx))
		case TTWord:
			kw := keywords.GetKeyWord(val)
			if kw == keywords.KwEOF {
				out = append(out, newToken(SttValue, val, keywords.KwEOF, curIdx))
				return nil
			}

			// todo 多个的组合key`

			switch kw {
			case keywords.KwIs:

			case keywords.KwNot:

			}

			out = append(out, newToken(SttKeyword, val, kw, curIdx))
			return nil

		case TTPunctuation:
			out = append(out, newToken(SttPunc, val, keywords.KwEOF, curIdx))
		case TTComment, TTSpace:
			// pass
		}
		return nil
	})
	return
}
