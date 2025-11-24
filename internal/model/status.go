package model

type Status string

const (
	StatusUnknown Status = "UNKNOWN"
	StatusOpen    Status = "OPEN"
	StatusMerged  Status = "MERGED"
)
