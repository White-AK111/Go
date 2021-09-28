package catchpanic

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test get and recover panic
func TestGetAndRecoverPanic(t *testing.T) {
	assert.NotNil(t, GetAndRecoverPanic(), "error on panic can't be nil")
}

func ExampleGetAndRecoverPanic() {
	err := GetAndRecoverPanic()
	if err != nil {
		fmt.Println(err)
	}
	// Output: error: integer divide by zero
	//trace:
	//goroutine 1 [running]:
	//runtime/debug.Stack(0x0, 0xc000086cb8, 0x40ce18)
	//        /usr/local/go/src/runtime/debug/stack.go:24 +0x9f
	//github.com/White-AK111/GB_Go_Level2/Lesson1/catchpanic.New(0x4cbfdc, 0x16, 0x557c70, 0x7efda5a40fff)
	//        /home/white/GolandProjects/Go/src/GB_Go_Level2/Lesson1/catchpanic/catchpanic.go:35 +0x26
	//github.com/White-AK111/GB_Go_Level2/Lesson1/catchpanic.GetAndRecoverPanic.func1(0xc000086f18)
	//        /home/white/GolandProjects/Go/src/GB_Go_Level2/Lesson1/catchpanic/catchpanic.go:61 +0x56
	//panic(0x4b3ea0, 0x557c70)
	//        /usr/local/go/src/runtime/panic.go:965 +0x1b9
	//github.com/White-AK111/GB_Go_Level2/Lesson1/catchpanic.getPanic()
	//        /home/white/GolandProjects/Go/src/GB_Go_Level2/Lesson1/catchpanic/catchpanic.go:51 +0x11
	//github.com/White-AK111/GB_Go_Level2/Lesson1/catchpanic.GetAndRecoverPanic(0x0, 0x0)
	//        /home/white/GolandProjects/Go/src/GB_Go_Level2/Lesson1/catchpanic/catchpanic.go:65 +0x52
	//main.main()
	//        /home/white/GolandProjects/Go/src/GB_Go_Level2/Lesson1/main.go:11 +0x26
	//
	//timestamp: 28.09.2021 07:56:44
}
