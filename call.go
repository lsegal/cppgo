package cpp

import "unsafe"

func (p ptr) getaddr(offset int) uintptr {
	paddr := indirect(uintptr(p)) + uintptr(offset)*unsafe.Sizeof(p)
	return indirect(paddr)
}
