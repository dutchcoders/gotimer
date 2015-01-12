package timer

import (
	"time"
)

// Func is the function prototype that will be called by Interval
type Func func() error

// Interval waits for duration and will call timerFn, an error returned by timerFn will quit the function
func Interval(timerFn Func, duration time.Duration) {
	for {
		time.Sleep(duration)

		if err := timerFn(); err != nil {
			break
		}
	}
}
