package main

import "C"

var P = &Plugin{}

type Plugin struct {
	Info          *PluginInfo
	GlobalVars    *GlobalVars
	EngineFuncs   *EngineFuncs
	MetaGlobals   *MetaGlobals
	MetaUtilFuncs *MUtilFuncs

	EngineHooks     *EngineHooks
	EngineHooksPost *EngineHooks
}

type MetaCallbacks struct {
	MetaInit   func()
	MetaDetach func(now int, reason int) int
}

type EngineHookResult struct {
	MetaRes MetaRes
	Result  interface{}
}

type EngineHooks struct {
	MessageBegin func(msgDest int, msgType int, pOrigin *float32, pEdict *Edict) EngineHookResult
	MessageEnd   func() EngineHookResult
}
