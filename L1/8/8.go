package _

func setBit(n int64, i uint, bit bool) int64 {
	if bit {
		n |= 1 << i
	} else {
		n &= ^(1 << i)
	}
	return n
}
