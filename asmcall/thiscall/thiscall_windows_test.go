package thiscall

import (
	"testing"

	"github.com/lsegal/cppgo/asmcall/internal/asmcalltest"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	obj := asmcalltest.GetThiscallObj()

	v, e := Call(asmcalltest.GetThiscallF0Addr(), obj)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(42), v)

	v, e = Call(asmcalltest.GetThiscallF1Addr(), obj, 16)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(17), v)

	v, e = Call(asmcalltest.GetThiscallF2Addr(), obj, 4, 2)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(2), v)

	v, e = Call(asmcalltest.GetThiscallF3Addr(), obj, 4, 2, 2)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(1), v)

	v, e = Call(asmcalltest.GetThiscallF4Addr(), obj, 16, 2, 4, 2)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(1), v)

	v, e = Call(asmcalltest.GetThiscallF5Addr(), obj, 99, 99, 99, 99, 12)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(12), v)
}
