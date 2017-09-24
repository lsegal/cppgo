package dl

import (
	"syscall"
	"unsafe"

	"github.com/lsegal/cppgo/asmcall/stdcall"
)

const (
	RTLD_NOW = iota
	RTLD_LAZY
	RTLD_GLOBAL
	RTLD_LOCAL
)

func Open(filename string, flags ...int) *Library {
	lib := syscall.MustLoadDLL(filename)
	return &Library{p: unsafe.Pointer(lib)}
}

func (l Library) Load(procname string) *Func {
	lib := (*syscall.DLL)(l.p)
	proc := lib.MustFindProc(procname)
	return &Func{p: uintptr(proc.Addr())}
}

func (l Library) Close() {
	(*syscall.DLL)(l.p).Release()
}

func (f Func) Stdcall(a ...uintptr) (uintptr, error) {
	return stdcall.Call(f.p, a...)
}
