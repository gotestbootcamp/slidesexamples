package sum

import "testing"

func TestCalculator(t *testing.T) {
	c := NewCalculator()
	t.Cleanup(func() {
		c.Unregister()
	})

	t.Run("sum 1+2", func(t *testing.T) {
		if c.Sum(1, 2) != 3 {
			t.Fail()
		}
	})
	t.Run("sum 1+3", func(t *testing.T) {
		if c.Sum(1, 3) != 4 {
			t.Fail()
		}
	})
}

func TestCalculatorTable(t *testing.T) {
	c := NewCalculator()
	t.Cleanup(func() {
		c.Unregister()
	})

	tests := []struct {
		name     string
		first    int
		second   int
		expected int
	}{
		{"1+2", 1, 2, 3},
		{"1+3", 1, 3, 4},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if c.Sum(tc.first, tc.second) != tc.expected {
				t.Fail()
			}
		})
	}
}
