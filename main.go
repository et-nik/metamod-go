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
import (
	"fmt"
	"strconv"
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
func Meta_Attach(now C.int, pFunctionTable *C.META_FUNCTIONS, pMGlobals *C.meta_globals_t, pGamedllFuncs *C.void) C.int {
	pFunctionTable.pfnGetEntityAPI2 = C.GETENTITYAPI2_FN(C.GetEntityAPI2)
	pFunctionTable.pfnGetEntityAPI2_Post = C.GETENTITYAPI2_FN(C.GetEntityAPI2_Post)
	pFunctionTable.pfnGetNewDLLFunctions = C.GETNEWDLLFUNCTIONS_FN(C.GetNewDLLFunctions)
	pFunctionTable.pfnGetEngineFunctions = C.GET_ENGINE_FUNCTIONS_FN(C.GetEngineFunctions)
	pFunctionTable.pfnGetEngineFunctions_Post = C.GET_ENGINE_FUNCTIONS_FN(C.GetEngineFunctions_Post)

	P.MetaGlobals = MetaGlobalsFromC(pMGlobals)

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

	P.EngineFuncs.AddServerCommand("entinfo", func(argc int, argv ...string) {
		if argc < 2 {
			fmt.Println("Usage: entinfo <entityIndex>")
			return
		}

		entityIndex, err := strconv.Atoi(argv[1])
		if err != nil {
			fmt.Println("Invalid entity index")
			return
		}

		edict := P.EngineFuncs.EntityOfEntIndex(entityIndex)
		if edict == nil {
			fmt.Println("Entity not found")
			return
		}

		entVars := edict.EntVars()

		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("Entity info")
		fmt.Println("Index:", entityIndex)
		fmt.Println("SerialNumber:", edict.SerialNumber())
		fmt.Println("Classname:", entVars.ClassName())
		fmt.Println("Globalname:", entVars.GlobalName())
		fmt.Println("Origin:", entVars.Origin())
		fmt.Println("Model:", entVars.Model())
		fmt.Println("ViewModel:", entVars.ViewModel())
		fmt.Println("WeaponModel:", entVars.WeaponModel())
		fmt.Println("Health:", entVars.Health())
		fmt.Println("Max Health:", entVars.MaxHealth())
	})

	P.EngineFuncs.AddServerCommand("traceplayers", func(argc int, argv ...string) {
		if argc < 3 {
			fmt.Println("Usage: traceplayers <player1> <player2>")
			return
		}

		player1, err := strconv.Atoi(argv[1])
		if err != nil {
			fmt.Println("Invalid player index")
			return
		}

		player2, err := strconv.Atoi(argv[2])
		if err != nil {
			fmt.Println("Invalid player index")
			return
		}

		edict1 := P.EngineFuncs.EntityOfEntIndex(player1)
		if edict1 == nil {
			fmt.Println("Player 1 not found")
			return
		}

		edict2 := P.EngineFuncs.EntityOfEntIndex(player2)
		if edict2 == nil {
			fmt.Println("Player 2 not found")
			return
		}

		entVars1 := edict1.EntVars()
		entVars2 := edict2.EntVars()

		result := P.EngineFuncs.TraceLine(entVars1.Origin(), entVars2.Origin(), 0, nil)

		fmt.Println()
		fmt.Println("=====================================")
		fmt.Println("Trace result")
		fmt.Println("Player 1 Origin:", entVars1.Origin())
		fmt.Println("Player 2 Origin:", entVars2.Origin())
		fmt.Println("Start:", entVars1.Origin())
		fmt.Println("End:", entVars2.Origin())
		fmt.Println("Fraction:", result.Fraction)
		fmt.Println("EndPos:", result.EndPos)
		fmt.Println("PlaneDist:", result.PlaneDist)
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

	P.MetaUtilFuncs.LogDeveloper("Meta_Query called")

	return 1
}

//export Meta_Detach
func Meta_Detach(now C.int, reason C.int) C.int {
	return 1
}

//export Meta_Init
func Meta_Init() {
}

//export GiveFnptrsToDll
func GiveFnptrsToDll(pengfuncsFromEngine *C.enginefuncs_t, pGlobals *C.globalvars_t) {
	globalVars := GlobalVarsFromC(pGlobals)
	P.GlobalVars = globalVars

	P.EngineFuncs = NewEngineFuncs(pengfuncsFromEngine, globalVars)
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

	C.SetHooks(pengfuncsFromEngine)

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
	C.SetDLLFunctions(pFunctionTable)

	return 1
}

//export GetEntityAPI2_Post
func GetEntityAPI2_Post(pFunctionTable *C.DLL_FUNCTIONS, interfaceVersion *C.int) C.int {
	C.SetDLLFunctionsPost(pFunctionTable)

	return 1
}
