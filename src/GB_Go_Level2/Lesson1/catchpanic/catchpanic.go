package catchpanic

import (
	"fmt"
	"runtime/debug"
	"time"
)

// ErrorWithTraceAndTime struct for custom error
type ErrorWithTraceAndTime struct {
	text      string
	trace     string
	timestamp time.Time
}

// New create new custom error
func New(text string) error {
	return &ErrorWithTraceAndTime{
		text:      text,
		trace:     string(debug.Stack()),
		timestamp: time.Now(),
	}
}

// Error print custom error
func (e *ErrorWithTraceAndTime) Error() string {
	return fmt.Sprintf("error: %s\ntrace:\n%s\ntimestamp: %v\n", e.text, e.trace, e.timestamp.Format("02.01.2006 15:04:05"))
}

// getPanic create panic
func getPanic() {
	num := 100
	var zero int

	// Division by zero, get panic
	result := num / zero

	fmt.Println(result)
}

// GetAndRecoverPanic create panic and catch her
func GetAndRecoverPanic() (err error) {
	// Recover panic
	defer func() {
		if r := recover(); r != nil {
			err = New("integer divide by zero")
		}
	}()

	getPanic()
	return err
}
