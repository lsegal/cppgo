package cpp

import (
	"unsafe"

	"github.com/lsegal/cppgo/internal/fcall"
)

func (p Ptr) getaddr(offset int) uintptr {
	paddr := indirect(uintptr(p)) + uintptr(offset)*unsafe.Sizeof(p)
	return indirect(paddr)
}

func (p Ptr) call(offset int, a ...uintptr) (uintptr, error) {
	return fcall.CallCdecl(p.getaddr(offset), a...)
}

func SetECX(this uint32)
