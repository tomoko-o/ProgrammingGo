package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	count := 0
	for x > 0 {
		if x != x&(x-1) {
			count++
			x = x & (x - 1)
		}
	}
	return count
}
