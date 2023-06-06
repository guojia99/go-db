/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package types

const (
	TypeUnspecified = iota
	TypeTiny        // int16 or bool
	TypeShort       // int32
	TypeLong        // int64
	TypeFloat       // float32
	TypeDouble      // float64
	TypeNull
	TypeTimestamp // 1234567890
	TypeDate      // 2006-01-02 15:04:05
	TypeDuration  // time.Duration
	TypeDatetime
	TypeYear
	TypeNewDate
	TypeString
	TypeVarchar // TypeString limit
	TypeVarString
	TypeBit        // 0b data
	TypeJSON       // json data
	TypeNewDecimal // decimal data
	TypeEnum       // enum data
	TypeSet
	TypeTinyBlob
	TypeMediumBlob
	TypeLongBlob
	TypeBlob
	TypeGeometry
)

type Flag uint

const (
	PrimaryKeyFlag Flag = 1 << iota
	NotNullFlag
	UniqueKeyFlag
	AutoIncrementFlag
	ZerofillFlag
	UnsignedFlag
	DefaultValueFlag
	NoDefaultValueFlag
	PreventNullInsertFlag
	TimestampFlag
	JsonFlag
)

func hasFlag(flag uint, other Flag) bool { return flag&uint(other) != 0 }

type Charset string
