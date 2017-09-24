package cpp_test

import "os/exec"
import "runtime"

var (
	compileCmd = exec.Command("cmd", "/c", "build.bat")
	libname    = `fixtures\dll.dll`
)

func init() {
	if runtime.GOARCH == "386" {
		procname += "@0"
	}
}
