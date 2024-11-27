package metamod_go

/*
#import <cvardef.h>

#define MAX_CVARS 128

cvar_t memlockedCVars[MAX_CVARS];
int memlockedCVarsCount = 0;

cvar_t * makeCvar(char *name, char *string, int flags, float value, cvar_t *next) {
	if (memlockedCVarsCount >= MAX_CVARS) {
		return NULL;
	}

	cvar_t *cvar = &memlockedCVars[memlockedCVarsCount++];

	cvar->name = name;
	cvar->string = string;
	cvar->flags = flags;
	cvar->value = value;
	cvar->next = next;

	return cvar;
}

*/
import "C"
import "unsafe"

const (
	CVarArchive       = 1 << iota // set to cause it to be saved to vars.rc
	CVarUserinfo                  // changes the client's info string
	CVarServer                    // notifies players when changed
	CVarExtdll                    // defined by external DLL
	CVarClientdll                 // defined by the client dll
	CVarProtected                 // It's a server cvar, but we don't send the data since it's a password, etc.  Sends 1 if it's not bland/zero, 0 otherwise as value
	CVarSponly                    // This cvar cannot be changed by clients connected to a multiplayer server.
	CVarPrintableonly             // This cvar's string cannot contain unprintable characters ( e.g., used for player name etc ).
	CVarUnlogged                  // If this is a FCVAR_SERVER, don't log changes to the log file / console if we are creating a log
)

type CVar struct {
	p *C.cvar_t

	next *CVar
}

func NewCVar(name, initvalue string, flags int) *CVar {
	return &CVar{
		p: C.makeCvar(C.CString(name), C.CString(initvalue), C.int(flags), 0, nil),
	}
}

func cvarFromC(c *C.cvar_t) *CVar {
	return &CVar{
		p: c,
	}
}

func (c *CVar) Name() string {
	return C.GoString(c.p.name)
}

func (c *CVar) String() string {
	return C.GoString(c.p.string)
}

func (c *CVar) SetString(value string) {
	cs := C.CString(value)
	defer C.free(unsafe.Pointer(cs))

	globalPluginState.engineFuncs.CvarDirectSet(c.p, cs)
}

func (c *CVar) Float() float32 {
	return float32(c.p.value)
}

func (c *CVar) SetFloat(value float32) {
	globalPluginState.engineFuncs.CVarSetFloat(c.p, C.float(value))
}

func (c *CVar) Int() int {
	return int(c.p.value)
}

func (c *CVar) SetInt(value int) {
	globalPluginState.engineFuncs.CVarSetFloat(c.p, C.float(float32(value)))
}

func (c *CVar) Value() float32 {
	return float32(c.p.value)
}

func (c *CVar) SetValue(value float32) {
	globalPluginState.engineFuncs.CVarSetFloat(c.p, C.float(value))
}

func (c *CVar) Flags() int {
	return int(c.p.flags)
}

func (c *CVar) Next() *CVar {
	return c.Next()
}
