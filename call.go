package cpp

import "unsafe"

func (p Ptr) getaddr(offset int) uintptr {
	paddr := indirect(uintptr(p)) + uintptr(offset)*unsafe.Sizeof(p)
	return indirect(paddr)
}
