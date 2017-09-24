package thiscall

import (
	"testing"

	"github.com/lsegal/cppgo/asmcall/internal/asmcalltest"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	v, e := Call(asmcalltest.GetThiscallF0Addr(), 0)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(42), v)

	v, e = Call(asmcalltest.GetThiscallF1Addr(), 0, 16)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(17), v)

	v, e = Call(asmcalltest.GetThiscallF2Addr(), 0, 4, 2)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(2), v)

	v, e = Call(asmcalltest.GetThiscallF3Addr(), 0, 4, 2, 2)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(1), v)

	v, e = Call(asmcalltest.GetThiscallF4Addr(), 0, 16, 2, 4, 2)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(1), v)

	v, e = Call(asmcalltest.GetThiscallF5Addr(), 0, 99, 99, 99, 99, 12)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(12), v)
}
