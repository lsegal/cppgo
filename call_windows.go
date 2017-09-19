package cpp

import (
	"errors"
	"syscall"
)

func (p Ptr) stdcall(offset int, a ...uintptr) (uintptr, error) {
	addr := p.getaddr(offset)
	switch l := len(a); l {
	case 0:
		return ret(syscall.Syscall(addr, uintptr(l), 0, 0, 0))
	case 1:
		return ret(syscall.Syscall(addr, uintptr(l), a[0], 0, 0))
	case 2:
		return ret(syscall.Syscall(addr, uintptr(l), a[0], a[1], 0))
	case 3:
		return ret(syscall.Syscall(addr, uintptr(l), a[0], a[1], a[2]))
	case 4:
		return ret(syscall.Syscall6(addr, uintptr(l), a[0], a[1], a[2], a[3], 0, 0))
	case 5:
		return ret(syscall.Syscall6(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], 0))
	case 6:
		return ret(syscall.Syscall6(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5]))
	case 7:
		return ret(syscall.Syscall9(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], 0, 0))
	case 8:
		return ret(syscall.Syscall9(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], 0))
	case 9:
		return ret(syscall.Syscall9(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8]))
	case 10:
		return ret(syscall.Syscall12(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], 0, 0))
	case 11:
		return ret(syscall.Syscall12(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], 0))
	case 12:
		return ret(syscall.Syscall12(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11]))
	case 13:
		return ret(syscall.Syscall15(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], 0, 0))
	case 14:
		return ret(syscall.Syscall15(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], 0))
	case 15:
		return ret(syscall.Syscall15(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14]))
	default:
		return 0, errors.New("too many arguments")
	}
}

func ret(v1, v2 uintptr, err error) (uintptr, error) {
	return v1, nil
}
