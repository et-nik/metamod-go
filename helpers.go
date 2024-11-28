package metamod_go

/*
#include <eiface.h>

int indexOfEntity (edict_t *ent, edict_t *startEntity) {
  return (ent - startEntity);
};
*/
import "C"
import "github.com/et-nik/metamod-go/engine"

func IsNullEntity(e *Edict) bool {
	if e == nil {
		return true
	}

	return isNullEntity(e.p)
}

func isNullEntity(e *C.edict_t) bool {
	return e == nil || entityIndex(e) == 0 || e.free == 1
}

func EntityIndex(e *Edict) int {
	if e == nil {
		return 0
	}

	return entityIndex(e.p)
}

func entityIndex(e *C.edict_t) int {
	if e == nil {
		return 0
	}

	if globalPluginState.startEntity == nil {
		return 0
	}

	return int(C.indexOfEntity(e, globalPluginState.startEntity.p))
}

func PlayerIndex(e *Edict) int {
	if e == nil {
		return 0
	}

	return EntityIndex(e) - 1
}

func IsAlive(e *Edict) bool {
	if e == nil {
		return false
	}

	if IsNullEntity(e) {
		return false
	}

	entVars := e.EntVars()

	return entVars.DeadFlag() == engine.DeadFlagNo &&
		entVars.Health() > 0 &&
		entVars.MoveType() != engine.MoveTypeNoclip
}
