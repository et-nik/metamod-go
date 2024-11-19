package metamod_go

import "C"
import "unsafe"

//export goGameDLLInit
func goGameDLLInit() {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.GameDLLInit != nil {
		metaResult := globalPluginState.apiCallbacks.GameDLLInit()
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goGameDLLInitPost
func goGameDLLInitPost() {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goGameDLLInitPost")
}

//export goSpawn
func goSpawn(pEntity *C.edict_t) C.int {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.Spawn != nil {
		metaResult, result := globalPluginState.apiCallbacks.Spawn(
			edictFromC(globalPluginState.globalVars.p, pEntity),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return C.int(0)
}

//export goSpawnPost
func goSpawnPost(pEntity *C.edict_t) C.int {
	return C.int(0)
}

//export goThink
func goThink(pEntity *C.edict_t) {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.Think != nil {
		metaResult := globalPluginState.apiCallbacks.Think(
			edictFromC(globalPluginState.globalVars.p, pEntity),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goThinkPost
func goThinkPost(pEntity *C.edict_t) {
}

//export goUse
func goUse(pEntity *C.edict_t, pOther *C.edict_t) {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.Use != nil {
		metaResult := globalPluginState.apiCallbacks.Use(
			edictFromC(globalPluginState.globalVars.p, pEntity),
			edictFromC(globalPluginState.globalVars.p, pOther),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goUsePost
func goUsePost(pEntity *C.edict_t, pOther *C.edict_t) {
}

//export goTouch
func goTouch(pEntity *C.edict_t, pOther *C.edict_t) {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.Touch != nil {
		metaResult := globalPluginState.apiCallbacks.Touch(
			edictFromC(globalPluginState.globalVars.p, pEntity),
			edictFromC(globalPluginState.globalVars.p, pOther),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goTouchPost
func goTouchPost(pEntity *C.edict_t, pOther *C.edict_t) {
}

//export goBlocked
func goBlocked(pEntity *C.edict_t, pOther *C.edict_t) {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.Blocked != nil {
		metaResult := globalPluginState.apiCallbacks.Blocked(
			edictFromC(globalPluginState.globalVars.p, pEntity),
			edictFromC(globalPluginState.globalVars.p, pOther),
		)

		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goBlockedPost
func goBlockedPost(pEntity *C.edict_t, pOther *C.edict_t) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goBlockedPost")
}

//func goKeyValue(pEntity *C.edict_t, pkvd *C.KeyValueData) {
//	globalPluginState.metaUtilFuncs.LogDeveloper("Called goKeyValue")
//}
//
//func goKeyValuePost(pEntity *C.edict_t, pkvd *C.KeyValueData) {
//	globalPluginState.metaUtilFuncs.LogDeveloper("Called goKeyValuePost")
//}

//export goClientConnect
func goClientConnect(pEntity *C.edict_t, name *C.char, address *C.char, reject *C.void) C.qboolean {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.ClientConnect != nil {
		metaResult, result, reason := globalPluginState.apiCallbacks.ClientConnect(
			edictFromC(globalPluginState.globalVars.p, pEntity),
			C.GoString(name),
			C.GoString(address),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if reason != "" {
			rejectString := C.CString(reason)
			defer C.free(unsafe.Pointer(rejectString))
			C.strcpy((*C.char)(unsafe.Pointer(reject)), rejectString)
		}

		boolResult := 0
		if result {
			boolResult = 1
		}

		return C.qboolean(C.int(boolResult))
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return C.qboolean(0)
}

//export goClientConnectPost
func goClientConnectPost(pEntity *C.edict_t, name *C.char, address *C.char, reject *C.void) C.qboolean {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientConnectPost")

	return C.qboolean(0)
}

//export goClientDisconnect
func goClientDisconnect(pEntity *C.edict_t) {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.ClientDisconnect != nil {
		metaResult := globalPluginState.apiCallbacks.ClientDisconnect(
			edictFromC(globalPluginState.globalVars.p, pEntity),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goClientDisconnectPost
func goClientDisconnectPost(pEntity *C.edict_t) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientDisconnectPost")
}

//export goClientKill
func goClientKill(pEntity *C.edict_t) {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.ClientKill != nil {
		metaResult := globalPluginState.apiCallbacks.ClientKill(
			edictFromC(globalPluginState.globalVars.p, pEntity),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goClientKillPost
func goClientKillPost(pEntity *C.edict_t) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientKillPost")
}

//export goClientPutInServer
func goClientPutInServer(pEntity *C.edict_t) {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.ClientPutInServer != nil {
		metaResult := globalPluginState.apiCallbacks.ClientPutInServer(
			edictFromC(globalPluginState.globalVars.p, pEntity),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goClientPutInServerPost
func goClientPutInServerPost(pEntity *C.edict_t) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientPutInServerPost")
}

//export goClientCommand
func goClientCommand(pEntity *C.edict_t) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientCommand")
}

//export goClientCommandPost
func goClientCommandPost(pEntity *C.edict_t) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientCommandPost")
}

//export goClientUserInfoChanged
func goClientUserInfoChanged(pEntity *C.edict_t, info *C.char) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientUserInfoChanged")
}

//export goClientUserInfoChangedPost
func goClientUserInfoChangedPost(pEntity *C.edict_t, info *C.char) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientUserInfoChangedPost")
}

//export goServerActivate
func goServerActivate(pEdictList *C.edict_t, edictCount C.int, clientMax C.int) {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.ServerActivate != nil {
		metaResult := globalPluginState.apiCallbacks.ServerActivate(
			edictFromC(globalPluginState.globalVars.p, pEdictList),
			int(edictCount),
			int(clientMax),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goServerActivatePost
func goServerActivatePost(pEdictList *C.edict_t, edictCount C.int, clientMax C.int) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goServerActivatePost")
}

//export goServerDeactivate
func goServerDeactivate() {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.ServerDeactivate != nil {
		metaResult := globalPluginState.apiCallbacks.ServerDeactivate()
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goServerDeactivatePost
func goServerDeactivatePost() {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goServerDeactivatePost")
}

//export goPlayerPreThink
func goPlayerPreThink(pEntity *C.edict_t) {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.PlayerPreThink != nil {
		metaResult := globalPluginState.apiCallbacks.PlayerPreThink(
			edictFromC(globalPluginState.globalVars.p, pEntity),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goPlayerPreThinkPost
func goPlayerPreThinkPost(pEntity *C.edict_t) {
}

//export goPlayerPostThink
func goPlayerPostThink(pEntity *C.edict_t) {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.PlayerPostThink != nil {
		metaResult := globalPluginState.apiCallbacks.PlayerPostThink(
			edictFromC(globalPluginState.globalVars.p, pEntity),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goPlayerPostThinkPost
func goPlayerPostThinkPost(pEntity *C.edict_t) {
}

//export goStartFrame
func goStartFrame() {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.StartFrame != nil {
		metaResult := globalPluginState.apiCallbacks.StartFrame()
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goStartFramePost
func goStartFramePost() {
}

//export goParmsNewLevel
func goParmsNewLevel() {
}

//export goParmsNewLevelPost
func goParmsNewLevelPost() {
}

//export goParmsChangeLevel
func goParmsChangeLevel() {
}

//export goParmsChangeLevelPost
func goParmsChangeLevelPost() {
}

//export goGetGameDescription
func goGetGameDescription() *C.char {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.GetGameDescription != nil {
		metaResult, result := globalPluginState.apiCallbacks.GetGameDescription()

		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.CString(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goGetGameDescriptionPost
func goGetGameDescriptionPost() *C.char {
	return nil
}

//export goSpectatorConnect
func goSpectatorConnect(pEntity *C.edict_t) {
}

//export goSpectatorConnectPost
func goSpectatorConnectPost(pEntity *C.edict_t) {
}

//export goSpectatorDisconnect
func goSpectatorDisconnect(pEntity *C.edict_t) {
}

//export goSpectatorDisconnectPost
func goSpectatorDisconnectPost(pEntity *C.edict_t) {
}

//export goSpectatorThink
func goSpectatorThink(pEntity *C.edict_t) {
}

//export goSpectatorThinkPost
func goSpectatorThinkPost(pEntity *C.edict_t) {
}

//export goSysError
func goSysError(errorString *C.char) {
	if globalPluginState.apiCallbacks != nil && globalPluginState.apiCallbacks.SysError != nil {
		metaResult := globalPluginState.apiCallbacks.SysError(C.GoString(errorString))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goSysErrorPost
func goSysErrorPost(errorString *C.char) {
}

//export goPMMove
//func goPMMove(pMove *C.playermove_t, server C.int) {
//	globalPluginState.metaUtilFuncs.LogDeveloper("Called goPMMove")
//}
//
////export goPMMovePost
//func goPMMovePost(pMove *C.playermove_t, server C.int) {
//	globalPluginState.metaUtilFuncs.LogDeveloper("Called goPMMovePost")
//}

//export goPMInit
//func goPMInit(pMove *C.playermove_t) {
//	globalPluginState.metaUtilFuncs.LogDeveloper("Called goPMInit")
//}

//export goPMInitPost
//func goPMInitPost(pMove *C.playermove_t) {
//	globalPluginState.metaUtilFuncs.LogDeveloper("Called goPMInitPost")
//}

// New DLL functions

//export goOnFreeEntPrivateDate
func goOnFreeEntPrivateDate(pEntity *C.edict_t) {
}

//export goOnFreeEntPrivateDatePost
func goOnFreeEntPrivateDatePost(pEntity *C.edict_t) {
}

//export goGameDLLShutdown
func goGameDLLShutdown() {
}

//export goGameDLLShutdownPost
func goGameDLLShutdownPost() {
}
