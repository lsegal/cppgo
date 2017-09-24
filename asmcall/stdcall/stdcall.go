// Package stdcall implements method call ABI for the __stdcall calling
// convention.
//
// This package is only supported on Windows.
package stdcall

// Call calls a stdcall style function at memory address addr with the arguments
// list a. The function result value is returned as a uintptr to be translated
// by the caller. If the function cannot be called (usually due to an invalid
// number of argument), an error is returned.
func Call(addr uintptr, a ...uintptr) (uintptr, error) {
	return call(addr, a...)
}
