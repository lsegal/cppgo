package cpp_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	cpp "github.com/lsegal/cppgo"
	"github.com/lsegal/cppgo/dl"
	"github.com/lsegal/cppgo/internal/cpptest"
	"github.com/stretchr/testify/assert"
)

var (
	dll        *dl.Library
	get_object *dl.Func
	procname   = "get_object"
)

type lib struct {
	GetInt     func() int        `call:"cdecl"`
	GetBool    func(b bool) bool `call:"this"`
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

func load() {
	dll = dl.Open(libname, dl.RTLD_NOW)
	get_object = dll.Load(procname)
}

func shutdown() {
	dll.Close()
}

func objref() uintptr {
	o, _ := get_object.Call()
	return o
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
	o := objref()
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

func ExampleConvertRef() {
	/*
		// with Go type `lib`:
		type lib struct {
			GetString func() string
		}

		// and C++ class Library:
		class Library {
		public:
			const char *get_string();
		}
	*/
	var l lib
	err := cpp.ConvertRef(objref(), &l)
	if err != nil {
		return
	}
	fmt.Println(l.GetString())
	// Output: hello world
}

func BenchmarkCppGo(b *testing.B) {
	var l lib
	cpp.ConvertRef(objref(), &l)

	for i := 0; i < b.N; i++ {
		l.GetString()
	}
}

func BenchmarkCgo(b *testing.B) {
	o := cpptest.GetObject()
	for i := 0; i < b.N; i++ {
		cpptest.GetString(o)
	}
}
