package cpp

type Ptr uintptr

func ConvertRef(ref uintptr, obj interface{}) error {
	return Ptr(ref).Convert(obj)
}

func (p Ptr) Convert(obj interface{}) error {
	return p.convert(obj)
}
