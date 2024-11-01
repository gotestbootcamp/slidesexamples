package sum

import (
	"fmt"
	"time"
)

type Calculator struct {
	t time.Duration
}

func NewCalculator() Calculator {
	return Calculator{t: time.Millisecond}
}

func (c Calculator) Sum(a, b int) int {
	time.Sleep(c.t)
	return a + b
}

func (c Calculator) Subtract(a, b int) int {
	time.Sleep(c.t)
	return a - b
}

func (c Calculator) Unregister() {
	fmt.Println("unregistered")
}
