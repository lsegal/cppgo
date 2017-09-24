// Package cpp allows methods on C++ objects to be called directly from the
// Go runtime without requiring cgo compilation.
//
// For more information on how to use this library, see the project README:
// https://github.com/lsegal/cppgo/blob/master/README.md
package cpp

// ConvertRef converts a C++ object ref into a wrapper obj that can call
// methods on the reference object. The obj interface type should be a struct
// containing function pointers matching the interface of the C++ class.
// For example, given the following class:
//
//		class Math {
//		public:
//			int multiply(int value, int times);
//		}
//
//
// You might create a struct type Math as follows:
//
//		type Math struct {
//			Multiply func(value, times int) int
//		}
//
// You would then call ConvertRef with a pointer to this structure.
func ConvertRef(ref uintptr, obj interface{}) error {
	return ptr(ref).convert(obj)
}
