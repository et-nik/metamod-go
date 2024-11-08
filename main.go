package main

/*
#include "metamod/plinfo.h"
*/
import "C"
import (
	"fmt"
	"github.com/et-nik/metamod-go/metamod"
)

func main() {}

var Plugin = &metamod.PluginInfo{
	InterfaceVersion: MetaInterfaceVersion,
	Name:             "GoPlugin",
	Version:          "1.0",
	Date:             "2024-11-08",
	Author:           "KNiK",
	Url:              "https://github.com/et-nik/metamod-go",
	Loadable:         metamod.PluginLoadTimeAnyTime,
	Unloadable:       metamod.PluginLoadTimeAnyTime,
}

//type PluginInfo struct {
//	p C.plugin_info_t
//}

//export Meta_Attach
func Meta_Attach(now C.int, pFunctionTable *C.void, pMGlobals *C.void, pGamedllFuncs *C.void) C.int {
	fmt.Println("=====================================")
	fmt.Println("(Meta_Attach) Hi from Go!")
	fmt.Println("=====================================")

	return 1
}

//export Meta_Query
func Meta_Query(interfaceVersion *C.char, plinfo **C.plugin_info_t, pMetaUtilFuncs *C.void) C.int {
	*plinfo = Plugin.ToC()

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
func GiveFnptrsToDll(pengfuncsFromEngine *C.void, pGlobals *C.void) {
	fmt.Println("=====================================")
	fmt.Println("(GiveFnptrsToDll) Hi from Go!")
	fmt.Println("=====================================")
}
