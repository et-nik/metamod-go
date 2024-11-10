package main

/*
#cgo CPPFLAGS: -I${SRCDIR}/hlsdk/dlls -I${SRCDIR}/hlsdk/engine -I${SRCDIR}/hlsdk/pm_shared -I${SRCDIR}/hlsdk/common -include ${SRCDIR}/hlsdk/public/basetypes.h -include ${SRCDIR}/hlsdk/common/const.h -include ${SRCDIR}/hlsdk/engine/edict.h

#include <stdio.h>

#include <eiface.h>

#include "metamod/index.h"

int GetNewDLLFunctions(NEW_DLL_FUNCTIONS *pNewFunctionTable, int *interfaceVersion);
int GetEntityAPI2(DLL_FUNCTIONS* pFunctionTable, int* interfaceVersion);
int GetEntityAPI2_Post(DLL_FUNCTIONS* pFunctionTable, int* interfaceVersion);

extern void SetDLLFunctions(DLL_FUNCTIONS *pFunctionTable);

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
	LogTag:           "GoPlugin",
	Loadable:         PluginLoadTimeAnyTime,
	Unloadable:       PluginLoadTimeAnyTime,
}

//export Meta_Attach
func Meta_Attach(now C.int, pFunctionTable *C.META_FUNCTIONS, pMGlobals *C.void, pGamedllFuncs *C.void) C.int {
	fmt.Println("=====================================")
	fmt.Println("(Meta_Attach) Hi from Go!")
	fmt.Println("=====================================")

	pFunctionTable.pfnGetEntityAPI2 = C.GETENTITYAPI2_FN(C.GetEntityAPI2)
	pFunctionTable.pfnGetEntityAPI2_Post = C.GETENTITYAPI2_FN(C.GetEntityAPI2_Post)
	pFunctionTable.pfnGetNewDLLFunctions = C.GETNEWDLLFUNCTIONS_FN(C.GetNewDLLFunctions)

	P.EngineFuncs.AddServerCommand("test", func(argc int, argv ...string) {
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("Server command test")
		fmt.Println("argc:", argc)
		fmt.Println("argv:", argv)
		fmt.Println("---- GlobalVars ----")
		fmt.Println("Time:", P.GlobalVars.Time())
		fmt.Println("Frame time:", P.GlobalVars.FrameTime())
		fmt.Println("Map name:", P.GlobalVars.MapName())
		fmt.Println("Deathmatch:", P.GlobalVars.Deathmatch())
		fmt.Println("Coop:", P.GlobalVars.Coop())
		fmt.Println("Max Clients:", P.GlobalVars.MaxClients())
		fmt.Println("Max Enties:", P.GlobalVars.MaxEntities())
		fmt.Println("=====================================")
		fmt.Println()
	})

	P.EngineFuncs.AddServerCommand("test2", func(argc int, argv ...string) {
		someVal := "someVal"
		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("Server command test2")
		fmt.Println("argc:", argc)
		fmt.Println("argv:", argv)
		fmt.Println("someVal", someVal)
		P.MetaUtilFuncs.LogMessage("Server command test")
		fmt.Println("=====================================")
		fmt.Println()

	})

	return 1
}

//export Meta_Query
func Meta_Query(interfaceVersion *C.char, plinfo **C.plugin_info_t, pMetaUtilFuncs *C.mutil_funcs_t) C.int {
	*plinfo = pluginInfo.ToCPluginInfo()

	setCGlobalPluginInfo(pluginInfo)

	P.Info = pluginInfo
	P.MetaUtilFuncs = &MUtilFuncs{
		p: pMetaUtilFuncs,
	}

	P.MetaUtilFuncs.LogConsole("Hello from Go!")
	P.MetaUtilFuncs.LogConsole("Hello from Go!")
	P.MetaUtilFuncs.LogConsole("Hello from Go!")
	P.MetaUtilFuncs.LogConsole("Hello from Go!")

	P.MetaUtilFuncs.LogMessage("(Message TEST) Hello from Go!")
	P.MetaUtilFuncs.LogMessage("(Message TEST) Hello from Go!")
	P.MetaUtilFuncs.LogMessage("(Message TEST) Hello from Go!")
	P.MetaUtilFuncs.LogMessage("(Message TEST) Hello from Go!")
	P.MetaUtilFuncs.LogMessage("(Message TEST) Hello from Go!")

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
func GiveFnptrsToDll(pengfuncsFromEngine *C.enginefuncs_t, pGlobals *C.globalvars_t) {
	fmt.Println("=====================================")
	fmt.Println("(GiveFnptrsToDll) Hi from Go!")
	fmt.Println("=====================================")

	gb := GlobalVarsFromC(pGlobals)
	fmt.Println("GlobalVars time:", gb.Time())

	P.GlobalVars = gb

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

//export GetEntityAPI2
func GetEntityAPI2(pFunctionTable *C.DLL_FUNCTIONS, interfaceVersion *C.int) C.int {
	fmt.Println("=====================================")
	fmt.Println("(GetEntityAPI2) Hi from Go!")
	fmt.Println("=====================================")

	C.SetDLLFunctions(pFunctionTable)

	return 1
}

//export GetEntityAPI2_Post
func GetEntityAPI2_Post(pFunctionTable *C.DLL_FUNCTIONS, interfaceVersion *C.int) C.int {
	fmt.Println("=====================================")
	fmt.Println("(GetEntityAPI2) Hi from Go!")
	fmt.Println("=====================================")

	return 1
}
