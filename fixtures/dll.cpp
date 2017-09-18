#include <string.h>

class Library
{
public:
  Library() {}
  virtual int __stdcall return_int() {
    return 42;
  }
  virtual const char __stdcall *return_string() {
    return "hello world";
  }
  virtual Library __stdcall *self() {
    return this;
  }
  virtual bool __stdcall accept_string_int_and_object(char *str, unsigned int val, Library *other) {
    return strcmp(str, "hello world") == 0 && val == (unsigned int)-1 &&
           this == other && other->return_int() == this->return_int();
  }
};

extern "C" __declspec(dllexport)
Library* __stdcall get_object() {
  return new Library();
}
