#pragma once

#include "dllapi.h"
#include "engine_api.h"

typedef struct META_FUNCTIONS
{
	GETENTITYAPI_FN         pfnGetEntityAPI;
	GETENTITYAPI_FN         pfnGetEntityAPI_Post;
	GETENTITYAPI2_FN        pfnGetEntityAPI2;
	GETENTITYAPI2_FN        pfnGetEntityAPI2_Post;
	GETNEWDLLFUNCTIONS_FN   pfnGetNewDLLFunctions;
	GETNEWDLLFUNCTIONS_FN   pfnGetNewDLLFunctions_Post;
	GET_ENGINE_FUNCTIONS_FN pfnGetEngineFunctions;
	GET_ENGINE_FUNCTIONS_FN pfnGetEngineFunctions_Post;
} META_FUNCTIONS;