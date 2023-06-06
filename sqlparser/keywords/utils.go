/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package keywords

import "strings"

type KeyWord int

func (kw KeyWord) String() string { return keyWordByLowerString[kw] }
func IsKeyWord(in string) bool    { return lowerStringByKeyWord[strings.ToLower(in)] != KwEOF }
func GetKeyWord(in string) KeyWord {
	val, ok := lowerStringByKeyWord[strings.ToLower(in)]
	if ok {
		return val
	}
	return KwEOF
}
