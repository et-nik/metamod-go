package main

/*
#include <eiface.h>
#include <com_model.h>

const char* ReadString(globalvars_t *gpGlobals, int offset) {
	return (const char *)(gpGlobals->pStringBase + (unsigned int)(offset));
}

int MakeString(globalvars_t *gpGlobals, char *str) {
	return ((uint64)(str) - (uint64)(ReadString(gpGlobals, 0)));
}

*/
import "C"

const (
	MaxEntLeafs = 48
)

type AlertType int

const (
	AlertTypeNotice    AlertType = iota
	AlertTypeConsole             // same as at_notice, but forces a ConPrintf, not a message box
	AlertTypeAIConsole           // same as at_console, but only shown if developer level is 2!
	AlertTypeWarning
	AlertTypeError
	AlertTypeLogged // Server print to console ( only in multiplayer games ).
)

type PrintType int

const (
	PrintTypeConsole PrintType = iota
	PrintTypeCenter
	PrintTypeChat
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
}

func edictFromC(globalVars *C.globalvars_t, e *C.edict_t) *Edict {
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
	return entVarsFromC(e.globalVars, &e.p.v)
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
}

func entVarsFromC(globalVars *C.globalvars_t, ev *C.entvars_t) *EntVars {
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
	if len(e.p.origin) == 0 {
		return [3]float32{}
	}

	return [3]float32{
		float32(e.p.origin[0]),
		float32(e.p.origin[1]),
		float32(e.p.origin[2]),
	}
}

func (e *EntVars) SetOrigin(origin [3]float32) {
	e.p.origin[0] = C.float(origin[0])
	e.p.origin[1] = C.float(origin[1])
	e.p.origin[2] = C.float(origin[2])
}

func (e *EntVars) OldOrigin() [3]float32 {
	return [3]float32{
		float32(e.p.oldorigin[0]),
		float32(e.p.oldorigin[1]),
		float32(e.p.oldorigin[2]),
	}
}

func (e *EntVars) SetOldOrigin(origin [3]float32) {
	e.p.oldorigin[0] = C.float(origin[0])
	e.p.oldorigin[1] = C.float(origin[1])
	e.p.oldorigin[2] = C.float(origin[2])
}

func (e *EntVars) Velocity() [3]float32 {
	return [3]float32{
		float32(e.p.velocity[0]),
		float32(e.p.velocity[1]),
		float32(e.p.velocity[2]),
	}
}

func (e *EntVars) SetVelocity(velocity [3]float32) {
	e.p.velocity[0] = C.float(velocity[0])
	e.p.velocity[1] = C.float(velocity[1])
	e.p.velocity[2] = C.float(velocity[2])
}

func (e *EntVars) BaseVelocity() [3]float32 {
	return [3]float32{
		float32(e.p.basevelocity[0]),
		float32(e.p.basevelocity[1]),
		float32(e.p.basevelocity[2]),
	}
}

func (e *EntVars) SetBaseVelocity(velocity [3]float32) {
	e.p.basevelocity[0] = C.float(velocity[0])
	e.p.basevelocity[1] = C.float(velocity[1])
	e.p.basevelocity[2] = C.float(velocity[2])
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

func (e *EntVars) SetMoveDir(dir [3]float32) {
	e.p.movedir[0] = C.float(dir[0])
	e.p.movedir[1] = C.float(dir[1])
	e.p.movedir[2] = C.float(dir[2])
}

// Angles Model angles
func (e *EntVars) Angles() [3]float32 {
	return [3]float32{
		float32(e.p.angles[0]),
		float32(e.p.angles[1]),
		float32(e.p.angles[2]),
	}
}

func (e *EntVars) SetAngles(angles [3]float32) {
	e.p.angles[0] = C.float(angles[0])
	e.p.angles[1] = C.float(angles[1])
	e.p.angles[2] = C.float(angles[2])
}

func (e *EntVars) Avelocity() [3]float32 {
	return [3]float32{
		float32(e.p.avelocity[0]),
		float32(e.p.avelocity[1]),
		float32(e.p.avelocity[2]),
	}
}

func (e *EntVars) SetAvelocity(avelocity [3]float32) {
	e.p.avelocity[0] = C.float(avelocity[0])
	e.p.avelocity[1] = C.float(avelocity[1])
	e.p.avelocity[2] = C.float(avelocity[2])
}

func (e *EntVars) PunchAngle() [3]float32 {
	return [3]float32{
		float32(e.p.punchangle[0]),
		float32(e.p.punchangle[1]),
		float32(e.p.punchangle[2]),
	}
}

func (e *EntVars) SetPunchAngle(punchangle [3]float32) {
	e.p.punchangle[0] = C.float(punchangle[0])
	e.p.punchangle[1] = C.float(punchangle[1])
	e.p.punchangle[2] = C.float(punchangle[2])
}

func (e *EntVars) VAngle() [3]float32 {
	return [3]float32{
		float32(e.p.v_angle[0]),
		float32(e.p.v_angle[1]),
		float32(e.p.v_angle[2]),
	}
}

func (e *EntVars) SetVAngle(vAngle [3]float32) {
	e.p.v_angle[0] = C.float(vAngle[0])
	e.p.v_angle[1] = C.float(vAngle[1])
	e.p.v_angle[2] = C.float(vAngle[2])
}

func (e *EntVars) EndPos() [3]float32 {
	return [3]float32{
		float32(e.p.endpos[0]),
		float32(e.p.endpos[1]),
		float32(e.p.endpos[2]),
	}
}

func (e *EntVars) SetEndPos(endPos [3]float32) {
	e.p.endpos[0] = C.float(endPos[0])
	e.p.endpos[1] = C.float(endPos[1])
	e.p.endpos[2] = C.float(endPos[2])
}

func (e *EntVars) StartPos() [3]float32 {
	return [3]float32{
		float32(e.p.startpos[0]),
		float32(e.p.startpos[1]),
		float32(e.p.startpos[2]),
	}
}

func (e *EntVars) SetStartPost(startPos [3]float32) {
	e.p.startpos[0] = C.float(startPos[0])
	e.p.startpos[1] = C.float(startPos[1])
	e.p.startpos[2] = C.float(startPos[2])
}

func (e *EntVars) ImpactTime() float32 {
	return float32(e.p.impacttime)
}

func (e *EntVars) SetImpactTime(impactTime float32) {
	e.p.impacttime = C.float(impactTime)
}

func (e *EntVars) StartTime() float32 {
	return float32(e.p.starttime)
}

func (e *EntVars) SetStartTime(startTime float32) {
	e.p.starttime = C.float(startTime)
}

func (e *EntVars) FixAngle() int {
	return int(e.p.fixangle)
}

func (e *EntVars) SetFixAngle(fixAngle int) {
	e.p.fixangle = C.int(fixAngle)
}

func (e *EntVars) IdealPitch() float32 {
	return float32(e.p.idealpitch)
}

func (e *EntVars) SetIdealPitch(idealPitch float32) {
	e.p.idealpitch = C.float(idealPitch)
}

func (e *EntVars) PitchSpeed() float32 {
	return float32(e.p.pitch_speed)
}

func (e *EntVars) SetPitchSpeed(pitchSpeed float32) {
	e.p.pitch_speed = C.float(pitchSpeed)
}

func (e *EntVars) IdealYaw() float32 {
	return float32(e.p.ideal_yaw)
}

func (e *EntVars) SetIdealYaw(idealYaw float32) {
	e.p.ideal_yaw = C.float(idealYaw)
}

func (e *EntVars) YawSpeed() float32 {
	return float32(e.p.yaw_speed)
}

func (e *EntVars) SetYawSpeed(yawSpeed float32) {
	e.p.yaw_speed = C.float(yawSpeed)
}

func (e *EntVars) ModelIndex() int {
	return int(e.p.modelindex)
}

func (e *EntVars) SetModelIndex(modelIndex int) {
	e.p.modelindex = C.int(modelIndex)
}

func (e *EntVars) Model() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.model))
}

func (e *EntVars) SetModel(model string) {
	e.p.model = C.int(allocString(model))
}

func (e *EntVars) ViewModel() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.viewmodel))
}

func (e *EntVars) SetViewModel(viewModel string) {
	e.p.viewmodel = C.int(allocString(viewModel))
}

func (e *EntVars) WeaponModel() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.weaponmodel))
}

func (e *EntVars) SetWeaponModel(weaponModel string) {
	e.p.weaponmodel = C.int(allocString(weaponModel))
}

func (e *EntVars) AbsMin() [3]float32 {
	return [3]float32{
		float32(e.p.absmin[0]),
		float32(e.p.absmin[1]),
		float32(e.p.absmin[2]),
	}
}

func (e *EntVars) SetAbsMin(absMin [3]float32) {
	e.p.absmin[0] = C.float(absMin[0])
	e.p.absmin[1] = C.float(absMin[1])
	e.p.absmin[2] = C.float(absMin[2])
}

func (e *EntVars) AbsMax() [3]float32 {
	return [3]float32{
		float32(e.p.absmax[0]),
		float32(e.p.absmax[1]),
		float32(e.p.absmax[2]),
	}
}

func (e *EntVars) SetAbsMax(absMax [3]float32) {
	e.p.absmax[0] = C.float(absMax[0])
	e.p.absmax[1] = C.float(absMax[1])
	e.p.absmax[2] = C.float(absMax[2])
}

func (e *EntVars) Mins() [3]float32 {
	return [3]float32{
		float32(e.p.mins[0]),
		float32(e.p.mins[1]),
		float32(e.p.mins[2]),
	}
}

func (e *EntVars) SetMins(mins [3]float32) {
	e.p.mins[0] = C.float(mins[0])
	e.p.mins[1] = C.float(mins[1])
	e.p.mins[2] = C.float(mins[2])
}

func (e *EntVars) Maxs() [3]float32 {
	return [3]float32{
		float32(e.p.maxs[0]),
		float32(e.p.maxs[1]),
		float32(e.p.maxs[2]),
	}
}

func (e *EntVars) SetMaxs(maxs [3]float32) {
	e.p.maxs[0] = C.float(maxs[0])
	e.p.maxs[1] = C.float(maxs[1])
	e.p.maxs[2] = C.float(maxs[2])
}

func (e *EntVars) Size() [3]float32 {
	return [3]float32{
		float32(e.p.size[0]),
		float32(e.p.size[1]),
		float32(e.p.size[2]),
	}
}

func (e *EntVars) SetSize(size [3]float32) {
	e.p.size[0] = C.float(size[0])
	e.p.size[1] = C.float(size[1])
	e.p.size[2] = C.float(size[2])
}

func (e *EntVars) Ltime() float32 {
	return float32(e.p.ltime)
}

func (e *EntVars) SetLtime(ltime float32) {
	e.p.ltime = C.float(ltime)
}

func (e *EntVars) NextThink() float32 {
	return float32(e.p.nextthink)
}

func (e *EntVars) SetNextThink(nextThink float32) {
	e.p.nextthink = C.float(nextThink)
}

func (e *EntVars) MoveType() int {
	return int(e.p.movetype)
}

func (e *EntVars) SetMoveType(moveType int) {
	e.p.movetype = C.int(moveType)
}

func (e *EntVars) Solid() int {
	return int(e.p.solid)
}

func (e *EntVars) SetSolid(solid int) {
	e.p.solid = C.int(solid)
}

func (e *EntVars) Skin() int {
	return int(e.p.skin)
}

func (e *EntVars) SetSkin(skin int) {
	e.p.skin = C.int(skin)
}

func (e *EntVars) Body() int {
	return int(e.p.body)
}

func (e *EntVars) SetBody(body int) {
	e.p.body = C.int(body)
}

func (e *EntVars) Effects() int {
	return int(e.p.effects)
}

func (e *EntVars) SetEffects(effects int) {
	e.p.effects = C.int(effects)
}

func (e *EntVars) Gravity() float32 {
	return float32(e.p.gravity)
}

func (e *EntVars) SetGravity(gravity float32) {
	e.p.gravity = C.float(gravity)
}

func (e *EntVars) Friction() float32 {
	return float32(e.p.friction)
}

func (e *EntVars) SetFriction(friction float32) {
	e.p.friction = C.float(friction)
}

func (e *EntVars) LightLevel() int {
	return int(e.p.light_level)
}

func (e *EntVars) SetLightLevel(lightLevel int) {
	e.p.light_level = C.int(lightLevel)
}

// Sequence animation sequence
func (e *EntVars) Sequence() int {
	return int(e.p.sequence)
}

func (e *EntVars) SetSequence(sequence int) {
	e.p.sequence = C.int(sequence)
}

// GaitSequence movement animation sequence for player (0 for none)
func (e *EntVars) GaitSequence() int {
	return int(e.p.gaitsequence)
}

func (e *EntVars) SetGaitSequence(gaitSequence int) {
	e.p.gaitsequence = C.int(gaitSequence)
}

// Frame % playback position in animation sequences (0..255)
func (e *EntVars) Frame() float32 {
	return float32(e.p.frame)
}

func (e *EntVars) SetFrame(frame float32) {
	e.p.frame = C.float(frame)
}

// AnymTime world time when frame was set
func (e *EntVars) AnymTime() float32 {
	return float32(e.p.animtime)
}

func (e *EntVars) SetAnimTime(animTime float32) {
	e.p.animtime = C.float(animTime)
}

// FrameRate animation playback rate (-8x to 8x)
func (e *EntVars) FrameRate() float32 {
	return float32(e.p.framerate)
}

func (e *EntVars) SetFrameRate(frameRate float32) {
	e.p.framerate = C.float(frameRate)
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

func (e *EntVars) SetController(controller [4]byte) {
	e.p.controller[0] = C.byte(controller[0])
	e.p.controller[1] = C.byte(controller[1])
	e.p.controller[2] = C.byte(controller[2])
	e.p.controller[3] = C.byte(controller[3])
}

// Blending blending amount between sub-sequences
func (e *EntVars) Blending() [2]byte {
	return [2]byte{
		byte(e.p.blending[0]),
		byte(e.p.blending[1]),
	}
}

func (e *EntVars) SetBlending(blending [2]byte) {
	e.p.blending[0] = C.byte(blending[0])
	e.p.blending[1] = C.byte(blending[1])
}

func (e *EntVars) Scale() float32 {
	return float32(e.p.scale)
}

func (e *EntVars) SetScale(scale float32) {
	e.p.scale = C.float(scale)
}

func (e *EntVars) RenderMode() int {
	return int(e.p.rendermode)
}

func (e *EntVars) SetRenderMode(renderMode int) {
	e.p.rendermode = C.int(renderMode)
}

func (e *EntVars) RenderAmt() float32 {
	return float32(e.p.renderamt)
}

func (e *EntVars) SetRenderAmt(renderAmt float32) {
	e.p.renderamt = C.float(renderAmt)
}

func (e *EntVars) RenderColor() [3]float32 {
	return [3]float32{
		float32(e.p.rendercolor[0]),
		float32(e.p.rendercolor[1]),
		float32(e.p.rendercolor[2]),
	}
}

func (e *EntVars) SetRenderColor(renderColor [3]float32) {
	e.p.rendercolor[0] = C.float(renderColor[0])
	e.p.rendercolor[1] = C.float(renderColor[1])
	e.p.rendercolor[2] = C.float(renderColor[2])
}

func (e *EntVars) RenderFx() int {
	return int(e.p.renderfx)
}

func (e *EntVars) SetRenderFx(renderFx int) {
	e.p.renderfx = C.int(renderFx)
}

func (e *EntVars) Health() float32 {
	return float32(e.p.health)
}

func (e *EntVars) SetHealth(health float32) {
	e.p.health = C.float(health)
}

func (e *EntVars) MaxHealth() float32 {
	return float32(e.p.max_health)
}

func (e *EntVars) SetMaxHealth(maxHealth float32) {
	e.p.max_health = C.float(maxHealth)
}

func (e *EntVars) Frags() float32 {
	return float32(e.p.frags)
}

func (e *EntVars) SetFrags(frags float32) {
	e.p.frags = C.float(frags)
}

func (e *EntVars) Weapons() int {
	return int(e.p.weapons)
}

func (e *EntVars) SetWeapons(weapons int) {
	e.p.weapons = C.int(weapons)
}

func (e *EntVars) TakeDamage() float32 {
	return float32(e.p.takedamage)
}

func (e *EntVars) SetTakeDamage(takeDamage float32) {
	e.p.takedamage = C.float(takeDamage)
}

func (e *EntVars) DeadFlag() int {
	return int(e.p.deadflag)
}

func (e *EntVars) SetDeadFlag(deadFlag int) {
	e.p.deadflag = C.int(deadFlag)
}

func (e *EntVars) ViewOfs() [3]float32 {
	return [3]float32{
		float32(e.p.view_ofs[0]),
		float32(e.p.view_ofs[1]),
		float32(e.p.view_ofs[2]),
	}
}

func (e *EntVars) SetViewOfs(viewOfs [3]float32) {
	e.p.view_ofs[0] = C.float(viewOfs[0])
	e.p.view_ofs[1] = C.float(viewOfs[1])
	e.p.view_ofs[2] = C.float(viewOfs[2])
}

func (e *EntVars) Button() int {
	return int(e.p.button)
}

func (e *EntVars) SetButton(button int) {
	e.p.button = C.int(button)
}

func (e *EntVars) Impulse() int {
	return int(e.p.impulse)
}

func (e *EntVars) SetImpulse(impulse int) {
	e.p.impulse = C.int(impulse)
}

func (e *EntVars) Chain() *Edict {
	return edictFromC(e.globalVars, e.p.chain)
}

func (e *EntVars) SetChain(chain *Edict) {
	e.p.chain = chain.p
}

func (e *EntVars) DmgInflictor() *Edict {
	return edictFromC(e.globalVars, e.p.dmg_inflictor)
}

func (e *EntVars) SetDmgInflictor(dmgInflictor *Edict) {
	e.p.dmg_inflictor = dmgInflictor.p
}

func (e *EntVars) Enemy() *Edict {
	return edictFromC(e.globalVars, e.p.enemy)
}

func (e *EntVars) SetEnemy(enemy *Edict) {
	e.p.enemy = enemy.p
}

func (e *EntVars) AimEnt() *Edict {
	return edictFromC(e.globalVars, e.p.aiment)
}

func (e *EntVars) SetAimEnt(aimEnt *Edict) {
	e.p.aiment = aimEnt.p
}

func (e *EntVars) Owner() *Edict {
	return edictFromC(e.globalVars, e.p.owner)
}

func (e *EntVars) SetOwner(owner *Edict) {
	e.p.owner = owner.p
}

func (e *EntVars) GroundEntity() *Edict {
	return edictFromC(e.globalVars, e.p.groundentity)
}

func (e *EntVars) SetGroundEntity(groundEntity *Edict) {
	e.p.groundentity = groundEntity.p
}

func (e *EntVars) SpawnFlags() int {
	return int(e.p.spawnflags)
}

func (e *EntVars) SetSpawnFlags(spawnFlags int) {
	e.p.spawnflags = C.int(spawnFlags)
}

func (e *EntVars) Flags() int {
	return int(e.p.flags)
}

func (e *EntVars) SetFlags(flags int) {
	e.p.flags = C.int(flags)
}

func (e *EntVars) Colormap() int {
	return int(e.p.colormap)
}

func (e *EntVars) SetColormap(colormap int) {
	e.p.colormap = C.int(colormap)
}

func (e *EntVars) Team() int {
	return int(e.p.team)
}

func (e *EntVars) SetTeam(team int) {
	e.p.team = C.int(team)
}

func (e *EntVars) TeleportTime() float32 {
	return float32(e.p.teleport_time)
}

func (e *EntVars) SetTeleportTime(teleportTime float32) {
	e.p.teleport_time = C.float(teleportTime)
}

func (e *EntVars) Armortype() float32 {
	return float32(e.p.armortype)
}

func (e *EntVars) SetArmortype(armorType float32) {
	e.p.armortype = C.float(armorType)
}

func (e *EntVars) Armorvalue() float32 {
	return float32(e.p.armorvalue)
}

func (e *EntVars) SetArmorvalue(armorValue float32) {
	e.p.armorvalue = C.float(armorValue)
}

func (e *EntVars) WaterLevel() int {
	return int(e.p.waterlevel)
}

func (e *EntVars) SetWaterLevel(waterLevel int) {
	e.p.waterlevel = C.int(waterLevel)
}

func (e *EntVars) WaterType() int {
	return int(e.p.watertype)
}

func (e *EntVars) SetWaterType(waterType int) {
	e.p.watertype = C.int(waterType)
}

func (e *EntVars) Target() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.target))
}

func (e *EntVars) SetTarget(target string) {
	e.p.target = C.int(allocString(target))
}

func (e *EntVars) TargetName() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.targetname))
}

func (e *EntVars) SetTargetName(targetName string) {
	e.p.targetname = C.int(allocString(targetName))
}

func (e *EntVars) NetName() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.netname))
}

func (e *EntVars) SetNetName(netName string) {
	e.p.netname = C.int(allocString(netName))
}

func (e *EntVars) Message() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.message))
}

func (e *EntVars) SetMessage(message string) {
	e.p.message = C.int(allocString(message))
}

func (e *EntVars) DmgTake() float32 {
	return float32(e.p.dmg_take)
}

func (e *EntVars) SetDmgTake(dmgTake float32) {
	e.p.dmg_take = C.float(dmgTake)
}

func (e *EntVars) DmgSave() float32 {
	return float32(e.p.dmg_save)
}

func (e *EntVars) SetDmgSave(dmgSave float32) {
	e.p.dmg_save = C.float(dmgSave)
}

func (e *EntVars) Dmg() float32 {
	return float32(e.p.dmg)
}

func (e *EntVars) SetDmg(dmg float32) {
	e.p.dmg = C.float(dmg)
}

func (e *EntVars) DmgTime() float32 {
	return float32(e.p.dmgtime)
}

func (e *EntVars) SetDmgTime(dmgTime float32) {
	e.p.dmgtime = C.float(dmgTime)
}

func (e *EntVars) Noise() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.noise))
}

func (e *EntVars) SetNoise(noise string) {
	e.p.noise = C.int(allocString(noise))
}

func (e *EntVars) Noise1() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.noise1))
}

func (e *EntVars) SetNoise1(noise1 string) {
	e.p.noise1 = C.int(allocString(noise1))
}

func (e *EntVars) Noise2() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.noise2))
}

func (e *EntVars) SetNoise2(noise2 string) {
	e.p.noise1 = C.int(allocString(noise2))
}

func (e *EntVars) Noise3() string {
	return C.GoString(C.ReadString(e.globalVars, e.p.noise3))
}

func (e *EntVars) SetNoise3(noise3 string) {
	e.p.noise1 = C.int(allocString(noise3))
}

func (e *EntVars) Speed() float32 {
	return float32(e.p.speed)
}

func (e *EntVars) SetSpeed(speed float32) {
	e.p.speed = C.float(speed)
}

func (e *EntVars) AirFinished() float32 {
	return float32(e.p.air_finished)
}

func (e *EntVars) SetAirFinished(airFinished float32) {
	e.p.air_finished = C.float(airFinished)
}

func (e *EntVars) PainFinished() float32 {
	return float32(e.p.pain_finished)
}

func (e *EntVars) SetPainFinished(painFinished float32) {
	e.p.pain_finished = C.float(painFinished)
}

func (e *EntVars) RadsuitFinished() float32 {
	return float32(e.p.radsuit_finished)
}

func (e *EntVars) SetRadsuitFinished(radsuitFinished float32) {
	e.p.radsuit_finished = C.float(radsuitFinished)
}

func (e *EntVars) PContainingEntity() *Edict {
	return edictFromC(e.globalVars, e.p.pContainingEntity)
}

func (e *EntVars) SetPContainingEntity(pContainingEntity *Edict) {
	e.p.pContainingEntity = pContainingEntity.p
}

func (e *EntVars) PlayerClass() int {
	return int(e.p.playerclass)
}

func (e *EntVars) SetPlayerClass(playerClass int) {
	e.p.playerclass = C.int(playerClass)
}

func (e *EntVars) MaxSpeed() float32 {
	return float32(e.p.maxspeed)
}

func (e *EntVars) SetMaxSpeed(maxSpeed float32) {
	e.p.maxspeed = C.float(maxSpeed)
}

func (e *EntVars) Fov() float32 {
	return float32(e.p.fov)
}

func (e *EntVars) SetFov(fov float32) {
	e.p.fov = C.float(fov)
}

func (e *EntVars) WeaponAnim() int {
	return int(e.p.weaponanim)
}

func (e *EntVars) SetWeaponAnim(weaponAnim int) {
	e.p.weaponanim = C.int(weaponAnim)
}

func (e *EntVars) PushmSec() int {
	return int(e.p.pushmsec)
}

func (e *EntVars) SetPushmSec(pushmSec int) {
	e.p.pushmsec = C.int(pushmSec)
}

func (e *EntVars) BInDuck() int {
	return int(e.p.bInDuck)
}

func (e *EntVars) SetBInDuck(bInDuck int) {
	e.p.bInDuck = C.int(bInDuck)
}

func (e *EntVars) TimeStepSound() int {
	return int(e.p.flTimeStepSound)
}

func (e *EntVars) SetTimeStepSound(timeStepSound int) {
	e.p.flTimeStepSound = C.int(timeStepSound)
}

func (e *EntVars) SwimTime() int {
	return int(e.p.flSwimTime)
}

func (e *EntVars) SetSwimTime(swimTime int) {
	e.p.flSwimTime = C.int(swimTime)
}

func (e *EntVars) DuckTime() int {
	return int(e.p.flDuckTime)
}

func (e *EntVars) SetDuckTime(duckTime int) {
	e.p.flDuckTime = C.int(duckTime)
}

func (e *EntVars) StepLeft() int {
	return int(e.p.iStepLeft)
}

func (e *EntVars) SetStepLeft(stepLeft int) {
	e.p.iStepLeft = C.int(stepLeft)
}

func (e *EntVars) FallVelocity() float32 {
	return float32(e.p.flFallVelocity)
}

func (e *EntVars) SetFallVelocity(fallVelocity float32) {
	e.p.flFallVelocity = C.float(fallVelocity)
}

func (e *EntVars) GameState() int {
	return int(e.p.gamestate)
}

func (e *EntVars) SetGameState(gameState int) {
	e.p.gamestate = C.int(gameState)
}

func (e *EntVars) OldButtons() int {
	return int(e.p.oldbuttons)
}

func (e *EntVars) SetOldButtons(oldButtons int) {
	e.p.oldbuttons = C.int(oldButtons)
}

func (e *EntVars) GroupInfo() int {
	return int(e.p.groupinfo)
}

func (e *EntVars) SetGroupInfo(groupInfo int) {
	e.p.groupinfo = C.int(groupInfo)
}

func (e *EntVars) IUser1() int {
	return int(e.p.iuser1)
}

func (e *EntVars) SetIUser1(iUser1 int) {
	e.p.iuser1 = C.int(iUser1)
}

func (e *EntVars) IUser2() int {
	return int(e.p.iuser2)
}

func (e *EntVars) SetIUser2(iUser2 int) {
	e.p.iuser2 = C.int(iUser2)
}

func (e *EntVars) IUser3() int {
	return int(e.p.iuser3)
}

func (e *EntVars) SetIUser3(iUser3 int) {
	e.p.iuser3 = C.int(iUser3)
}

func (e *EntVars) IUser4() int {
	return int(e.p.iuser4)
}

func (e *EntVars) SetIUser4(iUser4 int) {
	e.p.iuser4 = C.int(iUser4)
}

func (e *EntVars) FUser1() float32 {
	return float32(e.p.fuser1)
}

func (e *EntVars) SetFUser1(fUser1 float32) {
	e.p.fuser1 = C.float(fUser1)
}

func (e *EntVars) FUser2() float32 {
	return float32(e.p.fuser2)
}

func (e *EntVars) SetFUser2(fUser2 float32) {
	e.p.fuser2 = C.float(fUser2)
}

func (e *EntVars) FUser3() float32 {
	return float32(e.p.fuser3)
}

func (e *EntVars) SetFUser3(fUser3 float32) {
	e.p.fuser3 = C.float(fUser3)
}

func (e *EntVars) FUser4() float32 {
	return float32(e.p.fuser4)
}

func (e *EntVars) SetFUser4(fUser4 float32) {
	e.p.fuser4 = C.float(fUser4)
}

func (e *EntVars) VUser1() [3]float32 {
	return [3]float32{
		float32(e.p.vuser1[0]),
		float32(e.p.vuser1[1]),
		float32(e.p.vuser1[2]),
	}
}

func (e *EntVars) SetVUser1(vUser1 [3]float32) {
	e.p.vuser1[0] = C.float(vUser1[0])
	e.p.vuser1[1] = C.float(vUser1[1])
	e.p.vuser1[2] = C.float(vUser1[2])
}

func (e *EntVars) VUser2() [3]float32 {
	return [3]float32{
		float32(e.p.vuser2[0]),
		float32(e.p.vuser2[1]),
		float32(e.p.vuser2[2]),
	}
}

func (e *EntVars) SetVUser2(vUser2 [3]float32) {
	e.p.vuser2[0] = C.float(vUser2[0])
	e.p.vuser2[1] = C.float(vUser2[1])
	e.p.vuser2[2] = C.float(vUser2[2])
}

func (e *EntVars) VUser3() [3]float32 {
	return [3]float32{
		float32(e.p.vuser3[0]),
		float32(e.p.vuser3[1]),
		float32(e.p.vuser3[2]),
	}
}

func (e *EntVars) SetVUser3(vUser3 [3]float32) {
	e.p.vuser3[0] = C.float(vUser3[0])
	e.p.vuser3[1] = C.float(vUser3[1])
	e.p.vuser3[2] = C.float(vUser3[2])
}

func (e *EntVars) VUser4() [3]float32 {
	return [3]float32{
		float32(e.p.vuser4[0]),
		float32(e.p.vuser4[1]),
		float32(e.p.vuser4[2]),
	}
}

func (e *EntVars) SetVUser4(vUser4 [3]float32) {
	e.p.vuser4[0] = C.float(vUser4[0])
	e.p.vuser4[1] = C.float(vUser4[1])
	e.p.vuser4[2] = C.float(vUser4[2])
}

func (e *EntVars) EUser1() *Edict {
	return edictFromC(e.globalVars, e.p.euser1)
}

func (e *EntVars) SetEUser1(eUser1 *Edict) {
	e.p.euser1 = eUser1.p
}

func (e *EntVars) EUser2() *Edict {
	return edictFromC(e.globalVars, e.p.euser2)
}

func (e *EntVars) SetEUser2(eUser2 *Edict) {
	e.p.euser2 = eUser2.p
}

func (e *EntVars) EUser3() *Edict {
	return edictFromC(e.globalVars, e.p.euser3)
}

func (e *EntVars) SetEUser3(eUser3 *Edict) {
	e.p.euser3 = eUser3.p
}

func (e *EntVars) EUser4() *Edict {
	return edictFromC(e.globalVars, e.p.euser4)
}

func (e *EntVars) SetEUser4(eUser4 *Edict) {
	e.p.euser4 = eUser4.p
}

type TraceResult struct {
	AllSolid    bool       // if true, plane is not valid
	StartSolid  bool       // if true, the initial point was in a solid area
	InOpen      bool       // if true, the initial point was in empty space
	InWater     bool       // if true, the initial point was underwater
	Fraction    float32    // time completed, 1.0 = didn't hit anything
	EndPos      [3]float32 // final position
	PlaneDist   float32    // distance from the plane
	PlaneNormal [3]float32 // surface normal at impact
	Hit         *Edict     // entity the surface is on
	HitGroup    int        // 0 == generic, non-zero is specific body part
}

func traceResultFromC(globalVars *C.globalvars_t, tr C.TraceResult) *TraceResult {
	result := &TraceResult{
		AllSolid:   tr.fAllSolid == 1,
		StartSolid: tr.fStartSolid == 1,
		InOpen:     tr.fInOpen == 1,
		InWater:    tr.fInWater == 1,
		Fraction:   float32(tr.flFraction),
		PlaneDist:  float32(tr.flPlaneDist),
	}

	if len(tr.vecEndPos) == 3 {
		result.EndPos = [3]float32{
			float32(tr.vecEndPos[0]),
			float32(tr.vecEndPos[1]),
			float32(tr.vecEndPos[2]),
		}
	}

	if len(tr.vecPlaneNormal) == 3 {
		result.PlaneNormal = [3]float32{
			float32(tr.vecPlaneNormal[0]),
			float32(tr.vecPlaneNormal[1]),
			float32(tr.vecPlaneNormal[2]),
		}
	}

	if tr.pHit != nil {
		result.Hit = edictFromC(globalVars, tr.pHit)
	}

	return result
}

const (
	mipLevels = 4
)

type Texture struct {
	Name           string
	Width          uint32
	Height         uint32
	AnimTotal      int
	AnimMin        int
	AnimMax        int
	AnimNext       *Texture
	AlternateAnims *Texture
	Offsets        [mipLevels]uint32
	PalOffset      uint32
}

func textureFromC(t *C.texture_t) *Texture {
	texture := &Texture{
		Name:      C.GoString(&t.name[0]),
		Width:     uint32(t.width),
		Height:    uint32(t.height),
		AnimTotal: int(t.anim_total),
		AnimMin:   int(t.anim_min),
		AnimMax:   int(t.anim_max),
		PalOffset: uint32(t.paloffset),
	}

	if t.anim_next != nil {
		texture.AnimNext = textureFromC(t.anim_next)
	}

	if t.alternate_anims != nil {
		texture.AlternateAnims = textureFromC(t.alternate_anims)
	}

	for i := 0; i < mipLevels; i++ {
		texture.Offsets[i] = uint32(t.offsets[i])
	}

	return texture
}
