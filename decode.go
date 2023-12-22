package can

type ByteOrder uint8

const (
	ByteOrderBigEndian ByteOrder = iota
	ByteOrderLittleEndian
)

type ValueType uint8

const (
	ValueTypeSigned ValueType = iota
	ValueTypeUnsigned
)

func decode(data []byte, startBit uint32, bitSize uint32, byteOrder ByteOrder, valueType ValueType) uint64 {
	/* safety check */
	if bitSize == 0 {
		return 0
	}

	/* copy bits */
	retVal := uint64(0)
	if byteOrder == ByteOrderBigEndian {
		/* start with MSB */
		srcBit := startBit
		dstBit := bitSize - 1
		for i := uint32(0); i < bitSize; i++ {
			/* copy bit */
			if (data[srcBit/8] & (1 << (srcBit % 8))) != 0 {
				retVal |= uint64(1) << dstBit
			}
			/* calculate next position */
			if (srcBit % 8) == 0 {
				srcBit += 15
			} else {
				srcBit--
			}
			dstBit--
		}
	} else {
		/* start with LSB */
		srcBit := startBit
		dstBit := uint32(0)
		for i := uint32(0); i < bitSize; i++ {
			/* copy bit */
			if (data[srcBit/8] & (1 << (srcBit % 8))) != 0 {
				retVal |= uint64(1) << dstBit
			}

			/* calculate next position */
			srcBit++
			dstBit++
		}
	}

	/* if signed, then fill all bits above MSB with 1 */
	if valueType == ValueTypeSigned {
		if (retVal & (1 << (bitSize - 1))) != 0 {
			for i := bitSize; i < 64; i++ {
				retVal |= uint64(1) << i
			}
		}
	}

	return retVal
}

func encode(data []byte, startBit uint32, bitSize uint32, byteOrder ByteOrder, valueType ValueType, rawValue uint64) {
	/* safety check */
	if bitSize == 0 {
		return
	}

	/* copy bits */
	if byteOrder == ByteOrderBigEndian {
		/* start with MSB */
		srcBit := startBit
		dstBit := bitSize - 1
		for i := uint32(0); i < bitSize; i++ {
			/* copy bit */
			if (rawValue & (uint64(1) << dstBit)) != 0 {
				data[srcBit/8] |= 1 << (srcBit % 8)
			} else {
				data[srcBit/8] &= ^(1 << (srcBit % 8))
			}

			/* calculate next position */
			if (srcBit % 8) == 0 {
				srcBit += 15
			} else {
				srcBit--
			}
			dstBit--
		}
	} else {
		/* start with LSB */
		srcBit := startBit
		dstBit := uint32(0)
		for i := uint32(0); i < bitSize; i++ {
			/* copy bit */
			if (rawValue & (uint64(1) << dstBit)) != 0 {
				data[srcBit/8] |= 1 << (srcBit % 8)
			} else {
				data[srcBit/8] &= ^(1 << (srcBit % 8))
			}

			/* calculate next position */
			srcBit++
			dstBit++
		}
	}
}
