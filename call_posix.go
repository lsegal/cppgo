// +build !windows

package cpp

func (p Ptr) stdcall(offset int, a ...uintptr) (uintptr, error) {
	// ignore stdcalls on non-Windows platform
	return p.call(offset, a...)
}
