package withstring

//go:generate stringer -type=Status
//go:generate enummethods -type=Status
type Status int

const (
	New = Status(iota)
	Updated
	Ignored
	Imported
)
