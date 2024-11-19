package metamod_go

/*
#include <eiface.h>

extern const char* ReadString(globalvars_t *gpGlobals, int offset);
extern int MakeString(globalvars_t *gpGlobals, char *str);

*/
import "C"

// Engine Global Functions
// This functions uses global var globalPluginState and globalPluginState.engineFuncs to call the engine functions

func allocString(s string) int {
	if globalPluginState.engineFuncs == nil {
		return 0
	}

	return globalPluginState.engineFuncs.AllocString(s)
}

func readString(idx int) string {
	if globalPluginState.engineFuncs == nil {
		return ""
	}

	return C.GoString(C.ReadString(globalPluginState.engineFuncs.globalVars.p, C.int(idx)))
}
