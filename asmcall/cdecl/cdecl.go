// Package cdecl implements method call ABI for the POSIX C/C++ calling
// convention.
//
// Note that this package does not rely on cgo to implement calling, so a
// compiler is not needed to call into C functions using this library.
package cdecl

import "errors"

func Call(addr uintptr, a ...uintptr) (uintptr, error) {
	switch l := len(a); l {
	case 0:
		return call0(addr), nil
	case 1:
		return call1(addr, a[0]), nil
	case 2:
		return call2(addr, a[0], a[1]), nil
	case 3:
		return call3(addr, a[0], a[1], a[2]), nil
	case 4:
		return call4(addr, a[0], a[1], a[2], a[3]), nil
	case 5:
		return call5(addr, a[0], a[1], a[2], a[3], a[4]), nil
	case 6:
		return call6(addr, a[0], a[1], a[2], a[3], a[4], a[5]), nil
	default:
		return 0, errors.New("too many arguments")
	}
}
