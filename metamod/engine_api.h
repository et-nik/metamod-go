#pragma once

typedef struct enginefuncs_s;

// Plugin's GetEngineFunctions, called by metamod.
typedef int (*GET_ENGINE_FUNCTIONS_FN)(struct enginefuncs_s* pengfuncsFromEngine, int* interfaceVersion);

// According to SDK engine/eiface.h:
// ONLY ADD NEW FUNCTIONS TO THE END OF THIS STRUCT. INTERFACE VERSION IS FROZEN AT 138
#define ENGINE_INTERFACE_VERSION 138
