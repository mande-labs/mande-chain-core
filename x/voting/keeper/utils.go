package keeper

func intAbs(n int64) uint64 {
	y := n >> 63
	return uint64((n ^ y) - y)
}
