//go:build go1.17
// +build go1.17

// Package unsafe provides wrapper of standard library unsafe package to be imported natively in Yaegi.
package unsafe

import (
	"reflect"
	"runtime"
	"unsafe"
)

func init() {
	// Add builtin functions to unsafe.
	Symbols["unsafe/unsafe"]["Add"] = reflect.ValueOf(add)
	Symbols["unsafe/unsafe"]["Slice"] = reflect.ValueOf(slice)
}

func add(ptr unsafe.Pointer, l int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(ptr) + uintptr(l))
}

type emptyInterface struct {
	_    uintptr
	word unsafe.Pointer
}

func slice(ptr interface{}, l int) interface{} {
	val := reflect.ValueOf(ptr)
	if val.Type().Kind() != reflect.Ptr {
		panic("unsafe.Slice: first argument to unsafe.Slice must be pointer")
	}
	if l < 0 {
		panic("unsafe.Slice: negative len")
	}
	if val.IsNil() && l != 0 {
		panic("unsafe.Slice: ptr is nil and len is not zero")
	}
	typ := reflect.SliceOf(val.Type().Elem())
	s := reflect.MakeSlice(typ, l, l).Interface()
	addr := (*emptyInterface)(unsafe.Pointer(&s)).word
	sh := (*reflect.SliceHeader)(addr)
	if val.IsNil() {
		sh.Data = 0 // for type []int(nil)
	} else {
		sh.Data = val.Pointer() // for type []int{}
	}
	runtime.KeepAlive(&ptr)
	return s
}
