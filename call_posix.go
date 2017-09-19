// +build linux darwin

package cpp

import "github.com/lsegal/cppgo/internal/dl"

func (p Ptr) call(offset int, a ...uintptr) (uintptr, error) {
	return dl.FuncAt(p.getaddr(offset)).Call(a...)
}
