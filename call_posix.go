// +build !windows

package cpp

import (
	"errors"

	"github.com/lsegal/cppgo/asmcall/cdecl"
)

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
	switch l := len(a); l {
	case 0:
		return cdecl.Call0(addr), nil
	case 1:
		return cdecl.Call1(addr, a[0]), nil
	case 2:
		return cdecl.Call2(addr, a[0], a[1]), nil
	case 3:
		return cdecl.Call3(addr, a[0], a[1], a[2]), nil
	case 4:
		return cdecl.Call4(addr, a[0], a[1], a[2], a[3]), nil
	case 5:
		return cdecl.Call5(addr, a[0], a[1], a[2], a[3], a[4]), nil
	case 6:
		return cdecl.Call6(addr, a[0], a[1], a[2], a[3], a[4], a[5]), nil
	default:
		return 0, errors.New("too many arguments")
	}
}
