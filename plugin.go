package main

import "C"

var P = &Plugin{}

type Plugin struct {
	Info          *PluginInfo
	GlobalVars    *GlobalVars
	EngineFuncs   *EngineFuncs
	MetaUtilFuncs *MUtilFuncs
}

type MetaCallbacks struct {
	MetaInit   func()
	MetaDetach func(now int, reason int) int
}

type EngineCallbacks struct {
}
