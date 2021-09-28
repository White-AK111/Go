// Package catchpanic demonstrates interception of panic by using recover() mechanism.
//
// So, you can call implicit panic and recovered her by using GetAndRecoverPanic() function.
package catchpanic

import (
	"fmt"
	"runtime/debug"
	"time"
)

// ErrorWithTraceAndTime structure for custom error, error include trace and time of error.
type ErrorWithTraceAndTime struct {
	Text      string
	Trace     string
	Timestamp time.Time
}

// New create a new ErrorWithTraceAndTime error with "text" as string, return error.
func New(text string) error {
	return &ErrorWithTraceAndTime{
		Text:      text,
		Trace:     string(debug.Stack()),
		Timestamp: time.Now(),
	}
}

// Error implements method Error() error interface for print custom error ErrorWithTraceAndTime.
func (e *ErrorWithTraceAndTime) Error() string {
	return fmt.Sprintf("error: %s\ntrace:\n%s\ntimestamp: %v\n", e.Text, e.Trace, e.Timestamp.Format("02.01.2006 15:04:05"))
}

// getPanic create implicit panic by division by zero.
func getPanic() {
	num := 100
	var zero int

	// Division by zero, get panic.
	result := num / zero

	fmt.Println(result)
}

// GetAndRecoverPanic call and interception of implicit panic by division by zero, by using recover() mechanism, return error.
func GetAndRecoverPanic() (err error) {
	// Recover panic.
	defer func() {
		if r := recover(); r != nil {
			err = New("integer divide by zero")
		}
	}()

	getPanic()
	return err
}
