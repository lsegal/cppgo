package cpp

import "C"
import (
	"errors"
	"reflect"
	"runtime"
	"strings"
	"unsafe"
)

var (
	errMustBeStruct = errors.New("value must be a reference to struct")
)

const (
	callCdecl = iota
	callStdcall
	callThiscall
)

func (p Ptr) convert(obj interface{}) error {
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
		f := e.Field(i)
		ft := f.Type
		if ft.Kind() != reflect.Func {
			continue
		}
		if ft.NumOut() > 1 {
			return errors.New(e.Field(i).Name + ": more than 1 return value is unsupported")
		}

		// check for call convention (only affects Windows)
		calltype := callCdecl
		if runtime.GOOS == "windows" {
			if c := f.Tag.Get("call"); strings.HasPrefix(c, "std") {
				calltype = callStdcall
			} else if c == "" {
				calltype = callThiscall
			}
		}

		vfn := reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
			gchold := make([]interface{}, len(args))
			ins := make([]uintptr, len(args)+1)
			ins[0] = uintptr(p)
			var o interface{}
			for n, arg := range args {
				ins[n+1], o = toptr(arg)
				gchold[n] = o
			}

			var out uintptr
			switch calltype {
			case callStdcall:
				out, _ = p.stdcall(i, ins...)
			case callThiscall:
				out, _ = p.thiscall(i, ins...)
			default:
				out, _ = p.cdeclcall(i, ins...)
			}

			if ft.NumOut() == 0 {
				return []reflect.Value{}
			}
			return []reflect.Value{toval(ft.Out(0), out)}
		})
		reflect.ValueOf(obj).Elem().Field(i).Set(vfn)
	}

	return nil
}

func toptr(v reflect.Value) (uintptr, interface{}) {
	switch v.Type().Kind() {
	case reflect.Uintptr:
		return uintptr(v.Interface().(uintptr)), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uintptr(v.Int()), nil
	case reflect.Bool:
		if v.Bool() {
			return 1, nil
		}
		return 0, nil
	case reflect.String:
		return strtoptr(v.String())
	default:
		return v.Pointer(), nil
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

func strtoptr(s string) (uintptr, interface{}) {
	b := []byte(s)
	return uintptr(unsafe.Pointer(&b[0])), b
}

func strtoval(p uintptr) reflect.Value {
	b := *(**[1 << 20]byte)(unsafe.Pointer(&p))
	i := 0
	for b[i] != 0 {
		i++
	}
	return reflect.ValueOf(string(b[0:i]))
}
