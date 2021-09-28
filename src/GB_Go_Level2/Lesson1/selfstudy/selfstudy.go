package selfstudy

import (
	"errors"
	"fmt"
	"time"
)

// DontPanic don't panic in goroutine
func DontPanic() (err error) {
	go func() {
		// Move recover into goroutine
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered", r)
				err = errors.New("panic in goroutine")
			}
		}()

		panic("A-A-A!!!")
	}()
	time.Sleep(time.Second)
	return err
}
