package dl

/*
#include <dlfcn.h>
#include <stdlib.h>

typedef void *(fn0)();
typedef void *(fn1)(void *a1);
typedef void *(fn2)(void *a1, void *a2);
typedef void *(fn3)(void *a1, void *a2, void *a3);
typedef void *(fn4)(void *a1, void *a2, void *a3, void *a4);
typedef void *(fn5)(void *a1, void *a2, void *a3, void *a4, void *a5);
typedef void *(fn6)(void *a1, void *a2, void *a3, void *a4, void *a5, void *a6);

static void *call0(void *addr) { return ((fn0*)addr)(); }
static void *call1(void *addr, void *a1) { return ((fn1*)addr)(a1); }
static void *call2(void *addr, void *a1, void *a2) { return ((fn2*)addr)(a1, a2); }
static void *call3(void *addr, void *a1, void *a2, void *a3) { return ((fn3*)addr)(a1, a2, a3); }
static void *call4(void *addr, void *a1, void *a2, void *a3, void *a4) { return ((fn4*)addr)(a1, a2, a3, a4); }
static void *call5(void *addr, void *a1, void *a2, void *a3, void *a4, void *a5) { return ((fn5*)addr)(a1, a2, a3, a4, a5); }
static void *call6(void *addr, void *a1, void *a2, void *a3, void *a4, void *a5, void *a6) { return ((fn6*)addr)(a1, a2, a3, a4, a5, a6); }
*/
import "C"
import (
	"errors"
	"unsafe"
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
	p unsafe.Pointer
}

func FuncAt(addr uintptr) *Func {
	return &Func{p: unsafe.Pointer(*(**uintptr)(unsafe.Pointer(&addr)))}
}

func (f Func) Call(a ...uintptr) (uintptr, error) {
	switch len(a) {
	case 0:
		return uintptr(C.call0(f.p)), nil
	case 1:
		return uintptr(C.call1(f.p, ref(a[0]))), nil
	case 2:
		return uintptr(C.call2(f.p, ref(a[0]), ref(a[1]))), nil
	case 3:
		return uintptr(C.call3(f.p, ref(a[0]), ref(a[1]), ref(a[2]))), nil
	case 4:
		return uintptr(C.call4(f.p, ref(a[0]), ref(a[1]), ref(a[2]), ref(a[3]))), nil
	case 5:
		return uintptr(C.call5(f.p, ref(a[0]), ref(a[1]), ref(a[2]), ref(a[3]), ref(a[4]))), nil
	case 6:
		return uintptr(C.call6(f.p, ref(a[0]), ref(a[1]), ref(a[2]), ref(a[3]), ref(a[4]), ref(a[5]))), nil
	default:
		return 0, errors.New("too many arguments")
	}
}

func ref(u uintptr) unsafe.Pointer {
	return unsafe.Pointer(*(**uintptr)(unsafe.Pointer(&u)))
}

func Open(filename string, flags int) *Library {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return &Library{p: C.dlopen(cfilename, C.int(flags))}
}

func (l Library) Sym(sym string) *Func {
	csym := C.CString(sym)
	defer C.free(unsafe.Pointer(csym))
	return &Func{p: C.dlsym(l.p, csym)}
}

func (l Library) Close() {
	C.dlclose(l.p)
}
