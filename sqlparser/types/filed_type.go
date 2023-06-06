/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package types

type FiledType struct {
	Type    int
	Flag    Flag
	Len     int
	Decimal int
	Charset Charset
	Enum    []string // the enum like Enum('a', 'b') or Enum(1, 2, 3)
}
