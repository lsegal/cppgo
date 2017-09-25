// +build !windows

package cpp

import "github.com/lsegal/cppgo/asmcall/cdecl"

func (p ptr) stdcall(offset int, a ...uintptr) (uintptr, error) {
	// ignore stdcalls on non-Windows platform
	return p.cdeclcall(offset, a...)
}

func (p ptr) thiscall(offset int, a ...uintptr) (uintptr, error) {
	// ignore thiscalls on non-Windows platform
	return p.cdeclcall(offset, a...)
}

func (p ptr) cdeclcall(offset int, a ...uintptr) (uintptr, error) {
	addr := p.getaddr(offset)
	return cdecl.Call(addr, a...)
}
