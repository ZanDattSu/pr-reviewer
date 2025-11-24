package model

type Status int

const (
	StatusUnknown Status = iota
	StatusOpen
	StatusMerged
)

func (s Status) String() string {
	switch s {
	case StatusOpen:
		return "OPEN"
	case StatusMerged:
		return "MERGED"
	default:
		return "UNKNOWN"
	}
}
