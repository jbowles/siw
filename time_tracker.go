package siw

import (
	"fmt"
	"time"
)

func timeTracker(start time.Time, name string) {
	var elapsed = time.Since(start)
	fmt.Printf("For %q elapsed time = \t actual: %v\n", name, elapsed)
}
