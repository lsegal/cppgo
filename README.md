# cppgo

This library allows methods on C++ objects to be called directly from the
Go runtime without requiring cgo compilation.

**NOTE**: This library currently only works on Windows due to lack of `dlopen()`
support in the Go stdlib on Linux/Mac. It is possible to support other platforms,
but the work has not been done.

To set up a Go object that proxies method calls to a C++ object:

1. Define a Go struct type with function pointer field declarations that match
   the C++ class declaration,
2. Get the address of the C++ object in memory as a `uintptr` type,
3. Call `cpp.ConvertRef(addr, &o)` to point this proxy struct object (`o`) at
   the C++ object by its `addr`.
4. After this initial setup, the function pointers on the struct object will be
   ready to use like any Go method.

## Example Usage

The following example will call into a C++ class `Library` with a
`GetString(char *name)` object method prototype. See the "STEP X" comments
for usage guides:

```go
package main

var (
  dll = syscall.MustLoadLibrary("mylib.dll")
  // this may be "new_object@0" if your C++ compiler decides to mangle names.
  create = dll.MustFindProc("new_object")
)

// STEP 1. define our C++ proxy struct type with function pointers.
// The functions in this struct will be filled in by the `cpp.ConvertRef()`
// function, at which point this struct will proxy all method calls to the
// C++ object.
type Library struct {
  GetString(name string) string
}

func main() {
  // STEP 2. get an address for the C++ object
  // NOTE: you may need to free this later depending on call semantics.
  o, _, _ := create.Call()

  // STEP 3. point our proxy structure at the functions located in the object
  // that we got from step 2.
  if err := cpp.ConvertRef(o, &l); err != nil {
    panic(err)
  }

  // STEP 4. call the function with arguments
  fmt.Println(l.GetString("Loren"))
}

// Prints:
// Hello, Loren!
```

The C++ class for the above program could look something like:

```cpp
#include <stdio.h>

class Library {
public:
  virtual char __stdcall *GetString(char *name) {
    return sprintf("Hello, %s!", name);
  }
}

extern "C" __declspec(dllexport) Library* __stdcall new_object() {
  return new Library();
}
```

## Caveats & Gotchas

### Limited Type Conversion

You can define arbitrary function arguments and return types, but the internal
translation does not support a full range of types. Currently, only the
following types are supported:

* `uint*`, `int*`, `string`, `uintptr`

Any other type of pointer will be passed as a pointer directly, which may be
what you want, but may not be. For any type that isn't well converted by
the library, use `uintptr` to send its address.

Note that slices are not well supported due to the extra information
encoded in a Go slice.

Note also that `string` converts only to and from the `char*` C type, in other
words, C strings. The `std::cstring` or `wchar_t` types are not yet supported.

### Passing Objects as Arguments

When passing C++ objects to methods, you will want to use the `cpp.Ptr` or
`uintptr` value representing the address of the object. You cannot use the
pointer to the proxy struct, since this does not actually point to the
object, and cppgo does not know how to translate between the two.

For example, to send objects to methods, define a function and call it using
the bare address pointers:

```go
// C++ class defined as:
//
// class MyClass {
//   virtual void DoWork(MyClass *other);
// }
type MyClass struct {
  DoWork(obj uintptr)
}

func main() {
  var m1 MyClass
  var m2 MyClass
  o1 := get_object_address()
  o2 := get_object_address()

  cpp.ConvertRef(o, &m1)
  cpp.ConvertRef(o, &m2)

  // we may have m2 here, but we call DoWork()
  m1.DoWork(o2)
}
```

### Non-Virtual Functions Not Supported

This library does not yet support non-virtual functions. Only functions
defined with the `virtual` keyword are callable.

### Mac/Linux are Unsupported

As mentioned above, Mac and Linux platforms are not yet supported due to
different library loading mechanics.

## Author & License

Written by Loren Segal in 2017, licensed under MIT License (see LICENSE).
