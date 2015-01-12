package main

import (
	"fmt"
	"github.com/dutchcoders/gotimer"
	"time"
)

func main() {
	var i int = 0

	go timer.Interval(func() error {
		fmt.Printf("test %d\n", i)
		i++
		return nil
	}, 1000*time.Millisecond)

	time.Sleep(100000 * time.Millisecond)
}
