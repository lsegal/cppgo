// Package dl implements support for loading dynamic libraries at runtime
// without cgo.
package dl

import (
	"unsafe"

	"github.com/lsegal/cppgo/asmcall/cdecl"
)

// Library represents a dynamic library object.
type Library struct {
	p unsafe.Pointer
}

// Open will open a dynamic library by the given filename on disk and optional
// flags which are binary ANDed together. If no flags are passed, RTLD_NOW
// is assumed on POSIX systems.
//
// Note: The flags parameter is ignored on the Windows platform.
func Open(filename string, flags ...int) *Library {
	return open(filename, flags...)
}

// Load will search for a function by procname and return it if available.
// If the function is not available, this function may panic.
func (l Library) Load(procname string) *Func {
	return l.load(procname)
}

// Close closes a library and frees its resources. Calling any Func objects
// returned by the library after Close() will cause a panic.
func (l Library) Close() {
	l.close()
}

// Func represents a function exported by a library.
type Func struct {
	p uintptr
}

// FuncAt returns a bare function object at address addr. Use this function
// if you have loaded up a library using another mechanism but still want
// to take advantage of the Call abstraction method.
func FuncAt(addr uintptr) *Func {
	return &Func{p: addr}
}

// Call calls a function with argument list a. It returns the result as a
// uintptr, and a possible error if an unsupported number of arguments were
// passed.
func (f Func) Call(a ...uintptr) (uintptr, error) {
	return cdecl.Call(f.p, a...)
}

// Stdcall calls a function using the __stdcall calling convention with
// argument list a. It returns the result as a uintptr, and a possible error
// if an unsupported number of arguments were passed.
//
// Note: On non-Windows systems, this function is equivalent to Call().
func (f Func) Stdcall(a ...uintptr) (uintptr, error) {
	return f.stdcall(a...)
}
