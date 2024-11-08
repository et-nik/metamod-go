package main

/*
typedef struct plugin_info_t
{
    const char*     ifvers;     // meta_interface version
    const char*     name;       // full name of plugin
    const char*     version;    // version
    const char*     date;       // date
    const char*     author;     // author name/email
    const char*     url;        // URL
    const char*     logtag;     // log message prefix (unused right now)
    int   loadable;   // when loadable
    int   unloadable; // when unloadable
} plugin_info_t;
*/
import "C"

type PluginInfo struct {
	p C.plugin_info_t
}
