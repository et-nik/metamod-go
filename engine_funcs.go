package main

/*
#include <eiface.h>

*/
import "C"
import "unsafe"

//var engineFuncs *C.enginefuncs_t

type EngineFuncs struct {
	p *C.enginefuncs_t
}

func (ef *EngineFuncs) PrecacheModel(name string) int {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	//return int(ef.p.pfnPrecacheModel(cs))
	return 0
}

func (ef *EngineFuncs) PrecacheSound(name string) int {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	//return int(ef.p.pfnPrecacheSound(cs))
	return 0
}
