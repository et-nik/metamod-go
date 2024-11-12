package main

/*
#include <eiface.h>

extern const char* ReadString(globalvars_t *gpGlobals, int offset);
extern int MakeString(globalvars_t *gpGlobals, char *str);

*/
import "C"

// Engine Global Functions
// This functions uses global var P and P.EngineFuncs to call the engine functions

func allocString(s string) int {
	return P.EngineFuncs.AllocString(s)
}

func readString(idx int) string {
	return C.GoString(C.ReadString(P.EngineFuncs.globalVars.p, C.int(idx)))
}
