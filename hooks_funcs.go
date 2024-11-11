package main

/*
//extern void **GetVTable(void *pthis, int size);
*/
import "C"
import "fmt"

//export goHookMessageBegin
func goHookMessageBegin(msgDest C.int, msgType C.int, pOrigin *C.float, pEdict *C.edict_t) {
	if P.EngineHooks != nil && P.EngineHooks.MessageBegin != nil {

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
	if P.EngineHooks != nil && P.EngineHooks.MessageEnd != nil {
		r := P.EngineHooks.MessageEnd()
		P.MetaGlobals.SetMres(r.MetaRes)
	}
}

func findFuncs() {
	//h := C.dlopen(nil, 0)
	fmt.Println("findFuncs")

	//_ZN11CBasePlayer10TakeDamageEP9entvars_sS1_fi
	takeDamage := C.dlsym(C.RTLD_DEFAULT, C.CString("_ZN11CBasePlayer10TakeDamageEP9entvars_sS1_fi"))

	if takeDamage != nil {
		fmt.Println("Found TakeDamage")
	} else {
		fmt.Println("TakeDamage not found")
	}

	goHookMessageBeginFunc := C.dlsym(C.RTLD_DEFAULT, C.CString("goHookMessageBegin"))

	if goHookMessageBeginFunc != nil {
		fmt.Println("Found goHookMessageBegin")
	} else {
		fmt.Println("goHookMessageBegin not found")
	}
}
