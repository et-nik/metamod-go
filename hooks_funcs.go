package main

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
	if P.EngineHooks != nil && P.EngineHooks.PrecacheModel != nil {
		metaResult, result := P.EngineHooks.PrecacheModel(C.GoString(s))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookPrecacheSound
func goHookPrecacheSound(s *C.char) C.int {
	if P.EngineHooks != nil && P.EngineHooks.PrecacheSound != nil {
		metaResult, result := P.EngineHooks.PrecacheSound(C.GoString(s))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookSetModel
func goHookSetModel(pEdict *C.edict_t, s *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.SetModel != nil {
		r := P.EngineHooks.SetModel(edictFromC(P.GlobalVars.p, pEdict), C.GoString(s))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookModelIndex
func goHookModelIndex(s *C.char) C.int {
	if P.EngineHooks != nil && P.EngineHooks.ModelIndex != nil {
		metaResult, result := P.EngineHooks.ModelIndex(C.GoString(s))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookModelFrames
func goHookModelFrames(index C.int) C.int {
	if P.EngineHooks != nil && P.EngineHooks.ModelFrames != nil {
		metaResult, result := P.EngineHooks.ModelFrames(int(index))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookSetSize
func goHookSetSize(pEdict *C.edict_t, mins, maxs *C.float) {
	if P.EngineHooks != nil && P.EngineHooks.SetSize != nil {
		minsVec := C.castPtrToVec3(mins)
		maxsVec := C.castPtrToVec3(maxs)

		metaResult := P.EngineHooks.SetSize(
			edictFromC(P.GlobalVars.p, pEdict),
			[3]float32{float32(minsVec[0]), float32(minsVec[1]), float32(minsVec[2])},
			[3]float32{float32(maxsVec[0]), float32(maxsVec[1]), float32(maxsVec[2])},
		)

		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookChangeLevel
func goHookChangeLevel(levelname *C.char, landmark *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.ChangeLevel != nil {
		r := P.EngineHooks.ChangeLevel(C.GoString(levelname), C.GoString(landmark))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookVecToYaw
func goHookVecToYaw(vec *C.float) C.float {
	if P.EngineHooks != nil && P.EngineHooks.VecToYaw != nil {
		v := C.castPtrToVec3(vec)
		metaResult, result := P.EngineHooks.VecToYaw(
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
		)
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.float(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookVecToAngles
func goHookVecToAngles(vec *C.float) *C.float {
	if P.EngineHooks != nil && P.EngineHooks.VecToAngles != nil {
		v := C.castPtrToVec3(vec)

		metaResult, result := P.EngineHooks.VecToAngles(
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
		)
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		r := [3]C.float{C.float(result[0]), C.float(result[1]), C.float(result[2])}
		return &r[0]
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookMoveToOrigin
func goHookMoveToOrigin(pEdict *C.edict_t, goal *C.float, dist C.float, moveType C.int) {
	if P.EngineHooks != nil && P.EngineHooks.MoveToOrigin != nil {
		v := C.castPtrToVec3(goal)

		metaResult := P.EngineHooks.MoveToOrigin(
			edictFromC(P.GlobalVars.p, pEdict),
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
			float32(dist),
			MoveType(int(moveType)),
		)
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookChangeYaw
func goHookChangeYaw(pEdict *C.edict_t) {
	if P.EngineHooks != nil && P.EngineHooks.ChangeYaw != nil {
		metaResult := P.EngineHooks.ChangeYaw(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookChangePitch
func goHookChangePitch(pEdict *C.edict_t) {
	if P.EngineHooks != nil && P.EngineHooks.ChangePitch != nil {
		metaResult := P.EngineHooks.ChangePitch(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookFindEntityByString
func goHookFindEntityByString(pEdict *C.edict_t, field *C.char, s *C.char) *C.edict_t {
	if P.EngineHooks != nil && P.EngineHooks.FindEntityByString != nil {
		metaResult, result := P.EngineHooks.FindEntityByString(
			edictFromC(P.GlobalVars.p, pEdict),
			FindEntityField(C.GoString(field)),
			C.GoString(s),
		)

		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result.p != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookGetEntityIllum
func goHookGetEntityIllum(pEdict *C.edict_t) C.int {
	if P.EngineHooks != nil && P.EngineHooks.GetEntityIllum != nil {
		metaResult, result := P.EngineHooks.GetEntityIllum(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookFindEntityInSphere
func goHookFindEntityInSphere(pEdict *C.edict_t, origin *C.float, radius C.float) *C.edict_t {
	if P.EngineHooks != nil && P.EngineHooks.FindEntityInSphere != nil {
		v := C.castPtrToVec3(origin)

		metaResult, result := P.EngineHooks.FindEntityInSphere(
			edictFromC(P.GlobalVars.p, pEdict),
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
			float32(radius),
		)
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result.p != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookFindClientInPVS
func goHookFindClientInPVS(pEdict *C.edict_t) *C.edict_t {
	if P.EngineHooks != nil && P.EngineHooks.FindClientInPVS != nil {
		metaResult, result := P.EngineHooks.FindClientInPVS(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result.p != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookEntitiesInPVS
func goHookEntitiesInPVS(pEdict *C.edict_t) *C.edict_t {
	if P.EngineHooks != nil && P.EngineHooks.EntitiesInPVS != nil {
		metaResult, result := P.EngineHooks.EntitiesInPVS(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result.p != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookMakeVectors
func goHookMakeVectors(angles *C.float) {
	if P.EngineHooks != nil && P.EngineHooks.MakeVectors != nil {
		v := C.castPtrToVec3(angles)

		metaResult := P.EngineHooks.MakeVectors([3]float32{float32(v[0]), float32(v[1]), float32(v[2])})
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookAngleVectors
func goHookAngleVectors(vector *C.float, forward, right, up *C.float) {
	if P.EngineHooks != nil && P.EngineHooks.AngleVectors != nil {
		v := C.castPtrToVec3(vector)

		forwardVec := C.castPtrToVec3(forward)
		rightVec := C.castPtrToVec3(right)
		upVec := C.castPtrToVec3(up)

		metaResult := P.EngineHooks.AngleVectors(
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
			[3]float32{float32(forwardVec[0]), float32(forwardVec[1]), float32(forwardVec[2])},
			[3]float32{float32(rightVec[0]), float32(rightVec[1]), float32(rightVec[2])},
			[3]float32{float32(upVec[0]), float32(upVec[1]), float32(upVec[2])},
		)

		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookCreateEntity
func goHookCreateEntity() *C.edict_t {
	if P.EngineHooks != nil && P.EngineHooks.CreateEntity != nil {
		metaResult, result := P.EngineHooks.CreateEntity()
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result.p != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookCreateNamedEntity
func goHookCreateNamedEntity(s *C.char) *C.edict_t {
	if P.EngineHooks != nil && P.EngineHooks.CreateNamedEntity != nil {
		result := P.EngineHooks.CreateNamedEntity(C.GoString(s))

		if result.p != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookRemoveEntity
func goHookRemoveEntity(pEdict *C.edict_t) {
	if P.EngineHooks != nil && P.EngineHooks.RemoveEntity != nil {
		r := P.EngineHooks.RemoveEntity(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMakeStatic
func goHookMakeStatic(pEdict *C.edict_t) {
	if P.EngineHooks != nil && P.EngineHooks.MakeStatic != nil {
		r := P.EngineHooks.MakeStatic(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookEntIsOnFloor
func goHookEntIsOnFloor(pEdict *C.edict_t) C.int {
	if P.EngineHooks != nil && P.EngineHooks.EntIsOnFloor != nil {
		metaResult, result := P.EngineHooks.EntIsOnFloor(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result {
			return 1
		}

		return 0
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookDropToFloor
func goHookDropToFloor(pEdict *C.edict_t) C.int {
	if P.EngineHooks != nil && P.EngineHooks.DropToFloor != nil {
		metaResult, result := P.EngineHooks.DropToFloor(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookWalkMove
func goHookWalkMove(pEdict *C.edict_t, yaw C.float, dist C.float, mode C.int) C.int {
	if P.EngineHooks != nil && P.EngineHooks.WalkMove != nil {
		metaResult, result := P.EngineHooks.WalkMove(edictFromC(P.GlobalVars.p, pEdict), float32(yaw), float32(dist), WalkMoveMode(mode))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookSetOrigin
func goHookSetOrigin(pEdict *C.edict_t, origin *C.float) {
	if P.EngineHooks != nil && P.EngineHooks.SetOrigin != nil {
		v := C.castPtrToVec3(origin)

		metaResult := P.EngineHooks.SetOrigin(
			edictFromC(P.GlobalVars.p, pEdict),
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
		)
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookEmitSound
func goHookEmitSound(pEdict *C.edict_t, channel C.int, sample *C.char, volume C.float, attenuation C.int, fFlags C.int, pitch C.int) {
	if P.EngineHooks != nil && P.EngineHooks.EmitSound != nil {
		metaResult := P.EngineHooks.EmitSound(
			edictFromC(P.GlobalVars.p, pEdict),
			int(channel),
			C.GoString(sample),
			float32(volume),
			int(attenuation),
			int(fFlags),
			int(pitch),
		)
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
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
	if P.EngineHooks != nil && P.EngineHooks.EmitAmbientSound != nil {
		v := C.castPtrToVec3(position)

		metaResult := P.EngineHooks.EmitAmbientSound(
			edictFromC(P.GlobalVars.p, pEdict),
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
			C.GoString(sample),
			float32(volume),
			float32(attenuation),
			int(fFlags),
			int(pitch),
		)
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return
}

//export goHookTraceLine
func goHookTraceLine(v1, v2 *C.float, fNoMonsters C.int, pentToSkip *C.edict_t, ptr *C.TraceResult) {
	if P.EngineHooks != nil && P.EngineHooks.TraceLine != nil {
		v1Vec := C.castPtrToVec3(v1)
		v2Vec := C.castPtrToVec3(v2)

		metaResult, result := P.EngineHooks.TraceLine(
			[3]float32{float32(v1Vec[0]), float32(v1Vec[1]), float32(v1Vec[2])},
			[3]float32{float32(v2Vec[0]), float32(v2Vec[1]), float32(v2Vec[2])},
			int(fNoMonsters),
			edictFromC(P.GlobalVars.p, pentToSkip),
		)

		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			*ptr = *result.ToC()
		}

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookTraceToss
func goHookTraceToss(pent, pentToIgnore *C.edict_t, ptr *C.TraceResult) {
	if P.EngineHooks != nil && P.EngineHooks.TraceToss != nil {
		metaResult, result := P.EngineHooks.TraceToss(edictFromC(P.GlobalVars.p, pent), edictFromC(P.GlobalVars.p, pentToIgnore))

		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			*ptr = *result.ToC()
		}

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookTraceMonsterHull
func goHookTraceMonsterHull(pent *C.edict_t, v1, v2 *C.float, fNoMonsters C.int, pentToSkip *C.edict_t, ptr *C.TraceResult) C.int {
	if P.EngineHooks != nil && P.EngineHooks.TraceMonsterHull != nil {
		v1Vec := C.castPtrToVec3(v1)
		v2Vec := C.castPtrToVec3(v2)

		metaResult, result, hit := P.EngineHooks.TraceMonsterHull(
			edictFromC(P.GlobalVars.p, pent),
			[3]float32{float32(v1Vec[0]), float32(v1Vec[1]), float32(v1Vec[2])},
			[3]float32{float32(v2Vec[0]), float32(v2Vec[1]), float32(v2Vec[2])},
			int(fNoMonsters),
			edictFromC(P.GlobalVars.p, pentToSkip),
		)

		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			*ptr = *result.ToC()
		}

		return C.int(hit)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookTraceHull
func goHookTraceHull(v1, v2 *C.float, fNoMonsters, hullNumber C.int, pentToSkip *C.edict_t, ptr *C.TraceResult) {
	if P.EngineHooks != nil && P.EngineHooks.TraceHull != nil {
		v1Vec := C.castPtrToVec3(v1)
		v2Vec := C.castPtrToVec3(v2)

		metaResult, result := P.EngineHooks.TraceHull(
			[3]float32{float32(v1Vec[0]), float32(v1Vec[1]), float32(v1Vec[2])},
			[3]float32{float32(v2Vec[0]), float32(v2Vec[1]), float32(v2Vec[2])},
			int(fNoMonsters),
			int(hullNumber),
			edictFromC(P.GlobalVars.p, pentToSkip),
		)

		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			*ptr = *result.ToC()
		}

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookTraceModel
func goHookTraceModel(v1, v2 *C.float, hullNumber C.int, pent *C.edict_t, ptr *C.TraceResult) {
	if P.EngineHooks != nil && P.EngineHooks.TraceModel != nil {
		v1Vec := C.castPtrToVec3(v1)
		v2Vec := C.castPtrToVec3(v2)

		metaResult, result := P.EngineHooks.TraceModel(
			[3]float32{float32(v1Vec[0]), float32(v1Vec[1]), float32(v1Vec[2])},
			[3]float32{float32(v2Vec[0]), float32(v2Vec[1]), float32(v2Vec[2])},
			int(hullNumber),
			edictFromC(P.GlobalVars.p, pent),
		)

		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			*ptr = *result.ToC()
		}

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookTraceTexture
func goHookTraceTexture(pTextureEntity *C.edict_t, v1, v2 *C.float) *C.char {
	if P.EngineHooks != nil && P.EngineHooks.TraceTexture != nil {
		v1Vec := C.castPtrToVec3(v1)
		v2Vec := C.castPtrToVec3(v2)

		metaResult, result := P.EngineHooks.TraceTexture(
			edictFromC(P.GlobalVars.p, pTextureEntity),
			[3]float32{float32(v1Vec[0]), float32(v1Vec[1]), float32(v1Vec[2])},
			[3]float32{float32(v2Vec[0]), float32(v2Vec[1]), float32(v2Vec[2])},
		)

		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			return (*C.char)(unsafe.Pointer(result.ToC()))
		}
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookGetAimVector
func goHookGetAimVector(pent *C.edict_t, speed C.float, ptr *C.float) {
	if P.EngineHooks != nil && P.EngineHooks.GetAimVector != nil {
		metaResult, result := P.EngineHooks.GetAimVector(edictFromC(P.GlobalVars.p, pent), float32(speed))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		C.setVec3FloatPtr(C.float(result[0]), C.float(result[1]), C.float(result[2]), ptr)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookServerCommand
func goHookServerCommand(s *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.ServerCommand != nil {
		r := P.EngineHooks.ServerCommand(C.GoString(s))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookServerExecute
func goHookServerExecute() {
	if P.EngineHooks != nil && P.EngineHooks.ServerExecute != nil {
		r := P.EngineHooks.ServerExecute()
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookClientCommand
func goHookClientCommand(pEdict *C.edict_t, format *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.ClientCommand != nil {
		r := P.EngineHooks.ClientCommand(edictFromC(P.GlobalVars.p, pEdict), C.GoString(format))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookAddServerCommand
func goHookAddServerCommand(s *C.char, f unsafe.Pointer) {
	if P.EngineHooks != nil && P.EngineHooks.AddServerCommand != nil {
		r := P.EngineHooks.AddServerCommand(C.GoString(s), f)
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookParticleEffect
func goHookParticleEffect(origin, direction *C.float, color, count C.float) {
	if P.EngineHooks != nil && P.EngineHooks.ParticleEffect != nil {
		originVec := C.castPtrToVec3(origin)
		directionVec := C.castPtrToVec3(direction)

		r := P.EngineHooks.ParticleEffect(
			[3]float32{float32(originVec[0]), float32(originVec[1]), float32(originVec[2])},
			[3]float32{float32(directionVec[0]), float32(directionVec[1]), float32(directionVec[2])},
			float32(color),
			float32(count),
		)
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookLightStyle
func goHookLightStyle(style C.int, value *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.LightStyle != nil {
		r := P.EngineHooks.LightStyle(int(style), C.GoString(value))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookDecalIndex
func goHookDecalIndex(s *C.char) C.int {
	if P.EngineHooks != nil && P.EngineHooks.DecalIndex != nil {
		metaResult, result := P.EngineHooks.DecalIndex(C.GoString(s))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookPointContents
func goHookPointContents(v *C.float) C.int {
	if P.EngineHooks != nil && P.EngineHooks.PointContents != nil {
		vVec := C.castPtrToVec3(v)

		metaResult, result := P.EngineHooks.PointContents([3]float32{float32(vVec[0]), float32(vVec[1]), float32(vVec[2])})
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookMessageBegin
func goHookMessageBegin(msgDest C.int, msgType C.int, pOrigin *C.float, pEdict *C.edict_t) {
	if P.EngineHooks != nil && P.EngineHooks.MessageBegin != nil {
		var origin *float32
		if pOrigin != nil {
			*origin = float32(*pOrigin)
		}

		r := P.EngineHooks.MessageBegin(int(msgDest), int(msgType), origin, edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(r.MetaResult)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageEnd
func goHookMessageEnd() {
	if P.EngineHooks != nil && P.EngineHooks.MessageEnd != nil {
		r := P.EngineHooks.MessageEnd()
		P.MetaGlobals.SetMres(r.MetaResult)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteByte
func goHookMessageWriteByte(b C.int) {
	if P.EngineHooks != nil && P.EngineHooks.MessageWriteByte != nil {
		r := P.EngineHooks.MessageWriteByte(int(b))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteChar
func goHookMessageWriteChar(c C.int) {
	if P.EngineHooks != nil && P.EngineHooks.MessageWriteChar != nil {
		r := P.EngineHooks.MessageWriteChar(int(c))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteShort
func goHookMessageWriteShort(s C.int) {
	if P.EngineHooks != nil && P.EngineHooks.MessageWriteShort != nil {
		r := P.EngineHooks.MessageWriteShort(int(s))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteLong
func goHookMessageWriteLong(l C.int) {
	if P.EngineHooks != nil && P.EngineHooks.MessageWriteLong != nil {
		r := P.EngineHooks.MessageWriteLong(int(l))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteAngle
func goHookMessageWriteAngle(f C.float) {
	if P.EngineHooks != nil && P.EngineHooks.MessageWriteAngle != nil {
		r := P.EngineHooks.MessageWriteAngle(float32(f))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteCoord
func goHookMessageWriteCoord(f C.float) {
	if P.EngineHooks != nil && P.EngineHooks.MessageWriteCoord != nil {
		r := P.EngineHooks.MessageWriteCoord(float32(f))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteString
func goHookMessageWriteString(s *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.MessageWriteString != nil {
		r := P.EngineHooks.MessageWriteString(C.GoString(s))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookMessageWriteEntity
func goHookMessageWriteEntity(id C.int) {
	if P.EngineHooks != nil && P.EngineHooks.MessageWriteEntity != nil {
		r := P.EngineHooks.MessageWriteEntity(int(id))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookCVarRegister
func goHookCVarRegister(cvar *C.cvar_t) {
	if P.EngineHooks != nil && P.EngineHooks.CVarRegister != nil {
		metaResult := P.EngineHooks.CVarRegister(cvarFromC(cvar))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookCVarGetString
func goHookCVarGetString(s *C.char) *C.char {
	if P.EngineHooks != nil && P.EngineHooks.CVarGetString != nil {
		metaResult, result := P.EngineHooks.CVarGetString(C.GoString(s))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookCVarGetFloat
func goHookCVarGetFloat(s *C.char) C.float {
	if P.EngineHooks != nil && P.EngineHooks.CVarGetFloat != nil {
		metaResult, result := P.EngineHooks.CVarGetFloat(C.GoString(s))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.float(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookCVarSetFloat
func goHookCVarSetFloat(s *C.char, value C.float) {
	if P.EngineHooks != nil && P.EngineHooks.CVarSetFloat != nil {
		r := P.EngineHooks.CVarSetFloat(C.GoString(s), float32(value))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookCVarSetString
func goHookCVarSetString(s, value *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.CVarSetString != nil {
		r := P.EngineHooks.CVarSetString(C.GoString(s), C.GoString(value))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookAlertMessage
func goHookAlertMessage(level C.ALERT_TYPE, format *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.AlertMessage != nil {
		r := P.EngineHooks.AlertMessage(AlertType(level), C.GoString(format))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookPvAllocEntPrivateData
func goHookPvAllocEntPrivateData(pEdict *C.edict_t, cb C.int32) unsafe.Pointer {
	if P.EngineHooks != nil && P.EngineHooks.PvAllocEntPrivateData != nil {
		metaResult, result := P.EngineHooks.PvAllocEntPrivateData(edictFromC(P.GlobalVars.p, pEdict), int32(cb))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return result
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookPvEntPrivateData
func goHookPvEntPrivateData(pEdict *C.edict_t) unsafe.Pointer {
	if P.EngineHooks != nil && P.EngineHooks.PvEntPrivateData != nil {
		metaResult, result := P.EngineHooks.PvEntPrivateData(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return result
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookFreeEntPrivateData
func goHookFreeEntPrivateData(pEdict *C.edict_t) {
	if P.EngineHooks != nil && P.EngineHooks.FreeEntPrivateData != nil {
		r := P.EngineHooks.FreeEntPrivateData(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookGetVarsOfEnt
func goHookGetVarsOfEnt(pEdict *C.edict_t) *C.entvars_t {
	if P.EngineHooks != nil && P.EngineHooks.GetVarsOfEnt != nil {
		metaResult, result := P.EngineHooks.GetVarsOfEnt(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookIndexOfEdict
func goHookIndexOfEdict(pEdict *C.edict_t) C.int {
	if P.EngineHooks != nil && P.EngineHooks.IndexOfEdict != nil {
		metaResult, result := P.EngineHooks.IndexOfEdict(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookPEntityOfEntIndex
func goHookPEntityOfEntIndex(index C.int) *C.edict_t {
	if P.EngineHooks != nil && P.EngineHooks.PEntityOfEntIndex != nil {
		metaResult, result := P.EngineHooks.PEntityOfEntIndex(int(index))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookFindEntityByVars
func goHookFindEntityByVars(vars *C.entvars_t) *C.edict_t {
	if P.EngineHooks != nil && P.EngineHooks.FindEntityByVars != nil {
		metaResult, result := P.EngineHooks.FindEntityByVars(entVarsFromC(P.GlobalVars.p, vars))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookGetModelPtr
func goHookGetModelPtr(pEdict *C.edict_t) *C.void {
	if P.EngineHooks != nil && P.EngineHooks.GetModelPtr != nil {
		metaResult, result := P.EngineHooks.GetModelPtr(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return (*C.void)(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookRegUserMsg
func goHookRegUserMsg(name *C.char, size C.int) C.int {
	if P.EngineHooks != nil && P.EngineHooks.RegUserMsg != nil {
		metaResult, result := P.EngineHooks.RegUserMsg(C.GoString(name), int(size))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookFunctionFromName
func goHookFunctionFromName(name *C.char) C.uint32 {
	if P.EngineHooks != nil && P.EngineHooks.FunctionFromName != nil {
		metaResult, result := P.EngineHooks.FunctionFromName(C.GoString(name))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.uint32(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookNameForFunction
func goHookNameForFunction(function C.uint32) *C.char {
	if P.EngineHooks != nil && P.EngineHooks.NameForFunction != nil {
		metaResult, result := P.EngineHooks.NameForFunction(uint32(function))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookClientPrint
func goHookClientPrint(pEdict *C.edict_t, level C.PRINT_TYPE, msg *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.ClientPrint != nil {
		metaResult := P.EngineHooks.ClientPrint(edictFromC(P.GlobalVars.p, pEdict), PrintType(level), C.GoString(msg))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookServerPrint
func goHookServerPrint(msg *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.ServerPrint != nil {
		metaResult := P.EngineHooks.ServerPrint(C.GoString(msg))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookGetAttachment
func goHookGetAttachment(pEdict *C.edict_t, attachment C.int, origin, angles *C.float) C.int {
	if P.EngineHooks != nil && P.EngineHooks.GetAttachment != nil {

		gorigin := [3]float32{}
		gangles := [3]float32{}

		metaResult := P.EngineHooks.GetAttachment(
			edictFromC(P.GlobalVars.p, pEdict),
			int(attachment),
			&gorigin,
			&gangles,
		)

		P.MetaGlobals.SetMres(metaResult.MetaResult)

		C.setVec3FloatPtr(C.float(gorigin[0]), C.float(gorigin[1]), C.float(gorigin[2]), origin)
		C.setVec3FloatPtr(C.float(gangles[0]), C.float(gangles[1]), C.float(gangles[2]), angles)

		return 0
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookRandomLong
func goHookRandomLong(low, high C.int) C.int {
	if P.EngineHooks != nil && P.EngineHooks.RandomLong != nil {
		metaResult, result := P.EngineHooks.RandomLong(int32(low), int32(high))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookRandomFloat
func goHookRandomFloat(low, high C.float) C.float {
	if P.EngineHooks != nil && P.EngineHooks.RandomFloat != nil {
		metaResult, result := P.EngineHooks.RandomFloat(float32(low), float32(high))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.float(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookSetView
func goHookSetView(pClient, pEdict *C.edict_t) {
	if P.EngineHooks != nil && P.EngineHooks.SetView != nil {
		r := P.EngineHooks.SetView(edictFromC(P.GlobalVars.p, pClient), edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookTime
func goHookTime() C.float {
	if P.EngineHooks != nil && P.EngineHooks.Time != nil {
		metaResult, result := P.EngineHooks.Time()
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.float(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookCrosshairAngle
func goHookCrosshairAngle(pClient *C.edict_t, pitch, yaw C.float) {
	if P.EngineHooks != nil && P.EngineHooks.CrosshairAngle != nil {
		r := P.EngineHooks.CrosshairAngle(edictFromC(P.GlobalVars.p, pClient), float32(pitch), float32(yaw))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookLoadFileForMe
func goHookLoadFileForMe(name *C.char, _ *C.int) *C.byte {
	if P.EngineHooks != nil && P.EngineHooks.LoadFileForMe != nil {
		metaResult, result, err := P.EngineHooks.LoadFileForMe(C.GoString(name))
		if err != nil {
			return nil
		}

		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			return (*C.byte)(unsafe.Pointer(&result[0]))
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookFreeFile
//func goHookFreeFile(buffer *C.byte) {
//	if P.EngineHooks != nil && P.EngineHooks.FreeFile != nil {
//		r := P.EngineHooks.FreeFile(buffer)
//		P.MetaGlobals.SetMres(r.MetaResult)
//
//		return
//	}
//
//	P.MetaGlobals.SetMres(MetaResultIgnored)
//}

//export goHookGetGameDir
func goHookGetGameDir() *C.char {
	if P.EngineHooks != nil && P.EngineHooks.GetGameDir != nil {
		metaResult, result := P.EngineHooks.GetGameDir()
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookCvar_RegisterVariable
func goHookCvar_RegisterVariable(cvar *C.cvar_t) {
	if P.EngineHooks != nil && P.EngineHooks.CVarRegisterVariable != nil {
		metaResult := P.EngineHooks.CVarRegisterVariable(cvarFromC(cvar))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookFadeClientVolume
func goHookFadeClientVolume(pEdict *C.edict_t, fadePercent, fadeOutSeconds, holdTime, fadeInSeconds C.int) {
	if P.EngineHooks != nil && P.EngineHooks.FadeClientVolume != nil {
		r := P.EngineHooks.FadeClientVolume(
			edictFromC(P.GlobalVars.p, pEdict),
			int(fadePercent),
			int(fadeOutSeconds),
			int(holdTime),
			int(fadeInSeconds),
		)
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookSetClientMaxspeed
func goHookSetClientMaxspeed(pEdict *C.edict_t, speed C.float) {
	if P.EngineHooks != nil && P.EngineHooks.SetClientMaxspeed != nil {
		r := P.EngineHooks.SetClientMaxspeed(edictFromC(P.GlobalVars.p, pEdict), float32(speed))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookCreateFakeClient
func goHookCreateFakeClient(name *C.char) *C.edict_t {
	if P.EngineHooks != nil && P.EngineHooks.CreateFakeClient != nil {
		metaResult, result := P.EngineHooks.CreateFakeClient(C.GoString(name))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result.p != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookRunPlayerMove
func goHookRunPlayerMove(pEdict *C.edict_t, viewangles, forwardmove, sidemove, upmove *C.float, buttons C.ushort, impulse, msec C.byte) {
	if P.EngineHooks != nil && P.EngineHooks.RunPlayerMove != nil {
		v := C.castPtrToVec3(viewangles)

		r := P.EngineHooks.RunPlayerMove(
			edictFromC(P.GlobalVars.p, pEdict),
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
			float32(*forwardmove),
			float32(*sidemove),
			float32(*upmove),
			uint16(buttons),
			uint16(impulse),
			uint16(msec),
		)
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookNumberOfEntities
func goHookNumberOfEntities() C.int {
	if P.EngineHooks != nil && P.EngineHooks.NumberOfEntities != nil {
		metaResult, result := P.EngineHooks.NumberOfEntities()
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookGetInfoKeyBuffer
func goHookGetInfoKeyBuffer(pEdict *C.edict_t) *C.char {
	if P.EngineHooks != nil && P.EngineHooks.GetInfoKeyBuffer != nil {
		metaResult, result := P.EngineHooks.GetInfoKeyBuffer(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			return (*C.char)(unsafe.Pointer(&result[0]))
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookInfoKeyValue
func goHookInfoKeyValue(info, key *C.char) *C.char {
	if P.EngineHooks != nil && P.EngineHooks.InfoKeyValue != nil {
		metaResult, result := P.EngineHooks.InfoKeyValue([]byte(C.GoString(info)), C.GoString(key))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookSetKeyValue
func goHookSetKeyValue(info, key, value *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.SetKeyValue != nil {
		r := P.EngineHooks.SetKeyValue([]byte(C.GoString(info)), C.GoString(key), C.GoString(value))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookSetClientKeyValue
func goHookSetClientKeyValue(clientIndex C.int, key, value *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.SetClientKeyValue != nil {
		r := P.EngineHooks.SetClientKeyValue(int(clientIndex), C.GoString(key), C.GoString(value))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookIsMapValid
func goHookIsMapValid(mapName *C.char) C.int {
	if P.EngineHooks != nil && P.EngineHooks.IsMapValid != nil {
		metaResult, result := P.EngineHooks.IsMapValid(C.GoString(mapName))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result {
			return 1
		}

		return 0
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookStaticDecal
func goHookStaticDecal(origin *C.float, decalIndex, entityIndex, modelIndex C.int) {
	if P.EngineHooks != nil && P.EngineHooks.StaticDecal != nil {
		v := C.castPtrToVec3(origin)

		r := P.EngineHooks.StaticDecal(
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
			int(decalIndex),
			int(entityIndex),
			int(modelIndex),
		)
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookPrecacheGeneric
func goHookPrecacheGeneric(s *C.char) C.int {
	if P.EngineHooks != nil && P.EngineHooks.PrecacheGeneric != nil {
		metaResult, result := P.EngineHooks.PrecacheGeneric(C.GoString(s))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookGetPlayerUserId
func goHookGetPlayerUserId(pEdict *C.edict_t) C.int {
	if P.EngineHooks != nil && P.EngineHooks.GetPlayerUserId != nil {
		metaResult, result := P.EngineHooks.GetPlayerUserId(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookIsDedicatedServer
func goHookIsDedicatedServer() C.int {
	if P.EngineHooks != nil && P.EngineHooks.IsDedicatedServer != nil {
		metaResult, result := P.EngineHooks.IsDedicatedServer()
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result {
			return 1
		}

		return 0
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookCVarGetPointer
func goHookCVarGetPointer(name *C.char) *C.cvar_t {
	if P.EngineHooks != nil && P.EngineHooks.CVarGetPointer != nil {
		metaResult, result := P.EngineHooks.CVarGetPointer(C.GoString(name))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookGetPlayerWONId
func goHookGetPlayerWONId(pEdict *C.edict_t) C.uint {
	if P.EngineHooks != nil && P.EngineHooks.GetPlayerWONID != nil {
		metaResult, result := P.EngineHooks.GetPlayerWONID(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.uint(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return 0
}

//export goHookInfo_RemoveKey
func goHookInfo_RemoveKey(info, key *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.InfoRemoveKey != nil {
		r := P.EngineHooks.InfoRemoveKey(C.GoString(info), C.GoString(key))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookGetPhysicsKeyValue
func goHookGetPhysicsKeyValue(pEdict *C.edict_t, key *C.char) *C.char {
	if P.EngineHooks != nil && P.EngineHooks.GetPhysicsKeyValue != nil {
		metaResult, result := P.EngineHooks.GetPhysicsKeyValue(edictFromC(P.GlobalVars.p, pEdict), C.GoString(key))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookSetPhysicsKeyValue
func goHookSetPhysicsKeyValue(pEdict *C.edict_t, key, value *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.SetPhysicsKeyValue != nil {
		r := P.EngineHooks.SetPhysicsKeyValue(edictFromC(P.GlobalVars.p, pEdict), C.GoString(key), C.GoString(value))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookGetPhysicsInfoString
func goHookGetPhysicsInfoString(pEdict *C.edict_t) *C.char {
	if P.EngineHooks != nil && P.EngineHooks.GetPhysicsInfoString != nil {
		metaResult, result := P.EngineHooks.GetPhysicsInfoString(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookPrecacheEvent
func goHookPrecacheEvent(eventType C.int, eventName *C.char) C.int {
	if P.EngineHooks != nil && P.EngineHooks.PrecacheEvent != nil {
		metaResult, result := P.EngineHooks.PrecacheEvent(int(eventType), C.GoString(eventName))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

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
	if P.EngineHooks != nil && P.EngineHooks.PlaybackEvent != nil {
		originVec := C.castPtrToVec3(origin)
		anglesVec := C.castPtrToVec3(angles)

		boolParam1 := int(bparam1) == 1
		boolParam2 := int(bparam2) == 1

		metaResult := P.EngineHooks.PlaybackEvent(
			int(flags),
			edictFromC(P.GlobalVars.p, pInvoker),
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
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookSetFatPVS
func goHookSetFatPVS(origin *C.float) *C.byte {
	if P.EngineHooks != nil && P.EngineHooks.SetFatPVS != nil {
		v := C.castPtrToVec3(origin)

		metaResult, result := P.EngineHooks.SetFatPVS([3]float32{float32(v[0]), float32(v[1]), float32(v[2])})
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			return (*C.byte)(result)
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookSetFatPAS
func goHookSetFatPAS(origin *C.float) *C.byte {
	if P.EngineHooks != nil && P.EngineHooks.SetFatPAS != nil {
		v := C.castPtrToVec3(origin)

		metaResult, result := P.EngineHooks.SetFatPAS([3]float32{float32(v[0]), float32(v[1]), float32(v[2])})
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != nil {
			return (*C.byte)(result)
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookCvar_DirectSet
func goHookCvar_DirectSet(cvar *C.cvar_t, value *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.CvarDirectSet != nil {
		r := P.EngineHooks.CvarDirectSet(cvarFromC(cvar), C.GoString(value))
		P.MetaGlobals.SetMres(r.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

//export goHookGetPlayerStats
func goHookGetPlayerStats(pClient *C.edict_t, ping, packetLoss *C.int) {
	if P.EngineHooks != nil && P.EngineHooks.GetPlayerStats != nil {
		metaResult, pingResult, lossResult := P.EngineHooks.GetPlayerStats(edictFromC(P.GlobalVars.p, pClient))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		*ping = C.int(pingResult)
		*packetLoss = C.int(lossResult)
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}

// const char *(*pfnGetPlayerAuthId)(edict_t *e);
//
//export goHookGetPlayerAuthId
func goHookGetPlayerAuthId(pClient *C.edict_t) *C.char {
	if P.EngineHooks != nil && P.EngineHooks.GetPlayerAuthId != nil {
		metaResult, result := P.EngineHooks.GetPlayerAuthId(edictFromC(P.GlobalVars.p, pClient))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookQueryClientCvarValue
func goHookQueryClientCvarValue(pClient *C.edict_t, cvarName *C.char) *C.char {
	if P.EngineHooks != nil && P.EngineHooks.QueryClientCvarValue != nil {
		metaResult, result := P.EngineHooks.QueryClientCvarValue(edictFromC(P.GlobalVars.p, pClient), C.GoString(cvarName))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		if result != "" {
			return C.CString(result)
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)

	return nil
}

//export goHookQueryClientCvarValue2
func goHookQueryClientCvarValue2(pClient *C.edict_t, cvarName *C.char, requestID C.int) {
	if P.EngineHooks != nil && P.EngineHooks.QueryClientCvarValue2 != nil {
		metaResult := P.EngineHooks.QueryClientCvarValue2(edictFromC(P.GlobalVars.p, pClient), C.GoString(cvarName), int(requestID))
		P.MetaGlobals.SetMres(metaResult.MetaResult)

		return
	}

	P.MetaGlobals.SetMres(MetaResultIgnored)
}
