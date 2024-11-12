package main

/*
#import <cvardef.h>

cvar_t * makeCvar(char *name, char *string, int flags, float value, cvar_t *next) {
	cvar_t *cvar = malloc(sizeof(cvar_t));

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
	CvarArchive       = 1 << iota // set to cause it to be saved to vars.rc
	CvarUserinfo                  // changes the client's info string
	CvarServer                    // notifies players when changed
	CvarExtdll                    // defined by external DLL
	CvarClientdll                 // defined by the client dll
	CvarProtected                 // It's a server cvar, but we don't send the data since it's a password, etc.  Sends 1 if it's not bland/zero, 0 otherwise as value
	CvarSponly                    // This cvar cannot be changed by clients connected to a multiplayer server.
	CvarPrintableonly             // This cvar's string cannot contain unprintable characters ( e.g., used for player name etc ).
	CvarUnlogged                  // If this is a FCVAR_SERVER, don't log changes to the log file / console if we are creating a log
)

type Cvar struct {
	p *C.cvar_t

	next *Cvar
}

func NewCvar(name, str string, flags int, value float32, next *Cvar) *Cvar {
	var n *C.cvar_t
	if next != nil {
		n = next.p
	}

	cv := &Cvar{
		p: C.makeCvar(C.CString(name), C.CString(str), C.int(flags), C.float(value), n),
	}

	if next != nil {
		cv.next = next
	}

	return cv
}

func cvarFromC(c *C.cvar_t) *Cvar {
	return &Cvar{
		p: c,
	}
}

func (c *Cvar) Name() string {
	return C.GoString(c.p.name)
}

func (c *Cvar) String() string {
	return C.GoString(c.p.string)
}

func (c *Cvar) Flags() int {
	return int(c.p.flags)
}

func (c *Cvar) Value() float32 {
	return float32(c.p.value)
}

func (c *Cvar) Next() *Cvar {
	return c.Next()
}

func (c *Cvar) Free() {
	C.free(unsafe.Pointer(c.p))
}
