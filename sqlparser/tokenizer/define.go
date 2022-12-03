package tokenizer

// Punctuation copy like `https://www.cnblogs.com/moonSong/p/4989637.html`
// https://en.wikipedia.org/wiki/Punctuation
type Punctuation = string

const (
	NotPunctuation Punctuation = ``

	Quotes          Punctuation = `'`
	DoubleQuotes    Punctuation = `"`
	ApostropheQuote Punctuation = "`"
	ForwardSlash    Punctuation = `/'`
	BackSlash       Punctuation = `\`
	DoubleBackSlash Punctuation = `\\`

	ParenLeft    Punctuation = `(`
	ParenRight   Punctuation = `)`
	BraceLeft    Punctuation = `{`
	BraceRight   Punctuation = `}`
	BracketLeft  Punctuation = `[`
	BracketRight Punctuation = `]`

	MinusSign       Punctuation = `-`
	PlusSign        Punctuation = `+`
	ExclamationMark Punctuation = `!`
	Equal           Punctuation = `=`
	UnEqual         Punctuation = `!=`
	Less            Punctuation = `<`
	EqualLess       Punctuation = `<=`
	Greater         Punctuation = `>`
	EqualGreater    Punctuation = `>=`

	Space          Punctuation = ` `
	NewLine        Punctuation = "\n"
	JumpBeginLine  Punctuation = "\r"
	Abdication     Punctuation = "\b"
	HorizontalTabs Punctuation = "\t"
	VerticalTabs   Punctuation = "\v"
	FormFeed       Punctuation = "\f"

	Hashtag    Punctuation = `#`
	Colon      Punctuation = `:`
	Semicolon  Punctuation = `;`
	Question   Punctuation = `?`
	At         Punctuation = `@`
	Tilde      Punctuation = `~`
	Percent    Punctuation = `%`
	Caret      Punctuation = `^`
	Ampersand  Punctuation = `&`
	Asterisk   Punctuation = `*`
	Bar        Punctuation = `|`
	Comma      Punctuation = `,`
	Underscore Punctuation = `_`
	Point      Punctuation = `.`
	Dollar     Punctuation = `$`
)

type TokenType int

const (
	TokenTypeNot                   TokenType = iota // not
	TokenTypeSpace                                  // ` ` space
	TokenTypeNumber                                 // `123`
	TokenTypeDollarNumber                           // `$123` will int
	TokenTypeString                                 // `'abc'`
	TokenTypeDoubleQuoteString                      // `"abc"`
	TokenTypeApostropheQuoteString                  // ``abc``
	TokenTypeWord                                   // `abc`
	TokenTypePunctuation                            // like Punctuation
	TokenTypeBoolWord                               // `true` or `false`
	TokenTypeNullWord                               // `NULL` or `null`
	TokenTypeComment                                // `/* Abc */`
)

func (tkp TokenType) String() string {
	switch tkp {
	case TokenTypeSpace:
		return "space"
	case TokenTypeString:
		return "string"
	case TokenTypeDoubleQuoteString:
		return "dq_string"
	case TokenTypeApostropheQuoteString:
		return "aq_string"
	case TokenTypeNumber:
		return "number"
	case TokenTypePunctuation:
		return "punctuation"
	case TokenTypeDollarNumber:
		return "$number"
	case TokenTypeWord:
		return "word"
	case TokenTypeBoolWord:
		return "bool"
	case TokenTypeNullWord:
		return "null"
	case TokenTypeComment:
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

func (tk Token) ToPunctuation() Punctuation {
	if tk.Type != TokenTypePunctuation {
		return NotPunctuation
	}
	return Punctuation(tk.Value)
}

type Tokens []Token

func (ts Tokens) String() string {
	str := ""
	for _, v := range ts {
		str += v.Value
	}
	return str
}

func combineOkay(t TokenType) bool {
	// nolint:exhaustive
	switch t {
	case TokenTypeNumber, TokenTypeDollarNumber:
		return false
	}
	return true
}
