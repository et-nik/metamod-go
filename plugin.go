package main

var P = &Plugin{}

type Plugin struct {
	Info          *PluginInfo
	GlobalVars    *GlobalVars
	EngineFuncs   *EngineFuncs
	MetaUtilFuncs *MUtilFuncs
}
