package popcount_test

import (
	"fmt"
	"lessons/ex2/2_3"
	"testing"
)

func TestPopcount(t *testing.T) {
	i := popcount.PopCount(255)
	fmt.Println(i)
}
