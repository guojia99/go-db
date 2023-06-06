/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package sqlparser

type Type = int

const (
	// TypeNull are empty data type
	TypeNull Type = iota
	TypeBoolean
	TypeInteger
	TypeDecimal
	TypeString
	TypeTimestamp
	TypeDatetime
	TypeInterval
	TypeDuration
	TypeSet
	TypeJSON
	TypeEnum
	TypeArray
)

// Flag information.
type Flag uint64

const (
	FlagNotNull       Flag = 1 << 0
	FlagPriKey        Flag = 1 << 1
	FlagUniqueKey     Flag = 1 << 2
	FlagAutoIncrement Flag = 1 << 3
	FlagDefaultValue  Flag = 1 << 4
	FlagUnique        Flag = 1 << 5
	FlagJSON          Flag = 1 << 6
)

func (f Flag) Has(want Flag) bool    { return f&want > 0 }
func (f Flag) IsNotNull() bool       { return f&FlagNotNull > 0 }
func (f Flag) IsPriKey() bool        { return f&FlagPriKey > 0 }
func (f Flag) IsUniqueKey() bool     { return f&FlagUniqueKey > 0 }
func (f Flag) IsAutoIncrement() bool { return f&FlagAutoIncrement > 0 }
func (f Flag) IsDefaultValue() bool  { return f&FlagDefaultValue > 0 }
func (f Flag) IsUnique() bool        { return f&FlagUnique > 0 }
func (f Flag) IsJSON() bool          { return f&FlagJSON > 0 }
