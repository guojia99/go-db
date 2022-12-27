package tokenizer

import (
	"strings"
)

const (
	numberList         = "0123456789"
	hexNumberList      = numberList + "abcdefABCDEF"
	capitalizationList = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Punc copy like `https://www.cnblogs.com/moonSong/p/4989637.html`
// https://en.wikipedia.org/wiki/Punctuation
type Punc = string

const (
	NotPunc             Punc = ``
	PuncQuotes          Punc = `'`
	PuncDoubleQuotes    Punc = `"`
	PuncApostropheQuote Punc = "`"
	PuncForwardSlash    Punc = `/`
	PuncBackSlash       Punc = `\`
	PuncDoubleBackSlash Punc = `\\`
	PuncParenLeft       Punc = `(`
	PuncParenRight      Punc = `)`
	PuncBraceLeft       Punc = `{`
	PuncBraceRight      Punc = `}`
	PuncBracketLeft     Punc = `[`
	PuncBracketRight    Punc = `]`
	PuncMinusSign       Punc = `-`
	PuncPlusSign        Punc = `+`
	PuncExclamationMark Punc = `!`
	PuncEqual           Punc = `=`
	PuncUnEqual         Punc = `!=`
	PuncLess            Punc = `<`
	PuncEqualLess       Punc = `<=`
	PuncGreater         Punc = `>`
	PuncEqualGreater    Punc = `>=`
	PuncSpace           Punc = ` `
	PuncNewLine         Punc = "\n"
	PuncJumpBeginLine   Punc = "\r"
	PuncAbdication      Punc = "\b"
	PuncHorizontalTabs  Punc = "\t"
	PuncVerticalTabs    Punc = "\v"
	PuncFormFeed        Punc = "\f"
	PuncHashtag         Punc = `#`
	PuncColon           Punc = `:`
	PuncSemicolon       Punc = `;`
	PuncQuestion        Punc = `?`
	PuncAt              Punc = `@`
	PuncTilde           Punc = `~`
	PuncPercent         Punc = `%`
	PuncCaret           Punc = `^`
	PuncAmpersand       Punc = `&`
	PuncAsterisk        Punc = `*`
	PuncBar             Punc = `|`
	PuncComma           Punc = `,`
	PuncUnderscore      Punc = `_`
	PuncPoint           Punc = `.`
	PuncDollar          Punc = `$`
)

type TokenType int

const (
	TTSpace       TokenType = iota // ` ` or `\n`... space
	TTNumber                       // `123`
	TTString                       // `'abc'` or `"abc"` or ``abc``
	TTWord                         // `abc`
	TTPunctuation                  // like Punc
	TTComment                      // `/* Abc */`
)

func (tkp TokenType) String() string {
	switch tkp {
	case TTSpace:
		return "space"
	case TTString:
		return "string"
	case TTNumber:
		return "number"
	case TTPunctuation:
		return "punctuation"
	case TTWord:
		return "word"
	case TTComment:
		return "comment"
	}
	return "not"
}

// Token is the set of lexical tokens of the Go programming language.
type Token struct {
	Value string    `json:"v"`
	Type  TokenType `json:"t"`
}

func (tk Token) String() string { return tk.Value }

func (tk Token) ToPunctuation() Punc {
	if tk.Type != TTPunctuation {
		return NotPunc
	}
	return tk.Value
}

type Tokens []Token

func (ts Tokens) String() string {
	str := ""
	for _, v := range ts {
		str += v.Value
	}
	return str
}

type parserTokenFn func(typ TokenType, value string, curIdx int) error

func parserTokenIterator(str string, parserTokenFn parserTokenFn) {
	length := len(str)
	if length == 0 {
		return
	}

	var (
		tkStart, idx int
		lastString   Punc
		lasNumber    string
	)

	tokenDealWith := func(typ TokenType) {
		if idx-tkStart != 0 {
			if err := parserTokenFn(typ, str[tkStart:idx], idx); err != nil {
				// quit done.
				idx = length
				return
			}
			tkStart = idx
		}
	}

	getCurStrAndNext := func() string {
		char := string(str[idx])
		idx++
		return char
	}

	// Base to call all start punt
Base:
	for idx < length {
		lasNumber, lastString = "", ""
		char := getCurStrAndNext()
		switch char {
		case PuncForwardSlash:
			if idx < length {
				lastChar := string(str[idx])
				switch lastChar {
				case PuncAsterisk:
					goto Comment
				case PuncForwardSlash:
					goto ForwardSlashComment
				}
			}
			tokenDealWith(TTPunctuation)
		case PuncMinusSign:
			if idx < length {
				lastChar := string(str[idx])
				if lastChar == PuncMinusSign {
					goto MinusSignComment
					continue
				}
				tokenDealWith(TTPunctuation)
			}
		case PuncQuotes, PuncDoubleQuotes, PuncApostropheQuote:
			lastString = char
			goto QuoteString
		case PuncSpace, PuncNewLine, PuncJumpBeginLine, PuncAbdication, PuncHorizontalTabs, PuncVerticalTabs, PuncFormFeed:
			goto Spaces
		case PuncPoint:
			if idx < length {
				if lastChar := string(str[idx]); strings.Contains(numberList, lastChar) {
					goto NumberDot
				}
			}
			tokenDealWith(TTPunctuation)
		case "0":
			if idx < length {
				lastChar := string(str[idx])
				switch lastChar {
				case "x", "X":
					idx++
					goto HexNumber
				}
			}
			goto Number
		case PuncExclamationMark, PuncLess, PuncGreater:
			if idx < length {
				if lastChar := string(str[idx]); lastChar == PuncEqual {
					idx++
					tokenDealWith(TTPunctuation)
					continue
				}
			}
			tokenDealWith(TTPunctuation)
		case PuncParenLeft, PuncParenRight, PuncBraceLeft, PuncBraceRight, PuncBracketLeft, PuncBracketRight,
			PuncHashtag, PuncColon, PuncSemicolon, PuncQuestion, PuncTilde, PuncPercent, PuncCaret, PuncAmpersand,
			PuncBar, PuncComma, PuncUnderscore, PuncEqual, PuncAt, PuncDollar:
			tokenDealWith(TTPunctuation)
		default:
			if strings.Contains(capitalizationList, char) {
				goto Word
			}
			if strings.Contains(numberList, char) {
				goto Number
			}
			tokenDealWith(TTPunctuation)
		}
	}
	goto End

	// Comment to call like end `/* abc */`
Comment:
	for idx < length {
		switch getCurStrAndNext() {
		case PuncAsterisk:
			if idx < length && string(str[idx]) == PuncForwardSlash {
				idx++
				tokenDealWith(TTComment)
				goto Base
			}
		}
	}
	tokenDealWith(TTComment)
	goto End

MinusSignComment:
	for idx < length {
		switch getCurStrAndNext() {
		case PuncNewLine:
			tokenDealWith(TTComment)
			goto Base
		}
	}
	tokenDealWith(TTComment)
	goto End

ForwardSlashComment:
	for idx < length {
		switch getCurStrAndNext() {
		case PuncNewLine:
			tokenDealWith(TTComment)
			goto Base
		}
	}
	tokenDealWith(TTComment)
	goto End

QuoteString:
	for idx < length {
		switch getCurStrAndNext() {
		case lastString:
			tokenDealWith(TTString)
			goto Base
		case PuncDoubleBackSlash:
			if idx < length {
				idx++
				continue
			}
			tokenDealWith(TTString)
			goto End
		}
	}
	tokenDealWith(TTString)
	goto End

Number:
	for idx < length {
		switch lasNumber {
		case PuncMinusSign, PuncPlusSign:
			if !strings.Contains(numberList, string(str[idx])) {
				tokenDealWith(TTPunctuation)
				goto Base
			}
		}

		char := getCurStrAndNext()
		if strings.Contains(numberList, char) {
			continue
		}
		if char == PuncPoint {
			goto NumberDot
		}
		idx--
		tokenDealWith(TTNumber)
		goto Base
	}
	tokenDealWith(TTNumber)
	goto End

NumberDot:
	for idx < length {
		char := getCurStrAndNext()
		if strings.Contains(numberList, char) {
			continue
		}
		idx--
		tokenDealWith(TTNumber)
		goto Base
	}
	tokenDealWith(TTNumber)
	goto End

HexNumber:
	for idx < length {
		char := getCurStrAndNext()
		if strings.Contains(hexNumberList, char) {
			continue
		}
		idx--
		tokenDealWith(TTNumber)
		goto Base
	}
	tokenDealWith(TTNumber)
	goto End

Spaces:
	for idx < length {
		switch getCurStrAndNext() {
		case PuncSpace, PuncNewLine, PuncJumpBeginLine, PuncAbdication, PuncHorizontalTabs, PuncVerticalTabs, PuncFormFeed:
		default:
			idx--
			tokenDealWith(TTSpace)
			goto Base
		}
	}
	tokenDealWith(TTSpace)
	goto End

Word:
	for idx < length {
		char := getCurStrAndNext()
		if strings.Contains(numberList+capitalizationList+PuncUnderscore, char) {
			continue
		}
		idx--
		tokenDealWith(TTWord)
		goto Base
	}
	tokenDealWith(TTWord)
	goto End
End:
	return
}
