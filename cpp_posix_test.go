// +build !windows

package cpp_test

import "os/exec"

var (
	compileCmd = exec.Command("sh", "build.sh")
	libname    = `fixtures/dll.so`
)
