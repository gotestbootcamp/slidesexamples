package sum

import "time"

func Sum(a, b int) int {
	time.Sleep(500 * time.Millisecond)
	return a + b
}

type SumFixed struct {
	ToSum int
}
