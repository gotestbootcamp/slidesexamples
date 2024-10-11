package sum

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("with %d", i)
		t.Run(name, func(t *testing.T) {
			res := Sum(6, i)
			if 6+i != res {
				t.Errorf("Expected %d from %d, got %d", 6+i, i, res)
			}
		})
	}
}

func TestSetup(t *testing.T) {
	// Setup
	t.Run("first", func(t *testing.T) {
		//
	})
	t.Run("second", func(t *testing.T) {
		//
	})
	t.Run("third", func(t *testing.T) {
		//
	})
	// Teardown
}
