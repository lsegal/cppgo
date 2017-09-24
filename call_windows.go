package cpp

import (
	"github.com/lsegal/cppgo/internal/asmcall/stdcall"
	"github.com/lsegal/cppgo/internal/asmcall/thiscall"
)

func (p Ptr) cdeclcall(offset int, a ...uintptr) (uintptr, error) {
	return p.thiscall(offset, a...)
}

func (p Ptr) stdcall(offset int, a ...uintptr) (uintptr, error) {
	addr := p.getaddr(offset)
	return stdcall.Call(addr, a...)
}

func (p Ptr) thiscall(offset int, a ...uintptr) (uintptr, error) {
	addr := p.getaddr(offset)
	return thiscall.Call(addr, a...)
}
