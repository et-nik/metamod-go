package main

import "C"

//export goHookMessageBegin
func goHookMessageBegin(msgDest int, msgType int, pOrigin *float32, pEdict *C.edict_t) {
	//msgName := P.MetaUtilFuncs.GetUserMsgName(msgType, 0)
	//fmt.Println("msgName: ", msgName)
}

//export goHookMessageEnd
func goHookMessageEnd() {
}
