package metamod_go

import "C"
import "unsafe"

var globalPluginState = &Plugin{
	engineFuncs:   &EngineFuncs{},
	metaUtilFuncs: &MUtilFuncs{},

	engineHooks: &EngineHooks{},
}

type timelineStatus int

const (
	statusUnknown timelineStatus = iota
	statusLibLoaded
	statusMetaQueried
	statusMetaAttached
	statusGameStarted
	statusMetaDetached
)

type Plugin struct {
	globalVars     *GlobalVars
	metaGlobals    *MetaGlobals
	timelineStatus timelineStatus

	engineFuncs   *EngineFuncs
	metaCallbacks *MetaCallbacks
	metaUtilFuncs *MUtilFuncs

	engineHooks     *EngineHooks
	engineHooksPost *EngineHooks
}

func SetPluginInfo(info *PluginInfo) error {
	setCGlobalPluginInfo(info)

	return nil
}

func SetEngineHooks(hooks *EngineHooks) error {
	if globalPluginState.timelineStatus >= statusMetaQueried {
		return ErrMetaQueried
	}

	globalPluginState.engineHooks = hooks

	return nil
}

func GetEngineFuncs() (*EngineFuncs, error) {
	if globalPluginState.timelineStatus < statusLibLoaded {
		return nil, ErrLibNotLoaded
	}

	return globalPluginState.engineFuncs, nil
}

func GetMetaUtilFuncs() (*MUtilFuncs, error) {
	return globalPluginState.metaUtilFuncs, nil
}

func GetEngineHooks() *EngineHooks {
	return globalPluginState.engineHooks
}

func SetMetaCallbacks(callbacks *MetaCallbacks) error {
	if globalPluginState.timelineStatus >= statusLibLoaded {
		return ErrLibIsLoaded
	}

	globalPluginState.metaCallbacks = callbacks

	return nil
}

type MetaCallbacks struct {
	MetaInit   func()
	MetaQuery  func() int
	MetaAttach func(now int) int
	MetaDetach func(now int, reason int) int
}

type EngineHookResult struct {
	MetaResult
}

// EngineHookResultIgnored Plugin didn't take any action.
var EngineHookResultIgnored = EngineHookResult{MetaResult: MetaResultIgnored}

// EngineHookResultHandled Plugin did something, but real function should still be called.
var EngineHookResultHandled = EngineHookResult{MetaResult: MetaResultHandled}

// EngineHookResultOverride Call real function, but use my return value.
var EngineHookResultOverride = EngineHookResult{MetaResult: MetaResultOverride}

// EngineHookResultSupercede Skip real function; use my return value.
var EngineHookResultSupercede = EngineHookResult{MetaResult: MetaResultSupercede}

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
	TraceModel       func(v1, v2 [3]float32, hullNumber int, pent *Edict) (EngineHookResult, *TraceResult)
	TraceTexture     func(pent *Edict, v1, v2 [3]float32) (EngineHookResult, *Texture)
	GetAimVector     func(ent *Edict, speed float32) (EngineHookResult, [3]float32)
	ServerCommand    func(str string) EngineHookResult
	AddServerCommand func(name string, fn unsafe.Pointer) EngineHookResult
	ServerExecute    func() EngineHookResult
	ClientCommand    func(pEdict *Edict, format string) EngineHookResult
	ParticleEffect   func(
		origin, direction [3]float32,
		color, count float32,
	) EngineHookResult
	LightStyle            func(style int, value string) EngineHookResult
	DecalIndex            func(name string) (EngineHookResult, int)
	PointContents         func(v [3]float32) (EngineHookResult, int)
	MessageBegin          func(msgDest int, msgType int, pOrigin *float32, pEdict *Edict) EngineHookResult
	MessageEnd            func() EngineHookResult
	CVarRegister          func(cvar *Cvar) EngineHookResult
	CVarGetString         func(name string) (EngineHookResult, string)
	CVarGetFloat          func(name string) (EngineHookResult, float32)
	CVarSetFloat          func(name string, value float32) EngineHookResult
	CVarSetString         func(name, value string) EngineHookResult
	AlertMessage          func(alertType AlertType, msg string) EngineHookResult
	MessageWriteByte      func(b int) EngineHookResult
	MessageWriteChar      func(c int) EngineHookResult
	MessageWriteShort     func(s int) EngineHookResult
	MessageWriteLong      func(l int) EngineHookResult
	MessageWriteAngle     func(f float32) EngineHookResult
	MessageWriteCoord     func(f float32) EngineHookResult
	MessageWriteString    func(s string) EngineHookResult
	MessageWriteEntity    func(id int) EngineHookResult
	PvAllocEntPrivateData func(ent *Edict, size int32) (EngineHookResult, unsafe.Pointer)
	PvEntPrivateData      func(ent *Edict) (EngineHookResult, unsafe.Pointer)
	FreeEntPrivateData    func(ent *Edict) EngineHookResult
	GetVarsOfEnt          func(ent *Edict) (EngineHookResult, *EntVars)
	IndexOfEdict          func(ent *Edict) (EngineHookResult, int)
	PEntityOfEntIndex     func(index int) (EngineHookResult, *Edict)
	FindEntityByVars      func(vars *EntVars) (EngineHookResult, *Edict)
	GetModelPtr           func(pEdict *Edict) (EngineHookResult, unsafe.Pointer)
	RegUserMsg            func(name string, size int) (EngineHookResult, int)
	FunctionFromName      func(name string) (EngineHookResult, uint32)
	NameForFunction       func(fn uint32) (EngineHookResult, string)
	ClientPrint           func(pEdict *Edict, printType PrintType, msg string) EngineHookResult
	ServerPrint           func(msg string) EngineHookResult
	GetAttachment         func(pEdict *Edict, attachmentIndex int, rgflOrigin, rgflAngles *[3]float32) EngineHookResult
	RandomLong            func(low, high int32) (EngineHookResult, int32)
	RandomFloat           func(low, high float32) (EngineHookResult, float32)
	SetView               func(pEdict *Edict, pOther *Edict) EngineHookResult
	Time                  func() (EngineHookResult, float32)
	CrosshairAngle        func(pClient *Edict, pitch, yaw float32) EngineHookResult
	LoadFileForMe         func(filename string) (EngineHookResult, []byte, error)
	//FreeFile             func(buffer []byte) EngineHookResult
	GetGameDir           func() (EngineHookResult, string)
	CVarRegisterVariable func(variable *Cvar) EngineHookResult
	FadeClientVolume     func(pEdict *Edict, fadePercent, fadeOutSeconds, holdTime, fadeInSeconds int) EngineHookResult
	SetClientMaxspeed    func(e *Edict, maxSpeed float32) EngineHookResult
	CreateFakeClient     func(name string) (EngineHookResult, *Edict)
	RunPlayerMove        func(
		client *Edict,
		viewAngles [3]float32,
		forwardMove, sideMove, upMove float32,
		buttons uint16,
		impulse uint16,
		msec uint16,
	) EngineHookResult
	NumberOfEntities     func() (EngineHookResult, int)
	GetInfoKeyBuffer     func(e *Edict) (EngineHookResult, []byte)
	InfoKeyValue         func(infoBuffer []byte, key string) (EngineHookResult, string)
	SetKeyValue          func(infoBuffer []byte, key, value string) EngineHookResult
	SetClientKeyValue    func(clientIndex int, key, value string) EngineHookResult
	IsMapValid           func(filename string) (EngineHookResult, bool)
	StaticDecal          func(origin [3]float32, decalIndex int, entityIndex, modelIndex int) EngineHookResult
	PrecacheGeneric      func(modelName string) (EngineHookResult, int)
	GetPlayerUserId      func(e *Edict) (EngineHookResult, int)
	IsDedicatedServer    func() (EngineHookResult, bool)
	CVarGetPointer       func(name string) (EngineHookResult, *Cvar)
	GetPlayerWONID       func(e *Edict) (EngineHookResult, uint)
	InfoRemoveKey        func(infobuffer, key string) EngineHookResult
	GetPhysicsKeyValue   func(client *Edict, key string) (EngineHookResult, string)
	SetPhysicsKeyValue   func(client *Edict, key, value string) EngineHookResult
	GetPhysicsInfoString func(client *Edict) (EngineHookResult, string)
	PrecacheEvent        func(eventType int, eventName string) (EngineHookResult, int)
	PlaybackEvent        func(
		flags int,
		invoker *Edict,
		eventIndex uint16,
		delay float32,
		origin, angles [3]float32,
		fparam1, fparam2 float32,
		iparam1, iparam2 int,
		bparam1, bparam2 bool,
	) EngineHookResult
	SetFatPVS             func(origin [3]float32) (EngineHookResult, unsafe.Pointer)
	SetFatPAS             func(origin [3]float32) (EngineHookResult, unsafe.Pointer)
	CvarDirectSet         func(cvar *Cvar, value string) EngineHookResult
	GetPlayerStats        func(client *Edict) (EngineHookResult, int, int)
	GetPlayerAuthId       func(client *Edict) (EngineHookResult, string)
	QueryClientCvarValue  func(player *Edict, cvarName string) (EngineHookResult, string)
	QueryClientCvarValue2 func(player *Edict, cvarName string, requestID int) EngineHookResult
}
