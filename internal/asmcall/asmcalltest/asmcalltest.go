package asmcalltest

/*
#include <stdio.h>
#ifndef WIN32
#  define __cdecl
#endif

static int __cdecl f0() { return 42; }
static int __cdecl f1(int x) { return x + 1; }
static int __cdecl f2(int x, int y) { return x / y; }
static int __cdecl f3(int x, int y, int z) { return x / y / z; }
static int __cdecl f4(int a, int b, int c, int d) {  return a / b / c / d; }
static int __cdecl f5(int a, int b, int c, int d, int e) { return e; }
static int __cdecl f6(int a, int b, int c, int d, int e, int f) { return e * f; }

static void *cdeclcall_addr_f0() { return (void *)f0; }
static void *cdeclcall_addr_f1() { return (void *)f1; }
static void *cdeclcall_addr_f2() { return (void *)f2; }
static void *cdeclcall_addr_f3() { return (void *)f3; }
static void *cdeclcall_addr_f4() { return (void *)f4; }
static void *cdeclcall_addr_f5() { return (void *)f5; }
static void *cdeclcall_addr_f6() { return (void *)f6; }
*/
import "C"

func GetCdeclcallF0Addr() uintptr {
	return uintptr(C.cdeclcall_addr_f0())
}

func GetCdeclcallF1Addr() uintptr {
	return uintptr(C.cdeclcall_addr_f1())
}

func GetCdeclcallF2Addr() uintptr {
	return uintptr(C.cdeclcall_addr_f2())
}

func GetCdeclcallF3Addr() uintptr {
	return uintptr(C.cdeclcall_addr_f3())
}

func GetCdeclcallF4Addr() uintptr {
	return uintptr(C.cdeclcall_addr_f4())
}

func GetCdeclcallF5Addr() uintptr {
	return uintptr(C.cdeclcall_addr_f5())
}

func GetCdeclcallF6Addr() uintptr {
	return uintptr(C.cdeclcall_addr_f6())
}
