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

int getUserMsgID(struct mutil_funcs_t *t, const char *msgname, int *size) {
	return (*t->pfnGetUserMsgID)(PLID, msgname, size);
}

const char * getUserMsgName(struct mutil_funcs_t *t, int msgid, int *size) {
	return (*t->pfnGetUserMsgName)(PLID, msgid, size);
}


*/
import "C"

type MUtilFuncs struct {
	p *C.mutil_funcs_t
}

func newMUtilFuncsFromC(p *C.mutil_funcs_t) *MUtilFuncs {
	return &MUtilFuncs{p}
}

func (m *MUtilFuncs) LogConsole(msg string) {
	if m.p == nil {
		return
	}

	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))

	C.logConsole(m.p, cs)
}

func (m *MUtilFuncs) LogError(msg string) {
	if m.p == nil {
		return
	}

	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))

	C.logError(m.p, cs)
}

func (m *MUtilFuncs) LogMessage(msg string) {
	if m.p == nil {
		return
	}

	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))

	C.logMessage(m.p, cs)
}

func (m *MUtilFuncs) LogDeveloper(msg string) {
	if m.p == nil {
		return
	}

	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))

	C.logDeveloper(m.p, cs)
}

func (m *MUtilFuncs) LogCenterSay(msg string) {
	if m.p == nil {
		return
	}

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

func (m *MUtilFuncs) GetUserMsgID(msgname string, size int) int {
	if m.p == nil {
		return 0
	}

	cs := C.CString(msgname)
	defer C.free(unsafe.Pointer(cs))

	csize := C.int(size)

	return int(C.getUserMsgID(m.p, cs, &csize))
}

// (plid_t plid, int msgid, int *size);
func (m *MUtilFuncs) GetUserMsgName(msgid int, size int) string {
	if m.p == nil {
		return ""
	}

	csize := C.int(size)

	cs := C.getUserMsgName(m.p, C.int(msgid), &csize)

	return C.GoString(cs)
}
