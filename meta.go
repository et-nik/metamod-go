package main

/*
#include "metamod/index.h"
*/
import "C"

const MetaInterfaceVersion = "5:13"

type MetaRes int

// Flags returned by a plugin's api function.
// NOTE: order is crucial, as greater/less comparisons are made.
const (
	MetaResUnset     MetaRes = iota
	MetaResIgnored           // plugin didn't take any action
	MetaResHandled           // plugin did something, but real function should still be called
	MetaResOverride          // call real function, but use my return value
	MetaResSupercede         // skip real function; use my return value
)

type MetaGlobals struct {
	p *C.meta_globals_t
}

func MetaGlobalsFromC(p *C.meta_globals_t) *MetaGlobals {
	return &MetaGlobals{p}
}

func (m *MetaGlobals) Mres() MetaRes {
	return MetaRes(int(m.p.mres))
}

func (m *MetaGlobals) SetMres(mres MetaRes) {
	m.p.mres = C.META_RES(int(mres))
}

func (m *MetaGlobals) PrevMres() MetaRes {
	return MetaRes(int(m.p.prev_mres))
}

func (m *MetaGlobals) Status() MetaRes {
	return MetaRes(int(m.p.status))
}
