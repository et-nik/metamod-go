package main

import "C"

//export goHookMessageBegin
func goHookMessageBegin(msgDest C.int, msgType C.int, pOrigin *C.float, pEdict *C.edict_t) {
	if P.EngineHooks.MessageBegin != nil {

		var origin *float32
		if pOrigin != nil {
			*origin = float32(*pOrigin)
		}

		r := P.EngineHooks.MessageBegin(int(msgDest), int(msgType), origin, EdictFromC(P.GlobalVars.p, pEdict))
		P.MetaGlobals.SetMres(r.MetaRes)
	}
}

//export goHookMessageEnd
func goHookMessageEnd() {
	if P.EngineHooks.MessageEnd != nil {
		r := P.EngineHooks.MessageEnd()
		P.MetaGlobals.SetMres(r.MetaRes)
	}
}
