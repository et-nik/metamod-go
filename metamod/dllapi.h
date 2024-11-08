#pragma once

// Typedefs for these are provided in SDK engine/eiface.h, but I didn't
// like the names (APIFUNCTION, APIFUNCTION2, NEW_DLL_FUNCTIONS_FN).
typedef int (*GETENTITYAPI_FN)(DLL_FUNCTIONS* pFunctionTable, int interfaceVersion);
typedef int (*GETENTITYAPI2_FN)(DLL_FUNCTIONS* pFunctionTable, int* interfaceVersion);
typedef int (*GETNEWDLLFUNCTIONS_FN)(NEW_DLL_FUNCTIONS* pFunctionTable, int* interfaceVersion);