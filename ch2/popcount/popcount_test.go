package popcount_test

import (
	"fmt"
	"lessons/ch2/popcount"
	"testing"
)

func TestPopcount(t *testing.T) {
	i := popcount.PopCount(17)
	fmt.Println(i)
}
