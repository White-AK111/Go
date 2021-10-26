package race

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Test get and recover panic
func TestGoroutinePool_IncRace(t *testing.T) {
	var exp int64 = 1000
	rc := NewGoroutinePool(1000)
	ctxRc, cancelFuncMan := context.WithTimeout(context.Background(), time.Second)
	defer cancelFuncMan()
	res, _ := rc.IncRace(ctxRc)

	for i := 0; i < 10; i++ {
		assert.Equal(t, exp, res, "Incorrect result expected 1000 got %d \nCheck to race condition", res)
	}
}
