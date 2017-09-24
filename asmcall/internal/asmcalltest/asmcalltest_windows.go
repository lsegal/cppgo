package asmcalltest

/*
#include <stdio.h>
static int __stdcall f0() { return 42; }
static int __stdcall f1(int x) { return x + 1; }
static int __stdcall f2(int x, int y) { return x / y; }
static int __stdcall f3(int x, int y, int z) { return x / y / z; }
static int __stdcall f4(int a, int b, int c, int d) {  return a / b / c / d; }
static int __stdcall f5(int a, int b, int c, int d, int e) { return e; }
static int __stdcall f6(int a, int b, int c, int d, int e, int f) { return e * f; }

static void *stdcall_addr_f0() { return (void *)f0; }
static void *stdcall_addr_f1() { return (void *)f1; }
static void *stdcall_addr_f2() { return (void *)f2; }
static void *stdcall_addr_f3() { return (void *)f3; }
static void *stdcall_addr_f4() { return (void *)f4; }
static void *stdcall_addr_f5() { return (void *)f5; }
static void *stdcall_addr_f6() { return (void *)f6; }

extern void init();
extern void *thiscall_addr_f0();
extern void *thiscall_addr_f1();
extern void *thiscall_addr_f2();
extern void *thiscall_addr_f3();
extern void *thiscall_addr_f4();
extern void *thiscall_addr_f5();
extern void *thiscall_addr_f6();
*/
import "C"

func init() {
	C.init()
}

func GetStdcallF0Addr() uintptr {
	return uintptr(C.stdcall_addr_f0())
}

func GetStdcallF1Addr() uintptr {
	return uintptr(C.stdcall_addr_f1())
}

func GetStdcallF2Addr() uintptr {
	return uintptr(C.stdcall_addr_f2())
}

func GetStdcallF3Addr() uintptr {
	return uintptr(C.stdcall_addr_f3())
}

func GetStdcallF4Addr() uintptr {
	return uintptr(C.stdcall_addr_f4())
}

func GetStdcallF5Addr() uintptr {
	return uintptr(C.stdcall_addr_f5())
}

func GetStdcallF6Addr() uintptr {
	return uintptr(C.stdcall_addr_f6())
}

func GetThiscallF0Addr() uintptr {
	return uintptr(C.thiscall_addr_f0())
}

func GetThiscallF1Addr() uintptr {
	return uintptr(C.thiscall_addr_f1())
}

func GetThiscallF2Addr() uintptr {
	return uintptr(C.thiscall_addr_f2())
}

func GetThiscallF3Addr() uintptr {
	return uintptr(C.thiscall_addr_f3())
}

func GetThiscallF4Addr() uintptr {
	return uintptr(C.thiscall_addr_f4())
}

func GetThiscallF5Addr() uintptr {
	return uintptr(C.thiscall_addr_f5())
}
