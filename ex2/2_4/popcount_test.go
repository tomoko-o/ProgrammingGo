package popcount_test

import (
	"fmt"
	"lessons/ex2/2_4"
	"testing"
)

func TestPopcount(t *testing.T) {
	i := popcount.PopCount(256)
	fmt.Println(i)
}
