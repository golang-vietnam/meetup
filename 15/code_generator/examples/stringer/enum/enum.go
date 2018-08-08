package enum

//go:generate stringer -type=ChangeType -output=gen.go
type ChangeType int

const (
	UpsertJSON ChangeType = iota + 1
	UpsertText
	Remove
	Rename
	ApplyJSONPatch
	ApplyTextPatch
)
