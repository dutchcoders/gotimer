# gotimer
Timer library for go. Will call function after timeout.

### Example
```
package main

import (
    "fmt"
    "github.com/dutchcoders/gotimer"
    "time"
)

func main() {

    go timer.Interval(func(args ...interface{}) error {
        fmt.Println("test")
        return nil
    }, 1000*time.Millisecond)

    time.Sleep(100000 * time.Millisecond)
}
```
