package metamod_go

/*
#include <eiface.h>

#include <stdlib.h>
#include <stdio.h>
#include <com_model.h>

#define SERVER_COMMAND_CALLBACKS_CATEGORY "server_commands"

extern const char* ReadString(globalvars_t *gpGlobals, int offset);
extern int MakeString(globalvars_t *gpGlobals, char *str);

extern void callGoFunction(void *f, int argc, char **argv);
extern void getGoCallback(char *category, char *v, void **f);

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

int engineFuncsEntIsOnFloor(struct enginefuncs_s *t, edict_t *ent) {
	return (*t->pfnEntIsOnFloor)(ent);
}

int engineFuncsDropToFloor(struct enginefuncs_s *t, edict_t *ent) {
	(*t->pfnDropToFloor)(ent);
}

int engineFuncsWalkMove(struct enginefuncs_s *t, edict_t *ent, float yaw, float dist, int iMode) {
	return (*t->pfnWalkMove)(ent, yaw, dist, iMode);
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

texture_t* engineFuncsTraceTexture(struct enginefuncs_s *t, edict_t *pent, const float *v1, const float *v2) {
	return (texture_t*)(*t->pfnTraceTexture)(pent, v1, v2);
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

void* engineFuncsPvAllocEntPrivateData(struct enginefuncs_s *t, edict_t *ent, int32 cb) {
	return (*t->pfnPvAllocEntPrivateData)(ent, cb);
}

void* engineFuncsPvEntPrivateData(struct enginefuncs_s *t, edict_t *ent) {
	return (*t->pfnPvEntPrivateData)(ent);
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

void engineFuncsAddServerCommand(struct enginefuncs_s *t, char *cmd_name) {
	engineFuncs = t;

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

void* engineFuncsLoadFileForMe(struct enginefuncs_s *t, const char *filename, int *pLength) {
	return (void*)(*t->pfnLoadFileForMe)(filename, pLength);
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

void engineFuncsRunPlayerMove(struct enginefuncs_s *t, edict_t *fakeClient, float *viewangles, float forwardmove, float sidemove, float upmove, unsigned short buttons, byte impulse, byte msec) {
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

unsigned char* engineFuncsSetFatPVS(struct enginefuncs_s *t, float *origin) {
	return (*t->pfnSetFatPVS)(origin);
}

unsigned char* engineFuncsSetFatPAS(struct enginefuncs_s *t, float *origin) {
	return (*t->pfnSetFatPAS)(origin);
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
	"fmt"
	"github.com/et-nik/metamod-go/engine"
	"github.com/et-nik/metamod-go/vector"
	"strings"
	"unsafe"
)

const (
	maxUserMsgName = 12
	maxPath        = 260
)

type EngineFuncs struct {
	p          *C.enginefuncs_t
	globalVars *GlobalVars

	stringCache *inmemoryCache[string, int]
}

func newEngineFuncs(p *C.enginefuncs_t, globalVars *GlobalVars) *EngineFuncs {
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

func (ef *EngineFuncs) ModelIndex(name string) int {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	return int(C.engineFuncsModelIndex(ef.p, cs))
}

func (ef *EngineFuncs) ModelFrames(index int) int {
	return int(C.engineFuncsModelFrames(ef.p, C.int(index)))
}

func (ef *EngineFuncs) SetSize(e *Edict, mins, maxs vector.Vector) {
	C.engineFuncsSetSize(
		ef.p,
		e.p,
		(*C.float)(&mins[0]),
		(*C.float)(&maxs[0]),
	)
}

func (ef *EngineFuncs) ChangeLevel(levelName string, landmark string) {
	csLevelName := C.CString(levelName)
	defer C.free(unsafe.Pointer(csLevelName))

	var csLandmark *C.char
	if landmark != "" {
		csLandmark = C.CString(landmark)
		defer C.free(unsafe.Pointer(csLandmark))
	}

	C.engineFuncsChangeLevel(ef.p, csLevelName, csLandmark)
}

func (ef *EngineFuncs) VecToYaw(vec vector.Vector) float32 {
	return float32(C.engineFuncsVecToYaw(ef.p, (*C.float)(&vec[0])))
}

func (ef *EngineFuncs) VecToAngles(vec vector.Vector) vector.Vector {
	var angles vector.Vector
	C.engineFuncsVecToAngles(ef.p, (*C.float)(&vec[0]), (*C.float)(&angles[0]))

	return angles
}

func (ef *EngineFuncs) MoveToOrigin(e *Edict, goal vector.Vector, dist float32, moveType engine.MoveType) {
	C.engineFuncsMoveToOrigin(
		ef.p,
		e.p,
		(*C.float)(&goal[0]),
		C.float(dist),
		C.int(int(moveType)),
	)
}

func (ef *EngineFuncs) ChangeYaw(e *Edict) {
	C.engineFuncsChangeYaw(ef.p, e.p)
}

func (ef *EngineFuncs) ChangePitch(e *Edict) {
	C.engineFuncsChangePitch(ef.p, e.p)
}

func (ef *EngineFuncs) FindEntityByString(start *Edict, field engine.FindEntityField, value string) *Edict {
	csField := C.CString(string(field))
	defer C.free(unsafe.Pointer(csField))

	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))

	e := C.engineFuncsFindEntityByString(ef.p, start.ptr(), csField, csValue)

	return edictFromC(ef.globalVars.p, e)
}

func (ef *EngineFuncs) GetEntityIllum(e *Edict) int {
	return int(C.engineFuncsGetEntityIllum(ef.p, e.p))
}

func (ef *EngineFuncs) FindEntityInSphere(start *Edict, origin vector.Vector, radius float32) *Edict {
	e := C.engineFuncsFindEntityInSphere(
		ef.p,
		start.ptr(),
		(*C.float)(&origin[0]),
		C.float(radius),
	)

	return edictFromC(ef.globalVars.p, e)
}

// FindClientInPVS Finds a client in the Potentially Visible Set of the given entity.
func (ef *EngineFuncs) FindClientInPVS(e *Edict) *Edict {
	edict := C.engineFuncsFindClientInPVS(ef.p, e.p)

	return edictFromC(ef.globalVars.p, edict)
}

func (ef *EngineFuncs) EntitiesInPVS(e *Edict) *Edict {
	edict := C.engineFuncsEntitiesInPVS(ef.p, e.p)

	return edictFromC(ef.globalVars.p, edict)
}

func (ef *EngineFuncs) MakeVectors(angles vector.Vector) {
	C.engineFuncsMakeVectors(ef.p, (*C.float)(&angles[0]))
}

func (ef *EngineFuncs) AngleVectors(vector vector.Vector, forward, right, up vector.Vector) {
	C.engineFuncsAngleVectors(
		ef.p,
		(*C.float)(&vector[0]),
		(*C.float)(&forward[0]),
		(*C.float)(&right[0]),
		(*C.float)(&up[0]),
	)
}

func (ef *EngineFuncs) MakeStatic(e *Edict) {
	C.engineFuncsMakeStatic(ef.p, e.p)
}

func (ef *EngineFuncs) EntIsOnFloor(e *Edict) bool {
	return int(C.engineFuncsEntIsOnFloor(ef.p, e.p)) == 1
}

func (ef *EngineFuncs) DropToFloor(e *Edict) int {
	return int(C.engineFuncsDropToFloor(ef.p, e.p))
}

func (ef *EngineFuncs) WalkMove(e *Edict, yaw float32, dist float32, mode engine.WalkMoveMode) int {
	return int(C.engineFuncsWalkMove(ef.p, e.p, C.float(yaw), C.float(dist), C.int(mode)))
}

func (ef *EngineFuncs) SetOrigin(e *Edict, origin vector.Vector) {
	C.engineFuncsSetOrigin(ef.p, e.p, (*C.float)(&origin[0]))
}

func (ef *EngineFuncs) EmitSound(e *Edict, channel int, sample string, volume, attenuation float32, flags int, pitch int) {
	cs := C.CString(sample)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsEmitSound(
		ef.p,
		e.p,
		C.int(channel),
		cs,
		C.float(volume),
		C.float(attenuation),
		C.int(flags),
		C.int(pitch),
	)
}

func (ef *EngineFuncs) EmitAmbientSound(
	e *Edict,
	position vector.Vector,
	sample string,
	volume, attenuation float32,
	flags int,
	pitch int,
) {
	cs := C.CString(sample)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsEmitAmbientSound(
		ef.p,
		e.p,
		(*C.float)(&position[0]),
		cs,
		C.float(volume),
		C.float(attenuation),
		C.int(flags),
		C.int(pitch),
	)
}

func (ef *EngineFuncs) AddServerCommand(name string, callback func(int, ...string)) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	if ef.p == nil {
		panic("enginefuncs is nil, looks like the plugin is not running")
	}

	setGoCallback(C.SERVER_COMMAND_CALLBACKS_CATEGORY, name, callback)

	C.engineFuncsAddServerCommand(ef.p, cs)
}

// EntityOfEntIndex Gets the edict at the given entity index.
// If the given index is not valid, returns null.
//   - Otherwise, if the entity at the given index is not in use, returns null.
//   - Otherwise, if the entity at the given index is equal or more than svs.maxclients and does not have a CBaseEntity instance, returns null.
//   - Otherwise, returns the entity.
func (ef *EngineFuncs) EntityOfEntIndex(index int) *Edict {
	edict := C.engineFuncsEntityOfEntIndex(ef.p, C.int(index))

	return edictFromC(ef.globalVars.p, edict)
}

func (ef *EngineFuncs) MessageBegin(
	msgDest int,
	msgType int,
	pOrigin *vector.Vector,
	edict *Edict,
) {
	var orgPtr *C.float
	if pOrigin != nil {
		orgPtr = (*C.float)(&pOrigin[0])
	}

	var edictPtr *C.edict_t
	if edict != nil {
		edictPtr = edict.p
	}

	C.engineFuncsMessageBegin(
		ef.p,
		C.int(msgDest),
		C.int(msgType),
		orgPtr,
		edictPtr,
	)
}

func (ef *EngineFuncs) MessageEnd() {
	C.engineFuncsMessageEnd(ef.p)
}

func (ef *EngineFuncs) MessageWriteByte(b int) {
	C.engineFuncsWriteByte(ef.p, C.int(b))
}

func (ef *EngineFuncs) MessageWriteChar(c int) {
	C.engineFuncsWriteChar(ef.p, C.int(c))
}

func (ef *EngineFuncs) MessageWriteShort(s int) {
	C.engineFuncsWriteShort(ef.p, C.int(s))
}

func (ef *EngineFuncs) MessageWriteLong(l int) {
	C.engineFuncsWriteLong(ef.p, C.int(l))
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

// SzFromIndex Gets a string from an index.
func (ef *EngineFuncs) SzFromIndex(index int) string {
	return C.GoString(C.engineFuncsSzFromIndex(ef.p, C.int(index)))
}

func (ef *EngineFuncs) StringFromIndex(index int) string {
	return ef.SzFromIndex(index)
}

func (ef *EngineFuncs) TraceLine(
	v1, v2 vector.Vector,
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

	return traceResultFromC(globalPluginState.globalVars.p, tr)
}

// TraceToss Traces a toss.
// This simulates tossing the entity using its current origin, velocity, angular velocity, angles and gravity.
// Note that this does not use the same code as MOVETYPE_TOSS, and may return different results.
func (ef *EngineFuncs) TraceToss(pent, pentToIgnore *Edict) *TraceResult {
	var tr C.TraceResult

	var pentToIgnoreC *C.edict_t
	if pentToIgnore != nil {
		pentToIgnoreC = pentToIgnore.p
	}

	C.engineFuncsTraceToss(
		ef.p,
		pent.p,
		pentToIgnoreC,
		&tr,
	)

	return traceResultFromC(globalPluginState.globalVars.p, tr)
}

// TraceMonsterHull Performs a trace between a starting and ending position, using the given entity's mins and maxs.
// This can be any entity, not just monsters.
// Returns true if the trace was entirely in a solid object, or if it hit something.
func (ef *EngineFuncs) TraceMonsterHull(
	pent *Edict,
	v1, v2 vector.Vector,
	noMonsters int,
	pentToSkip *Edict,
) (*TraceResult, int) {
	var tr C.TraceResult

	var pentToSkipC *C.edict_t
	if pentToSkip != nil {
		pentToSkipC = pentToSkip.p
	}

	result := C.engineFuncsTraceMonsterHull(
		ef.p,
		pent.p,
		(*C.float)(&v1[0]),
		(*C.float)(&v2[0]),
		C.int(noMonsters),
		pentToSkipC,
		&tr,
	)

	return traceResultFromC(globalPluginState.globalVars.p, tr), int(result)
}

// TraceHull Performs a trace between a starting and ending position, using the given hull number.
func (ef *EngineFuncs) TraceHull(
	v1, v2 vector.Vector,
	noMonsters, hullNumber int,
	pentToSkip *Edict,
) *TraceResult {
	var tr C.TraceResult

	var pentToSkipC *C.edict_t
	if pentToSkip != nil {
		pentToSkipC = pentToSkip.p
	}

	C.engineFuncsTraceHull(
		ef.p,
		(*C.float)(&v1[0]),
		(*C.float)(&v2[0]),
		C.int(noMonsters),
		C.int(hullNumber),
		pentToSkipC,
		&tr,
	)

	return traceResultFromC(globalPluginState.globalVars.p, tr)
}

// TraceModel Performs a trace between a starting and ending position.
// Similar to TraceHull, but will instead perform a trace in the given world hull using the given entity's model's hulls.
// For studio models this will use the model's hitboxes.
// If the given entity's model is a studio model, uses its hitboxes.
// If it's a brush model, the brush model's hull for the given hull number is used (this may differ if custom brush hull sizes are in use).
// Otherwise, the entity bounds are converted into a hull.
func (ef *EngineFuncs) TraceModel(
	v1, v2 vector.Vector,
	hullNumber int,
	pent *Edict,
) *TraceResult {
	var tr C.TraceResult

	C.engineFuncsTraceModel(
		ef.p,
		(*C.float)(&v1[0]),
		(*C.float)(&v2[0]),
		C.int(hullNumber),
		pent.p,
		&tr,
	)

	return traceResultFromC(globalPluginState.globalVars.p, tr)
}

// TraceTexture Used to get texture info.
// The given entity must have a brush model set.
// If the traceline intersects the model, the texture of the surface it intersected is returned.
func (ef *EngineFuncs) TraceTexture(
	pent *Edict,
	v1, v2 vector.Vector,
) *Texture {
	texture := C.engineFuncsTraceTexture(
		ef.p,
		pent.p,
		(*C.float)(&v1[0]),
		(*C.float)(&v2[0]),
	)

	return textureFromC(texture)
}

func (ef *EngineFuncs) GetAimVector(ent *Edict, speed float32) vector.Vector {
	var vec vector.Vector
	C.engineFuncsGetAimVector(ef.p, ent.p, C.float(speed), (*C.float)(&vec[0]))

	return vec
}

func (ef *EngineFuncs) ServerCommand(str string) {
	if str == "" {
		return
	}

	if !strings.HasSuffix(str, "\n") || !strings.HasSuffix(str, ";") {
		str += "\n"
	}

	cs := C.CString(str)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsServerCommand(ef.p, cs)
}

func (ef *EngineFuncs) ServerExecute() {
	C.engineFuncsServerExecute(ef.p)
}

// ClientCommand Sends a command to the client.
func (ef *EngineFuncs) ClientCommand(pEdict *Edict, str string) {
	if pEdict == nil {
		return
	}

	if str == "" {
		return
	}

	if !strings.HasSuffix(str, "\n") || !strings.HasSuffix(str, ";") {
		str += "\n"
	}

	cs := C.CString(str)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsClientCommand(ef.p, pEdict.p, cs)
}

// ParticleEffect Creates a particle effect.
func (ef *EngineFuncs) ParticleEffect(
	origin, direction vector.Vector,
	color, count float32,
) {
	C.engineFuncsParticleEffect(
		ef.p,
		(*C.float)(&origin[0]),
		(*C.float)(&direction[0]),
		C.float(color),
		C.float(count),
	)
}

// LightStyle Sets a light style.
func (ef *EngineFuncs) LightStyle(style int, value string) {
	cs := C.CString(value)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsLightStyle(ef.p, C.int(style), cs)
}

// DecalIndex Returns the index of a decal.
func (ef *EngineFuncs) DecalIndex(name string) int {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	return int(C.engineFuncsDecalIndex(ef.p, cs))
}

// PointContents Gets the contents of the given location in the world.
func (ef *EngineFuncs) PointContents(v vector.Vector) int {
	return int(C.engineFuncsPointContents(
		ef.p,
		(*C.float)(&v[0]),
	))
}

// CVarRegister Registers a cvar.
// Sets the flag FCVAR_EXTDLL on the cvar.
// It returns the relevant cvar pointer you can use to get the value.
func (ef *EngineFuncs) CVarRegister(cvar *CVar) *CVar {
	C.engineFuncsCVarRegister(ef.p, cvar.p)

	return ef.CVarGetPointer(cvar.Name())
}

// CVarGetString Gets the value of a cvar as a string.
func (ef *EngineFuncs) CVarGetString(name string) string {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	return C.GoString(C.engineFuncsCVarGetString(ef.p, cs))
}

// CVarGetFloat Gets the value of a cvar as a float.
func (ef *EngineFuncs) CVarGetFloat(name string) float32 {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	return float32(C.engineFuncsCVarGetFloat(ef.p, cs))
}

// CVarSetFloat Sets the value of a cvar as a float.
func (ef *EngineFuncs) CVarSetFloat(name string, value float32) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsCVarSetFloat(ef.p, cs, C.float(value))
}

// CVarSetString Sets the value of a cvar as a string.
func (ef *EngineFuncs) CVarSetString(name, value string) {
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))

	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))

	C.engineFuncsCVarSetString(ef.p, csName, csValue)
}

// AlertMessage Sends a message to the server console.
func (ef *EngineFuncs) AlertMessage(alertType engine.AlertType, msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsAlertMessage(ef.p, C.ALERT_TYPE(alertType), cs)
}

// PvAllocEntPrivateData Allocates memory for CBaseEntity instances.
// The memory is freed when the entity is removed.
// It returns a pointer to the allocated memory.
func (ef *EngineFuncs) PvAllocEntPrivateData(ent *Edict, size int32) unsafe.Pointer {
	return C.engineFuncsPvAllocEntPrivateData(ef.p, ent.p, C.int32(size))
}

// PvEntPrivateData Gets the private data of an entity.
func (ef *EngineFuncs) PvEntPrivateData(ent *Edict) unsafe.Pointer {
	return C.engineFuncsPvEntPrivateData(ef.p, ent.p)
}

// FreeEntPrivateData Frees the private data of an entity.
func (ef *EngineFuncs) FreeEntPrivateData(ent *Edict) {
	C.engineFuncsFreeEntPrivateData(ef.p, ent.p)
}

// GetVarsOfEnt Gets the entvars_t of an entity.
func (ef *EngineFuncs) GetVarsOfEnt(ent *Edict) *EntVars {
	return entVarsFromC(ef.globalVars.p, C.engineFuncsGetVarsOfEnt(ef.p, ent.p))
}

// IndexOfEdict Gets the index of an entity.
func (ef *EngineFuncs) IndexOfEdict(pEdict *Edict) int {
	return int(C.engineFuncsIndexOfEdict(ef.p, pEdict.p))
}

// PEntityOfEntIndex Gets the edict at the given entity index.
// If the given index is not valid, returns null.
//   - Otherwise, if the entity at the given index is not in use, returns null.
//   - Otherwise, if the entity at the given index is equal or more than svs.maxclients and does not have a CBaseEntity instance, returns null.
//   - Otherwise, returns the entity.
func (ef *EngineFuncs) PEntityOfEntIndex(index int) *Edict {
	return ef.EntityOfEntIndex(index)
}

// FindEntityByVars Finds an entity by its vars.
// This will enumerate all entities, so this operation can be very expensive.
func (ef *EngineFuncs) FindEntityByVars(vars *EntVars) *Edict {
	return edictFromC(ef.globalVars.p, C.engineFuncsFindEntityByVars(ef.p, vars.p))
}

// GetModelPtr Gets the model pointer of an entity.
func (ef *EngineFuncs) GetModelPtr(pEdict *Edict) unsafe.Pointer {
	return C.engineFuncsGetModelPtr(ef.p, pEdict.p)
}

// RegUserMsg Registers a user message.
func (ef *EngineFuncs) RegUserMsg(name string, size int) int {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	return int(C.engineFuncsRegUserMsg(ef.p, cs, C.int(size)))
}

// FunctionFromName Gets the function ID from a name.
func (ef *EngineFuncs) FunctionFromName(name string) uint32 {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	return uint32(C.engineFuncsFunctionFromName(ef.p, cs))
}

// NameForFunction Gets the name of a function.
func (ef *EngineFuncs) NameForFunction(function uint32) string {
	return C.GoString(C.engineFuncsNameForFunction(ef.p, C.uint32(function)))
}

// ClientPrint Prints a message to a client.
func (ef *EngineFuncs) ClientPrint(pEdict *Edict, ptype engine.PrintType, msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsClientPrintf(ef.p, pEdict.p, C.PRINT_TYPE(ptype), cs)
}

// ServerPrint Prints a message to the server console.
func (ef *EngineFuncs) ServerPrint(msg string) {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsServerPrint(ef.p, cs)
}

func (ef *EngineFuncs) ServerPrintf(format string, args ...interface{}) {
	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}

	ef.ServerPrint(msg)
}

func (ef *EngineFuncs) GetAttachment(pEdict *Edict, attachmentIndex int, rgflOrigin, rgflAngles *vector.Vector) {
	C.engineFuncsGetAttachment(ef.p, pEdict.p, C.int(attachmentIndex), (*C.float)(&rgflOrigin[0]), (*C.float)(&rgflAngles[0]))
}

// RandomLong Generates a random long number between low and high.
func (ef *EngineFuncs) RandomLong(low, high int32) int32 {
	return int32(C.engineFuncsRandomLong(ef.p, C.int32(low), C.int32(high)))
}

func (ef *EngineFuncs) RandomFloat(low, high float32) float32 {
	return float32(C.engineFuncsRandomFloat(ef.p, C.float(low), C.float(high)))
}

func (ef *EngineFuncs) SetView(pClient, pViewent *Edict) {
	C.engineFuncsSetView(ef.p, pClient.p, pViewent.p)
}

func (ef *EngineFuncs) Time() float32 {
	return float32(C.engineFuncsTime(ef.p))
}

// CrosshairAngle Sets the angles of the given player's crosshairs to the given settings.
// Set both to 0 to disable.
func (ef *EngineFuncs) CrosshairAngle(pClient *Edict, pitch, yaw float32) {
	C.engineFuncsCrosshairAngle(ef.p, pClient.p, C.float(pitch), C.float(yaw))
}

// LoadFileForMe Loads a file from disk.
// filename Name of the file. Path starts in the game directory.
func (ef *EngineFuncs) LoadFileForMe(filename string) ([]byte, error) {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))

	var length C.int
	bufferPtr := C.engineFuncsLoadFileForMe(ef.p, cs, &length)
	if bufferPtr == nil {
		return nil, nil
	}

	defer C.engineFuncsFreeFile(ef.p, bufferPtr)

	return C.GoBytes(bufferPtr, length), nil
}

func (ef *EngineFuncs) FreeFile(buffer []byte) {
	C.engineFuncsFreeFile(ef.p, unsafe.Pointer(&buffer[0]))
}

func (ef *EngineFuncs) EndSection(pszSectionName string) {
	cs := C.CString(pszSectionName)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsEndSection(ef.p, cs)
}

// GetGameDir Gets the game directory.
func (ef *EngineFuncs) GetGameDir() string {
	var szGetGameDir [maxPath]byte
	C.engineFuncsGetGameDir(ef.p, (*C.char)(unsafe.Pointer(&szGetGameDir[0])))

	return C.GoString((*C.char)(unsafe.Pointer(&szGetGameDir[0])))
}

// CVarRegisterVariable Registers a cvar.
// Identical to CVarRegister, except it doesn't set the CVarExtdll flag.
func (ef *EngineFuncs) CVarRegisterVariable(variable *CVar) {
	C.engineFuncsCvar_RegisterVariable(ef.p, variable.p)
}

// FadeClientVolume Fades the volume of a client.
func (ef *EngineFuncs) FadeClientVolume(pEdict *Edict, fadePercent, fadeOutSeconds, holdTime, fadeInSeconds int) {
	C.engineFuncsFadeClientVolume(ef.p, pEdict.p, C.int(fadePercent), C.int(fadeOutSeconds), C.int(holdTime), C.int(fadeInSeconds))
}

// SetClientMaxspeed Sets the max speed of a client.
func (ef *EngineFuncs) SetClientMaxspeed(e *Edict, maxSpeed float32) {
	C.engineFuncsSetClientMaxspeed(ef.p, e.p, C.float(maxSpeed))
}

// CreateFakeClient Creates a fake client.
func (ef *EngineFuncs) CreateFakeClient(netname string) *Edict {
	cs := C.CString(netname)
	defer C.free(unsafe.Pointer(cs))

	return edictFromC(ef.globalVars.p, C.engineFuncsCreateFakeClient(ef.p, cs))
}

// RunPlayerMove Runs player movement for a fake client.
func (ef *EngineFuncs) RunPlayerMove(
	fakeClient *Edict,
	viewAngles vector.Vector,
	forwardMove, sideMove, upMove float32,
	buttons uint16,
	impulse uint16,
	msec uint16,
) {
	C.engineFuncsRunPlayerMove(
		ef.p,
		fakeClient.p,
		(*C.float)(&viewAngles[0]),
		C.float(forwardMove),
		C.float(sideMove),
		C.float(upMove),
		C.ushort(buttons),
		C.byte(impulse),
		C.byte(msec),
	)
}

// NumberOfEntities Gets the number of entities.
func (ef *EngineFuncs) NumberOfEntities() int {
	return int(C.engineFuncsNumberOfEntities(ef.p))
}

// GetInfoKeyBuffer Gets the info key buffer of an entity.
func (ef *EngineFuncs) GetInfoKeyBuffer(e *Edict) *InfoBuffer {
	c := C.engineFuncsGetInfoKeyBuffer(ef.p, e.p)

	return infoBufferFromC(c)
}

// InfoKeyValue Gets the value of a key in an info buffer.
func (ef *EngineFuncs) InfoKeyValue(infobuffer InfoBuffer, key string) string {
	csKey := C.CString(key)
	defer C.free(unsafe.Pointer(csKey))

	return C.GoString(C.engineFuncsInfoKeyValue(ef.p, infobuffer.p, csKey))
}

// SetKeyValue Sets the value of a key in an info buffer.
func (ef *EngineFuncs) SetKeyValue(infobuffer InfoBuffer, key, value string) {
	csKey := C.CString(key)
	defer C.free(unsafe.Pointer(csKey))

	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))

	C.engineFuncsSetKeyValue(ef.p, infobuffer.p, csKey, csValue)
}

// SetClientKeyValue Sets the value of a key in an info buffer for a client.
func (ef *EngineFuncs) SetClientKeyValue(clientIndex int, infobuffer InfoBuffer, key, value string) {
	csKey := C.CString(key)
	defer C.free(unsafe.Pointer(csKey))

	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))

	C.engineFuncsSetClientKeyValue(ef.p, C.int(clientIndex), infobuffer.p, csKey, csValue)
}

// IsMapValid Checks if a map is valid.
func (ef *EngineFuncs) IsMapValid(filename string) bool {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))

	return int(C.engineFuncsIsMapValid(ef.p, cs)) == 1
}

// StaticDecal Creates a static decal.
func (ef *EngineFuncs) StaticDecal(origin vector.Vector, decalIndex, entityIndex, modelIndex int) {
	C.engineFuncsStaticDecal(
		ef.p,
		(*C.float)(&origin[0]),
		C.int(decalIndex),
		C.int(entityIndex),
		C.int(modelIndex),
	)
}

// PrecacheGeneric Precaches a generic.
// If this is called after ServerActivate, triggers a host error.
func (ef *EngineFuncs) PrecacheGeneric(s string) int {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	return int(C.engineFuncsPrecacheGeneric(ef.p, cs))
}

// GetPlayerUserId Returns the server assigned userid for this player.
func (ef *EngineFuncs) GetPlayerUserId(e *Edict) int {
	return int(C.engineFuncsGetPlayerUserId(ef.p, e.p))
}

// BuildSoundMsg Builds a sound message.
// entity 	- Entity that is playing the sound.
// channel 	- Channel to play the sound on.
// sample 	- Sound to play.
// volume 	- Volume of the sound. Must be in the range [ 0, 1 ].
// attenuation - Attenuation.
// flags - Sound flags.
// pitch - Pitch. Must be in the range [ 0, 255 ].
// msgDest - Message type.
// msgID - Message ID.
// origin - Origin in the world to use for PAS and PVS messages.
// ed - Client to send the message to for message types that target one client.
func (ef *EngineFuncs) BuildSoundMsg(
	entity *Edict,
	channel int,
	sample string,
	volume, attenuation float32,
	flags, pitch, msgType, msgID int,
	origin vector.Vector,
	ed *Edict,
) {
	csSample := C.CString(sample)
	defer C.free(unsafe.Pointer(csSample))

	C.engineFuncsBuildSoundMsg(
		ef.p,
		entity.p,
		C.int(channel),
		csSample,
		C.float(volume),
		C.float(attenuation),
		C.int(flags),
		C.int(pitch),
		C.int(msgType),
		C.int(msgID),
		(*C.float)(&origin[0]),
		ed.p,
	)
}

// IsDedicatedServer Checks if the server is dedicated.
func (ef *EngineFuncs) IsDedicatedServer() bool {
	return int(C.engineFuncsIsDedicatedServer(ef.p)) == 1
}

// CVarGetPointer Gets the pointer to a cvar.
func (ef *EngineFuncs) CVarGetPointer(name string) *CVar {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	return cvarFromC(C.engineFuncsCVarGetPointer(ef.p, cs))
}

// GetPlayerWONID Gets the WON ID of a player.
func (ef *EngineFuncs) GetPlayerWONID(e *Edict) uint {
	return uint(C.engineFuncsGetPlayerWONId(ef.p, e.p))
}

// InfoRemoveKey Removes a key from an info buffer.
func (ef *EngineFuncs) InfoRemoveKey(infobuffer InfoBuffer, key string) {
	csKey := C.CString(key)
	defer C.free(unsafe.Pointer(csKey))

	C.engineFuncsInfo_RemoveKey(ef.p, infobuffer.p, csKey)
}

// GetPhysicsKeyValue Gets the value of a key in a physics keyvalue buffer.
func (ef *EngineFuncs) GetPhysicsKeyValue(client *Edict, key string) *InfoBuffer {
	csKey := C.CString(key)
	defer C.free(unsafe.Pointer(csKey))

	return infoBufferFromC(C.engineFuncsGetPhysicsKeyValue(ef.p, client.p, csKey))
}

// SetPhysicsKeyValue Sets the value of a key in a physics keyvalue buffer.
func (ef *EngineFuncs) SetPhysicsKeyValue(client *Edict, key, value string) {
	csKey := C.CString(key)
	defer C.free(unsafe.Pointer(csKey))

	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))

	C.engineFuncsSetPhysicsKeyValue(ef.p, client.p, csKey, csValue)
}

// GetPhysicsInfoString Gets the physics info string of an entity.
func (ef *EngineFuncs) GetPhysicsInfoString(client *Edict) string {
	return C.GoString(C.engineFuncsGetPhysicsInfoString(ef.p, client.p))
}

// PrecacheEvent Precaches an event.
// It returns the event index.
func (ef *EngineFuncs) PrecacheEvent(eventType int, eventName string) int {
	csEventName := C.CString(eventName)
	defer C.free(unsafe.Pointer(csEventName))

	return int(C.engineFuncsPrecacheEvent(ef.p, C.int(eventType), csEventName))
}

// PlaybackEvent Plays an event.
// flags - Event flags.
// invoker - Client that triggered the event.
// eventIndex - Event index. @see PrecacheEvent
// delay - Delay before the event should be run.
// origin - If not g_vecZero, this is the origin parameter sent to the clients.
// angles - If not g_vecZero, this is the angles parameter sent to the clients.
// fparam1 - Float parameter 1.
// fparam2 - Float parameter 2.
// iparam1 - Integer parameter 1.
// iparam2 - Integer parameter 2.
// bparam1 - Boolean parameter 1.
// bparam2 - Boolean parameter 2.
func (ef *EngineFuncs) PlaybackEvent(
	flags int,
	invoker *Edict,
	eventIndex uint16,
	delay float32,
	origin, angles vector.Vector,
	fparam1, fparam2 float32,
	iparam1, iparam2 int,
	bparam1, bparam2 bool,
) {
	bp1 := 0
	if bparam1 {
		bp1 = 1
	}

	bp2 := 0
	if bparam2 {
		bp2 = 1
	}

	C.engineFuncsPlaybackEvent(
		ef.p,
		C.int(flags),
		invoker.p,
		C.uint16(eventIndex),
		C.float(delay),
		(*C.float)(&origin[0]),
		(*C.float)(&angles[0]),
		C.float(fparam1),
		C.float(fparam2),
		C.int(iparam1),
		C.int(iparam2),
		C.int(bp1),
		C.int(bp2),
	)
}

// SetFatPVS Adds the given origin to the current PVS.
func (ef *EngineFuncs) SetFatPVS(origin vector.Vector) unsafe.Pointer {
	return unsafe.Pointer(C.engineFuncsSetFatPVS(
		ef.p,
		(*C.float)(&origin[0]),
	))
}

// SetFatPAS Adds the given origin to the current PAS.
func (ef *EngineFuncs) SetFatPAS(origin vector.Vector) unsafe.Pointer {
	return unsafe.Pointer(C.engineFuncsSetFatPAS(
		ef.p,
		(*C.float)(&origin[0]),
	))
}

// CvarDirectSet Directly sets a cvar value.
func (ef *EngineFuncs) CvarDirectSet(cvar *CVar, value string) {
	csValue := C.CString(value)
	defer C.free(unsafe.Pointer(csValue))

	C.engineFuncsCvar_DirectSet(ef.p, cvar.p, csValue)
}

// GetPlayerStats Gets ping and packet loss of a player.
func (ef *EngineFuncs) GetPlayerStats(client *Edict) (ping, packetLoss int) {
	var cp, cl C.int
	C.engineFuncsGetPlayerStats(ef.p, client.p, &cp, &cl)
	return int(cp), int(cl)
}

// GetPlayerAuthId Gets the auth ID of a player.
func (ef *EngineFuncs) GetPlayerAuthId(client *Edict) string {
	return C.GoString(C.engineFuncsGetPlayerAuthId(ef.p, client.p))
}

func (ef *EngineFuncs) QueryClientCvarValue(player *Edict, cvarName string) {
	cs := C.CString(cvarName)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsQueryClientCvarValue(ef.p, player.p, cs)
}

func (ef *EngineFuncs) QueryClientCvarValue2(player *Edict, cvarName string, requestID int) {
	cs := C.CString(cvarName)
	defer C.free(unsafe.Pointer(cs))

	C.engineFuncsQueryClientCvarValue2(ef.p, player.p, cs, C.int(requestID))
}
