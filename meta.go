package metamod_go

/*
#include "metamod/index.h"
*/
import "C"

const MetaInterfaceVersion = "5:13"

type MetaResult int

// Flags returned by a plugin's api function.
// NOTE: order is crucial, as greater/less comparisons are made.
const (
	MetaResultUnset     MetaResult = iota
	MetaResultIgnored              // plugin didn't take any action
	MetaResultHandled              // plugin did something, but real function should still be called
	MetaResultOverride             // call real function, but use my return value
	MetaResultSupercede            // skip real function; use my return value
)

type MetaGlobals struct {
	p *C.meta_globals_t
}

func MetaGlobalsFromC(p *C.meta_globals_t) *MetaGlobals {
	return &MetaGlobals{p}
}

func (m *MetaGlobals) Mres() MetaResult {
	return MetaResult(int(m.p.mres))
}

func (m *MetaGlobals) SetMres(mres MetaResult) {
	m.p.mres = C.META_RES(int(mres))
}

func (m *MetaGlobals) PrevMres() MetaResult {
	return MetaResult(int(m.p.prev_mres))
}

func (m *MetaGlobals) Status() MetaResult {
	return MetaResult(int(m.p.status))
}

const Success = 1
const Failure = 0
