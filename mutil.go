package main

import "C"
import "unsafe"

/*
#include "metamod/mutil.h"

void logConsole(struct mutil_funcs_t *t, plid_t plid, const char *msg) {
	(*t->pfnLogConsole)(plid, msg);
};

*/
import "C"

type MUtilFuncs struct {
	info *PluginInfo
	p    *C.mutil_funcs_t
}

func (muf *MUtilFuncs) LogConsole(msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))

	C.logConsole(muf.p, muf.info.ToC(), cs)
}
