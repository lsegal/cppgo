package cpp

import "unsafe"

func indirect(ptr uintptr) uintptr {
	if ptr == 0 {
		return 0
	}
	return **(**uintptr)(unsafe.Pointer(&ptr))
}
