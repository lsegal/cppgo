#include <string.h>

#ifdef WIN32
#define CALLFMT __stdcall
#define DLLEXPORT __declspec(dllexport)
#else
#define CALLFMT
#define DLLEXPORT
#endif

class Library
{
public:
  Library() {}
  virtual int CALLFMT return_int() {
    return 42;
  }
  virtual bool CALLFMT return_bool(bool in) {
    return !in;
  }
  virtual void CALLFMT flip_bool(bool *in) {
    *in = !*in;
  }
  virtual const char CALLFMT *return_string() {
    return "hello world";
  }
  virtual Library CALLFMT *self() {
    return this;
  }
  virtual bool CALLFMT accept_string_int_and_object(char *str, unsigned int val, Library *other) {
    return strcmp(str, "hello world") == 0 && val == (unsigned int)-1 &&
           this == other && other->return_int() == this->return_int();
  }
};

extern "C" DLLEXPORT Library* CALLFMT get_object() {
  return new Library();
}
