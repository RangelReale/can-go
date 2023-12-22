package can

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDecode(t *testing.T) {
	/* extract negative signal */
	value := decode(&Data{0x10}, 1, 4, ByteOrderLittleEndian, ValueTypeSigned)
	assert.Check(t, is.Equal(uint64(0xFFFFFFFFFFFFFFF8), value))

	/* extract positive signal */
	value = decode(&Data{0x0E}, 1, 4, ByteOrderLittleEndian, ValueTypeSigned)
	assert.Check(t, is.Equal(uint64(0x07), value))
}
