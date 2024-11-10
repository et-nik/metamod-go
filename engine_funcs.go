package main

/*
#include <eiface.h>

#include <stdlib.h>
#include <stdio.h>

#define MAX_SERVER_COMMAND_CALLBACKS 128
#define SERVER_COMMAND_CALLBACKS_CATEGORY "server_commands"

extern const char* ReadString(globalvars_t *gpGlobals, int offset);

extern void callGoFunction(void *f, int argc, char **argv);
extern void getGoCallback(char *category, char *v, void **f);
extern void setGoCallback(char *category, char *v, void *f);

void* engineFuncsGoCallbacks[MAX_SERVER_COMMAND_CALLBACKS];
int engineFuncsGoCallbacksCount = 0;

struct enginefuncs_s *engineFuncs;

typedef void (*server_command_callback_t)(void);

// Engine funcs
int engineFuncsPrecacheModel(struct enginefuncs_s *t, char *s) {
	return (*t->pfnPrecacheModel)(s);
}

int engineFuncsPrecacheSound(struct enginefuncs_s *t, char *s) {
	return (*t->pfnPrecacheSound)(s);
}

const char* engineFuncsCmd_Args(struct enginefuncs_s *t) {
	return (*t->pfnCmd_Args)();
}

const char* engineFuncsCmd_Argv(struct enginefuncs_s *t, int argc) {
	return (*t->pfnCmd_Argv)(argc);
}

int engineFuncsCmd_Argc(struct enginefuncs_s *t) {
	return (*t->pfnCmd_Argc)();
}

void engineFuncsAddServerCommand(struct enginefuncs_s *t, char *cmd_name, void *f) {
	engineFuncs = t;

	setGoCallback(SERVER_COMMAND_CALLBACKS_CATEGORY, cmd_name, f);

	void ff(void) {
		void *f;
		getGoCallback(SERVER_COMMAND_CALLBACKS_CATEGORY, engineFuncsCmd_Argv(engineFuncs, 0), &f);

		int argc = engineFuncsCmd_Argc(engineFuncs);

		char **argv = (char**)malloc(argc * sizeof(char*));

		for (int i = 0; i < argc; i++) {
			argv[i] = engineFuncsCmd_Argv(engineFuncs, i);
		}

		callGoFunction(f, argc, argv);
		free(argv);
    }

	(*t->pfnAddServerCommand)(cmd_name, &ff);
}

edict_t* engineFuncsEntityOfEntIndex(struct enginefuncs_s *t, int index) {
	return (*t->pfnPEntityOfEntIndex)(index);
}

void engineMessageBegin(struct enginefuncs_s *t, int msg_dest, int msg_type, const float *pOrigin, edict_t *ed) {
	(*t->pfnMessageBegin)(msg_dest, msg_type, pOrigin, ed);
}

void engineMessageEnd(struct enginefuncs_s *t) {
	(*t->pfnMessageEnd)();
}

void engineWriteByte(struct enginefuncs_s *t, int i) {
	(*t->pfnWriteByte)(i);
}

void engineWriteChar(struct enginefuncs_s *t, int i) {
	(*t->pfnWriteChar)(i);
}

void engineWriteShort(struct enginefuncs_s *t, int i) {
	(*t->pfnWriteShort)(i);
}

void engineWriteLong(struct enginefuncs_s *t, int i) {
	(*t->pfnWriteLong)(i);
}

void engineWriteAngle(struct enginefuncs_s *t, float f) {
	(*t->pfnWriteAngle)(f);
}

void engineWriteCoord(struct enginefuncs_s *t, float f) {
	(*t->pfnWriteCoord)(f);
}

void engineWriteString(struct enginefuncs_s *t, const char *s) {
	(*t->pfnWriteString)(s);
}

void engineWriteEntity(struct enginefuncs_s *t, int i) {
	(*t->pfnWriteEntity)(i);
}

*/
import "C"
import "C"

import (
	"unsafe"
)

type EngineFuncs struct {
	p          *C.enginefuncs_t
	globalVars *GlobalVars
}

func NewEngineFuncs(p *C.enginefuncs_t, globalVars *GlobalVars) *EngineFuncs {
	return &EngineFuncs{
		p:          p,
		globalVars: globalVars,
	}
}

func (ef *EngineFuncs) PrecacheModel(name string) int {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	return int(C.engineFuncsPrecacheModel(ef.p, cs))
}

func (ef *EngineFuncs) PrecacheSound(name string) int {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	return int(C.engineFuncsPrecacheSound(ef.p, cs))
}

func (ef *EngineFuncs) AddServerCommand(name string, callback func(int, ...string)) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	f := unsafe.Pointer(&callback)

	C.engineFuncsAddServerCommand(ef.p, cs, f)
}

func (ef *EngineFuncs) EntityOfEntIndex(index int) *Edict {
	edict := C.engineFuncsEntityOfEntIndex(ef.p, C.int(index))

	return EdictFromC(ef.globalVars.p, edict)
}

func (ef *EngineFuncs) MessageBegin(
	msgDest int,
	msgType int,
	pOrigin float32,
	edict *Edict,
) {
	C.engineMessageBegin(
		ef.p,
		C.int(msgDest),
		C.int(msgType),
		(*C.float)(&pOrigin),
		edict.p,
	)
}

func (ef *EngineFuncs) MessageEnd() {
	C.engineMessageEnd(ef.p)
}

func (ef *EngineFuncs) MessageWriteByte(i int) {
	C.engineWriteByte(ef.p, C.int(i))
}

func (ef *EngineFuncs) MessageWriteChar(i int) {
	C.engineWriteChar(ef.p, C.int(i))
}

func (ef *EngineFuncs) MessageWriteShort(i int) {
	C.engineWriteShort(ef.p, C.int(i))
}

func (ef *EngineFuncs) MessageWriteLong(i int) {
	C.engineWriteLong(ef.p, C.int(i))
}

func (ef *EngineFuncs) MessageWriteAngle(f float32) {
	C.engineWriteAngle(ef.p, C.float(f))
}

func (ef *EngineFuncs) MessageWriteCoord(f float32) {
	C.engineWriteCoord(ef.p, C.float(f))
}

func (ef *EngineFuncs) MessageWriteString(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	C.engineWriteString(ef.p, cs)
}

func (ef *EngineFuncs) MessageWriteEntity(e *Edict) {
	C.engineWriteEntity(ef.p, C.int(e.SerialNumber()))
}
