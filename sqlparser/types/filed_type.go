package types

type FiledType struct {
	Type    int
	Flag    Flag
	Len     int
	Decimal int
	Charset Charset
	Enum    []string // the enum like Enum('a', 'b') or Enum(1, 2, 3)
}
