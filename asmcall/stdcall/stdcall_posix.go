// +build !windows

package stdcall

import "errors"

func call(addr uintptr, a ...uintptr) (uintptr, error) {
	return 0, errors.New("unsupported on this platform")
}
