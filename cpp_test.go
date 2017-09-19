package cpp_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	cpp "github.com/lsegal/cppgo"
	"github.com/stretchr/testify/assert"
)

type lib struct {
	GetInt     func() int
	GetBool    func(b bool) bool
	FlipBool   func(b *bool)
	GetString  func() string
	GetSelf    func() *lib
	TestValues func(string, int, uintptr) bool
	Add        func(n, m int) int `call:"std"`
}

func compile() {
	_, f, _, _ := runtime.Caller(0)
	compileCmd.Dir = filepath.Join(filepath.Dir(f), "fixtures")
	b, err := compileCmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("ERROR: %v\n----\n%s\n----\n", err, string(b)))
	}
}

func TestMain(m *testing.M) {
	compile()
	load()
	defer shutdown()

	code := m.Run()
	os.Exit(code)
}

func TestCpp(t *testing.T) {
	var l lib
	o := objptr()
	cpp.ConvertRef(o, &l)

	assert.Equal(t, 42, l.GetInt())
	assert.Equal(t, "hello world", l.GetString())
	assert.True(t, l.GetBool(false))
	assert.False(t, l.GetBool(true))
	b := false
	l.FlipBool(&b)
	assert.True(t, b)
	if assert.NotNil(t, l.GetSelf()) {
		assert.Equal(t, 42, l.GetSelf().GetInt())
	}
	assert.Equal(t, true, l.TestValues("hello world", -1, o))
	assert.Equal(t, 13, l.Add(11, 2))
}
