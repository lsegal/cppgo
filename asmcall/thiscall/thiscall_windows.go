package thiscall

import "errors"

func call0(addr uintptr, a uintptr) uintptr
func call1(addr uintptr, a uintptr, b uintptr) uintptr
func call2(addr uintptr, a uintptr, b uintptr, c uintptr) uintptr
func call3(addr uintptr, a uintptr, b uintptr, c uintptr, d uintptr) uintptr
func call4(addr uintptr, a uintptr, b uintptr, c uintptr, d uintptr, e uintptr) uintptr
func call5(addr uintptr, a uintptr, b uintptr, c uintptr, d uintptr, e uintptr, f uintptr) uintptr

func Call(addr uintptr, a ...uintptr) (uintptr, error) {
	switch l := len(a); l {
	case 0:
		return 0, errors.New("must pass this arg")
	case 1:
		return call0(addr, a[0]), nil
	case 2:
		return call1(addr, a[0], a[1]), nil
	case 3:
		return call2(addr, a[0], a[1], a[2]), nil
	case 4:
		return call3(addr, a[0], a[1], a[2], a[3]), nil
	case 5:
		return call4(addr, a[0], a[1], a[2], a[3], a[4]), nil
	case 6:
		return call5(addr, a[0], a[1], a[2], a[3], a[4], a[5]), nil
	default:
		return 0, errors.New("too many arguments")
	}
}
