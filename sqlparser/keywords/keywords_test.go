package keywords

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestNewKeyWordFile(t *testing.T) {
	temp := `
package keywords

import strings "strings"

type KeyWord int

const (
	KwEOF KeyWord = iota
	%s
)

var keyWordByLowerString = map[KeyWord]string{
	%s
}

var lowerStringByKeyWord = map[string]KeyWord{
	%s
}

func (kw KeyWord) String() string { return keyWordByLowerString[kw] }
func IsKeyWord(in string) bool    { return lowerStringByKeyWord[strings.ToLower(in)] != KwEOF }
func GetKeyWord(in string) KeyWord {
	val, ok := lowerStringByKeyWord[strings.ToLower(in)]
	if ok {
		return val
	}
	return KwEOF
}

`

	data, _ := os.ReadFile("./keyword.txt")
	dataList := strings.Split(string(data), "\n")

	constKeyWordStr := ""
	keyWordByLowerStringStr := ""
	lowerStringByKeyWordStr := ""
	for _, val := range dataList {
		var key = val[:1] + strings.ToLower(val[1:])

		if keyUnIndex := strings.Index(key, "_"); keyUnIndex != -1 {
			key = key[:keyUnIndex] + strings.ToUpper(string(key[keyUnIndex+1])) + key[keyUnIndex+2:]
		}
		key = fmt.Sprintf("Kw%s", key)

		constKeyWordStr += fmt.Sprintf("\t%s\n", key)
		keyWordByLowerStringStr += fmt.Sprintf("\t%s:\"%s\",\n", key, strings.ToLower(val))
		lowerStringByKeyWordStr += fmt.Sprintf("\t\"%s\":%s,\n", strings.ToLower(val), key)
	}

	setData := fmt.Sprintf(temp, constKeyWordStr, keyWordByLowerStringStr, lowerStringByKeyWordStr)
	_ = os.WriteFile("./keywords.go", []byte(setData), 0600)
}
