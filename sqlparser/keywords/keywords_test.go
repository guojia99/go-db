/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

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

		// double key word
		if keyUnIndex := strings.Index(key, " "); keyUnIndex != -1 {
			key = key[:keyUnIndex] + strings.ToUpper(string(key[keyUnIndex+1])) + key[keyUnIndex+2:]
		}

		key = fmt.Sprintf("Kw%s", key)

		constKeyWordStr += fmt.Sprintf("\t%s\n", key)
		keyWordByLowerStringStr += fmt.Sprintf("\t%s:\"%s\",\n", key, strings.ToUpper(val))
		lowerStringByKeyWordStr += fmt.Sprintf("\t\"%s\":%s,\n", strings.ToUpper(val), key)
	}

	setData := fmt.Sprintf(temp, constKeyWordStr, keyWordByLowerStringStr, lowerStringByKeyWordStr)
	_ = os.WriteFile("./keywords.go", []byte(setData), 0600)
}
