package main

/*
#cgo CPPFLAGS: -I${SRCDIR}/hlsdk/dlls -I${SRCDIR}/hlsdk/engine -I${SRCDIR}/hlsdk/pm_shared -I${SRCDIR}/hlsdk/common -include ${SRCDIR}/hlsdk/public/basetypes.h -include ${SRCDIR}/hlsdk/common/const.h -include ${SRCDIR}/hlsdk/engine/edict.h

#include <stdio.h>

#include <eiface.h>

#include "metamod/plinfo.h"
#include "metamod/meta_api.h"
#include "metamod/mutil.h"

int GetNewDLLFunctions(NEW_DLL_FUNCTIONS *pNewFunctionTable, int *interfaceVersion);
*/
import "C"
import (
	"fmt"
)

func main() {}

var pluginInfo = &PluginInfo{
	InterfaceVersion: MetaInterfaceVersion,
	Name:             "GoPlugin",
	Version:          "1.0",
	Date:             "2024-11-08",
	Author:           "KNiK",
	Url:              "https://github.com/et-nik/metamod-go",
	Loadable:         PluginLoadTimeAnyTime,
	Unloadable:       PluginLoadTimeAnyTime,
}

//export Meta_Attach
func Meta_Attach(now C.int, pFunctionTable *C.META_FUNCTIONS, pMGlobals *C.void, pGamedllFuncs *C.void) C.int {
	fmt.Println("=====================================")
	fmt.Println("(Meta_Attach) Hi from Go!")
	fmt.Println("=====================================")

	//var gMetaFunctionTable = new(C.META_FUNCTIONS)

	//C.GetNewDLLFunctions(nil, nil)
	//
	//gMetaFunctionTable = &C.META_FUNCTIONS{
	//	pfnGetEntityAPI:            nil,
	//	pfnGetEntityAPI_Post:       nil,
	//	pfnGetEntityAPI2:           nil,
	//	pfnGetEntityAPI2_Post:      nil,
	//	pfnGetNewDLLFunctions:      C.GETNEWDLLFUNCTIONS_FN(C.GetNewDLLFunctions),
	//	pfnGetNewDLLFunctions_Post: nil,
	//	pfnGetEngineFunctions:      nil,
	//	pfnGetEngineFunctions_Post: nil,
	//}

	pFunctionTable.pfnGetNewDLLFunctions = C.GETNEWDLLFUNCTIONS_FN(C.GetNewDLLFunctions)

	return 1
}

//export Meta_Query
func Meta_Query(interfaceVersion *C.char, plinfo **C.plugin_info_t, pMetaUtilFuncs *C.mutil_funcs_t) C.int {
	*plinfo = pluginInfo.ToC()

	P.Info = pluginInfo
	P.MetaUtilFuncs = &MUtilFuncs{
		info: pluginInfo,
		p:    pMetaUtilFuncs,
	}

	P.MetaUtilFuncs.LogConsole("Hello from Go!")
	P.MetaUtilFuncs.LogConsole("Hello from Go!")
	P.MetaUtilFuncs.LogConsole("Hello from Go!")
	P.MetaUtilFuncs.LogConsole("Hello from Go!")

	return 1
}

//export Meta_Detach
func Meta_Detach(now C.int, reason C.int) C.int {
	fmt.Println("=====================================")
	fmt.Println("(Meta_Detach) Hi from Go!")
	fmt.Println("=====================================")

	return 1
}

//export Meta_Init
func Meta_Init() {
	fmt.Println("=====================================")
	fmt.Println("(Meta_Init) Hi from Go!")
	fmt.Println("=====================================")
}

//export GiveFnptrsToDll
func GiveFnptrsToDll(pengfuncsFromEngine *C.enginefuncs_t, pGlobals *C.void) {
	fmt.Println("=====================================")
	fmt.Println("(GiveFnptrsToDll) Hi from Go!")
	fmt.Println("=====================================")

	P.EngineFuncs = &EngineFuncs{
		p: pengfuncsFromEngine,
	}
}

//export GetNewDLLFunctions
func GetNewDLLFunctions(pNewFunctionTable *C.NEW_DLL_FUNCTIONS, interfaceVersion *C.int) C.int {
	fmt.Println("=====================================")
	fmt.Println("(GetNewDLLFunctions) Hi from Go!")
	fmt.Println("=====================================")

	return 1
}

//export GetNewDLLFunctions_Post
func GetNewDLLFunctions_Post(pNewFunctionTable *C.NEW_DLL_FUNCTIONS, interfaceVersion *C.int) C.int {
	fmt.Println("=====================================")
	fmt.Println("(GetNewDLLFunctions_Post) Hi from Go!")
	fmt.Println("=====================================")

	return 1
}

//export GetEntityAPI2
func GetEntityAPI2(pFunctionTable *C.DLL_FUNCTIONS, interfaceVersion *C.int) C.int {
	fmt.Println("=====================================")
	fmt.Println("(GetEntityAPI2) Hi from Go!")
	fmt.Println("=====================================")

	return 1
}

//export GetEntityAPI2_Post
func GetEntityAPI2_Post(pFunctionTable *C.DLL_FUNCTIONS, interfaceVersion *C.int) C.int {
	fmt.Println("=====================================")
	fmt.Println("(GetEntityAPI2) Hi from Go!")
	fmt.Println("=====================================")

	return 1
}

//export GetEngineFunctions
func GetEngineFunctions(pengfuncsFromEngine *C.enginefuncs_t, interfaceVersion *C.int) C.int {
	fmt.Println("=====================================")
	fmt.Println("(GetEngineFunctions) Hi from Go!")
	fmt.Println("=====================================")

	return 1
}

//export GetEngineFunctions_Post
func GetEngineFunctions_Post(pengfuncsFromEngine *C.enginefuncs_t, interfaceVersion *C.int) C.int {
	fmt.Println("=====================================")
	fmt.Println("(GetEngineFunctions) Hi from Go!")
	fmt.Println("=====================================")

	return 1
}
