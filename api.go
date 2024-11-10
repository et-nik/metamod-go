package main

/*
#include <eiface.h>

void GoGameDLLInit(void);

static void GameDLLInit( void ) {
	printf("\n\n\nGameDLLInit\n\n\n");
	GoGameDLLInit();
}

void SetDLLFunctions(DLL_FUNCTIONS *pFunctionTable) {
	printf("\n\n\nSetDLLFunctions\n\n\n");

	pFunctionTable->pfnGameInit = GameDLLInit;
}

*/
import "C"
