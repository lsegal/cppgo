package dl

import (
	"syscall"
	"unsafe"

	"github.com/lsegal/cppgo/asmcall/stdcall"
)

const (
	// RTLD flags are used on POSIX systems to provide extra attributes for
	// a loaded library. See man dlopen(3) for more information on flags.
	RTLD_NOW = iota
	RTLD_LAZY
	RTLD_GLOBAL
	RTLD_LOCAL
)

func open(filename string, flags ...int) *Library {
	lib := syscall.MustLoadDLL(filename)
	return &Library{p: unsafe.Pointer(lib)}
}

func (l Library) load(procname string) *Func {
	lib := (*syscall.DLL)(l.p)
	proc := lib.MustFindProc(procname)
	return &Func{p: uintptr(proc.Addr())}
}

func (l Library) close() {
	(*syscall.DLL)(l.p).Release()
}

func (f Func) stdcall(a ...uintptr) (uintptr, error) {
	return stdcall.Call(f.p, a...)
}
