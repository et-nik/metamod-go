package main

import "C"

import (
	"unsafe"
)

//export callGoFunction
func callGoFunction(f unsafe.Pointer, argc C.int, argv **C.char) {
	gargc := int(argc)
	gargv := make([]string, int(argc))

	cstrings := (*[1 << 28]*C.char)(unsafe.Pointer(argv))[:argc:argc]

	for i, cString := range cstrings {
		gargv[i] = C.GoString(cString)
	}

	(*(*func(int, ...string))(f))(gargc, gargv...)
}

var callbacks = make(map[string]map[string]unsafe.Pointer)

//export getGoCallback
func getGoCallback(category *C.char, v *C.char, f *unsafe.Pointer) {
	cat := C.GoString(category)
	key := C.GoString(v)

	val, ok := callbacks[cat][key]
	if !ok {
		emptyFn := func(int, ...string) {}
		*f = unsafe.Pointer(&emptyFn)

		return
	}

	*f = val
}

//export setGoCallback
func setGoCallback(category *C.char, v *C.char, f unsafe.Pointer) {
	cat := C.GoString(category)
	key := C.GoString(v)

	if callbacks[cat] == nil {
		callbacks[cat] = make(map[string]unsafe.Pointer, 64)
	}

	callbacks[cat][key] = f
}
