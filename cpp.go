package cpp

import (
	"errors"
	"reflect"
	"syscall"
	"unsafe"
)

var (
	errMustBeStruct = errors.New("value must be a reference to struct")
)

type Ptr uintptr

func ConvertRef(ref uintptr, obj interface{}) error {
	return Ptr(ref).Convert(obj)
}

func (p Ptr) Convert(obj interface{}) error {
	if p == 0 {
		return errors.New("invalid address")
	}

	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr {
		return errMustBeStruct
	}
	e := t.Elem()
	if e.Kind() != reflect.Struct {
		return errMustBeStruct
	}

	for idx := 0; idx < e.NumField(); idx++ {
		i := idx
		ft := e.Field(i).Type
		if ft.Kind() != reflect.Func {
			continue
		}
		if ft.NumOut() > 1 {
			return errors.New(e.Field(i).Name + ": more than 1 return value is unsupported")
		}

		vfn := reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
			ins := make([]uintptr, len(args)+1)
			ins[0] = uintptr(p)
			for n, arg := range args {
				ins[n+1] = toptr(arg)
			}
			out, _ := p.call(i, ins...)

			if ft.NumOut() == 0 {
				return []reflect.Value{}
			}
			return []reflect.Value{toval(ft.Out(0), out)}
		})
		reflect.ValueOf(obj).Elem().Field(i).Set(vfn)
	}

	return nil
}

func toptr(v reflect.Value) uintptr {
	switch v.Type().Kind() {
	case reflect.Uintptr:
		return uintptr(v.Interface().(uintptr))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uintptr(v.Int())
	case reflect.Bool:
		if v.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		return strtoptr(v.String())
	default:
		return indirect(v.Pointer())
	}
}

func isCppObj(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return false
	}
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.Kind() == reflect.Func {
			return true
		}
	}
	return false
}

func toval(t reflect.Type, p uintptr) reflect.Value {
	if isCppObj(t) {
		v := reflect.New(t.Elem())
		err := ConvertRef(p, v.Interface())
		if err != nil {
			return reflect.Zero(t)
		}
		return v
	}

	switch t.Kind() {
	case reflect.String:
		return strtoval(p)
	default:
		return reflect.NewAt(t, unsafe.Pointer(&p)).Elem()
	}
}

func strtoptr(s string) uintptr {
	b := []byte(s)
	return uintptr(unsafe.Pointer(&b[0]))
}

func strtoval(p uintptr) reflect.Value {
	b := *(**[1 << 20]byte)(unsafe.Pointer(&p))
	i := 0
	for b[i] != 0 {
		i++
	}
	return reflect.ValueOf(string(b[0:i]))
}

func (p Ptr) call(offset int, a ...uintptr) (uintptr, error) {
	paddr := indirect(uintptr(p)) + uintptr(offset)*unsafe.Sizeof(p)
	addr := indirect(paddr)
	switch l := len(a); l {
	case 0:
		return ret(syscall.Syscall(addr, uintptr(l), 0, 0, 0))
	case 1:
		return ret(syscall.Syscall(addr, uintptr(l), a[0], 0, 0))
	case 2:
		return ret(syscall.Syscall(addr, uintptr(l), a[0], a[1], 0))
	case 3:
		return ret(syscall.Syscall(addr, uintptr(l), a[0], a[1], a[2]))
	case 4:
		return ret(syscall.Syscall6(addr, uintptr(l), a[0], a[1], a[2], a[3], 0, 0))
	case 5:
		return ret(syscall.Syscall6(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], 0))
	case 6:
		return ret(syscall.Syscall6(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5]))
	case 7:
		return ret(syscall.Syscall9(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], 0, 0))
	case 8:
		return ret(syscall.Syscall9(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], 0))
	case 9:
		return ret(syscall.Syscall9(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8]))
	case 10:
		return ret(syscall.Syscall12(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], 0, 0))
	case 11:
		return ret(syscall.Syscall12(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], 0))
	case 12:
		return ret(syscall.Syscall12(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11]))
	case 13:
		return ret(syscall.Syscall15(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], 0, 0))
	case 14:
		return ret(syscall.Syscall15(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], 0))
	case 15:
		return ret(syscall.Syscall15(addr, uintptr(l), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14]))
	default:
		return 0, errors.New("too many arguments")
	}
}

func ret(v1, v2 uintptr, err error) (uintptr, error) {
	return v1, nil
}

func indirect(ptr uintptr) uintptr {
	if ptr == 0 {
		return 0
	}
	return **(**uintptr)(unsafe.Pointer(&ptr))
}
