package popcount_test

import (
	"fmt"
	"lessons/ex2/2_5"
	"testing"
)

func TestPopcount(t *testing.T) {
	i := popcount.PopCount(1)
	fmt.Println(i)
}
