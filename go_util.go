package metamod_go

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

type callbackFn func(int, ...string)

var callbacks = make(map[string]map[string]callbackFn)

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

	*f = unsafe.Pointer(&val)
}

func setGoCallback(category string, key string, f callbackFn) {
	if callbacks[category] == nil {
		callbacks[category] = make(map[string]callbackFn, 64)
	}

	callbacks[category][key] = f
}
