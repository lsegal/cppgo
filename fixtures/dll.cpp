#include <string.h>

#ifndef WIN32
#  define __cdecl
#  define __stdcall
#  define __thiscall
#  define __declspec(dllexport)
#endif

class Library
{
public:
  Library() {}
  virtual int __cdecl return_int() {
    return 42;
  }
  virtual bool __thiscall return_bool(bool in) {
    return !in;
  }
  virtual void flip_bool(bool *in) {
    *in = !*in;
  }
  virtual const char *return_string() {
    return "hello world";
  }
  virtual Library *self() {
    return this;
  }
  virtual bool accept_string_int_and_object(char *str, unsigned int val, Library *other) {
    return strcmp(str, "hello world") == 0 && val == (unsigned int)-1 &&
           this == other && other->return_int() == this->return_int();
  }
  virtual int __stdcall stdcall_add(int n, int m) {
    return n+m;
  }
};

extern "C" __declspec(dllexport) Library* __stdcall get_object() {
  return new Library();
}
