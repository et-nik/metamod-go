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
