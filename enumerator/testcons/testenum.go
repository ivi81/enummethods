package testcons

//go:generate stringer -type=TestEnum
//go:generate enummethods -type=TestEnum
type TestEnum int

const (
	First = TestEnum(iota)
	Second
	Third
	Fourth
)
