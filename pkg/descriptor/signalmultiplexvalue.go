package descriptor

type SignalMultiplexValue struct {
	MultiplexerSwitch string
	Ranges            []SignalMultiplexRangeValue
}

type SignalMultiplexRangeValue struct {
	RangeStart uint64
	RangeEnd   uint64
}
