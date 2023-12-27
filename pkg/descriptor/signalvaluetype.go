package descriptor

import "strings"

// SignalValueType represents an extended signal value type.
type SignalValueType uint8

//go:generate stringer -type SignalValueType -trimprefix SignalValueType

const (
	SignalValueTypeNone SignalValueType = iota
	SignalValueTypeInt
	SignalValueTypeFloat32
	SignalValueTypeFloat64
)

// UnmarshalString sets the value of *s from the provided string.
func (s *SignalValueType) UnmarshalString(str string) error {
	// TODO: Decide on conventions and make this more strict
	switch strings.ToLower(str) {
	case "int", "integer":
		*s = SignalValueTypeInt
	case "float32":
		*s = SignalValueTypeFloat32
	case "float64", "float":
		*s = SignalValueTypeFloat64
	default:
		*s = SignalValueTypeNone
	}
	return nil
}
