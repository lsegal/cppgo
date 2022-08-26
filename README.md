# cppgo

[![Build Status](https://img.shields.io/travis/lsegal/cppgo.svg)](https://travis-ci.org/lsegal/cppgo)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/lsegal/cppgo)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/lsegal/cppgo/blob/master/LICENSE.txt)

This library allows methods on C++ objects to be called directly from the
Go runtime without requiring cgo compilation.

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
  lib = dl.Open("mylib.dll")
  create = lib.Load("new_object")
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
  o, _ := create.Call() // o is a uintptr

  // STEP 3. point our proxy structure at the functions located in the object
  // that we got from step 2.
  if err := cpp.ConvertRef(o, &lib); err != nil {
    panic(err)
  }

  // STEP 4. call the function with arguments
  fmt.Println(lib.GetString("Loren"))

  // Clean up library
  lib.Close()
}

// Prints:
// Hello, Loren!
```

The C++ class for the above program could look something like:

```cpp
#include <stdio.h>

#ifndef WIN32
#  define __declspec(x)
#endif

class Library {
public:
  virtual char *GetString(char *name) {
    return sprintf("Hello, %s!", name);
  }
}

extern "C" __declspec(dllexport) Library* new_object() {
  return new Library();
}
```

## Getting a C++ Object Pointer

In some cases you may also need to figure out how to get access to the C++
object you want to call into. Although you may already have the object pointer
in memory, sometimes this object will come from a library call. This library
also abstracts the loading of dynamic libraries through the `dl` package:

```go
dll := dl.Open("path/to/library")
fn := dll.Load("get_object")
result, err := fn.Call()

// result now olds a uintptr to your object

// call this when you are done with the C++ object.
// this may not be at the end of the load function.
dll.Close()
```

See documentation for `dl` for more information on advanced usage.

## Using `__stdcall` & `__cdecl` Calling Convention on Windows

By default, this library will use the "default" calling convention for
C++ methods on that platform. For POSIX systems, this is `__cdecl`, but on
Windows, the default calling convention is `__thiscall`. 

In short, if you are using the default calling convention for your C++ methods,
you do not have to do anything differently from the above example. However,
if you encounter methods tagged with `__stdcall` or `__cdecl`, you can
support these conventions by adding a `call:"std"` or `call:"cdecl"`
tag on the field declaration respectively:

```go
type Library struct {
  GetID()   int    `call:"std"`   // "stdcall" is also valid here
  GetName() string `call:"cdecl"`
}
```

This will ensure that the function pointer is compatible with the library's
calling convention.

## Caveats & Gotchas

### Limited Type Conversion

You can define arbitrary function arguments and return types, but the internal
translation does not support a full range of types. Currently, only the
following types are supported:

* `uint*`, `int*`, `string`, `uintptr`

Any other type of pointer will be passed as a pointer directly, which may be
what you want, but may not be. For any type that isn't well converted by
the library, use `uintptr` to send its address. Note also that floating points
are explicitly unsupported.

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
  o1 := get_object_address() // uintptr
  o2 := get_object_address() // uintptr

  cpp.ConvertRef(o, &m1)
  cpp.ConvertRef(o, &m2)

  // we may have m2 here, but we call DoWork() with the uintptr address.
  m1.DoWork(o2)
}
```

### Non-Virtual Functions Not Supported

This library does not yet support non-virtual functions. Only functions
defined with the `virtual` keyword are callable.

## Author & License

Written by Loren Segal in 2017, licensed under MIT License (see LICENSE).
