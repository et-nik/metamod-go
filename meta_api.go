package main

/*
#cgo CPPFLAGS: -I${SRCDIR}/hlsdk/dlls -I${SRCDIR}/hlsdk/engine -I${SRCDIR}/hlsdk/pm_shared -I${SRCDIR}/hlsdk/common -include ${SRCDIR}/hlsdk/public/basetypes.h -include ${SRCDIR}/hlsdk/common/const.h -include ${SRCDIR}/hlsdk/engine/edict.h

#include <stdio.h>

#include <eiface.h>

#include "metamod/index.h"

int GetNewDLLFunctions(NEW_DLL_FUNCTIONS *pNewFunctionTable, int *interfaceVersion);
int GetEntityAPI2(DLL_FUNCTIONS* pFunctionTable, int* interfaceVersion);
int GetEntityAPI2_Post(DLL_FUNCTIONS* pFunctionTable, int* interfaceVersion);
int GetEngineFunctions(enginefuncs_t *pengfuncsFromEngine, int *interfaceVersion);
int GetEngineFunctions_Post(enginefuncs_t *pengfuncsFromEngine, int *interfaceVersion);

// Lib functions
extern void SetDLLFunctions(DLL_FUNCTIONS *pFunctionTable);
extern void SetDLLFunctionsPost(DLL_FUNCTIONS *pFunctionTable);
extern void SetHooks(enginefuncs_t *pengfuncsFromEngine);

*/
import "C"

var pluginInfo = &PluginInfo{
	InterfaceVersion: MetaInterfaceVersion,
	Name:             "Metamod Go",
	Version:          "1.0",
	Date:             "2024-11-08",
	Author:           "KNiK",
	Url:              "https://github.com/et-nik/metamod-go",
	LogTag:           "GoPlugin",
	Loadable:         PluginLoadTimeAnyTime,
	Unloadable:       PluginLoadTimeAnyTime,
}

//export Meta_Init
func Meta_Init() {
	if globalPluginState.metaCallbacks != nil && globalPluginState.metaCallbacks.MetaInit != nil {
		globalPluginState.metaCallbacks.MetaInit()
	}

	globalPluginState.timelineStatus = statusLibLoaded
}

//export Meta_Query
func Meta_Query(interfaceVersion *C.char, plinfo **C.plugin_info_t, pMetaUtilFuncs *C.mutil_funcs_t) C.int {
	*plinfo = pluginInfo.toCPluginInfo()

	setCGlobalPluginInfo(pluginInfo)

	globalPluginState.metaUtilFuncs = newMUtilFuncsFromC(pMetaUtilFuncs)

	if globalPluginState.metaCallbacks != nil && globalPluginState.metaCallbacks.MetaQuery != nil {
		result := globalPluginState.metaCallbacks.MetaQuery()

		return C.int(result)
	}

	globalPluginState.metaUtilFuncs.LogDeveloper("Meta_Query called")

	globalPluginState.timelineStatus = statusMetaQueried

	return 1
}

//export Meta_Attach
func Meta_Attach(now C.int, pFunctionTable *C.META_FUNCTIONS, pMGlobals *C.meta_globals_t, pGamedllFuncs *C.void) C.int {
	pFunctionTable.pfnGetEntityAPI2 = C.GETENTITYAPI2_FN(C.GetEntityAPI2)
	pFunctionTable.pfnGetEntityAPI2_Post = C.GETENTITYAPI2_FN(C.GetEntityAPI2_Post)
	pFunctionTable.pfnGetNewDLLFunctions = C.GETNEWDLLFUNCTIONS_FN(C.GetNewDLLFunctions)
	pFunctionTable.pfnGetEngineFunctions = C.GET_ENGINE_FUNCTIONS_FN(C.GetEngineFunctions)
	pFunctionTable.pfnGetEngineFunctions_Post = C.GET_ENGINE_FUNCTIONS_FN(C.GetEngineFunctions_Post)

	globalPluginState.metaGlobals = MetaGlobalsFromC(pMGlobals)

	if globalPluginState.metaCallbacks != nil && globalPluginState.metaCallbacks.MetaAttach != nil {
		result := globalPluginState.metaCallbacks.MetaAttach(int(now))

		return C.int(result)
	}

	globalPluginState.timelineStatus = statusMetaAttached

	return 1
}

//export Meta_Detach
func Meta_Detach(now C.int, reason C.int) C.int {
	if globalPluginState.metaCallbacks != nil && globalPluginState.metaCallbacks.MetaDetach != nil {
		result := globalPluginState.metaCallbacks.MetaDetach(int(now), int(reason))

		return C.int(result)
	}

	globalPluginState.timelineStatus = statusMetaDetached

	return 1
}

//export GiveFnptrsToDll
func GiveFnptrsToDll(pengfuncsFromEngine *C.enginefuncs_t, pGlobals *C.globalvars_t) {
	globalVars := GlobalVarsFromC(pGlobals)
	globalPluginState.globalVars = globalVars

	globalPluginState.engineFuncs = NewEngineFuncs(pengfuncsFromEngine, globalVars)
}

//export GetNewDLLFunctions
func GetNewDLLFunctions(pNewFunctionTable *C.NEW_DLL_FUNCTIONS, interfaceVersion *C.int) C.int {
	return 1
}

//export GetNewDLLFunctions_Post
func GetNewDLLFunctions_Post(pNewFunctionTable *C.NEW_DLL_FUNCTIONS, interfaceVersion *C.int) C.int {
	return 1
}

//export GetEngineFunctions
func GetEngineFunctions(pengfuncsFromEngine *C.enginefuncs_t, interfaceVersion *C.int) C.int {
	C.SetHooks(pengfuncsFromEngine)

	return 1
}

//export GetEngineFunctions_Post
func GetEngineFunctions_Post(pengfuncsFromEngine *C.enginefuncs_t, interfaceVersion *C.int) C.int {
	return 1
}

//export GetEntityAPI2
func GetEntityAPI2(pFunctionTable *C.DLL_FUNCTIONS, interfaceVersion *C.int) C.int {
	C.SetDLLFunctions(pFunctionTable)

	return 1
}

//export GetEntityAPI2_Post
func GetEntityAPI2_Post(pFunctionTable *C.DLL_FUNCTIONS, interfaceVersion *C.int) C.int {
	C.SetDLLFunctionsPost(pFunctionTable)

	return 1
}
