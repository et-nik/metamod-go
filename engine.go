package main

/*
const char* ReadString(globalvars_t *gpGlobals, int offset) {
	return (const char *)(gpGlobals->pStringBase + (unsigned int)(offset));
}
*/
import "C"

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
	p          *C.edict_t
	globalVars *C.globalvars_t

	//Free         bool
	//SerialNumber int
	//Area         *Link
	//
	//HeadNode int
	//NumLeafs int
	//LeafNums [MaxEntLeafs]int16
	//
	//FreeTime float64
	//
	//PvPrivateData unsafe.Pointer
	//
	//V EntVars
}

func EdictFromC(globalVars *C.globalvars_t, e *C.edict_t) *Edict {
	if e == nil {
		return nil
	}

	edict := &Edict{
		p:          e,
		globalVars: globalVars,
	}

	return edict
}

func (e *Edict) Free() int {
	return int(e.p.free)
}

func (e *Edict) SerialNumber() int {
	return int(e.p.serialnumber)
}

func (e *Edict) EntVars() *EntVars {
	return EntVarsFromC(e.globalVars, &e.p.v)
}

//func (e *Edict) ToCEdictT() *C.edict_t {
//	free := 0
//	if e.Free {
//		free = 1
//	}
//
//	leafNums := [MaxEntLeafs]C.short{}
//	for i, leafNum := range e.LeafNums {
//		leafNums[i] = C.short(leafNum)
//	}
//
//	return &C.edict_t{
//		free:         C.qboolean(C.int(free)),
//		serialnumber: C.int(e.SerialNumber),
//		area:         *e.Area.ToCLinkT(),
//
//		headnode:  C.int(e.HeadNode),
//		num_leafs: C.int(e.NumLeafs),
//		leafnums:  leafNums,
//
//		freetime:      C.float(e.FreeTime),
//		pvPrivateData: e.PvPrivateData,
//		//v:             *e.V.ToCEntVarsT(),
//	}
//}

type GlobalVars struct {
	p *C.globalvars_t
}

func GlobalVarsFromC(g *C.globalvars_t) *GlobalVars {
	if g == nil {
		return nil
	}

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
	p          *C.entvars_t
	globalVars *C.globalvars_t

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

func EntVarsFromC(globalVars *C.globalvars_t, ev *C.entvars_t) *EntVars {
	if ev == nil {
		return nil
	}

	e := &EntVars{
		p:          ev,
		globalVars: globalVars,
	}

	return e
}

func (e *EntVars) ClassName() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.classname))
}

func (e *EntVars) GlobalName() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.globalname))
}

func (e *EntVars) Origin() [3]float32 {
	return [3]float32{
		float32(e.p.origin[0]),
		float32(e.p.origin[1]),
		float32(e.p.origin[2]),
	}
}

func (e *EntVars) OldOrigin() [3]float32 {
	return [3]float32{
		float32(e.p.oldorigin[0]),
		float32(e.p.oldorigin[1]),
		float32(e.p.oldorigin[2]),
	}
}

func (e *EntVars) Velocity() [3]float32 {
	return [3]float32{
		float32(e.p.velocity[0]),
		float32(e.p.velocity[1]),
		float32(e.p.velocity[2]),
	}
}

func (e *EntVars) BaseVelocity() [3]float32 {
	return [3]float32{
		float32(e.p.basevelocity[0]),
		float32(e.p.basevelocity[1]),
		float32(e.p.basevelocity[2]),
	}
}

// ClBaseVelocity Base velocity that was passed in to server physics so
// client can predict conveyors correctly.
// Server zeroes it, so we need to store here, too.
func (e *EntVars) ClBaseVelocity() [3]float32 {
	return [3]float32{
		float32(e.p.clbasevelocity[0]),
		float32(e.p.clbasevelocity[1]),
		float32(e.p.clbasevelocity[2]),
	}
}

func (e *EntVars) MoveDir() [3]float32 {
	return [3]float32{
		float32(e.p.movedir[0]),
		float32(e.p.movedir[1]),
		float32(e.p.movedir[2]),
	}
}

// Angles Model angles
func (e *EntVars) Angles() [3]float32 {
	return [3]float32{
		float32(e.p.angles[0]),
		float32(e.p.angles[1]),
		float32(e.p.angles[2]),
	}
}

func (e *EntVars) Avelocity() [3]float32 {
	return [3]float32{
		float32(e.p.avelocity[0]),
		float32(e.p.avelocity[1]),
		float32(e.p.avelocity[2]),
	}
}

func (e *EntVars) PunchAngle() [3]float32 {
	return [3]float32{
		float32(e.p.punchangle[0]),
		float32(e.p.punchangle[1]),
		float32(e.p.punchangle[2]),
	}
}

func (e *EntVars) VAngle() [3]float32 {
	return [3]float32{
		float32(e.p.v_angle[0]),
		float32(e.p.v_angle[1]),
		float32(e.p.v_angle[2]),
	}
}

func (e *EntVars) EndPos() [3]float32 {
	return [3]float32{
		float32(e.p.endpos[0]),
		float32(e.p.endpos[1]),
		float32(e.p.endpos[2]),
	}
}

func (e *EntVars) StartPos() [3]float32 {
	return [3]float32{
		float32(e.p.startpos[0]),
		float32(e.p.startpos[1]),
		float32(e.p.startpos[2]),
	}
}

func (e *EntVars) ImpactTime() float32 {
	return float32(e.p.impacttime)
}

func (e *EntVars) StartTime() float32 {
	return float32(e.p.starttime)
}

func (e *EntVars) FixAngle() int {
	return int(e.p.fixangle)
}

func (e *EntVars) IdealPitch() float32 {
	return float32(e.p.idealpitch)
}

func (e *EntVars) PitchSpeed() float32 {
	return float32(e.p.pitch_speed)
}

func (e *EntVars) IdealYaw() float32 {
	return float32(e.p.ideal_yaw)
}

func (e *EntVars) YawSpeed() float32 {
	return float32(e.p.yaw_speed)
}

func (e *EntVars) ModelIndex() int {
	return int(e.p.modelindex)
}

func (e *EntVars) Model() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.model))
}

func (e *EntVars) ViewModel() int {
	return int(e.p.viewmodel)
}

func (e *EntVars) WeaponModel() int {
	return int(e.p.weaponmodel)
}

func (e *EntVars) AbsMin() [3]float32 {
	return [3]float32{
		float32(e.p.absmin[0]),
		float32(e.p.absmin[1]),
		float32(e.p.absmin[2]),
	}
}

func (e *EntVars) AbsMax() [3]float32 {
	return [3]float32{
		float32(e.p.absmax[0]),
		float32(e.p.absmax[1]),
		float32(e.p.absmax[2]),
	}
}

func (e *EntVars) Mins() [3]float32 {
	return [3]float32{
		float32(e.p.mins[0]),
		float32(e.p.mins[1]),
		float32(e.p.mins[2]),
	}
}

func (e *EntVars) Maxs() [3]float32 {
	return [3]float32{
		float32(e.p.maxs[0]),
		float32(e.p.maxs[1]),
		float32(e.p.maxs[2]),
	}
}

func (e *EntVars) Size() [3]float32 {
	return [3]float32{
		float32(e.p.size[0]),
		float32(e.p.size[1]),
		float32(e.p.size[2]),
	}
}

func (e *EntVars) Ltime() float32 {
	return float32(e.p.ltime)
}

func (e *EntVars) NextThink() float32 {
	return float32(e.p.nextthink)
}

func (e *EntVars) MoveType() int {
	return int(e.p.movetype)
}

func (e *EntVars) Solid() int {
	return int(e.p.solid)
}

func (e *EntVars) Skin() int {
	return int(e.p.skin)
}

func (e *EntVars) Body() int {
	return int(e.p.body)
}

func (e *EntVars) Effects() int {
	return int(e.p.effects)
}

func (e *EntVars) Gravity() float32 {
	return float32(e.p.gravity)
}

func (e *EntVars) Friction() float32 {
	return float32(e.p.friction)
}

func (e *EntVars) LightLevel() int {
	return int(e.p.light_level)
}

// Sequence animation sequence
func (e *EntVars) Sequence() int {
	return int(e.p.sequence)
}

// GaitSequence movement animation sequence for player (0 for none)
func (e *EntVars) GaitSequence() int {
	return int(e.p.gaitsequence)
}

// Frame % playback position in animation sequences (0..255)
func (e *EntVars) Frame() float32 {
	return float32(e.p.frame)
}

// AnymTime world time when frame was set
func (e *EntVars) AnymTime() float32 {
	return float32(e.p.animtime)
}

// FrameRate animation playback rate (-8x to 8x)
func (e *EntVars) FrameRate() float32 {
	return float32(e.p.framerate)
}

// Controller bone controller setting
func (e *EntVars) Controller() [4]byte {
	return [4]byte{
		byte(e.p.controller[0]),
		byte(e.p.controller[1]),
		byte(e.p.controller[2]),
		byte(e.p.controller[3]),
	}
}

// Blending blending amount between sub-sequences
func (e *EntVars) Blending() [2]byte {
	return [2]byte{
		byte(e.p.blending[0]),
		byte(e.p.blending[1]),
	}
}

func (e *EntVars) Scale() float32 {
	return float32(e.p.scale)
}

func (e *EntVars) RenderMode() int {
	return int(e.p.rendermode)
}

func (e *EntVars) RenderAmt() float32 {
	return float32(e.p.renderamt)
}

func (e *EntVars) RenderColor() [3]float32 {
	return [3]float32{
		float32(e.p.rendercolor[0]),
		float32(e.p.rendercolor[1]),
		float32(e.p.rendercolor[2]),
	}
}

func (e *EntVars) RenderFx() int {
	return int(e.p.renderfx)
}

func (e *EntVars) Health() float32 {
	return float32(e.p.health)
}

func (e *EntVars) MaxHealth() float32 {
	return float32(e.p.max_health)
}

func (e *EntVars) Frags() float32 {
	return float32(e.p.frags)
}

func (e *EntVars) Weapons() int {
	return int(e.p.weapons)
}

func (e *EntVars) TakeDamage() int {
	return int(e.p.takedamage)
}

func (e *EntVars) DeadFlag() int {
	return int(e.p.deadflag)
}

func (e *EntVars) ViewOfs() [3]float32 {
	return [3]float32{
		float32(e.p.view_ofs[0]),
		float32(e.p.view_ofs[1]),
		float32(e.p.view_ofs[2]),
	}
}

func (e *EntVars) Button() int {
	return int(e.p.button)
}

func (e *EntVars) Impulse() int {
	return int(e.p.impulse)
}

func (e *EntVars) Chain() *Edict {
	return EdictFromC(e.globalVars, e.p.chain)
}

func (e *EntVars) DmgInflictor() *Edict {
	return EdictFromC(e.globalVars, e.p.dmg_inflictor)
}

func (e *EntVars) Enemy() *Edict {
	return EdictFromC(e.globalVars, e.p.enemy)
}

func (e *EntVars) AimEnt() *Edict {
	return EdictFromC(e.globalVars, e.p.aiment)
}

func (e *EntVars) Owner() *Edict {
	return EdictFromC(e.globalVars, e.p.owner)
}

func (e *EntVars) GroundEntity() *Edict {
	return EdictFromC(e.globalVars, e.p.groundentity)
}

func (e *EntVars) SpawnFlags() int {
	return int(e.p.spawnflags)
}

func (e *EntVars) Flags() int {
	return int(e.p.flags)
}

func (e *EntVars) Colormap() int {
	return int(e.p.colormap)
}

func (e *EntVars) Team() int {
	return int(e.p.team)
}

func (e *EntVars) TeleportTime() float32 {
	return float32(e.p.teleport_time)
}

func (e *EntVars) Armortype() float32 {
	return float32(e.p.armortype)
}

func (e *EntVars) Armorvalue() float32 {
	return float32(e.p.armorvalue)
}

func (e *EntVars) WaterLevel() int {
	return int(e.p.waterlevel)
}

func (e *EntVars) WaterType() int {
	return int(e.p.watertype)
}

func (e *EntVars) Target() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.target))
}

func (e *EntVars) TargetName() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.targetname))
}

func (e *EntVars) NetName() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.netname))
}

func (e *EntVars) Message() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.message))
}

func (e *EntVars) DmgTake() float32 {
	return float32(e.p.dmg_take)
}

func (e *EntVars) DmgSave() float32 {
	return float32(e.p.dmg_save)
}

func (e *EntVars) Dmg() float32 {
	return float32(e.p.dmg)
}

func (e *EntVars) DmgTime() float32 {
	return float32(e.p.dmgtime)
}

func (e *EntVars) Noise() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.noise))
}

func (e *EntVars) Noise1() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.noise1))
}

func (e *EntVars) Noise2() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.noise2))
}

func (e *EntVars) Noise3() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.noise3))
}

func (e *EntVars) Speed() float32 {
	return float32(e.p.speed)
}

func (e *EntVars) AirFinished() float32 {
	return float32(e.p.air_finished)
}

func (e *EntVars) PainFinished() float32 {
	return float32(e.p.pain_finished)
}

func (e *EntVars) RadsuitFinished() float32 {
	return float32(e.p.radsuit_finished)
}

func (e *EntVars) PContainingEntity() *Edict {
	return EdictFromC(e.globalVars, e.p.pContainingEntity)
}

func (e *EntVars) PlayerClass() int {
	return int(e.p.playerclass)
}

func (e *EntVars) MaxSpeed() float32 {
	return float32(e.p.maxspeed)
}

func (e *EntVars) Fov() float32 {
	return float32(e.p.fov)
}

func (e *EntVars) WeaponAnim() int {
	return int(e.p.weaponanim)
}

func (e *EntVars) PushmSec() int {
	return int(e.p.pushmsec)
}

func (e *EntVars) BInDuck() int {
	return int(e.p.bInDuck)
}

func (e *EntVars) TimeStepSound() int {
	return int(e.p.flTimeStepSound)
}

func (e *EntVars) SwimTime() int {
	return int(e.p.flSwimTime)
}

func (e *EntVars) DuckTime() int {
	return int(e.p.flDuckTime)
}

func (e *EntVars) StepLeft() int {
	return int(e.p.iStepLeft)
}

func (e *EntVars) FallVelocity() int {
	return int(e.p.flFallVelocity)
}

func (e *EntVars) GameState() int {
	return int(e.p.gamestate)
}

func (e *EntVars) OldButtons() int {
	return int(e.p.oldbuttons)
}

func (e *EntVars) GroupInfo() int {
	return int(e.p.groupinfo)
}

func (e *EntVars) IUser1() int {
	return int(e.p.iuser1)
}

func (e *EntVars) IUser2() int {
	return int(e.p.iuser2)
}

func (e *EntVars) IUser3() int {
	return int(e.p.iuser3)
}

func (e *EntVars) IUser4() int {
	return int(e.p.iuser4)
}

func (e *EntVars) FUser1() float32 {
	return float32(e.p.fuser1)
}

func (e *EntVars) FUser2() float32 {
	return float32(e.p.fuser2)
}

func (e *EntVars) FUser3() float32 {
	return float32(e.p.fuser3)
}

func (e *EntVars) FUser4() float32 {
	return float32(e.p.fuser4)
}

func (e *EntVars) VUser1() [3]float32 {
	return [3]float32{
		float32(e.p.vuser1[0]),
		float32(e.p.vuser1[1]),
		float32(e.p.vuser1[2]),
	}
}

func (e *EntVars) VUser2() [3]float32 {
	return [3]float32{
		float32(e.p.vuser2[0]),
		float32(e.p.vuser2[1]),
		float32(e.p.vuser2[2]),
	}
}

func (e *EntVars) VUser3() [3]float32 {
	return [3]float32{
		float32(e.p.vuser3[0]),
		float32(e.p.vuser3[1]),
		float32(e.p.vuser3[2]),
	}
}

func (e *EntVars) VUser4() [3]float32 {
	return [3]float32{
		float32(e.p.vuser4[0]),
		float32(e.p.vuser4[1]),
		float32(e.p.vuser4[2]),
	}
}

func (e *EntVars) EUser1() *Edict {
	return EdictFromC(e.globalVars, e.p.euser1)
}

func (e *EntVars) EUser2() *Edict {
	return EdictFromC(e.globalVars, e.p.euser2)
}

func (e *EntVars) EUser3() *Edict {
	return EdictFromC(e.globalVars, e.p.euser3)
}

func (e *EntVars) EUser4() *Edict {
	return EdictFromC(e.globalVars, e.p.euser4)
}
