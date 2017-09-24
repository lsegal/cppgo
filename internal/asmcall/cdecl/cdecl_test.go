package cdecl

import (
	"testing"

	"github.com/lsegal/cppgo/internal/asmcall/asmcalltest"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	v, e := Call(asmcalltest.GetCdeclcallF0Addr())
	assert.NoError(t, e)
	assert.Equal(t, uintptr(42), v)

	v, e = Call(asmcalltest.GetCdeclcallF1Addr(), 16)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(17), v)

	v, e = Call(asmcalltest.GetCdeclcallF2Addr(), 4, 2)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(2), v)

	v, e = Call(asmcalltest.GetCdeclcallF3Addr(), 4, 2, 2)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(1), v)

	v, e = Call(asmcalltest.GetCdeclcallF4Addr(), 16, 2, 4, 2)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(1), v)

	v, e = Call(asmcalltest.GetCdeclcallF5Addr(), 99, 99, 99, 99, 12)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(12), v)

	v, e = Call(asmcalltest.GetCdeclcallF6Addr(), 99, 99, 99, 99, 2, 3)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(6), v)
}
