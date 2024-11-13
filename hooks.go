package main

/*
#include <eiface.h>

void goHookMessageBegin(int msg_dest, int msg_type, const float *pOrigin, edict_t *ed);
void goHookMessageEnd(void);

void SetHooks(enginefuncs_t *pengfuncsFromEngine) {
	memset(pengfuncsFromEngine, 0, sizeof(enginefuncs_t));

	pengfuncsFromEngine->pfnMessageBegin = goHookMessageBegin;
   	pengfuncsFromEngine->pfnMessageEnd = goHookMessageEnd;
}
*/
import "C"
