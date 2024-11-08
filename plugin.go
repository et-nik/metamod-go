package main

var P *Plugin

type Plugin struct {
	Info        *PluginInfo
	EngineFuncs *EngineFuncs
}
