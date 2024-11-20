#pragma once

#include "dllapi.h"
#include "engine_api.h"

// Flags returned by a plugin's api function.
// NOTE: order is crucial, as greater/less comparisons are made.
typedef enum META_RES
{
	MRES_UNSET = 0,
	MRES_IGNORED,       // plugin didn't take any action
	MRES_HANDLED,       // plugin did something, but real function should still be called
	MRES_OVERRIDE,      // call real function, but use my return value
	MRES_SUPERCEDE,     // skip real function; use my return value
} META_RES;

// Variables provided to plugins.
typedef struct meta_globals_t
{
	META_RES mres;          // writable; plugin's return flag
	META_RES prev_mres;     // readable; return flag of the previous plugin called
	META_RES status;        // readable; "highest" return flag so far
	void *orig_ret;         // readable; return value from "real" function
	void *override_ret;     // readable; return value from overriding/superceding plugin

#ifdef METAMOD_CORE
	uint32* esp_save;
#endif
} meta_globals_t;

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

// Pair of function tables provided by game DLL.
typedef struct {
	DLL_FUNCTIONS *dllapi_table;
	NEW_DLL_FUNCTIONS *newapi_table;
} gamedll_funcs_t;