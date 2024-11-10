package main

/*
#include <eiface.h>

void goGameDLLInit(void);
void goGameDLLInitPost(void);

void goSpawn(edict_t *pEntity);
void goSpawnPost(edict_t *pEntity);

void goThink(edict_t *pEntity);
void goThinkPost(edict_t *pEntity);

void goUse(edict_t *pEntity, edict_t *pOther, edict_t *pActivator, float useType, float value);
void goUsePost(edict_t *pEntity, edict_t *pOther, edict_t *pActivator, float useType, float value);

void goTouch(edict_t *pEntity, edict_t *pOther);
void goTouchPost(edict_t *pEntity, edict_t *pOther);

void goBlocked(edict_t *pEntity, edict_t *pOther);
void goBlockedPost(edict_t *pEntity, edict_t *pOther);

qboolean goClientConnect(edict_t *pEntity, const char *pszName, const char *pszAddress, char szRejectReason[ 128 ]);
qboolean goClientConnectPost(edict_t *pEntity, const char *pszName, const char *pszAddress, char szRejectReason[ 128 ]);

void goClientDisconnect(edict_t *pEntity);
void goClientDisconnectPost(edict_t *pEntity);

void goClientKill(edict_t *pEntity);
void goClientKillPost(edict_t *pEntity);

void goClientPutInServer(edict_t *pEntity);
void goClientPutInServerPost(edict_t *pEntity);

void goClientCommand(edict_t *pEntity);
void goClientCommandPost(edict_t *pEntity);

void goClientUserInfoChanged(edict_t *pEntity, char *infobuffer);
void goClientUserInfoChangedPost(edict_t *pEntity, char *infobuffer);

void goServerActivate(edict_t *pEdictList, int edictCount, int clientMax);
void goServerActivatePost(edict_t *pEdictList, int edictCount, int clientMax);

void goServerDeactivate(void);
void goServerDeactivatePost(void);

void goPlayerPreThink(edict_t *pEntity);
void goPlayerPreThinkPost(edict_t *pEntity);

void goPlayerPostThink(edict_t *pEntity);
void goPlayerPostThinkPost(edict_t *pEntity);

void goStartFrame(void);
void goStartFramePost(void);

void goParmsNewLevel(void);
void goParmsNewLevelPost(void);

void goParmsChangeLevel(void);
void goParmsChangeLevelPost(void);

void goSpectatorConnect(edict_t *pEntity);
void goSpectatorConnectPost(edict_t *pEntity);

void goSpectatorDisconnect(edict_t *pEntity);
void goSpectatorDisconnectPost(edict_t *pEntity);

void goSpectatorThink(edict_t *pEntity);
void goSpectatorThinkPost(edict_t *pEntity);

void goSysError(const char *error_string);
void goSysErrorPost(const char *error_string);

void SetDLLFunctions(DLL_FUNCTIONS *pFunctionTable) {
	pFunctionTable->pfnGameInit = goGameDLLInit;
	pFunctionTable->pfnSpawn = goSpawn;
	pFunctionTable->pfnThink = goThink;
	pFunctionTable->pfnUse = goUse;
	pFunctionTable->pfnTouch = goTouch;
	pFunctionTable->pfnBlocked = goBlocked;
	pFunctionTable->pfnClientConnect = goClientConnect;
	pFunctionTable->pfnClientDisconnect = goClientDisconnect;
	pFunctionTable->pfnClientKill = goClientKill;
	pFunctionTable->pfnClientPutInServer = goClientPutInServer;
	pFunctionTable->pfnServerActivate = goServerActivate;
	pFunctionTable->pfnServerDeactivate = goServerDeactivate;
	pFunctionTable->pfnPlayerPreThink = goPlayerPreThink;
	pFunctionTable->pfnPlayerPostThink = goPlayerPostThink;
	pFunctionTable->pfnStartFrame = goStartFrame;
	pFunctionTable->pfnParmsNewLevel = goParmsNewLevel;
	pFunctionTable->pfnParmsChangeLevel = goParmsChangeLevel;
	pFunctionTable->pfnSpectatorConnect = goSpectatorConnect;
	pFunctionTable->pfnSpectatorDisconnect = goSpectatorDisconnect;
	pFunctionTable->pfnSpectatorThink = goSpectatorThink;
	pFunctionTable->pfnSys_Error = goSysError;
}

void SetDLLFunctionsPost(DLL_FUNCTIONS *pFunctionTable) {
	pFunctionTable->pfnGameInit = goGameDLLInitPost;
	pFunctionTable->pfnSpawn = goSpawnPost;
	pFunctionTable->pfnThink = goThinkPost;
	pFunctionTable->pfnUse = goUsePost;
	pFunctionTable->pfnTouch = goTouchPost;
	pFunctionTable->pfnBlocked = goBlockedPost;
	pFunctionTable->pfnClientConnect = goClientConnectPost;
	pFunctionTable->pfnClientDisconnect = goClientDisconnectPost;
	pFunctionTable->pfnClientKill = goClientKillPost;
	pFunctionTable->pfnClientPutInServer = goClientPutInServerPost;
	pFunctionTable->pfnServerActivate = goServerActivatePost;
	pFunctionTable->pfnServerDeactivate = goServerDeactivatePost;
	pFunctionTable->pfnPlayerPreThink = goPlayerPreThinkPost;
	pFunctionTable->pfnPlayerPostThink = goPlayerPostThinkPost;
	pFunctionTable->pfnStartFrame = goStartFramePost;
	pFunctionTable->pfnParmsNewLevel = goParmsNewLevelPost;
	pFunctionTable->pfnParmsChangeLevel = goParmsChangeLevelPost;
	pFunctionTable->pfnSpectatorConnect = goSpectatorConnectPost;
	pFunctionTable->pfnSpectatorDisconnect = goSpectatorDisconnectPost;
	pFunctionTable->pfnSpectatorThink = goSpectatorThinkPost;
	pFunctionTable->pfnSys_Error = goSysErrorPost;
}



*/
import "C"
