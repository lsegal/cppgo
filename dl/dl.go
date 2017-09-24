package dl

import (
	"unsafe"

	"github.com/lsegal/cppgo/asmcall/cdecl"
)

type Library struct {
	p unsafe.Pointer
}

type Func struct {
	p uintptr
}

func FuncAt(addr uintptr) *Func {
	return &Func{p: addr}
}

func (f Func) Call(a ...uintptr) (uintptr, error) {
	return cdecl.Call(f.p, a...)
}
