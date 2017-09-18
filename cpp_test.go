package cpp_test

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
	"testing"

	cpp "github.com/lsegal/cppgo"
	"github.com/stretchr/testify/assert"
)

func init() {
	_, f, _, _ := runtime.Caller(0)
	cmd := exec.Command("cmd", "/c", "build.bat")
	cmd.Dir = filepath.Join(filepath.Dir(f), "fixtures")
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("ERROR:", err, string(b))
	}
}

type lib struct {
	GetInt     func() int
	GetString  func() string
	GetSelf    func() *lib
	TestValues func(string, int, uintptr) bool
}

func TestCpp(t *testing.T) {
	dll := syscall.MustLoadDLL(`fixtures\dll.dll`)
	get_object := dll.MustFindProc("get_object@0")
	defer dll.Release()

	var l lib
	o, _, _ := get_object.Call()
	cpp.ConvertRef(o, &l)

	assert.Equal(t, 42, l.GetInt())
	assert.Equal(t, "hello world", l.GetString())
	if assert.NotNil(t, l.GetSelf()) {
		assert.Equal(t, 42, l.GetSelf().GetInt())
	}
	assert.Equal(t, true, l.TestValues("hello world", -1, o))
}
