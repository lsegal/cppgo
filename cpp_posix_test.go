// +build linux darwin

package cpp_test

import (
	"os/exec"

	"github.com/lsegal/cppgo/internal/dl"
)

var (
	compileCmd = exec.Command("sh", "build.sh")

	dll        *dl.Library
	get_object *dl.Func
)

func load() {
	dll = dl.Open("fixtures/dll.so", dl.RTLD_NOW)
	get_object = dll.Sym("get_object")
}

func shutdown() {
	dll.Close()
}

func objptr() uintptr {
	o, _ := get_object.Call()
	return o
}
