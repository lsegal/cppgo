package cpptest

// extern void *get_object();
// extern char *get_string(void *obj);
import "C"
import "unsafe"

func GetObject() unsafe.Pointer {
	return C.get_object()
}

func GetString(obj unsafe.Pointer) string {
	return C.GoString(C.get_string(obj))
}
