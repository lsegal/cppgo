package dl

// #include <dlfcn.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"

	"github.com/lsegal/cppgo/internal/asmcall/cdecl"
)

var (
	RTLD_NOW    = int(C.RTLD_NOW)
	RTLD_LAZY   = int(C.RTLD_LAZY)
	RTLD_GLOBAL = int(C.RTLD_GLOBAL)
	RTLD_LOCAL  = int(C.RTLD_LOCAL)
)

type Library struct {
	p unsafe.Pointer
}

type Func struct {
	p uintptr
}

func FuncAt(addr uintptr) *Func {
	return &Func{p: addr}
}

func (f Func) Call(a ...uintptr) (uintptr, error) {
	return cdecl.Call(f.p, a...)
}

func Open(filename string, flags int) *Library {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return &Library{p: C.dlopen(cfilename, C.int(flags))}
}

func (l Library) Sym(sym string) *Func {
	csym := C.CString(sym)
	defer C.free(unsafe.Pointer(csym))
	return &Func{p: uintptr(C.dlsym(l.p, csym))}
}

func (l Library) Close() {
	C.dlclose(l.p)
}
