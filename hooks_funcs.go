package main

/*
#include <eiface.h>

extern vec3_t* castPtrToVec3(float *ptr);
*/
import "C"

//export goHookPrecacheModel
func goHookPrecacheModel(s *C.char) C.int {
	if P.EngineHooks != nil && P.EngineHooks.PrecacheModel != nil {
		metaResult, result := P.EngineHooks.PrecacheModel(C.GoString(s))
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

	return 0
}

//export goHookPrecacheSound
func goHookPrecacheSound(s *C.char) C.int {
	if P.EngineHooks != nil && P.EngineHooks.PrecacheSound != nil {
		metaResult, result := P.EngineHooks.PrecacheSound(C.GoString(s))
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

	return 0
}

//export goHookSetModel
func goHookSetModel(pEdict *C.edict_t, s *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.SetModel != nil {
		r := P.EngineHooks.SetModel(edictFromC(P.GlobalVars.p, pEdict), C.GoString(s))
		P.MetaGlobals.SetMres(r.MetaRes)

		return
	}

	P.MetaGlobals.SetMres(MetaResIgnored)
}

//export goHookModelIndex
func goHookModelIndex(s *C.char) C.int {
	if P.EngineHooks != nil && P.EngineHooks.ModelIndex != nil {
		metaResult, result := P.EngineHooks.ModelIndex(C.GoString(s))
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

	return 0
}

//export goHookModelFrames
func goHookModelFrames(index C.int) C.int {
	if P.EngineHooks != nil && P.EngineHooks.ModelFrames != nil {
		metaResult, result := P.EngineHooks.ModelFrames(int(index))
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

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

		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return
	}

	P.MetaGlobals.SetMres(MetaResIgnored)
}

//export goHookChangeLevel
func goHookChangeLevel(levelname *C.char, landmark *C.char) {
	if P.EngineHooks != nil && P.EngineHooks.ChangeLevel != nil {
		r := P.EngineHooks.ChangeLevel(C.GoString(levelname), C.GoString(landmark))
		P.MetaGlobals.SetMres(r.MetaRes)

		return
	}

	P.MetaGlobals.SetMres(MetaResIgnored)
}

//export goHookVecToYaw
func goHookVecToYaw(vec *C.float) C.float {
	if P.EngineHooks != nil && P.EngineHooks.VecToYaw != nil {
		v := C.castPtrToVec3(vec)
		metaResult, result := P.EngineHooks.VecToYaw(
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
		)
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return C.float(result)
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

	return 0
}

//export goHookVecToAngles
func goHookVecToAngles(vec *C.float) *C.float {
	if P.EngineHooks != nil && P.EngineHooks.VecToAngles != nil {
		v := C.castPtrToVec3(vec)

		metaResult, result := P.EngineHooks.VecToAngles(
			[3]float32{float32(v[0]), float32(v[1]), float32(v[2])},
		)
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		r := [3]C.float{C.float(result[0]), C.float(result[1]), C.float(result[2])}
		return &r[0]
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

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
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return
	}

	P.MetaGlobals.SetMres(MetaResIgnored)
}

//export goHookChangeYaw
func goHookChangeYaw(pEdict *C.edict_t) {
	if P.EngineHooks != nil && P.EngineHooks.ChangeYaw != nil {
		metaResult := P.EngineHooks.ChangeYaw(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return
	}

	P.MetaGlobals.SetMres(MetaResIgnored)
}

//export goHookChangePitch
func goHookChangePitch(pEdict *C.edict_t) {
	if P.EngineHooks != nil && P.EngineHooks.ChangePitch != nil {
		metaResult := P.EngineHooks.ChangePitch(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return
	}

	P.MetaGlobals.SetMres(MetaResIgnored)
}

//export goHookFindEntityByString
func goHookFindEntityByString(pEdict *C.edict_t, field *C.char, s *C.char) *C.edict_t {
	if P.EngineHooks != nil && P.EngineHooks.FindEntityByString != nil {
		metaResult, result := P.EngineHooks.FindEntityByString(
			edictFromC(P.GlobalVars.p, pEdict),
			FindEntityField(C.GoString(field)),
			C.GoString(s),
		)

		P.MetaGlobals.SetMres(metaResult.MetaRes)

		if result.p != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

	return nil
}

//export goHookGetEntityIllum
func goHookGetEntityIllum(pEdict *C.edict_t) C.int {
	if P.EngineHooks != nil && P.EngineHooks.GetEntityIllum != nil {
		metaResult, result := P.EngineHooks.GetEntityIllum(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

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
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		if result.p != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

	return nil
}

//export goHookFindClientInPVS
func goHookFindClientInPVS(pEdict *C.edict_t) *C.edict_t {
	if P.EngineHooks != nil && P.EngineHooks.FindClientInPVS != nil {
		metaResult, result := P.EngineHooks.FindClientInPVS(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		if result.p != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

	return nil
}

//export goHookEntitiesInPVS
func goHookEntitiesInPVS(pEdict *C.edict_t) *C.edict_t {
	if P.EngineHooks != nil && P.EngineHooks.EntitiesInPVS != nil {
		metaResult, result := P.EngineHooks.EntitiesInPVS(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		if result.p != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

	return nil
}

//export goHookMakeVectors
func goHookMakeVectors(angles *C.float) {
	if P.EngineHooks != nil && P.EngineHooks.MakeVectors != nil {
		v := C.castPtrToVec3(angles)

		metaResult := P.EngineHooks.MakeVectors([3]float32{float32(v[0]), float32(v[1]), float32(v[2])})
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return
	}

	P.MetaGlobals.SetMres(MetaResIgnored)
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

		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return
	}

	P.MetaGlobals.SetMres(MetaResIgnored)
}

//export goHookCreateEntity
func goHookCreateEntity() *C.edict_t {
	if P.EngineHooks != nil && P.EngineHooks.CreateEntity != nil {
		metaResult, result := P.EngineHooks.CreateEntity()
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		if result.p != nil {
			return result.p
		}

		return nil
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

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

	P.MetaGlobals.SetMres(MetaResIgnored)

	return nil
}

//export goHookRemoveEntity
func goHookRemoveEntity(pEdict *C.edict_t) {
	if P.EngineHooks != nil && P.EngineHooks.RemoveEntity != nil {
		r := P.EngineHooks.RemoveEntity(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(r.MetaRes)

		return
	}

	P.MetaGlobals.SetMres(MetaResIgnored)
}

//export goHookMakeStatic
func goHookMakeStatic(pEdict *C.edict_t) {
	if P.EngineHooks != nil && P.EngineHooks.MakeStatic != nil {
		r := P.EngineHooks.MakeStatic(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(r.MetaRes)

		return
	}

	P.MetaGlobals.SetMres(MetaResIgnored)
}

//export goHookEntIsOnFloor
func goHookEntIsOnFloor(pEdict *C.edict_t) C.int {
	if P.EngineHooks != nil && P.EngineHooks.EntIsOnFloor != nil {
		metaResult, result := P.EngineHooks.EntIsOnFloor(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		if result {
			return 1
		}

		return 0
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

	return 0
}

//export goHookDropToFloor
func goHookDropToFloor(pEdict *C.edict_t) C.int {
	if P.EngineHooks != nil && P.EngineHooks.DropToFloor != nil {
		metaResult, result := P.EngineHooks.DropToFloor(edictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

	return 0
}

//export goHookWalkMove
func goHookWalkMove(pEdict *C.edict_t, yaw C.float, dist C.float, mode C.int) C.int {
	if P.EngineHooks != nil && P.EngineHooks.WalkMove != nil {
		metaResult, result := P.EngineHooks.WalkMove(edictFromC(P.GlobalVars.p, pEdict), float32(yaw), float32(dist), WalkMoveMode(mode))
		P.MetaGlobals.SetMres(metaResult.MetaRes)

		return C.int(result)
	}

	P.MetaGlobals.SetMres(MetaResIgnored)

	return 0
}

////export goHookAddServerCommand
//func goHookAddServerCommand(s *C.char, f unsafe.Pointer) {
//	if P.EngineHooks != nil && P.EngineHooks.AddServerCommand != nil {
//		r := P.EngineHooks.AddServerCommand(C.GoString(s), f)
//		P.MetaGlobals.SetMres(r.MetaRes)
//
//		return
//	}
//
//	P.MetaGlobals.SetMres(MetaResIgnored)
//}
//
////export goHookMessageBegin
//func goHookMessageBegin(msgDest C.int, msgType C.int, pOrigin *C.float, pEdict *C.edict_t) {
//	if P.EngineHooks != nil && P.EngineHooks.MessageBegin != nil {
//		var origin *float32
//		if pOrigin != nil {
//			*origin = float32(*pOrigin)
//		}
//
//		r := P.EngineHooks.MessageBegin(int(msgDest), int(msgType), origin, edictFromC(P.GlobalVars.p, pEdict))
//		P.MetaGlobals.SetMres(r.MetaRes)
//	}
//
//	P.MetaGlobals.SetMres(MetaResIgnored)
//}
//
////export goHookMessageEnd
//func goHookMessageEnd() {
//	if P.EngineHooks != nil && P.EngineHooks.MessageEnd != nil {
//		r := P.EngineHooks.MessageEnd()
//		P.MetaGlobals.SetMres(r.MetaRes)
//	}
//
//	P.MetaGlobals.SetMres(MetaResIgnored)
//}
