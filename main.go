// Package timeecho provides a function to echo the current time.
package timeecho

import (
	"fmt"
	"time"
)

// EchoTime prints the current time to the console.
func EchoTime() {
	fmt.Println("Current time:", time.Now().Format("2006-01-02 15:04:05"))
}
