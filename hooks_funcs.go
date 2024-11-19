package metamod_go

/*
#include <eiface.h>

extern vec3_t* castPtrToVec3(float *ptr);
extern void setVec3FloatPtr(float x, float y, float z, float* ptr);

*/
import "C"
import (
	"unsafe"
)

//export goHookPrecacheModel
func goHookPrecacheModel(s *C.char) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.PrecacheModel != nil {
		metaResult, result := globalPluginState.engineHooks.PrecacheModel(C.GoString(s))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookPrecacheSound
func goHookPrecacheSound(s *C.char) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.PrecacheSound != nil {
		metaResult, result := globalPluginState.engineHooks.PrecacheSound(C.GoString(s))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookSetModel
func goHookSetModel(pEdict *C.edict_t, s *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.SetModel != nil {
		r := globalPluginState.engineHooks.SetModel(edictFromC(globalPluginState.globalVars.p, pEdict), C.GoString(s))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookModelIndex
func goHookModelIndex(s *C.char) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.ModelIndex != nil {
		metaResult, result := globalPluginState.engineHooks.ModelIndex(C.GoString(s))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookModelFrames
func goHookModelFrames(index C.int) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.ModelFrames != nil {
		metaResult, result := globalPluginState.engineHooks.ModelFrames(int(index))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookSetSize
func goHookSetSize(pEdict *C.edict_t, mins, maxs *C.float) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.SetSize != nil {
		minsVec := C.castPtrToVec3(mins)
		maxsVec := C.castPtrToVec3(maxs)

		metaResult := globalPluginState.engineHooks.SetSize(
			edictFromC(globalPluginState.globalVars.p, pEdict),
			[3]float32{float32(minsVec[0]), float32(minsVec[1]), float32(minsVec[2])},
			[3]float32{float32(maxsVec[0]), float32(maxsVec[1]), float32(maxsVec[2])},
		)

		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookChangeLevel
func goHookChangeLevel(levelname *C.char, landmark *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.ChangeLevel != nil {
		r := globalPluginState.engineHooks.ChangeLevel(C.GoString(levelname), C.GoString(landmark))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookVecToYaw
func goHookVecToYaw(vec *C.float) C.float {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.VecToYaw != nil {
		v := C.castPtrToVec3(vec)
		metaResult, result := globalPluginState.engineHooks.VecToYaw(
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.float(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookVecToAngles
func goHookVecToAngles(vec *C.float) *C.float {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.VecToAngles != nil {
		v := C.castPtrToVec3(vec)

		metaResult, result := globalPluginState.engineHooks.VecToAngles(
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		r := [3]C.float{C.float(result[0]), C.float(result[1]), C.float(result[2])}
		return &r[0]
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookMoveToOrigin
func goHookMoveToOrigin(pEdict *C.edict_t, goal *C.float, dist C.float, moveType C.int) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.MoveToOrigin != nil {
		v := C.castPtrToVec3(goal)

		metaResult := globalPluginState.engineHooks.MoveToOrigin(
			edictFromC(globalPluginState.globalVars.p, pEdict),
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
			float32(dist),
			MoveType(int(moveType)),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookChangeYaw
func goHookChangeYaw(pEdict *C.edict_t) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.ChangeYaw != nil {
		metaResult := globalPluginState.engineHooks.ChangeYaw(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookChangePitch
func goHookChangePitch(pEdict *C.edict_t) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.ChangePitch != nil {
		metaResult := globalPluginState.engineHooks.ChangePitch(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookFindEntityByString
func goHookFindEntityByString(pEdict *C.edict_t, field *C.char, s *C.char) *C.edict_t {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.FindEntityByString != nil {
		metaResult, result := globalPluginState.engineHooks.FindEntityByString(
			edictFromC(globalPluginState.globalVars.p, pEdict),
			FindEntityField(C.GoString(field)),
			C.GoString(s),
		)

		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result.p != nil {
			return result.p
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookGetEntityIllum
func goHookGetEntityIllum(pEdict *C.edict_t) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.GetEntityIllum != nil {
		metaResult, result := globalPluginState.engineHooks.GetEntityIllum(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookFindEntityInSphere
func goHookFindEntityInSphere(pEdict *C.edict_t, origin *C.float, radius C.float) *C.edict_t {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.FindEntityInSphere != nil {
		v := C.castPtrToVec3(origin)

		metaResult, result := globalPluginState.engineHooks.FindEntityInSphere(
			edictFromC(globalPluginState.globalVars.p, pEdict),
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
			float32(radius),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result.p != nil {
			return result.p
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookFindClientInPVS
func goHookFindClientInPVS(pEdict *C.edict_t) *C.edict_t {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.FindClientInPVS != nil {
		metaResult, result := globalPluginState.engineHooks.FindClientInPVS(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result.p != nil {
			return result.p
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookEntitiesInPVS
func goHookEntitiesInPVS(pEdict *C.edict_t) *C.edict_t {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.EntitiesInPVS != nil {
		metaResult, result := globalPluginState.engineHooks.EntitiesInPVS(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result.p != nil {
			return result.p
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookMakeVectors
func goHookMakeVectors(angles *C.float) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.MakeVectors != nil {
		v := C.castPtrToVec3(angles)

		metaResult := globalPluginState.engineHooks.MakeVectors([3]float32{float32(v[0]), float32(v[1]), float32(v[2])})
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookAngleVectors
func goHookAngleVectors(vector *C.float, forward, right, up *C.float) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.AngleVectors != nil {
		v := C.castPtrToVec3(vector)

		forwardVec := C.castPtrToVec3(forward)
		rightVec := C.castPtrToVec3(right)
		upVec := C.castPtrToVec3(up)

		metaResult := globalPluginState.engineHooks.AngleVectors(
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
			[3]float32{float32(forwardVec[0]), float32(forwardVec[1]), float32(forwardVec[2])},
			[3]float32{float32(rightVec[0]), float32(rightVec[1]), float32(rightVec[2])},
			[3]float32{float32(upVec[0]), float32(upVec[1]), float32(upVec[2])},
		)

		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookCreateEntity
func goHookCreateEntity() *C.edict_t {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.CreateEntity != nil {
		metaResult, result := globalPluginState.engineHooks.CreateEntity()
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result.p != nil {
			return result.p
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookCreateNamedEntity
func goHookCreateNamedEntity(s *C.char) *C.edict_t {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.CreateNamedEntity != nil {
		result := globalPluginState.engineHooks.CreateNamedEntity(C.GoString(s))

		if result.p != nil {
			return result.p
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookRemoveEntity
func goHookRemoveEntity(pEdict *C.edict_t) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.RemoveEntity != nil {
		r := globalPluginState.engineHooks.RemoveEntity(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMakeStatic
func goHookMakeStatic(pEdict *C.edict_t) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.MakeStatic != nil {
		r := globalPluginState.engineHooks.MakeStatic(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookEntIsOnFloor
func goHookEntIsOnFloor(pEdict *C.edict_t) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.EntIsOnFloor != nil {
		metaResult, result := globalPluginState.engineHooks.EntIsOnFloor(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result {
			return 1
		}

		return 0
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookDropToFloor
func goHookDropToFloor(pEdict *C.edict_t) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.DropToFloor != nil {
		metaResult, result := globalPluginState.engineHooks.DropToFloor(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookWalkMove
func goHookWalkMove(pEdict *C.edict_t, yaw C.float, dist C.float, mode C.int) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.WalkMove != nil {
		metaResult, result := globalPluginState.engineHooks.WalkMove(edictFromC(globalPluginState.globalVars.p, pEdict), float32(yaw), float32(dist), WalkMoveMode(mode))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookSetOrigin
func goHookSetOrigin(pEdict *C.edict_t, origin *C.float) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.SetOrigin != nil {
		v := C.castPtrToVec3(origin)

		metaResult := globalPluginState.engineHooks.SetOrigin(
			edictFromC(globalPluginState.globalVars.p, pEdict),
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookEmitSound
func goHookEmitSound(pEdict *C.edict_t, channel C.int, sample *C.char, volume C.float, attenuation C.int, fFlags C.int, pitch C.int) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.EmitSound != nil {
		metaResult := globalPluginState.engineHooks.EmitSound(
			edictFromC(globalPluginState.globalVars.p, pEdict),
			int(channel),
			C.GoString(sample),
			float32(volume),
			int(attenuation),
			int(fFlags),
			int(pitch),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookEmitAmbientSound
func goHookEmitAmbientSound(
	pEdict *C.edict_t,
	position *C.float,
	sample *C.char,
	volume C.float,
	attenuation C.float,
	fFlags C.int,
	pitch C.int,
) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.EmitAmbientSound != nil {
		v := C.castPtrToVec3(position)

		metaResult := globalPluginState.engineHooks.EmitAmbientSound(
			edictFromC(globalPluginState.globalVars.p, pEdict),
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
			C.GoString(sample),
			float32(volume),
			float32(attenuation),
			int(fFlags),
			int(pitch),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return
}

//export goHookTraceLine
func goHookTraceLine(v1, v2 *C.float, fNoMonsters C.int, pentToSkip *C.edict_t, ptr *C.TraceResult) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.TraceLine != nil {
		v1Vec := C.castPtrToVec3(v1)
		v2Vec := C.castPtrToVec3(v2)

		metaResult, result := globalPluginState.engineHooks.TraceLine(
			[3]float32{float32(v1Vec[0]), float32(v1Vec[1]), float32(v1Vec[2])},
			[3]float32{float32(v2Vec[0]), float32(v2Vec[1]), float32(v2Vec[2])},
			int(fNoMonsters),
			edictFromC(globalPluginState.globalVars.p, pentToSkip),
		)

		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			*ptr = *result.ToC()
		}

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookTraceToss
func goHookTraceToss(pent, pentToIgnore *C.edict_t, ptr *C.TraceResult) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.TraceToss != nil {
		metaResult, result := globalPluginState.engineHooks.TraceToss(edictFromC(globalPluginState.globalVars.p, pent), edictFromC(globalPluginState.globalVars.p, pentToIgnore))

		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			*ptr = *result.ToC()
		}

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookTraceMonsterHull
func goHookTraceMonsterHull(pent *C.edict_t, v1, v2 *C.float, fNoMonsters C.int, pentToSkip *C.edict_t, ptr *C.TraceResult) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.TraceMonsterHull != nil {
		v1Vec := C.castPtrToVec3(v1)
		v2Vec := C.castPtrToVec3(v2)

		metaResult, result, hit := globalPluginState.engineHooks.TraceMonsterHull(
			edictFromC(globalPluginState.globalVars.p, pent),
			[3]float32{float32(v1Vec[0]), float32(v1Vec[1]), float32(v1Vec[2])},
			[3]float32{float32(v2Vec[0]), float32(v2Vec[1]), float32(v2Vec[2])},
			int(fNoMonsters),
			edictFromC(globalPluginState.globalVars.p, pentToSkip),
		)

		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			*ptr = *result.ToC()
		}

		return C.int(hit)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookTraceHull
func goHookTraceHull(v1, v2 *C.float, fNoMonsters, hullNumber C.int, pentToSkip *C.edict_t, ptr *C.TraceResult) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.TraceHull != nil {
		v1Vec := C.castPtrToVec3(v1)
		v2Vec := C.castPtrToVec3(v2)

		metaResult, result := globalPluginState.engineHooks.TraceHull(
			[3]float32{float32(v1Vec[0]), float32(v1Vec[1]), float32(v1Vec[2])},
			[3]float32{float32(v2Vec[0]), float32(v2Vec[1]), float32(v2Vec[2])},
			int(fNoMonsters),
			int(hullNumber),
			edictFromC(globalPluginState.globalVars.p, pentToSkip),
		)

		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			*ptr = *result.ToC()
		}

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookTraceModel
func goHookTraceModel(v1, v2 *C.float, hullNumber C.int, pent *C.edict_t, ptr *C.TraceResult) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.TraceModel != nil {
		v1Vec := C.castPtrToVec3(v1)
		v2Vec := C.castPtrToVec3(v2)

		metaResult, result := globalPluginState.engineHooks.TraceModel(
			[3]float32{float32(v1Vec[0]), float32(v1Vec[1]), float32(v1Vec[2])},
			[3]float32{float32(v2Vec[0]), float32(v2Vec[1]), float32(v2Vec[2])},
			int(hullNumber),
			edictFromC(globalPluginState.globalVars.p, pent),
		)

		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			*ptr = *result.ToC()
		}

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookTraceTexture
func goHookTraceTexture(pTextureEntity *C.edict_t, v1, v2 *C.float) *C.char {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.TraceTexture != nil {
		v1Vec := C.castPtrToVec3(v1)
		v2Vec := C.castPtrToVec3(v2)

		metaResult, result := globalPluginState.engineHooks.TraceTexture(
			edictFromC(globalPluginState.globalVars.p, pTextureEntity),
			[3]float32{float32(v1Vec[0]), float32(v1Vec[1]), float32(v1Vec[2])},
			[3]float32{float32(v2Vec[0]), float32(v2Vec[1]), float32(v2Vec[2])},
		)

		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			return (*C.char)(unsafe.Pointer(result.ToC()))
		}
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookGetAimVector
func goHookGetAimVector(pent *C.edict_t, speed C.float, ptr *C.float) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.GetAimVector != nil {
		metaResult, result := globalPluginState.engineHooks.GetAimVector(edictFromC(globalPluginState.globalVars.p, pent), float32(speed))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		C.setVec3FloatPtr(C.float(result[0]), C.float(result[1]), C.float(result[2]), ptr)

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookServerCommand
func goHookServerCommand(s *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.ServerCommand != nil {
		r := globalPluginState.engineHooks.ServerCommand(C.GoString(s))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookServerExecute
func goHookServerExecute() {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.ServerExecute != nil {
		r := globalPluginState.engineHooks.ServerExecute()
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookClientCommand
func goHookClientCommand(pEdict *C.edict_t, format *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.ClientCommand != nil {
		r := globalPluginState.engineHooks.ClientCommand(edictFromC(globalPluginState.globalVars.p, pEdict), C.GoString(format))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookAddServerCommand
func goHookAddServerCommand(s *C.char, f unsafe.Pointer) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.AddServerCommand != nil {
		r := globalPluginState.engineHooks.AddServerCommand(C.GoString(s), f)
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookParticleEffect
func goHookParticleEffect(origin, direction *C.float, color, count C.float) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.ParticleEffect != nil {
		originVec := C.castPtrToVec3(origin)
		directionVec := C.castPtrToVec3(direction)

		r := globalPluginState.engineHooks.ParticleEffect(
			[3]float32{float32(originVec[0]), float32(originVec[1]), float32(originVec[2])},
			[3]float32{float32(directionVec[0]), float32(directionVec[1]), float32(directionVec[2])},
			float32(color),
			float32(count),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookLightStyle
func goHookLightStyle(style C.int, value *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.LightStyle != nil {
		r := globalPluginState.engineHooks.LightStyle(int(style), C.GoString(value))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookDecalIndex
func goHookDecalIndex(s *C.char) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.DecalIndex != nil {
		metaResult, result := globalPluginState.engineHooks.DecalIndex(C.GoString(s))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookPointContents
func goHookPointContents(v *C.float) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.PointContents != nil {
		vVec := C.castPtrToVec3(v)

		metaResult, result := globalPluginState.engineHooks.PointContents([3]float32{float32(vVec[0]), float32(vVec[1]), float32(vVec[2])})
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookMessageBegin
func goHookMessageBegin(msgDest C.int, msgType C.int, pOrigin *C.float, pEdict *C.edict_t) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.MessageBegin != nil {
		var origin *float32
		if pOrigin != nil {
			*origin = float32(*pOrigin)
		}

		r := globalPluginState.engineHooks.MessageBegin(int(msgDest), int(msgType), origin, edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageEnd
func goHookMessageEnd() {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.MessageEnd != nil {
		r := globalPluginState.engineHooks.MessageEnd()
		globalPluginState.metaGlobals.SetMres(MetaResult(r))
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteByte
func goHookMessageWriteByte(b C.int) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.MessageWriteByte != nil {
		r := globalPluginState.engineHooks.MessageWriteByte(int(b))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteChar
func goHookMessageWriteChar(c C.int) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.MessageWriteChar != nil {
		r := globalPluginState.engineHooks.MessageWriteChar(int(c))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteShort
func goHookMessageWriteShort(s C.int) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.MessageWriteShort != nil {
		r := globalPluginState.engineHooks.MessageWriteShort(int(s))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteLong
func goHookMessageWriteLong(l C.int) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.MessageWriteLong != nil {
		r := globalPluginState.engineHooks.MessageWriteLong(int(l))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteAngle
func goHookMessageWriteAngle(f C.float) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.MessageWriteAngle != nil {
		r := globalPluginState.engineHooks.MessageWriteAngle(float32(f))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteCoord
func goHookMessageWriteCoord(f C.float) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.MessageWriteCoord != nil {
		r := globalPluginState.engineHooks.MessageWriteCoord(float32(f))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteString
func goHookMessageWriteString(s *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.MessageWriteString != nil {
		r := globalPluginState.engineHooks.MessageWriteString(C.GoString(s))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteEntity
func goHookMessageWriteEntity(id C.int) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.MessageWriteEntity != nil {
		r := globalPluginState.engineHooks.MessageWriteEntity(int(id))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookCVarRegister
func goHookCVarRegister(cvar *C.cvar_t) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.CVarRegister != nil {
		metaResult := globalPluginState.engineHooks.CVarRegister(cvarFromC(cvar))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookCVarGetString
func goHookCVarGetString(s *C.char) *C.char {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.CVarGetString != nil {
		metaResult, result := globalPluginState.engineHooks.CVarGetString(C.GoString(s))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookCVarGetFloat
func goHookCVarGetFloat(s *C.char) C.float {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.CVarGetFloat != nil {
		metaResult, result := globalPluginState.engineHooks.CVarGetFloat(C.GoString(s))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.float(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookCVarSetFloat
func goHookCVarSetFloat(s *C.char, value C.float) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.CVarSetFloat != nil {
		r := globalPluginState.engineHooks.CVarSetFloat(C.GoString(s), float32(value))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookCVarSetString
func goHookCVarSetString(s, value *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.CVarSetString != nil {
		r := globalPluginState.engineHooks.CVarSetString(C.GoString(s), C.GoString(value))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookAlertMessage
func goHookAlertMessage(level C.ALERT_TYPE, format *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.AlertMessage != nil {
		r := globalPluginState.engineHooks.AlertMessage(AlertType(level), C.GoString(format))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookPvAllocEntPrivateData
func goHookPvAllocEntPrivateData(pEdict *C.edict_t, cb C.int32) unsafe.Pointer {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.PvAllocEntPrivateData != nil {
		metaResult, result := globalPluginState.engineHooks.PvAllocEntPrivateData(edictFromC(globalPluginState.globalVars.p, pEdict), int32(cb))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return result
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookPvEntPrivateData
func goHookPvEntPrivateData(pEdict *C.edict_t) unsafe.Pointer {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.PvEntPrivateData != nil {
		metaResult, result := globalPluginState.engineHooks.PvEntPrivateData(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return result
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookFreeEntPrivateData
func goHookFreeEntPrivateData(pEdict *C.edict_t) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.FreeEntPrivateData != nil {
		r := globalPluginState.engineHooks.FreeEntPrivateData(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookGetVarsOfEnt
func goHookGetVarsOfEnt(pEdict *C.edict_t) *C.entvars_t {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.GetVarsOfEnt != nil {
		metaResult, result := globalPluginState.engineHooks.GetVarsOfEnt(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			return result.p
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookIndexOfEdict
func goHookIndexOfEdict(pEdict *C.edict_t) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.IndexOfEdict != nil {
		metaResult, result := globalPluginState.engineHooks.IndexOfEdict(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookPEntityOfEntIndex
func goHookPEntityOfEntIndex(index C.int) *C.edict_t {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.PEntityOfEntIndex != nil {
		metaResult, result := globalPluginState.engineHooks.PEntityOfEntIndex(int(index))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			return result.p
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookFindEntityByVars
func goHookFindEntityByVars(vars *C.entvars_t) *C.edict_t {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.FindEntityByVars != nil {
		metaResult, result := globalPluginState.engineHooks.FindEntityByVars(entVarsFromC(globalPluginState.globalVars.p, vars))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			return result.p
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookGetModelPtr
func goHookGetModelPtr(pEdict *C.edict_t) *C.void {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.GetModelPtr != nil {
		metaResult, result := globalPluginState.engineHooks.GetModelPtr(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return (*C.void)(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookRegUserMsg
func goHookRegUserMsg(name *C.char, size C.int) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.RegUserMsg != nil {
		metaResult, result := globalPluginState.engineHooks.RegUserMsg(C.GoString(name), int(size))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookFunctionFromName
func goHookFunctionFromName(name *C.char) C.uint32 {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.FunctionFromName != nil {
		metaResult, result := globalPluginState.engineHooks.FunctionFromName(C.GoString(name))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.uint32(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookNameForFunction
func goHookNameForFunction(function C.uint32) *C.char {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.NameForFunction != nil {
		metaResult, result := globalPluginState.engineHooks.NameForFunction(uint32(function))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookClientPrint
func goHookClientPrint(pEdict *C.edict_t, level C.PRINT_TYPE, msg *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.ClientPrint != nil {
		metaResult := globalPluginState.engineHooks.ClientPrint(edictFromC(globalPluginState.globalVars.p, pEdict), PrintType(level), C.GoString(msg))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookServerPrint
func goHookServerPrint(msg *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.ServerPrint != nil {
		metaResult := globalPluginState.engineHooks.ServerPrint(C.GoString(msg))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookGetAttachment
func goHookGetAttachment(pEdict *C.edict_t, attachment C.int, origin, angles *C.float) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.GetAttachment != nil {

		gorigin := [3]float32{}
		gangles := [3]float32{}

		metaResult := globalPluginState.engineHooks.GetAttachment(
			edictFromC(globalPluginState.globalVars.p, pEdict),
			int(attachment),
			&gorigin,
			&gangles,
		)

		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		C.setVec3FloatPtr(C.float(gorigin[0]), C.float(gorigin[1]), C.float(gorigin[2]), origin)
		C.setVec3FloatPtr(C.float(gangles[0]), C.float(gangles[1]), C.float(gangles[2]), angles)

		return 0
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookRandomLong
func goHookRandomLong(low, high C.int) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.RandomLong != nil {
		metaResult, result := globalPluginState.engineHooks.RandomLong(int32(low), int32(high))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookRandomFloat
func goHookRandomFloat(low, high C.float) C.float {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.RandomFloat != nil {
		metaResult, result := globalPluginState.engineHooks.RandomFloat(float32(low), float32(high))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.float(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookSetView
func goHookSetView(pClient, pEdict *C.edict_t) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.SetView != nil {
		r := globalPluginState.engineHooks.SetView(edictFromC(globalPluginState.globalVars.p, pClient), edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookTime
func goHookTime() C.float {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.Time != nil {
		metaResult, result := globalPluginState.engineHooks.Time()
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.float(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookCrosshairAngle
func goHookCrosshairAngle(pClient *C.edict_t, pitch, yaw C.float) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.CrosshairAngle != nil {
		r := globalPluginState.engineHooks.CrosshairAngle(edictFromC(globalPluginState.globalVars.p, pClient), float32(pitch), float32(yaw))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookLoadFileForMe
func goHookLoadFileForMe(name *C.char, _ *C.int) *C.byte {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.LoadFileForMe != nil {
		metaResult, result, err := globalPluginState.engineHooks.LoadFileForMe(C.GoString(name))
		if err != nil {
			return nil
		}

		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			return (*C.byte)(unsafe.Pointer(&result[0]))
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookFreeFile
//func goHookFreeFile(buffer *C.byte) {
//	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.FreeFile != nil {
//		r := globalPluginState.engineHooks.FreeFile(buffer)
//		globalPluginState.metaGlobals.SetMres(MetaResult(r))
//
//		return
//	}
//
//	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
//}

//export goHookGetGameDir
func goHookGetGameDir() *C.char {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.GetGameDir != nil {
		metaResult, result := globalPluginState.engineHooks.GetGameDir()
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookCvar_RegisterVariable
func goHookCvar_RegisterVariable(cvar *C.cvar_t) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.CVarRegisterVariable != nil {
		metaResult := globalPluginState.engineHooks.CVarRegisterVariable(cvarFromC(cvar))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookFadeClientVolume
func goHookFadeClientVolume(pEdict *C.edict_t, fadePercent, fadeOutSeconds, holdTime, fadeInSeconds C.int) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.FadeClientVolume != nil {
		r := globalPluginState.engineHooks.FadeClientVolume(
			edictFromC(globalPluginState.globalVars.p, pEdict),
			int(fadePercent),
			int(fadeOutSeconds),
			int(holdTime),
			int(fadeInSeconds),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookSetClientMaxspeed
func goHookSetClientMaxspeed(pEdict *C.edict_t, speed C.float) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.SetClientMaxspeed != nil {
		r := globalPluginState.engineHooks.SetClientMaxspeed(edictFromC(globalPluginState.globalVars.p, pEdict), float32(speed))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookCreateFakeClient
func goHookCreateFakeClient(name *C.char) *C.edict_t {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.CreateFakeClient != nil {
		metaResult, result := globalPluginState.engineHooks.CreateFakeClient(C.GoString(name))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result.p != nil {
			return result.p
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookRunPlayerMove
func goHookRunPlayerMove(pEdict *C.edict_t, viewangles, forwardmove, sidemove, upmove *C.float, buttons C.ushort, impulse, msec C.byte) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.RunPlayerMove != nil {
		v := C.castPtrToVec3(viewangles)

		r := globalPluginState.engineHooks.RunPlayerMove(
			edictFromC(globalPluginState.globalVars.p, pEdict),
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
			float32(*forwardmove),
			float32(*sidemove),
			float32(*upmove),
			uint16(buttons),
			uint16(impulse),
			uint16(msec),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookNumberOfEntities
func goHookNumberOfEntities() C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.NumberOfEntities != nil {
		metaResult, result := globalPluginState.engineHooks.NumberOfEntities()
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookGetInfoKeyBuffer
func goHookGetInfoKeyBuffer(pEdict *C.edict_t) *C.char {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.GetInfoKeyBuffer != nil {
		metaResult, result := globalPluginState.engineHooks.GetInfoKeyBuffer(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			return (*C.char)(unsafe.Pointer(&result[0]))
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookInfoKeyValue
func goHookInfoKeyValue(info, key *C.char) *C.char {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.InfoKeyValue != nil {
		metaResult, result := globalPluginState.engineHooks.InfoKeyValue([]byte(C.GoString(info)), C.GoString(key))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookSetKeyValue
func goHookSetKeyValue(info, key, value *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.SetKeyValue != nil {
		r := globalPluginState.engineHooks.SetKeyValue([]byte(C.GoString(info)), C.GoString(key), C.GoString(value))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookSetClientKeyValue
func goHookSetClientKeyValue(clientIndex C.int, key, value *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.SetClientKeyValue != nil {
		r := globalPluginState.engineHooks.SetClientKeyValue(int(clientIndex), C.GoString(key), C.GoString(value))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookIsMapValid
func goHookIsMapValid(mapName *C.char) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.IsMapValid != nil {
		metaResult, result := globalPluginState.engineHooks.IsMapValid(C.GoString(mapName))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result {
			return 1
		}

		return 0
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookStaticDecal
func goHookStaticDecal(origin *C.float, decalIndex, entityIndex, modelIndex C.int) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.StaticDecal != nil {
		v := C.castPtrToVec3(origin)

		r := globalPluginState.engineHooks.StaticDecal(
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
			int(decalIndex),
			int(entityIndex),
			int(modelIndex),
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookPrecacheGeneric
func goHookPrecacheGeneric(s *C.char) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.PrecacheGeneric != nil {
		metaResult, result := globalPluginState.engineHooks.PrecacheGeneric(C.GoString(s))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookGetPlayerUserId
func goHookGetPlayerUserId(pEdict *C.edict_t) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.GetPlayerUserId != nil {
		metaResult, result := globalPluginState.engineHooks.GetPlayerUserId(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookIsDedicatedServer
func goHookIsDedicatedServer() C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.IsDedicatedServer != nil {
		metaResult, result := globalPluginState.engineHooks.IsDedicatedServer()
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result {
			return 1
		}

		return 0
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookCVarGetPointer
func goHookCVarGetPointer(name *C.char) *C.cvar_t {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.CVarGetPointer != nil {
		metaResult, result := globalPluginState.engineHooks.CVarGetPointer(C.GoString(name))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			return result.p
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookGetPlayerWONId
func goHookGetPlayerWONId(pEdict *C.edict_t) C.uint {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.GetPlayerWONID != nil {
		metaResult, result := globalPluginState.engineHooks.GetPlayerWONID(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.uint(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookInfo_RemoveKey
func goHookInfo_RemoveKey(info, key *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.InfoRemoveKey != nil {
		r := globalPluginState.engineHooks.InfoRemoveKey(C.GoString(info), C.GoString(key))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookGetPhysicsKeyValue
func goHookGetPhysicsKeyValue(pEdict *C.edict_t, key *C.char) *C.char {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.GetPhysicsKeyValue != nil {
		metaResult, result := globalPluginState.engineHooks.GetPhysicsKeyValue(edictFromC(globalPluginState.globalVars.p, pEdict), C.GoString(key))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookSetPhysicsKeyValue
func goHookSetPhysicsKeyValue(pEdict *C.edict_t, key, value *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.SetPhysicsKeyValue != nil {
		r := globalPluginState.engineHooks.SetPhysicsKeyValue(edictFromC(globalPluginState.globalVars.p, pEdict), C.GoString(key), C.GoString(value))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookGetPhysicsInfoString
func goHookGetPhysicsInfoString(pEdict *C.edict_t) *C.char {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.GetPhysicsInfoString != nil {
		metaResult, result := globalPluginState.engineHooks.GetPhysicsInfoString(edictFromC(globalPluginState.globalVars.p, pEdict))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookPrecacheEvent
func goHookPrecacheEvent(eventType C.int, eventName *C.char) C.int {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.PrecacheEvent != nil {
		metaResult, result := globalPluginState.engineHooks.PrecacheEvent(int(eventType), C.GoString(eventName))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return C.int(result)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookPlaybackEvent
func goHookPlaybackEvent(
	flags C.int,
	pInvoker *C.edict_t,
	eventIndex C.ushort,
	delay C.float,
	origin, angles *C.float,
	fparam1, fparam2 C.float,
	iparam1, iparam2 C.int,
	bparam1, bparam2 C.int,
) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.PlaybackEvent != nil {
		originVec := C.castPtrToVec3(origin)
		anglesVec := C.castPtrToVec3(angles)

		boolParam1 := int(bparam1) == 1
		boolParam2 := int(bparam2) == 1

		metaResult := globalPluginState.engineHooks.PlaybackEvent(
			int(flags),
			edictFromC(globalPluginState.globalVars.p, pInvoker),
			uint16(eventIndex),
			float32(delay),
			[3]float32{float32(originVec[0]), float32(originVec[1]), float32(originVec[2])},
			[3]float32{float32(anglesVec[0]), float32(anglesVec[1]), float32(anglesVec[2])},
			float32(fparam1),
			float32(fparam2),
			int(iparam1),
			int(iparam2),
			boolParam1,
			boolParam2,
		)
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookSetFatPVS
func goHookSetFatPVS(origin *C.float) *C.byte {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.SetFatPVS != nil {
		v := C.castPtrToVec3(origin)

		metaResult, result := globalPluginState.engineHooks.SetFatPVS([3]float32{float32(v[0]), float32(v[1]), float32(v[2])})
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			return (*C.byte)(result)
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookSetFatPAS
func goHookSetFatPAS(origin *C.float) *C.byte {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.SetFatPAS != nil {
		v := C.castPtrToVec3(origin)

		metaResult, result := globalPluginState.engineHooks.SetFatPAS([3]float32{float32(v[0]), float32(v[1]), float32(v[2])})
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != nil {
			return (*C.byte)(result)
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookCvar_DirectSet
func goHookCvar_DirectSet(cvar *C.cvar_t, value *C.char) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.CvarDirectSet != nil {
		r := globalPluginState.engineHooks.CvarDirectSet(cvarFromC(cvar), C.GoString(value))
		globalPluginState.metaGlobals.SetMres(MetaResult(r))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

//export goHookGetPlayerStats
func goHookGetPlayerStats(pClient *C.edict_t, ping, packetLoss *C.int) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.GetPlayerStats != nil {
		metaResult, pingResult, lossResult := globalPluginState.engineHooks.GetPlayerStats(edictFromC(globalPluginState.globalVars.p, pClient))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		*ping = C.int(pingResult)
		*packetLoss = C.int(lossResult)
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}

// const char *(*pfnGetPlayerAuthId)(edict_t *e);
//
//export goHookGetPlayerAuthId
func goHookGetPlayerAuthId(pClient *C.edict_t) *C.char {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.GetPlayerAuthId != nil {
		metaResult, result := globalPluginState.engineHooks.GetPlayerAuthId(edictFromC(globalPluginState.globalVars.p, pClient))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookQueryClientCvarValue
func goHookQueryClientCvarValue(pClient *C.edict_t, cvarName *C.char) *C.char {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.QueryClientCvarValue != nil {
		metaResult, result := globalPluginState.engineHooks.QueryClientCvarValue(edictFromC(globalPluginState.globalVars.p, pClient), C.GoString(cvarName))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookQueryClientCvarValue2
func goHookQueryClientCvarValue2(pClient *C.edict_t, cvarName *C.char, requestID C.int) {
	if globalPluginState.engineHooks != nil && globalPluginState.engineHooks.QueryClientCvarValue2 != nil {
		metaResult := globalPluginState.engineHooks.QueryClientCvarValue2(edictFromC(globalPluginState.globalVars.p, pClient), C.GoString(cvarName), int(requestID))
		globalPluginState.metaGlobals.SetMres(MetaResult(metaResult))

		return
	}

	globalPluginState.metaGlobals.SetMres(MetaResultIgnored)
}
