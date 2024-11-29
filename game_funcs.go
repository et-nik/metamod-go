package metamod_go

/*
#include <eiface.h>

int gameFuncsClientConnect(
	struct dll_functions_s *t,
	edict_t *pEntity,
	const char *pszName,
	const char *pszAddress,
	char  *szRejectReason
) {
	char ptr[128];
	int result;

	result = (*t->pfnClientConnect)(pEntity, pszName, pszAddress, ptr);

	strcpy(szRejectReason, ptr);

	return result;
}

int gameFuncsSpawn(struct dll_functions_s *t, edict_t *pEntity) {
	return (*t->pfnSpawn)(pEntity);
}

void gameFuncsClientDisconnect(struct dll_functions_s *t, edict_t *pEntity) {
	(*t->pfnClientDisconnect)(pEntity);
}

void gameFuncsClientKill(struct dll_functions_s *t, edict_t *pEntity) {
	(*t->pfnClientKill)(pEntity);
}

void gameFuncsClientPutInServer(struct dll_functions_s *t, edict_t *pEntity) {
	(*t->pfnClientPutInServer)(pEntity);
}

void gameFuncsClientCommand(struct dll_functions_s *t, edict_t *pEntity) {
	(*t->pfnClientCommand)(pEntity);
}

const char *gameFuncsGetGameDescription(struct dll_functions_s *t) {
	return (*t->pfnGetGameDescription)();
}

void gameFuncsPlayerCustomization(struct dll_functions_s *t, edict_t *pEntity, customization_t *pCustom) {
	(*t->pfnPlayerCustomization)(pEntity, pCustom);
}

void gameFuncsSetupVisibility(struct dll_functions_s *t, edict_t *pViewEntity, edict_t *pClient, unsigned char **pvs, unsigned char **pas) {
	(*t->pfnSetupVisibility)(pViewEntity, pClient, pvs, pas);
}

*/
import "C"
import "unsafe"

type GameDLLFuncs struct {
	p *C.DLL_FUNCTIONS
}

func newGameDLLFuncs(p *C.DLL_FUNCTIONS) *GameDLLFuncs {
	return &GameDLLFuncs{p}
}

func (g *GameDLLFuncs) Spawn(entity *Edict) int {
	return int(C.gameFuncsSpawn(g.p, entity.ptr()))
}

func (g *GameDLLFuncs) ClientConnect(
	entity *Edict,
	name string,
	address string,
) (int, string) {
	csName := C.CString(name)
	defer C.free(unsafe.Pointer(csName))

	csAddress := C.CString(address)
	defer C.free(unsafe.Pointer(csAddress))

	szRejectReason := make([]byte, 128)

	result := C.gameFuncsClientConnect(
		g.p,
		entity.ptr(),
		csName,
		csAddress,
		(*C.char)(unsafe.Pointer(&szRejectReason[0])),
	)

	r := int(result)

	if r == Success {
		return r, ""
	}

	return r, string(szRejectReason)
}

func (g *GameDLLFuncs) ClientDisconnect(entity *Edict) {
	C.gameFuncsClientDisconnect(g.p, entity.ptr())
}

func (g *GameDLLFuncs) ClientKill(entity *Edict) {
	C.gameFuncsClientKill(g.p, entity.ptr())
}

func (g *GameDLLFuncs) ClientPutInServer(entity *Edict) {
	C.gameFuncsClientPutInServer(g.p, entity.ptr())
}

func (g *GameDLLFuncs) ClientCommand(entity *Edict) {
	C.gameFuncsClientCommand(g.p, entity.ptr())
}

func (g *GameDLLFuncs) GetGameDescription() string {
	return C.GoString(C.gameFuncsGetGameDescription(g.p))
}
