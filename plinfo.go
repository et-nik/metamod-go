package main

/*
#include "metamod/plinfo.h"

plugin_info_t Plugin_info =
{
	"",								// ifvers
	"",								// name
	"",								// version
	"",								// date
	"",								// author
	"",								// url
	"",								// logtag
	PT_ANYTIME,						// (when) loadable
	PT_ANYTIME,						// (when) unloadable
};

#define PLID &Plugin_info
*/
import "C"

type PluginLoadTime int

const (
	PluginLoadTimeNever       PluginLoadTime = iota
	PluginLoadTimeStartup                    // should only be loaded/unloaded at initial hlds execution
	PluginLoadTimeChangeLevel                // can be loaded/unloaded between maps
	PluginLoadTimeAnyTime                    // can be loaded/unloaded at any time
	PluginLoadTimeAnyPause                   // can be loaded/unloaded at any time, and can be "paused" during a map
)

type PluginUnloadReason int

const (
	PluginUnloadReasonNull       PluginUnloadReason = iota
	PluginUnloadReasonIniDeleted                    // was deleted from plugins.ini
	PluginUnloadReasonFileNewer                     // file on disk is newer than last load
	PluginUnloadReasonCommand                       // requested by server/console command
	PluginUnloadReasonCmdForced                     // forced by server/console command
	PluginUnloadReasonDelayed                       // delayed from previous request; can't tell origin
	// only used for 'real_reason' on MPlugin::unload()
	PluginUnloadReasonPlugin    // requested by plugin function call
	PluginUnloadReasonPlgForced // forced by plugin function call
	PluginUnloadReasonReload    // forced unload by reload()
)

type PluginInfo struct {
	InterfaceVersion string
	Name             string
	Version          string
	Date             string
	Author           string
	Url              string
	LogTag           string
	Loadable         PluginLoadTime
	Unloadable       PluginLoadTime
}

func (p *PluginInfo) toCPluginInfo() *C.plugin_info_t {
	return &C.plugin_info_t{
		ifvers:     C.CString(p.InterfaceVersion),
		name:       C.CString(p.Name),
		version:    C.CString(p.Version),
		date:       C.CString(p.Date),
		author:     C.CString(p.Author),
		url:        C.CString(p.Url),
		logtag:     C.CString(p.LogTag),
		loadable:   C.enum_PLUG_LOADTIME(p.Loadable),
		unloadable: C.enum_PLUG_LOADTIME(p.Unloadable),
	}
}

func setCGlobalPluginInfo(p *PluginInfo) {
	C.Plugin_info.ifvers = C.CString(p.InterfaceVersion)
	C.Plugin_info.name = C.CString(p.Name)
	C.Plugin_info.version = C.CString(p.Version)
	C.Plugin_info.date = C.CString(p.Date)
	C.Plugin_info.author = C.CString(p.Author)
	C.Plugin_info.url = C.CString(p.Url)
	C.Plugin_info.logtag = C.CString(p.LogTag)
	C.Plugin_info.loadable = C.enum_PLUG_LOADTIME(p.Loadable)
	C.Plugin_info.unloadable = C.enum_PLUG_LOADTIME(p.Unloadable)
}
