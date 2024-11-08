package metamod

// #include "plinfo.h"
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

func (p *PluginInfo) ToC() *C.plugin_info_t {
	return &C.plugin_info_t{
		ifvers:     C.CString(p.InterfaceVersion),
		name:       C.CString(p.Name),
		version:    C.CString(p.Version),
		date:       C.CString(p.Date),
		author:     C.CString(p.Author),
		url:        C.CString(p.Url),
		logtag:     C.CString(p.LogTag),
		loadable:   C.int(p.Loadable),
		unloadable: C.int(p.Unloadable),
	}
}
