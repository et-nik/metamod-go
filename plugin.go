package main

import "C"

var P = &Plugin{
	EngineFuncs:   &EngineFuncs{},
	MetaUtilFuncs: &MUtilFuncs{},
}

type Plugin struct {
	Info *PluginInfo

	GlobalVars  *GlobalVars
	MetaGlobals *MetaGlobals

	EngineFuncs   *EngineFuncs
	MetaUtilFuncs *MUtilFuncs

	EngineHooks     *EngineHooks
	EngineHooksPost *EngineHooks
}

type MetaCallbacks struct {
	MetaInit   func()
	MetaDetach func(now int, reason int) int
}

type EngineHookResult struct {
	MetaRes
}

type EngineHooks struct {
	PrecacheModel      func(modelName string) (EngineHookResult, int)
	PrecacheSound      func(soundName string) (EngineHookResult, int)
	SetModel           func(e *Edict, model string) EngineHookResult
	ModelIndex         func(name string) (EngineHookResult, int)
	ModelFrames        func(index int) (EngineHookResult, int)
	SetSize            func(e *Edict, mins, maxs [3]float32) EngineHookResult
	ChangeLevel        func(mapName string, landmark string) EngineHookResult
	VecToYaw           func(vec [3]float32) (EngineHookResult, float32)
	VecToAngles        func(vec [3]float32) (EngineHookResult, [3]float32)
	MoveToOrigin       func(e *Edict, goal [3]float32, dist float32, moveType MoveType) EngineHookResult
	ChangeYaw          func(e *Edict) EngineHookResult
	ChangePitch        func(e *Edict) EngineHookResult
	FindEntityByString func(start *Edict, field FindEntityField, value string) (EngineHookResult, *Edict)
	GetEntityIllum     func(e *Edict) (EngineHookResult, int)
	FindEntityInSphere func(start *Edict, origin [3]float32, radius float32) (EngineHookResult, *Edict)
	FindClientInPVS    func(e *Edict) (EngineHookResult, *Edict)
	EntitiesInPVS      func(e *Edict) (EngineHookResult, *Edict)
	MakeVectors        func(angles [3]float32) EngineHookResult
	AngleVectors       func(vector [3]float32, forward, right, up [3]float32) EngineHookResult
	CreateEntity       func() (EngineHookResult, *Edict)
	CreateNamedEntity  func(className string) *Edict
	RemoveEntity       func(e *Edict) EngineHookResult
	MakeStatic         func(e *Edict) EngineHookResult
	EntIsOnFloor       func(e *Edict) (EngineHookResult, bool)
	DropToFloor        func(e *Edict) (EngineHookResult, int)
	WalkMove           func(e *Edict, yaw float32, dist float32, mode WalkMoveMode) (EngineHookResult, int)
	SetOrigin          func(e *Edict, origin [3]float32) EngineHookResult
	EmitSound          func(e *Edict, channel int, sample string, volume float32, attenuation int, fFlags int, pitch int) EngineHookResult
	EmitAmbientSound   func(
		e *Edict,
		origin [3]float32,
		sample string,
		volume, attenuation float32,
		flags int,
		pitch int,
	) EngineHookResult
	TraceLine func(
		v1, v2 [3]float32,
		noMonsters int,
		pentToSkip *Edict,
	) (EngineHookResult, *TraceResult)
	TraceToss        func(pent, pentToIgnore *Edict) (EngineHookResult, *TraceResult)
	TraceMonsterHull func(
		pent *Edict,
		v1, v2 [3]float32,
		noMonsters int,
		pentToSkip *Edict,
	) (EngineHookResult, *TraceResult, int)
	TraceHull func(
		v1, v2 [3]float32,
		noMonsters, hullNumber int,
		pentToSkip *Edict,
	) (EngineHookResult, *TraceResult)
	TraceModel    func(v1, v2 [3]float32, hullNumber int, pent *Edict) (EngineHookResult, *TraceResult)
	TraceTexture  func(pent *Edict, v1, v2 [3]float32) (EngineHookResult, *Texture)
	GetAimVector  func(ent *Edict, speed float32) (EngineHookResult, [3]float32)
	ServerCommand func(str string) EngineHookResult
	ServerExecute func() EngineHookResult
	ClientCommand func(pEdict *Edict, format string) EngineHookResult

	// --

	//AddServerCommand func(name string, fn unsafe.Pointer) EngineHookResult
	//EntityOfEntIndex func(index int) EngineHookResult
	//MessageBegin func(msgDest int, msgType int, pOrigin *float32, pEdict *Edict) EngineHookResult
	//MessageEnd   func() EngineHookResult
}
