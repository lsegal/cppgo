package fcall

/*
typedef unsigned long uintptr;
typedef uintptr (fn0)();
typedef uintptr (fn1)(uintptr a1);
typedef uintptr (fn2)(uintptr a1, uintptr a2);
typedef uintptr (fn3)(uintptr a1, uintptr a2, uintptr a3);
typedef uintptr (fn4)(uintptr a1, uintptr a2, uintptr a3, uintptr a4);
typedef uintptr (fn5)(uintptr a1, uintptr a2, uintptr a3, uintptr a4, uintptr a5);
typedef uintptr (fn6)(uintptr a1, uintptr a2, uintptr a3, uintptr a4, uintptr a5, uintptr a6);

static uintptr call0(uintptr addr) { return ((fn0*)addr)(); }
static uintptr call1(uintptr addr, uintptr a1) { return ((fn1*)addr)(a1); }
static uintptr call2(uintptr addr, uintptr a1, uintptr a2) { return ((fn2*)addr)(a1, a2); }
static uintptr call3(uintptr addr, uintptr a1, uintptr a2, uintptr a3) { return ((fn3*)addr)(a1, a2, a3); }
static uintptr call4(uintptr addr, uintptr a1, uintptr a2, uintptr a3, uintptr a4) { return ((fn4*)addr)(a1, a2, a3, a4); }
static uintptr call5(uintptr addr, uintptr a1, uintptr a2, uintptr a3, uintptr a4, uintptr a5) { return ((fn5*)addr)(a1, a2, a3, a4, a5); }
static uintptr call6(uintptr addr, uintptr a1, uintptr a2, uintptr a3, uintptr a4, uintptr a5, uintptr a6) { return ((fn6*)addr)(a1, a2, a3, a4, a5, a6); }
*/
import "C"
import "errors"

func CallCdecl(addr uintptr, a ...uintptr) (uintptr, error) {
	paddr := ref(addr)
	switch len(a) {
	case 0:
		return uintptr(C.call0(paddr)), nil
	case 1:
		return uintptr(C.call1(paddr, ref(a[0]))), nil
	case 2:
		return uintptr(C.call2(paddr, ref(a[0]), ref(a[1]))), nil
	case 3:
		return uintptr(C.call3(paddr, ref(a[0]), ref(a[1]), ref(a[2]))), nil
	case 4:
		return uintptr(C.call4(paddr, ref(a[0]), ref(a[1]), ref(a[2]), ref(a[3]))), nil
	case 5:
		return uintptr(C.call5(paddr, ref(a[0]), ref(a[1]), ref(a[2]), ref(a[3]), ref(a[4]))), nil
	case 6:
		return uintptr(C.call6(paddr, ref(a[0]), ref(a[1]), ref(a[2]), ref(a[3]), ref(a[4]), ref(a[5]))), nil
	default:
		return 0, errors.New("too many arguments")
	}
}

func ref(u uintptr) C.uintptr {
	return C.uintptr(u)
}
