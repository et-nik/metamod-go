package main

/*
#include <eiface.h>

#include <stdlib.h>
#include <stdio.h>

#define MAX_SERVER_COMMAND_CALLBACKS 128
#define SERVER_COMMAND_CALLBACKS_CATEGORY "server_commands"

extern const char* ReadString(globalvars_t *gpGlobals, int offset);
extern int MakeString(globalvars_t *gpGlobals, char *str);

extern void callGoFunction(void *f, int argc, char **argv);
extern void getGoCallback(char *category, char *v, void **f);
extern void setGoCallback(char *category, char *v, void *f);

void* engineFuncsGoCallbacks[MAX_SERVER_COMMAND_CALLBACKS];
int engineFuncsGoCallbacksCount = 0;

struct enginefuncs_s *engineFuncs;

typedef void (*server_command_callback_t)(void);

// Engine funcs
int engineFuncsPrecacheModel(struct enginefuncs_s *t, char *s) {
	return (*t->pfnPrecacheModel)(s);
}

int engineFuncsPrecacheSound(struct enginefuncs_s *t, char *s) {
	return (*t->pfnPrecacheSound)(s);
}

void engineFuncsSetModel(struct enginefuncs_s *t, edict_t *e, const char *m) {
	(*t->pfnSetModel)(e, m);
}

int engineFuncsModelIndex(struct enginefuncs_s *t, const char *m) {
	return (*t->pfnModelIndex)(m);
}

int engineFuncsModelFrames(struct enginefuncs_s *t, int modelIndex) {
	return (*t->pfnModelFrames)(modelIndex);
}

void engineFuncsSetSize(struct enginefuncs_s *t, edict_t *e, const float *mins, const float *maxs) {
	(*t->pfnSetSize)(e, mins, maxs);
}

void engineFuncsChangeLevel(struct enginefuncs_s *t, const char *s1, const char *s2) {
	(*t->pfnChangeLevel)(s1, s2);
}

void engineFuncsGetSpawnParms(struct enginefuncs_s *t, edict_t *ent) {
	(*t->pfnGetSpawnParms)(ent);
}

void engineFuncsSaveSpawnParms(struct enginefuncs_s *t, edict_t *ent) {
	(*t->pfnSaveSpawnParms)(ent);
}

float engineFuncsVecToYaw(struct enginefuncs_s *t, const float *vec) {
	return (*t->pfnVecToYaw)(vec);
}

void engineFuncsVecToAngles(struct enginefuncs_s *t, const float *vecIn, float *vecOut) {
	(*t->pfnVecToAngles)(vecIn, vecOut);
}

void engineFuncsMoveToOrigin(struct enginefuncs_s *t, edict_t *ent, const float *goal, float dist, int moveType) {
	(*t->pfnMoveToOrigin)(ent, goal, dist, moveType);
}

void engineFuncsChangeYaw(struct enginefuncs_s *t, edict_t *ent) {
	(*t->pfnChangeYaw)(ent);
}

void engineFuncsChangePitch(struct enginefuncs_s *t, edict_t *ent) {
	(*t->pfnChangePitch)(ent);
}

edict_t* engineFuncsFindEntityByString(struct enginefuncs_s *t, edict_t *eStartSearchAfter, const char *field, const char *value) {
	return (*t->pfnFindEntityByString)(eStartSearchAfter, field, value);
}

int engineFuncsGetEntityIllum(struct enginefuncs_s *t, edict_t *ent) {
	return (*t->pfnGetEntityIllum)(ent);
}

edict_t* engineFuncsFindEntityInSphere(struct enginefuncs_s *t, edict_t *eStartSearchAfter, const float *pos, float rad) {
	return (*t->pfnFindEntityInSphere)(eStartSearchAfter, pos, rad);
}

edict_t* engineFuncsFindClientInPVS(struct enginefuncs_s *t, edict_t *e) {
	return (*t->pfnFindClientInPVS)(e);
}

edict_t* engineFuncsEntitiesInPVS(struct enginefuncs_s *t, edict_t *e) {
	return (*t->pfnEntitiesInPVS)(e);
}

void engineFuncsMakeVectors(struct enginefuncs_s *t, const float *vec) {
	(*t->pfnMakeVectors)(vec);
}

void engineFuncsAngleVectors(struct enginefuncs_s *t, const float *vec, float *forward, float *right, float *up) {
	(*t->pfnAngleVectors)(vec, forward, right, up);
}

edict_t* engineFuncsCreateEntity(struct enginefuncs_s *t) {
	return (*t->pfnCreateEntity)();
}

void engineFuncsRemoveEntity(struct enginefuncs_s *t, edict_t *e) {
	(*t->pfnRemoveEntity)(e);
}

edict_t* engineFuncsCreateNamedEntity(struct enginefuncs_s *t, int className) {
	return (*t->pfnCreateNamedEntity)(className);
}

void engineFuncsMakeStatic(struct enginefuncs_s *t, edict_t *ent) {
	(*t->pfnMakeStatic)(ent);
}

void engineFuncsEntIsOnFloor(struct enginefuncs_s *t, edict_t *ent) {
	(*t->pfnEntIsOnFloor)(ent);
}

void engineFuncsDropToFloor(struct enginefuncs_s *t, edict_t *ent) {
	(*t->pfnDropToFloor)(ent);
}

void engineFuncsWalkMove(struct enginefuncs_s *t, edict_t *ent, float yaw, float dist, int iMode) {
	(*t->pfnWalkMove)(ent, yaw, dist, iMode);
}

void engineFuncsSetOrigin(struct enginefuncs_s *t, edict_t *e, const float *rgflOrigin) {
	(*t->pfnSetOrigin)(e, rgflOrigin);
}

void engineFuncsEmitSound(struct enginefuncs_s *t, edict_t *entity, int channel, const char *sample, float volume, float attenuation, int fFlags, int pitch) {
	(*t->pfnEmitSound)(entity, channel, sample, volume, attenuation, fFlags, pitch);
}

void engineFuncsEmitAmbientSound(struct enginefuncs_s *t, edict_t *entity, float *pos, const char *samp, float vol, float attenuation, int fFlags, int pitch) {
	(*t->pfnEmitAmbientSound)(entity, pos, samp, vol, attenuation, fFlags, pitch);
}

void engineFuncsTraceLine(struct enginefuncs_s *t, const float *v1, const float *v2, int fNoMonsters, edict_t *pentToSkip, TraceResult *ptr) {
	(*t->pfnTraceLine)(v1, v2, fNoMonsters, pentToSkip, ptr);
}

void engineFuncsTraceToss(struct enginefuncs_s *t, edict_t *pent, edict_t *pentToIgnore, TraceResult *ptr) {
	(*t->pfnTraceToss)(pent, pentToIgnore, ptr);
}

int engineFuncsTraceMonsterHull(struct enginefuncs_s *t, edict_t *pent, const float *v1, const float *v2, int fNoMonsters, edict_t *pentToSkip, TraceResult *ptr) {
	return (*t->pfnTraceMonsterHull)(pent, v1, v2, fNoMonsters, pentToSkip, ptr);
}

void engineFuncsTraceHull(struct enginefuncs_s *t, const float *v1, const float *v2, int fNoMonsters, int hullNumber, edict_t *pentToSkip, TraceResult *ptr) {
	(*t->pfnTraceHull)(v1, v2, fNoMonsters, hullNumber, pentToSkip, ptr);
}

void engineFuncsTraceModel(struct enginefuncs_s *t, const float *v1, const float *v2, int hullNumber, edict_t *pent, TraceResult *ptr) {
	(*t->pfnTraceModel)(v1, v2, hullNumber, pent, ptr);
}

const char* engineFuncsTraceTexture(struct enginefuncs_s *t, edict_t *pent, const float *v1, const float *v2) {
	return (*t->pfnTraceTexture)(pent, v1, v2);
}

void engineFuncsGetAimVector(struct enginefuncs_s *t, edict_t *ent, float speed, float *rgflReturn) {
	(*t->pfnGetAimVector)(ent, speed, rgflReturn);
}

void engineFuncsServerCommand(struct enginefuncs_s *t, char *str) {
	(*t->pfnServerCommand)(str);
}

void engineFuncsServerExecute(struct enginefuncs_s *t) {
	(*t->pfnServerExecute)();
}

void engineFuncsClientCommand(struct enginefuncs_s *t, edict_t *pEdict, const char *szFmt) {
	(*t->pfnClientCommand)(pEdict, szFmt);
}

void engineFuncsParticleEffect(struct enginefuncs_s *t, const float *org, const float *dir, float color, float count) {
	(*t->pfnParticleEffect)(org, dir, color, count);
}

void engineFuncsLightStyle(struct enginefuncs_s *t, int style, char *val) {
	(*t->pfnLightStyle)(style, val);
}

int engineFuncsDecalIndex(struct enginefuncs_s *t, const char *name) {
	return (*t->pfnDecalIndex)(name);
}

int engineFuncsPointContents(struct enginefuncs_s *t, const float *rgflVector) {
	return (*t->pfnPointContents)(rgflVector);
}

void engineFuncsMessageBegin(struct enginefuncs_s *t, int msg_dest, int msg_type, const float *pOrigin, edict_t *ed) {
	(*t->pfnMessageBegin)(msg_dest, msg_type, pOrigin, ed);
}

void engineFuncsMessageEnd(struct enginefuncs_s *t) {
	(*t->pfnMessageEnd)();
}

void engineFuncsWriteByte(struct enginefuncs_s *t, int i) {
	(*t->pfnWriteByte)(i);
}

void engineFuncsWriteChar(struct enginefuncs_s *t, int i) {
	(*t->pfnWriteChar)(i);
}

void engineFuncsWriteShort(struct enginefuncs_s *t, int i) {
	(*t->pfnWriteShort)(i);
}

void engineFuncsWriteLong(struct enginefuncs_s *t, int i) {
	(*t->pfnWriteLong)(i);
}

void engineFuncsWriteAngle(struct enginefuncs_s *t, float f) {
	(*t->pfnWriteAngle)(f);
}

void engineFuncsWriteCoord(struct enginefuncs_s *t, float f) {
	(*t->pfnWriteCoord)(f);
}

void engineFuncsWriteString(struct enginefuncs_s *t, const char *s) {
	(*t->pfnWriteString)(s);
}

void engineFuncsWriteEntity(struct enginefuncs_s *t, int i) {
	(*t->pfnWriteEntity)(i);
}

void engineFuncsCVarRegister(struct enginefuncs_s *t, cvar_t *pCvar) {
	(*t->pfnCVarRegister)(pCvar);
}

const char* engineFuncsCVarGetString(struct enginefuncs_s *t, const char *szVarName) {
	return (*t->pfnCVarGetString)(szVarName);
}

float engineFuncsCVarGetFloat(struct enginefuncs_s *t, const char *szVarName) {
	return (*t->pfnCVarGetFloat)(szVarName);
}

void engineFuncsCVarSetFloat(struct enginefuncs_s *t, const char *szVarName, float flValue) {
	(*t->pfnCVarSetFloat)(szVarName, flValue);
}

void engineFuncsCVarSetString(struct enginefuncs_s *t, const char *szVarName, const char *szValue) {
	(*t->pfnCVarSetString)(szVarName, szValue);
}

void engineFuncsAlertMessage(struct enginefuncs_s *t, ALERT_TYPE atype, char *szFmt) {
	(*t->pfnAlertMessage)(atype, szFmt);
}

void engineFuncsEngineFprintf(struct enginefuncs_s *t, FILE *pfile, char *szFmt) {
	(*t->pfnEngineFprintf)(pfile, szFmt);
}

void engineFuncsPvAllocEntPrivateData(struct enginefuncs_s *t, edict_t *ent, int32 cb) {
	(*t->pfnPvAllocEntPrivateData)(ent, cb);
}

void engineFuncsPvEntPrivateData(struct enginefuncs_s *t, edict_t *ent) {
	(*t->pfnPvEntPrivateData)(ent);
}

void engineFuncsFreeEntPrivateData(struct enginefuncs_s *t, edict_t *ent) {
	(*t->pfnFreeEntPrivateData)(ent);
}

const char* engineFuncsSzFromIndex(struct enginefuncs_s *t, int iString) {
	return (*t->pfnSzFromIndex)(iString);
}

int engineFuncsAllocString(struct enginefuncs_s *t, const char *str) {
	return (*t->pfnAllocString)(str);
}

struct entvars_s* engineFuncsGetVarsOfEnt(struct enginefuncs_s *t, edict_t *ent) {
	return (*t->pfnGetVarsOfEnt)(ent);
}

int engineFuncsIndexOfEdict(struct enginefuncs_s *t, const edict_t *pEdict) {
	return (*t->pfnIndexOfEdict)(pEdict);
}

edict_t* engineFuncsEntityOfEntIndex(struct enginefuncs_s *t, int index) {
	return (*t->pfnPEntityOfEntIndex)(index);
}

edict_t* engineFuncsFindEntityByVars(struct enginefuncs_s *t, struct entvars_s *pvars) {
	return (*t->pfnFindEntityByVars)(pvars);
}

void* engineFuncsGetModelPtr(struct enginefuncs_s *t, edict_t *pEdict) {
	return (*t->pfnGetModelPtr)(pEdict);
}

int engineFuncsRegUserMsg(struct enginefuncs_s *t, const char *pszName, int iSize) {
	return (*t->pfnRegUserMsg)(pszName, iSize);
}

void engineFuncsAnimationAutomove(struct enginefuncs_s *t, const edict_t *pEdict, float flTime) {
	(*t->pfnAnimationAutomove)(pEdict, flTime);
}

void engineFuncsGetBonePosition(struct enginefuncs_s *t, const edict_t *pEdict, int iBone, float *rgflOrigin, float *rgflAngles) {
	(*t->pfnGetBonePosition)(pEdict, iBone, rgflOrigin, rgflAngles);
}

uint32 engineFuncsFunctionFromName(struct enginefuncs_s *t, const char *pName) {
	return (*t->pfnFunctionFromName)(pName);
}

const char* engineFuncsNameForFunction(struct enginefuncs_s *t, uint32 function) {
	return (*t->pfnNameForFunction)(function);
}

void engineFuncsClientPrintf(struct enginefuncs_s *t, edict_t *pEdict, PRINT_TYPE ptype, const char *szMsg) {
	(*t->pfnClientPrintf)(pEdict, ptype, szMsg);
}

void engineFuncsServerPrint(struct enginefuncs_s *t, const char *szMsg) {
	(*t->pfnServerPrint)(szMsg);
}

const char* engineFuncsCmd_Args(struct enginefuncs_s *t) {
	return (*t->pfnCmd_Args)();
}

const char* engineFuncsCmd_Argv(struct enginefuncs_s *t, int argc) {
	return (*t->pfnCmd_Argv)(argc);
}

int engineFuncsCmd_Argc(struct enginefuncs_s *t) {
	return (*t->pfnCmd_Argc)();
}

void engineFuncsAddServerCommand(struct enginefuncs_s *t, char *cmd_name, void *f) {
	engineFuncs = t;

	setGoCallback(SERVER_COMMAND_CALLBACKS_CATEGORY, cmd_name, f);

	void ff(void) {
		void *f;
		getGoCallback(SERVER_COMMAND_CALLBACKS_CATEGORY, engineFuncsCmd_Argv(engineFuncs, 0), &f);

		int argc = engineFuncsCmd_Argc(engineFuncs);

		char **argv = (char**)malloc(argc * sizeof(char*));

		for (int i = 0; i < argc; i++) {
			argv[i] = engineFuncsCmd_Argv(engineFuncs, i);
		}

		callGoFunction(f, argc, argv);
		free(argv);
    }

	(*t->pfnAddServerCommand)(cmd_name, &ff);
}

void engineFuncsGetAttachment(struct enginefuncs_s *t, const edict_t *pEdict, int iAttachment, float *rgflOrigin, float *rgflAngles) {
	(*t->pfnGetAttachment)(pEdict, iAttachment, rgflOrigin, rgflAngles);
}

void engineFuncsCRC32_Init(struct enginefuncs_s *t, CRC32_t *pulCRC) {
	(*t->pfnCRC32_Init)(pulCRC);
}

void engineFuncsCRC32_ProcessBuffer(struct enginefuncs_s *t, CRC32_t *pulCRC, void *p, int len) {
	(*t->pfnCRC32_ProcessBuffer)(pulCRC, p, len);
}

void engineFuncsCRC32_ProcessByte(struct enginefuncs_s *t, CRC32_t *pulCRC, unsigned char ch) {
	(*t->pfnCRC32_ProcessByte)(pulCRC, ch);
}

CRC32_t engineFuncsCRC32_Final(struct enginefuncs_s *t, CRC32_t pulCRC) {
	return (*t->pfnCRC32_Final)(pulCRC);
}

int engineFuncsRandomLong(struct enginefuncs_s *t, int lLow, int lHigh) {
	return (*t->pfnRandomLong)(lLow, lHigh);
}

float engineFuncsRandomFloat(struct enginefuncs_s *t, float flLow, float flHigh) {
	return (*t->pfnRandomFloat)(flLow, flHigh);
}

void engineFuncsSetView(struct enginefuncs_s *t, const edict_t *pClient, const edict_t *pViewent) {
	(*t->pfnSetView)(pClient, pViewent);
}

float engineFuncsTime(struct enginefuncs_s *t) {
	return (*t->pfnTime)();
}

void engineFuncsCrosshairAngle(struct enginefuncs_s *t, const edict_t *pClient, float pitch, float yaw) {
	(*t->pfnCrosshairAngle)(pClient, pitch, yaw);
}

byte* engineFuncsLoadFileForMe(struct enginefuncs_s *t, const char *filename, int *pLength) {
	return (*t->pfnLoadFileForMe)(filename, pLength);
}

void engineFuncsFreeFile(struct enginefuncs_s *t, void *buffer) {
	(*t->pfnFreeFile)(buffer);
}

void engineFuncsEndSection(struct enginefuncs_s *t, const char *pszSectionName) {
	(*t->pfnEndSection)(pszSectionName);
}

int engineFuncsCompareFileTime(struct enginefuncs_s *t, const char *filename1, const char *filename2, int *iCompare) {
	return (*t->pfnCompareFileTime)(filename1, filename2, iCompare);
}

void engineFuncsGetGameDir(struct enginefuncs_s *t, char *szGetGameDir) {
	(*t->pfnGetGameDir)(szGetGameDir);
}

void engineFuncsCvar_RegisterVariable(struct enginefuncs_s *t, cvar_t *variable) {
	(*t->pfnCvar_RegisterVariable)(variable);
}

void engineFuncsFadeClientVolume(struct enginefuncs_s *t, const edict_t *pEdict, int fadePercent, int fadeOutSeconds, int holdTime, int fadeInSeconds) {
	(*t->pfnFadeClientVolume)(pEdict, fadePercent, fadeOutSeconds, holdTime, fadeInSeconds);
}

void engineFuncsSetClientMaxspeed(struct enginefuncs_s *t, const edict_t *pEdict, float fNewMaxspeed) {
	(*t->pfnSetClientMaxspeed)(pEdict, fNewMaxspeed);
}

edict_t* engineFuncsCreateFakeClient(struct enginefuncs_s *t, const char *netname) {
	return (*t->pfnCreateFakeClient)(netname);
}

void engineFuncsRunPlayerMove(struct enginefuncs_s *t, edict_t *fakeClient, const float *viewangles, float forwardmove, float sidemove, float upmove, unsigned short buttons, byte impulse, byte msec) {
	(*t->pfnRunPlayerMove)(fakeClient, viewangles, forwardmove, sidemove, upmove, buttons, impulse, msec);
}

int engineFuncsNumberOfEntities(struct enginefuncs_s *t) {
	return (*t->pfnNumberOfEntities)();
}

char* engineFuncsGetInfoKeyBuffer(struct enginefuncs_s *t, edict_t *e) {
	return (*t->pfnGetInfoKeyBuffer)(e);
}

char* engineFuncsInfoKeyValue(struct enginefuncs_s *t, char *infobuffer, char *key) {
	return (*t->pfnInfoKeyValue)(infobuffer, key);
}

void engineFuncsSetKeyValue(struct enginefuncs_s *t, char *infobuffer, char *key, char *value) {
	(*t->pfnSetKeyValue)(infobuffer, key, value);
}

void engineFuncsSetClientKeyValue(struct enginefuncs_s *t, int clientIndex, char *infobuffer, char *key, char *value) {
	(*t->pfnSetClientKeyValue)(clientIndex, infobuffer, key, value);
}

int engineFuncsIsMapValid(struct enginefuncs_s *t, char *filename) {
	return (*t->pfnIsMapValid)(filename);
}

void engineFuncsStaticDecal(struct enginefuncs_s *t, const float *origin, int decalIndex, int entityIndex, int modelIndex) {
	(*t->pfnStaticDecal)(origin, decalIndex, entityIndex, modelIndex);
}

int engineFuncsPrecacheGeneric(struct enginefuncs_s *t, char *s) {
	return (*t->pfnPrecacheGeneric)(s);
}

int engineFuncsGetPlayerUserId(struct enginefuncs_s *t, edict_t *e) {
	return (*t->pfnGetPlayerUserId)(e);
}

void engineFuncsBuildSoundMsg(struct enginefuncs_s *t, edict_t *entity, int channel, const char *sample, float volume, float attenuation, int fFlags, int pitch, int msg_dest, int msg_type, const float *pOrigin, edict_t *ed) {
	(*t->pfnBuildSoundMsg)(entity, channel, sample, volume, attenuation, fFlags, pitch, msg_dest, msg_type, pOrigin, ed);
}

int engineFuncsIsDedicatedServer(struct enginefuncs_s *t) {
	return (*t->pfnIsDedicatedServer)();
}

cvar_t* engineFuncsCVarGetPointer(struct enginefuncs_s *t, const char *szVarName) {
	return (*t->pfnCVarGetPointer)(szVarName);
}

unsigned int engineFuncsGetPlayerWONId(struct enginefuncs_s *t, edict_t *e) {
	return (*t->pfnGetPlayerWONId)(e);
}

void engineFuncsInfo_RemoveKey(struct enginefuncs_s *t, char *s, char *key) {
	(*t->pfnInfo_RemoveKey)(s, key);
}

char* engineFuncsGetPhysicsKeyValue(struct enginefuncs_s *t, const edict_t *pClient, const char *key) {
	return (*t->pfnGetPhysicsKeyValue)(pClient, key);
}

void engineFuncsSetPhysicsKeyValue(struct enginefuncs_s *t, const edict_t *pClient, const char *key, const char *value) {
	(*t->pfnSetPhysicsKeyValue)(pClient, key, value);
}

const char* engineFuncsGetPhysicsInfoString(struct enginefuncs_s *t, const edict_t *pClient) {
	return (*t->pfnGetPhysicsInfoString)(pClient);
}

unsigned short engineFuncsPrecacheEvent(struct enginefuncs_s *t, int type, const char *psz) {
	return (*t->pfnPrecacheEvent)(type, psz);
}

void engineFuncsPlaybackEvent(struct enginefuncs_s *t, int flags, const edict_t *pInvoker, unsigned short eventindex, float delay, float *origin, float *angles, float fparam1, float fparam2, int iparam1, int iparam2, int bparam1, int bparam2) {
	(*t->pfnPlaybackEvent)(flags, pInvoker, eventindex, delay, origin, angles, fparam1, fparam2, iparam1, iparam2, bparam1, bparam2);
}

unsigned char* engineFuncsSetFatPVS(struct enginefuncs_s *t, float *org) {
	return (*t->pfnSetFatPVS)(org);
}

unsigned char* engineFuncsSetFatPAS(struct enginefuncs_s *t, float *org) {
	return (*t->pfnSetFatPAS)(org);
}

int engineFuncsCheckVisibility(struct enginefuncs_s *t, const edict_t *entity, unsigned char *pset) {
	return (*t->pfnCheckVisibility)(entity, pset);
}

void engineFuncsDeltaSetField(struct enginefuncs_s *t, struct delta_s *pFields, const char *fieldname) {
	(*t->pfnDeltaSetField)(pFields, fieldname);
}

void engineFuncsDeltaUnsetField(struct enginefuncs_s *t, struct delta_s *pFields, const char *fieldname) {
	(*t->pfnDeltaUnsetField)(pFields, fieldname);
}

void engineFuncsDeltaAddEncoder(struct enginefuncs_s *t, char *name, void (*conditionalencode)(struct delta_s *pFields, const unsigned char *from, const unsigned char *to)) {
	(*t->pfnDeltaAddEncoder)(name, conditionalencode);
}

int engineFuncsGetCurrentPlayer(struct enginefuncs_s *t) {
	return (*t->pfnGetCurrentPlayer)();
}

int engineFuncsCanSkipPlayer(struct enginefuncs_s *t, const edict_t *player) {
	return (*t->pfnCanSkipPlayer)(player);
}

int engineFuncsDeltaFindField(struct enginefuncs_s *t, struct delta_s *pFields, const char *fieldname) {
	return (*t->pfnDeltaFindField)(pFields, fieldname);
}

void engineFuncsDeltaSetFieldByIndex(struct enginefuncs_s *t, struct delta_s *pFields, int fieldNumber) {
	(*t->pfnDeltaSetFieldByIndex)(pFields, fieldNumber);
}

void engineFuncsDeltaUnsetFieldByIndex(struct enginefuncs_s *t, struct delta_s *pFields, int fieldNumber) {
	(*t->pfnDeltaUnsetFieldByIndex)(pFields, fieldNumber);
}

void engineFuncsSetGroupMask(struct enginefuncs_s *t, int mask, int op) {
	(*t->pfnSetGroupMask)(mask, op);
}

int engineFuncsCreateInstancedBaseline(struct enginefuncs_s *t, int classname, struct entity_state_s *baseline) {
	return (*t->pfnCreateInstancedBaseline)(classname, baseline);
}

void engineFuncsCvar_DirectSet(struct enginefuncs_s *t, cvar_t *var, char *value) {
	(*t->pfnCvar_DirectSet)(var, value);
}

void engineFuncsForceUnmodified(struct enginefuncs_s *t, FORCE_TYPE type, float *rgflOrigin, float *rgflAngles, const char *szKeyName) {
	(*t->pfnForceUnmodified)(type, rgflOrigin, rgflAngles, szKeyName);
}

void engineFuncsGetPlayerStats(struct enginefuncs_s *t, const edict_t *pClient, int *ping, int *packet_loss) {
	(*t->pfnGetPlayerStats)(pClient, ping, packet_loss);
}

int engineFuncsVoice_GetClientListening(struct enginefuncs_s *t, int iReceiver, int iSender) {
	return (*t->pfnVoice_GetClientListening)(iReceiver, iSender);
}

int engineFuncsVoice_SetClientListening(struct enginefuncs_s *t, int iReceiver, int iSender, int bListen) {
	return (*t->pfnVoice_SetClientListening)(iReceiver, iSender, bListen);
}

const char* engineFuncsGetPlayerAuthId(struct enginefuncs_s *t, edict_t *e) {
	return (*t->pfnGetPlayerAuthId)(e);
}

sequenceEntry_s engineFuncsSequenceGet(struct enginefuncs_s *t, const char *fileName, const char *entryName) {
	(*t->pfnSequenceGet)(fileName, entryName);
}

sentenceEntry_s engineFuncsSequencePickSentence(struct enginefuncs_s *t, const char *groupName, int pickMethod, int *picked) {
	(*t->pfnSequencePickSentence)(groupName, pickMethod, picked);
}

int engineFuncsGetFileSize(struct enginefuncs_s *t, char *filename) {
	return (*t->pfnGetFileSize)(filename);
}

unsigned int engineFuncsGetApproxWavePlayLen(struct enginefuncs_s *t, const char *filepath) {
	return (*t->pfnGetApproxWavePlayLen)(filepath);
}

int engineFuncsIsCareerMatch(struct enginefuncs_s *t) {
	return (*t->pfnIsCareerMatch)();
}

int engineFuncsGetLocalizedStringLength(struct enginefuncs_s *t, const char *label) {
	return (*t->pfnGetLocalizedStringLength)(label);
}

void engineFuncsRegisterTutorMessageShown(struct enginefuncs_s *t, int mid) {
	(*t->pfnRegisterTutorMessageShown)(mid);
}

int engineFuncsGetTimesTutorMessageShown(struct enginefuncs_s *t, int mid) {
	return (*t->pfnGetTimesTutorMessageShown)(mid);
}

void engineFuncsProcessTutorMessageDecayBuffer(struct enginefuncs_s *t, int *buffer, int bufferLength) {
	(*t->ProcessTutorMessageDecayBuffer)(buffer, bufferLength);
}

void engineFuncsConstructTutorMessageDecayBuffer(struct enginefuncs_s *t, int *buffer, int bufferLength) {
	(*t->ConstructTutorMessageDecayBuffer)(buffer, bufferLength);
}

void engineFuncsResetTutorMessageDecayData(struct enginefuncs_s *t) {
	(*t->ResetTutorMessageDecayData)();
}

void engineFuncsQueryClientCvarValue(struct enginefuncs_s *t, edict_t *player, const char *cvarName) {
	(*t->pfnQueryClientCvarValue)(player, cvarName);
}

void engineFuncsQueryClientCvarValue2(struct enginefuncs_s *t, edict_t *player, const char *cvarName, int requestID) {
	(*t->pfnQueryClientCvarValue2)(player, cvarName, requestID);
}

void engineFuncsCheckParm(struct enginefuncs_s *t, const char *psz, char *pchParam) {
	(*t->pfnCheckParm)(psz, pchParam);
}

edict_t* engineFuncsPEntityOfEntIndexAllEntities(struct enginefuncs_s *t, int iEntIndex) {
	return (*t->pfnPEntityOfEntIndexAllEntities)(iEntIndex);
}

*/
import "C"

import (
	"unsafe"
)

type EngineFuncs struct {
	p          *C.enginefuncs_t
	globalVars *GlobalVars

	stringCache *inmemoryCache[string, int]
}

func NewEngineFuncs(p *C.enginefuncs_t, globalVars *GlobalVars) *EngineFuncs {
	return &EngineFuncs{
		p:          p,
		globalVars: globalVars,

		stringCache: newInmemoryCache[string, int](),
	}
}

func (ef *EngineFuncs) PrecacheModel(name string) int {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	return int(C.engineFuncsPrecacheModel(ef.p, cs))
}

func (ef *EngineFuncs) PrecacheSound(name string) int {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	return int(C.engineFuncsPrecacheSound(ef.p, cs))
}

func (ef *EngineFuncs) SetModel(e *Edict, model string) {
	cs := C.CString(model)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsSetModel(ef.p, e.p, cs)
}

func (ef *EngineFuncs) AddServerCommand(name string, callback func(int, ...string)) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	f := unsafe.Pointer(&callback)

	C.engineFuncsAddServerCommand(ef.p, cs, f)
}

func (ef *EngineFuncs) EntityOfEntIndex(index int) *Edict {
	edict := C.engineFuncsEntityOfEntIndex(ef.p, C.int(index))

	return edictFromC(ef.globalVars.p, edict)
}

func (ef *EngineFuncs) MessageBegin(
	msgDest int,
	msgType int,
	pOrigin float32,
	edict *Edict,
) {
	C.engineFuncsMessageBegin(
		ef.p,
		C.int(msgDest),
		C.int(msgType),
		(*C.float)(&pOrigin),
		edict.p,
	)
}

func (ef *EngineFuncs) MessageEnd() {
	C.engineFuncsMessageEnd(ef.p)
}

func (ef *EngineFuncs) MessageWriteByte(i int) {
	C.engineFuncsWriteByte(ef.p, C.int(i))
}

func (ef *EngineFuncs) MessageWriteChar(i int) {
	C.engineFuncsWriteChar(ef.p, C.int(i))
}

func (ef *EngineFuncs) MessageWriteShort(i int) {
	C.engineFuncsWriteShort(ef.p, C.int(i))
}

func (ef *EngineFuncs) MessageWriteLong(i int) {
	C.engineFuncsWriteLong(ef.p, C.int(i))
}

func (ef *EngineFuncs) MessageWriteAngle(f float32) {
	C.engineFuncsWriteAngle(ef.p, C.float(f))
}

func (ef *EngineFuncs) MessageWriteCoord(f float32) {
	C.engineFuncsWriteCoord(ef.p, C.float(f))
}

func (ef *EngineFuncs) MessageWriteString(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsWriteString(ef.p, cs)
}

func (ef *EngineFuncs) MessageWriteEntity(id int) {
	C.engineFuncsWriteEntity(ef.p, C.int(id))
}

func (ef *EngineFuncs) CreateEntity() *Edict {
	e := C.engineFuncsCreateEntity(ef.p)
	return edictFromC(ef.globalVars.p, e)
}

func (ef *EngineFuncs) CreateNamedEntity(className string) *Edict {
	cs := C.CString(className)
	defer C.free(unsafe.Pointer(cs))

	engineString := ef.AllocString(className)

	e := C.engineFuncsCreateNamedEntity(ef.p, C.int(engineString))
	return edictFromC(ef.globalVars.p, e)
}

func (ef *EngineFuncs) RemoveEntity(e *Edict) {
	C.engineFuncsRemoveEntity(ef.p, e.p)
}

func (ef *EngineFuncs) AllocString(s string) int {
	if v, ok := ef.stringCache.Get(s); ok {
		return v
	}

	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	alloc := int(C.engineFuncsAllocString(ef.p, cs))
	ef.stringCache.Set(s, alloc)

	return alloc
}

func (ef *EngineFuncs) TraceLine(
	v1, v2 [3]float32,
	noMonsters int,
	pentToSkip *Edict,
) *TraceResult {
	var tr C.TraceResult

	var pent *C.edict_t
	if pentToSkip != nil {
		pent = pentToSkip.p
	}

	C.engineFuncsTraceLine(
		ef.p,
		(*C.float)(&v1[0]),
		(*C.float)(&v2[0]),
		C.int(noMonsters),
		pent,
		&tr,
	)

	return traceResultFromC(P.GlobalVars.p, tr)
}
