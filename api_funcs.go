package main

import "C"

//export goGameDLLInit
func goGameDLLInit() {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goGameDLLInit")
}

//export goGameDLLInitPost
func goGameDLLInitPost() {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goGameDLLInitPost")
}

//export goSpawn
func goSpawn(pEntity *C.edict_t) {
}

//export goSpawnPost
func goSpawnPost(pEntity *C.edict_t) {
}

//export goThink
func goThink(pEntity *C.edict_t) {
}

//export goThinkPost
func goThinkPost(pEntity *C.edict_t) {
}

//export goUse
func goUse(pEntity *C.edict_t, pOther *C.edict_t, pActivator *C.edict_t, useType C.float, value C.float) {
}

//export goUsePost
func goUsePost(pEntity *C.edict_t, pOther *C.edict_t, pActivator *C.edict_t, useType C.float, value C.float) {
}

//export goTouch
func goTouch(pEntity *C.edict_t, pOther *C.edict_t) {
}

//export goTouchPost
func goTouchPost(pEntity *C.edict_t, pOther *C.edict_t) {
}

//export goBlocked
func goBlocked(pEntity *C.edict_t, pOther *C.edict_t) {
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
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientConnect")

	return C.qboolean(0)
}

//export goClientConnectPost
func goClientConnectPost(pEntity *C.edict_t, name *C.char, address *C.char, reject *C.void) C.qboolean {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientConnectPost")

	return C.qboolean(0)
}

//export goClientDisconnect
func goClientDisconnect(pEntity *C.edict_t) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientDisconnect")
}

//export goClientDisconnectPost
func goClientDisconnectPost(pEntity *C.edict_t) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientDisconnectPost")
}

//export goClientKill
func goClientKill(pEntity *C.edict_t) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientKill")
}

//export goClientKillPost
func goClientKillPost(pEntity *C.edict_t) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientKillPost")
}

//export goClientPutInServer
func goClientPutInServer(pEntity *C.edict_t) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goClientPutInServer")
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
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goServerActivate")
}

//export goServerActivatePost
func goServerActivatePost(pEdictList *C.edict_t, edictCount C.int, clientMax C.int) {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goServerActivatePost")
}

//export goServerDeactivate
func goServerDeactivate() {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goServerDeactivate")
}

//export goServerDeactivatePost
func goServerDeactivatePost() {
	globalPluginState.metaUtilFuncs.LogDeveloper("Called goServerDeactivatePost")
}

//export goPlayerPreThink
func goPlayerPreThink(pEntity *C.edict_t) {
}

//export goPlayerPreThinkPost
func goPlayerPreThinkPost(pEntity *C.edict_t) {
}

//export goPlayerPostThink
func goPlayerPostThink(pEntity *C.edict_t) {
}

//export goPlayerPostThinkPost
func goPlayerPostThinkPost(pEntity *C.edict_t) {
}

//export goStartFrame
func goStartFrame() {
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
