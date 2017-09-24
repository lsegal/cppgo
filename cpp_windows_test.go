package cpp_test

import (
	"os/exec"
	"runtime"
	"syscall"
)

var (
	compileCmd = exec.Command("cmd", "/c", "build.bat")

	dll        *syscall.DLL
	get_object *syscall.Proc
)

func load() {
	dll = syscall.MustLoadDLL(`fixtures\dll.dll`)
	procname := "get_object"
	if runtime.GOARCH == "386" {
		procname += "@0"
	}
	get_object = dll.MustFindProc(procname)
}

func shutdown() {
	dll.Release()
}

func objptr() uintptr {
	o, _, _ := get_object.Call()
	return o
}
