package main

/*
const char* ReadString(globalvars_t *gpGlobals, int offset) {
	return (const char *)(gpGlobals->pStringBase + (unsigned int)(offset));
}
*/
import "C"
import "unsafe"

const (
	MaxEntLeafs = 48
)

type Link struct {
	Next *Link
	Prev *Link
}

func (l *Link) ToCLinkT() *C.link_t {
	var link *C.link_t

	if l.Next != nil {
		link.next = l.Next.ToCLinkT()
	}

	if l.Prev != nil {
		link.prev = l.Prev.ToCLinkT()
	}

	return link
}

type Edict struct {
	Free         bool
	SerialNumber int
	Area         *Link

	HeadNode int
	NumLeafs int
	LeafNums [MaxEntLeafs]int16

	FreeTime float64

	PvPrivateData unsafe.Pointer

	V EntVars
}

func (e *Edict) ToCEdictT() *C.edict_t {
	free := 0
	if e.Free {
		free = 1
	}

	leafNums := [MaxEntLeafs]C.short{}
	for i, leafNum := range e.LeafNums {
		leafNums[i] = C.short(leafNum)
	}

	return &C.edict_t{
		free:         C.qboolean(C.int(free)),
		serialnumber: C.int(e.SerialNumber),
		area:         *e.Area.ToCLinkT(),

		headnode:  C.int(e.HeadNode),
		num_leafs: C.int(e.NumLeafs),
		leafnums:  leafNums,

		freetime:      C.float(e.FreeTime),
		pvPrivateData: e.PvPrivateData,
		//v:             *e.V.ToCEntVarsT(),
	}
}

type GlobalVars struct {
	p *C.globalvars_t
}

func GlobalVarsFromC(g *C.globalvars_t) *GlobalVars {
	gv := &GlobalVars{
		p: g,
	}

	return gv
}

func (gv *GlobalVars) Time() float32 {
	return float32(gv.p.time)
}

func (gv *GlobalVars) FrameTime() float32 {
	return float32(gv.p.frametime)
}

func (gv *GlobalVars) ForceRetouch() float32 {
	return float32(gv.p.force_retouch)
}

func (gv *GlobalVars) MapName() string {
	return C.GoString(C.ReadString(gv.p, gv.p.mapname))
}

func (gv *GlobalVars) StartSpot() string {
	return C.GoString(C.ReadString(gv.p, gv.p.startspot))
}

func (gv *GlobalVars) Deathmatch() bool {
	return gv.p.deathmatch == 1
}

func (gv *GlobalVars) Coop() bool {
	return gv.p.coop == 1
}

func (gv *GlobalVars) Teamplay() float32 {
	return float32(gv.p.teamplay)
}

func (gv *GlobalVars) MaxClients() int {
	return int(gv.p.maxClients)
}

func (gv *GlobalVars) MaxEntities() int {
	return int(gv.p.maxEntities)
}

type EntVars struct {
	p *C.entvars_t

	//Classname  int
	//GlobalName int
	//
	//Origin       [3]float32
	//OldOrigin    [3]float32
	//Velocity     [3]float32
	//BaseVelocity [3]float32
	//
	//// Base velocity that was passed in to server physics so
	//// client can predict conveyors correctly.
	//// Server zeroes it, so we need to store here, too.
	//ClBaseVelocity [3]float32
	//
	//MoveDir [3]float32
	//
	//Angles     [3]float32 // Model angles
	//Avelocity  [3]float32 // angle velocity (degrees per second)
	//PunchAngle [3]float32 // auto-decaying view angle adjustment
	//VAngle     [3]float32 // Viewing angle (player only)
	//
	//// For parametric entities
	//EndPos     [3]float32
	//StartPos   [3]float32
	//ImpactTime float32
	//StartTime  float32
	//
	//FixAngle   int // 0:nothing, 1:force view angles, 2:add avelocity
	//IdealPitch float32
	//PitchSpeed float32
	//IdealYaw   float32
	//YawSpeed   float32
	//
	//ModelIndex int
	//Model      string
	//
	//ViewModel   int // player's viewmodel
	//WeaponModel int // what other players see
	//
	//AbsMin [3]float32 // BB min translated to world coord
	//AbsMax [3]float32 // BB max translated to world coord
	//Mins   [3]float32 // local BB min
	//Maxs   [3]float32 // local BB max
	//Size   [3]float32 // maxs - mins
	//
	//Ltime     float32
	//NextThink float32
	//
	//MoveType int
	//Solid    int
	//
	//Skin    int
	//Body    int // sub-model selection for studiomodels
	//Effects int
	//
	//Gravity  float32 // % of "normal" gravity
	//Friction float32 // inverse elasticity of MOVETYPE_BOUNCE
	//
	//LightLevel int
	//
	//Sequence   int     // animation sequence
	//GaitSeq    int     // movement animation sequence for player (0 for none)
	//Frame      float32 // % playback position in animation sequences (0..255)
	//AnymTime   float32 // world time when frame was set
	//FrameRate  float32 // animation playback rate (-8x to 8x)
	//Controller [4]byte // bone controller setting
	//Blending   [2]byte // blending amount between sub-sequences
	//
	//Scale float32 // sprite rendering scale (0..255)
	//
	//RenderMode  int
	//RenderAmt   float32
	//RenderColor [3]float32
	//RenderFx    int
	//
	//Health     float32
	//Frags      float32
	//Weapons    int // bit mask for available weapons
	//TakeDamage int
	//
	//DeadFlag int
	//ViewOfs  [3]float32 // eye position
	//
	//Button  int
	//Impulse int
	//
	//Chain        *Edict
	//DmgInflictor *Edict
	//Enemy        *Edict
	//AimEnt       *Edict
	//Owner        *Edict
	//GroundEntity *Edict
	//
	//SpawnFlags int
	//Flags      int
	//
	//Colormap int
	//Team     int
	//
	//MaxHealth    float32
	//TeleportTime float32
	//Armortype    float32
	//Armorvalue   float32
	//WaterLevel   int
	//WaterType    int
	//
	//Target     string
	//TargetName string
	//NetName    string
	//Message    string
	//
	//DmgTake float32
	//DmgSave float32
	//Dmg     float32
	//DmgTime float32
	//
	//Noise  string
	//Noise1 string
	//Noise2 string
	//Noise3 string
	//
	//Speed           float32
	//AirFinished     float32
	//PainFinished    float32
	//RadsuitFinished float32
	//
	//PContainingEntity *Edict
	//
	//PlayerClass int
	//MaxSpeed    float32
	//
	//Fov        float32
	//WeaponAnim int
	//
	//PushmSec int
	//
	//BInDuck       int
	//TimeStepSound int
	//SwimTime      int
	//DuckTime      int
	//StepLeft      int
	//FallVelocity  int
	//
	//GameState int
	//
	//OldButtons int
	//
	//GroupInfo int
	//
	//// For mods
	//IUser1 int
	//IUser2 int
	//IUser3 int
	//IUser4 int
	//FUser1 float32
	//FUser2 float32
	//FUser3 float32
	//FUser4 float32
	//VUser1 [3]float32
	//VUser2 [3]float32
	//VUser3 [3]float32
	//VUser4 [3]float32
	//EUser1 *Edict
	//EUser2 *Edict
	//EUser3 *Edict
	//EUser4 *Edict
}

func EntVarsFromC(ev *C.entvars_t) *EntVars {
	e := &EntVars{
		p: ev,
	}

	return e
}

func (e *EntVars) ClassName() string {
	return C.GoString(C.ReadString(e.p, e.p.classname))
}

func (e *EntVars) GlobalName() string {
	return C.GoString(C.ReadString(e.p, e.p.globalname))
}
