class BenchmarkLibrary
{
public:
  BenchmarkLibrary() {}
  virtual const char *get_string() {
    return "hello world";
  }
};

extern "C" const BenchmarkLibrary *get_object() {
  return new BenchmarkLibrary();
}

extern "C" const char *get_string(BenchmarkLibrary *obj) {
  return obj->get_string();
}
