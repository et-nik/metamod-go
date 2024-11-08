/* Code generated by cmd/cgo; DO NOT EDIT. */

/* package github.com/et-nik/metamod-go */


#line 1 "cgo-builtin-export-prolog"

#include <stddef.h>

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#endif

#endif

/* Start of preamble from import "C" comments.  */


#line 3 "main.go"



#include <stdio.h>

#include <eiface.h>

#include "metamod/plinfo.h"
#include "metamod/meta_api.h"

int GetNewDLLFunctions(NEW_DLL_FUNCTIONS *pNewFunctionTable, int *interfaceVersion);

#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt32 GoInt;
typedef GoUint32 GoUint;
typedef size_t GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
#ifdef _MSC_VER
#include <complex.h>
typedef _Fcomplex GoComplex64;
typedef _Dcomplex GoComplex128;
#else
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;
#endif

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_32_bit_pointer_matching_GoInt[sizeof(void*)==32/8 ? 1:-1];

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef _GoString_ GoString;
#endif
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif

extern int Meta_Attach(int now, META_FUNCTIONS* pFunctionTable, void* pMGlobals, void* pGamedllFuncs);
extern int Meta_Query(char* interfaceVersion, plugin_info_t** plinfo, void* pMetaUtilFuncs);
extern int Meta_Detach(int now, int reason);
extern void Meta_Init();
extern void GiveFnptrsToDll(enginefuncs_t* pengfuncsFromEngine, void* pGlobals);
extern int GetNewDLLFunctions(NEW_DLL_FUNCTIONS* pNewFunctionTable, int* interfaceVersion);
extern int GetNewDLLFunctions_Post(NEW_DLL_FUNCTIONS* pNewFunctionTable, int* interfaceVersion);
extern int GetEntityAPI2(DLL_FUNCTIONS* pFunctionTable, int* interfaceVersion);
extern int GetEntityAPI2_Post(DLL_FUNCTIONS* pFunctionTable, int* interfaceVersion);
extern int GetEngineFunctions(enginefuncs_t* pengfuncsFromEngine, int* interfaceVersion);
extern int GetEngineFunctions_Post(enginefuncs_t* pengfuncsFromEngine, int* interfaceVersion);

#ifdef __cplusplus
}
#endif
