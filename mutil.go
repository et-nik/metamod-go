package main

import (
	"fmt"
	"unsafe"
)

/*
#include "metamod/mutil.h"
#include "metamod/plinfo.h"

void logConsole(struct mutil_funcs_t *t, const char *msg) {
	(*t->pfnLogConsole)(PLID, msg);
};

void logMessage(struct mutil_funcs_t *t, const char *msg) {
	(*t->pfnLogMessage)(PLID, msg);
};

void logError(struct mutil_funcs_t *t, const char *msg) {
	(*t->pfnLogError)(PLID, msg);
};

void logDeveloper(struct mutil_funcs_t *t, const char *msg) {
	(*t->pfnLogDeveloper)(PLID, msg);
};

void logCenterSay(struct mutil_funcs_t *t, const char *msg) {
	(*t->pfnCenterSay)(PLID, msg);
};


*/
import "C"

type MUtilFuncs struct {
	p *C.mutil_funcs_t
}

func (m *MUtilFuncs) LogConsole(msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))

	C.logConsole(m.p, cs)
}

func (m *MUtilFuncs) LogError(msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))

	C.logError(m.p, cs)
}

func (m *MUtilFuncs) LogMessage(msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))

	C.logMessage(m.p, cs)
}

func (m *MUtilFuncs) LogDeveloper(msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))

	C.logDeveloper(m.p, cs)
}

func (m *MUtilFuncs) LogCenterSay(msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))

	C.logCenterSay(m.p, cs)
}

func (m *MUtilFuncs) LogConsolef(format string, args ...interface{}) {
	m.LogConsole(fmt.Sprintf(format, args...))
}

func (m *MUtilFuncs) LogErrorf(format string, args ...interface{}) {
	m.LogError(fmt.Sprintf(format, args...))
}

func (m *MUtilFuncs) LogMessagef(format string, args ...interface{}) {
	m.LogMessage(fmt.Sprintf(format, args...))
}

func (m *MUtilFuncs) LogDeveloperf(format string, args ...interface{}) {
	m.LogDeveloper(fmt.Sprintf(format, args...))
}

func (m *MUtilFuncs) LogCenterSayf(format string, args ...interface{}) {
	m.LogCenterSay(fmt.Sprintf(format, args...))
}
