package reflection

import "reflect"

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

type walkElements struct {
	numElements   int
	getElementVia func(int) reflect.Value
	walkFn        func(input string)
}

func (w walkElements) walkThem() {
	for i := 0; i < w.numElements; i++ {
		walk(w.getElementVia(i).Interface(), w.walkFn)
	}
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice, reflect.Array:
		walkElements{
			numElements:   val.Len(),
			getElementVia: val.Index,
			walkFn:        fn,
		}.walkThem()
	case reflect.Struct:
		walkElements{
			numElements:   val.NumField(),
			getElementVia: val.Field,
			walkFn:        fn,
		}.walkThem()
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}
}
