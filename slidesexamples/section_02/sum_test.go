package sum

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Run(fmt.Sprintf("test sum %d", i), func(t *testing.T) {
			res := Sum(6, i)
			if 6+i != res {
				t.Errorf("Expected %d from %d, got %d", 6+i, i, res)
			}
		})
	}
}
