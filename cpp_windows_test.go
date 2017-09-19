package cpp_test

import (
	"os/exec"
	"syscall"
)

var (
	compileCmd = exec.Command("cmd", "/c", "build.bat")

	dll        *syscall.DLL
	get_object *syscall.Proc
)

func load() {
	dll = syscall.MustLoadDLL(`fixtures\dll.dll`)
	get_object = dll.MustFindProc("get_object@0")
}

func shutdown() {
	dll.Release()
}

func objptr() uintptr {
	o, _, _ := get_object.Call()
	return o
}
