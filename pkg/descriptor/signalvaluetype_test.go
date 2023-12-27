package descriptor

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestSignalValueType_UnmarshalString(t *testing.T) {
	for _, tt := range []struct {
		str      string
		expected SignalValueType
	}{
		{str: "int", expected: SignalValueTypeInt},
		{str: "float32", expected: SignalValueTypeFloat32},
		{str: "float64", expected: SignalValueTypeFloat64},
	} {
		tt := tt
		t.Run(tt.str, func(t *testing.T) {
			var actual SignalValueType
			assert.NilError(t, actual.UnmarshalString(tt.str))
			assert.Equal(t, tt.expected, actual)
		})
	}
}
