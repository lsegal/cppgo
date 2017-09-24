class Test {
public:
	int f0() { return 42; }
	int f1(int x) { return x + 1; }
	int f2(int x, int y) { return x / y; }
	int f3(int x, int y, int z) { return x / y / z; }
	int f4(int a, int b, int c, int d) {  return a / b / c / d; }
	int f5(int a, int b, int c, int d, int e) { return e; }
	int f6(int a, int b, int c, int d, int e, int f) { return e * f; }
};

#pragma GCC diagnostic ignored "-Wpmf-conversions"
static Test *t;
extern "C" {
	void init() { t = new Test(); }
	void *thiscall_addr_f0() { return (void *)&t->f0; }
	void *thiscall_addr_f1() { return (void *)&t->f1; }
	void *thiscall_addr_f2() { return (void *)&t->f2; }
	void *thiscall_addr_f3() { return (void *)&t->f3; }
	void *thiscall_addr_f4() { return (void *)&t->f4; }
	void *thiscall_addr_f5() { return (void *)&t->f5; }
	void *thiscall_addr_f6() { return (void *)&t->f6; }
}
