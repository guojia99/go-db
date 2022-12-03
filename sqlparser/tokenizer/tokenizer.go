package tokenizer

import (
	"strings"
)

const (
	numberList         = "0123456789"
	hexNumberList      = numberList + "abcdefABCDEF"
	capitalizationList = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func ParserTokens(str string) (tokens Tokens) {
	length := len(str)

	if length == 0 {
		return
	}
	tokens = make(Tokens, 0, length/5)
	var (
		tkStart, idx int
		lastString   Punctuation
		lasNumber    string
	)

	tokenDealWith := func(typ TokenType) {
		if typ == TokenTypeString {
			switch lastString {
			case DoubleQuotes:
				typ = TokenTypeDoubleQuoteString
			case ApostropheQuote:
				typ = TokenTypeApostropheQuoteString
			}
		}
		if idx-tkStart != 0 {
			if len(tokens) > 0 && tokens[len(tokens)-1].Type == typ && combineOkay(typ) {
				tokens[len(tokens)-1].Value = str[tkStart-len(tokens[len(tokens)-1].Value) : idx]
				tkStart = idx
				return
			}
			tk := Token{Value: str[tkStart:idx], Type: typ}
			tokens = append(tokens, tk)
			tkStart = idx
		}
	}

	addStr := func() string {
		char := string(str[idx])
		idx++
		return char
	}

Base:
	for idx < length {
		lasNumber, lastString = "", ""
		char := addStr()
		switch char {
		case ForwardSlash:
			if idx < length && string(str[idx]) == Asterisk {
				goto Comment
			}
			tokenDealWith(TokenTypePunctuation)
		case Quotes, DoubleQuotes, ApostropheQuote:
			lastString = char
			goto QuoteString
		case MinusSign:
			if idx < length {
				lastChar := string(str[idx])
				if strings.Contains(numberList, lastChar) {
					goto Number
				}
				if lastChar == MinusSign {
					goto MinusSignEOL
				}
			}
			tokenDealWith(TokenTypePunctuation)
		case Space, NewLine, JumpBeginLine, Abdication, HorizontalTabs, VerticalTabs, FormFeed:
			goto Spaces
		case Point:
			if idx < length {
				if lastChar := string(str[idx]); strings.Contains(numberList, lastChar) {
					goto NumberDot
				}
			}
			tokenDealWith(TokenTypePunctuation)
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
		case ExclamationMark, Less, EqualGreater:
			if idx < length {
				if lastChar := string(str[idx]); lastChar == Equal {
					idx++
					tokenDealWith(TokenTypePunctuation)
					continue
				}
			}
			tokenDealWith(TokenTypePunctuation)
		case ParenLeft, ParenRight, BraceLeft, BraceRight, BracketLeft, BracketRight,
			Hashtag, Colon, Semicolon, Question, Tilde, Percent, Caret, Ampersand,
			Bar, Comma, Underscore, Equal, At, BackSlash, Dollar:
			tokenDealWith(TokenTypePunctuation)
		default:
			if strings.Contains(capitalizationList, char) {
				goto Word
			}
			if strings.Contains(numberList, char) {
				goto Number
			}
			tokenDealWith(TokenTypePunctuation)
		}
	}
	goto End

Comment:
	for idx < length {
		switch addStr() {
		case Asterisk:
			if idx < length && string(str[idx]) == ForwardSlash {
				idx++
				tokenDealWith(TokenTypeComment)
				goto Base
			}
		}
	}
	tokenDealWith(TokenTypeComment)
	goto End

QuoteString:
	for idx < length {
		switch addStr() {
		case lastString:
			tokenDealWith(TokenTypeString)
			goto Base
		case DoubleBackSlash:
			if idx < length {
				idx++
				continue
			}
			tokenDealWith(TokenTypeString)
			goto End
		}
	}
	tokenDealWith(TokenTypeString)
	goto End

MinusSignEOL:
	for idx < length {
		switch addStr() {
		case NewLine:
			tokenDealWith(TokenTypeComment)
			goto Base
		}
	}
	tokenDealWith(TokenTypeComment)
	goto End

Number:
	for idx < length {
		switch lasNumber {
		case MinusSign, PlusSign:
			if !strings.Contains(numberList, string(str[idx])) {
				tokenDealWith(TokenTypePunctuation)
				goto Base
			}
		}

		char := addStr()
		if strings.Contains(numberList, char) {
			continue
		}
		if char == Point {
			goto NumberDot
		}
		idx--
		tokenDealWith(TokenTypeNumber)
		goto Base
	}
	tokenDealWith(TokenTypeNumber)
	goto End

NumberDot:
	for idx < length {
		char := addStr()
		if strings.Contains(numberList, char) {
			continue
		}
		idx--
		tokenDealWith(TokenTypeNumber)
		goto Base
	}
	tokenDealWith(TokenTypeNumber)
	goto End

HexNumber:
	for idx < length {
		char := addStr()
		if strings.Contains(hexNumberList, char) {
			continue
		}
		idx--
		tokenDealWith(TokenTypeNumber)
		goto Base
	}
	tokenDealWith(TokenTypeNumber)
	goto End

Spaces:
	for idx < length {
		switch addStr() {
		case Space, NewLine, JumpBeginLine, Abdication, HorizontalTabs, VerticalTabs, FormFeed:
		default:
			idx--
			tokenDealWith(TokenTypeSpace)
			goto Base
		}
	}
	tokenDealWith(TokenTypeSpace)
	goto End

Word:
	for idx < length {
		char := addStr()
		if strings.Contains(numberList+capitalizationList+Underscore, char) {
			continue
		}
		idx--
		tokenDealWith(TokenTypeWord)
		goto Base
	}
	tokenDealWith(TokenTypeWord)
	goto End
End:
	return tokens
}
