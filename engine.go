package metamod_go

/*
#include <eiface.h>
#include <com_model.h>

const char emptyString[] = "";

const char* ReadString(globalvars_t *gpGlobals, int offset) {
	unsigned int addr = (gpGlobals->pStringBase + (unsigned int)(offset));
	if ((void *)addr == NULL) {
		return emptyString;
	}

	return (const char *)addr;
}

int MakeString(globalvars_t *gpGlobals, char *str) {
	return ((uint64)(str) - (uint64)(ReadString(gpGlobals, 0)));
}

*/
import "C"
import (
	"github.com/et-nik/metamod-go/engine"
	"github.com/et-nik/metamod-go/vector"
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

func (e *Edict) ptr() *C.edict_t {
	if e == nil {
		return nil
	}

	return e.p
}

func (e *Edict) SerialNumber() int {
	return int(e.p.serialnumber)
}

func (e *Edict) EntVars() *EntVars {
	return entVarsFromC(e.globalVars, &e.p.v)
}

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

func (e *EntVars) IsValid() bool {
	if e.p == nil {
		return false
	}

	if e.p.pContainingEntity == nil {
		return false
	}

	if e.p.flags&engine.EdictFlagKillMe != 0 {
		return false
	}

	return true
}

func (e *EntVars) ClassName() string {
	if !e.IsValid() {
		return ""
	}

	return C.GoString(C.ReadString(e.globalVars, e.p.classname))
}

func (e *EntVars) GlobalName() string {
	if !e.IsValid() {
		return ""
	}

	return C.GoString(C.ReadString(e.globalVars, e.p.globalname))
}

func (e *EntVars) Origin() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	if len(e.p.origin) == 0 {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.origin[0]),
		float32(e.p.origin[1]),
		float32(e.p.origin[2]),
	}
}

func (e *EntVars) SetOrigin(origin vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.origin[0] = C.float(origin[0])
	e.p.origin[1] = C.float(origin[1])
	e.p.origin[2] = C.float(origin[2])
}

func (e *EntVars) OldOrigin() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.oldorigin[0]),
		float32(e.p.oldorigin[1]),
		float32(e.p.oldorigin[2]),
	}
}

func (e *EntVars) SetOldOrigin(origin vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.oldorigin[0] = C.float(origin[0])
	e.p.oldorigin[1] = C.float(origin[1])
	e.p.oldorigin[2] = C.float(origin[2])
}

func (e *EntVars) Velocity() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.velocity[0]),
		float32(e.p.velocity[1]),
		float32(e.p.velocity[2]),
	}
}

func (e *EntVars) SetVelocity(velocity vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.velocity[0] = C.float(velocity[0])
	e.p.velocity[1] = C.float(velocity[1])
	e.p.velocity[2] = C.float(velocity[2])
}

func (e *EntVars) BaseVelocity() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.basevelocity[0]),
		float32(e.p.basevelocity[1]),
		float32(e.p.basevelocity[2]),
	}
}

func (e *EntVars) SetBaseVelocity(velocity vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.basevelocity[0] = C.float(velocity[0])
	e.p.basevelocity[1] = C.float(velocity[1])
	e.p.basevelocity[2] = C.float(velocity[2])
}

// ClBaseVelocity Base velocity that was passed in to server physics so
// client can predict conveyors correctly.
// Server zeroes it, so we need to store here, too.
func (e *EntVars) ClBaseVelocity() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.clbasevelocity[0]),
		float32(e.p.clbasevelocity[1]),
		float32(e.p.clbasevelocity[2]),
	}
}

func (e *EntVars) MoveDir() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.movedir[0]),
		float32(e.p.movedir[1]),
		float32(e.p.movedir[2]),
	}
}

func (e *EntVars) SetMoveDir(dir vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.movedir[0] = C.float(dir[0])
	e.p.movedir[1] = C.float(dir[1])
	e.p.movedir[2] = C.float(dir[2])
}

// Angles Model angles
func (e *EntVars) Angles() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.angles[0]),
		float32(e.p.angles[1]),
		float32(e.p.angles[2]),
	}
}

func (e *EntVars) SetAngles(angles vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.angles[0] = C.float(angles[0])
	e.p.angles[1] = C.float(angles[1])
	e.p.angles[2] = C.float(angles[2])
}

func (e *EntVars) Avelocity() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.avelocity[0]),
		float32(e.p.avelocity[1]),
		float32(e.p.avelocity[2]),
	}
}

func (e *EntVars) SetAvelocity(avelocity vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.avelocity[0] = C.float(avelocity[0])
	e.p.avelocity[1] = C.float(avelocity[1])
	e.p.avelocity[2] = C.float(avelocity[2])
}

func (e *EntVars) PunchAngle() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.punchangle[0]),
		float32(e.p.punchangle[1]),
		float32(e.p.punchangle[2]),
	}
}

func (e *EntVars) SetPunchAngle(punchangle vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.punchangle[0] = C.float(punchangle[0])
	e.p.punchangle[1] = C.float(punchangle[1])
	e.p.punchangle[2] = C.float(punchangle[2])
}

func (e *EntVars) VAngle() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.v_angle[0]),
		float32(e.p.v_angle[1]),
		float32(e.p.v_angle[2]),
	}
}

func (e *EntVars) SetVAngle(vAngle vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.v_angle[0] = C.float(vAngle[0])
	e.p.v_angle[1] = C.float(vAngle[1])
	e.p.v_angle[2] = C.float(vAngle[2])
}

func (e *EntVars) EndPos() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.endpos[0]),
		float32(e.p.endpos[1]),
		float32(e.p.endpos[2]),
	}
}

func (e *EntVars) SetEndPos(endPos vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.endpos[0] = C.float(endPos[0])
	e.p.endpos[1] = C.float(endPos[1])
	e.p.endpos[2] = C.float(endPos[2])
}

func (e *EntVars) StartPos() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.startpos[0]),
		float32(e.p.startpos[1]),
		float32(e.p.startpos[2]),
	}
}

func (e *EntVars) SetStartPost(startPos vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.startpos[0] = C.float(startPos[0])
	e.p.startpos[1] = C.float(startPos[1])
	e.p.startpos[2] = C.float(startPos[2])
}

func (e *EntVars) ImpactTime() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.impacttime)
}

func (e *EntVars) SetImpactTime(impactTime float32) {
	if !e.IsValid() {
		return
	}

	e.p.impacttime = C.float(impactTime)
}

func (e *EntVars) StartTime() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.starttime)
}

func (e *EntVars) SetStartTime(startTime float32) {
	if !e.IsValid() {
		return
	}

	e.p.starttime = C.float(startTime)
}

func (e *EntVars) FixAngle() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.fixangle)
}

func (e *EntVars) SetFixAngle(fixAngle int) {
	if !e.IsValid() {
		return
	}

	e.p.fixangle = C.int(fixAngle)
}

func (e *EntVars) IdealPitch() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.idealpitch)
}

func (e *EntVars) SetIdealPitch(idealPitch float32) {
	if !e.IsValid() {
		return
	}

	e.p.idealpitch = C.float(idealPitch)
}

func (e *EntVars) PitchSpeed() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.pitch_speed)
}

func (e *EntVars) SetPitchSpeed(pitchSpeed float32) {
	if !e.IsValid() {
		return
	}

	e.p.pitch_speed = C.float(pitchSpeed)
}

func (e *EntVars) IdealYaw() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.ideal_yaw)
}

func (e *EntVars) SetIdealYaw(idealYaw float32) {
	if !e.IsValid() {
		return
	}

	e.p.ideal_yaw = C.float(idealYaw)
}

func (e *EntVars) YawSpeed() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.yaw_speed)
}

func (e *EntVars) SetYawSpeed(yawSpeed float32) {
	if !e.IsValid() {
		return
	}

	e.p.yaw_speed = C.float(yawSpeed)
}

func (e *EntVars) ModelIndex() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.modelindex)
}

func (e *EntVars) SetModelIndex(modelIndex int) {
	if !e.IsValid() {
		return
	}

	e.p.modelindex = C.int(modelIndex)
}

func (e *EntVars) Model() string {
	if !e.IsValid() {
		return ""
	}

	return C.GoString(C.ReadString(e.globalVars, e.p.model))
}

func (e *EntVars) SetModel(model string) {
	if !e.IsValid() {
		return
	}

	e.p.model = C.int(allocString(model))
}

func (e *EntVars) ViewModel() string {
	if !e.IsValid() {
		return ""
	}

	return C.GoString(C.ReadString(e.globalVars, e.p.viewmodel))
}

func (e *EntVars) SetViewModel(viewModel string) {
	if !e.IsValid() {
		return
	}

	e.p.viewmodel = C.int(allocString(viewModel))
}

func (e *EntVars) WeaponModel() string {
	if !e.IsValid() {
		return ""
	}

	return C.GoString(C.ReadString(e.globalVars, e.p.weaponmodel))
}

func (e *EntVars) SetWeaponModel(weaponModel string) {
	if !e.IsValid() {
		return
	}

	e.p.weaponmodel = C.int(allocString(weaponModel))
}

func (e *EntVars) AbsMin() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.absmin[0]),
		float32(e.p.absmin[1]),
		float32(e.p.absmin[2]),
	}
}

func (e *EntVars) SetAbsMin(absMin vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.absmin[0] = C.float(absMin[0])
	e.p.absmin[1] = C.float(absMin[1])
	e.p.absmin[2] = C.float(absMin[2])
}

func (e *EntVars) AbsMax() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.absmax[0]),
		float32(e.p.absmax[1]),
		float32(e.p.absmax[2]),
	}
}

func (e *EntVars) SetAbsMax(absMax vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.absmax[0] = C.float(absMax[0])
	e.p.absmax[1] = C.float(absMax[1])
	e.p.absmax[2] = C.float(absMax[2])
}

func (e *EntVars) Mins() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.mins[0]),
		float32(e.p.mins[1]),
		float32(e.p.mins[2]),
	}
}

func (e *EntVars) SetMins(mins vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.mins[0] = C.float(mins[0])
	e.p.mins[1] = C.float(mins[1])
	e.p.mins[2] = C.float(mins[2])
}

func (e *EntVars) Maxs() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.maxs[0]),
		float32(e.p.maxs[1]),
		float32(e.p.maxs[2]),
	}
}

func (e *EntVars) SetMaxs(maxs vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.maxs[0] = C.float(maxs[0])
	e.p.maxs[1] = C.float(maxs[1])
	e.p.maxs[2] = C.float(maxs[2])
}

func (e *EntVars) Size() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.size[0]),
		float32(e.p.size[1]),
		float32(e.p.size[2]),
	}
}

func (e *EntVars) SetSize(size vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.size[0] = C.float(size[0])
	e.p.size[1] = C.float(size[1])
	e.p.size[2] = C.float(size[2])
}

func (e *EntVars) Ltime() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.ltime)
}

func (e *EntVars) SetLtime(ltime float32) {
	if !e.IsValid() {
		return
	}

	e.p.ltime = C.float(ltime)
}

func (e *EntVars) NextThink() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.nextthink)
}

func (e *EntVars) SetNextThink(nextThink float32) {
	if !e.IsValid() {
		return
	}

	e.p.nextthink = C.float(nextThink)
}

func (e *EntVars) MoveType() engine.MoveType {
	if !e.IsValid() {
		return 0
	}

	return engine.MoveType(int(e.p.movetype))
}

func (e *EntVars) SetMoveType(moveType engine.MoveType) {
	if !e.IsValid() {
		return
	}

	e.p.movetype = C.int(moveType)
}

func (e *EntVars) Solid() engine.SolidType {
	if !e.IsValid() {
		return 0
	}

	return engine.SolidType(int(e.p.solid))
}

func (e *EntVars) SetSolid(solid int) {
	if !e.IsValid() {
		return
	}

	e.p.solid = C.int(solid)
}

func (e *EntVars) Skin() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.skin)
}

func (e *EntVars) SetSkin(skin int) {
	if !e.IsValid() {
		return
	}

	e.p.skin = C.int(skin)
}

func (e *EntVars) Body() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.body)
}

func (e *EntVars) SetBody(body int) {
	if !e.IsValid() {
		return
	}

	e.p.body = C.int(body)
}

func (e *EntVars) Effects() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.effects)
}

func (e *EntVars) SetEffects(effects int) {
	if !e.IsValid() {
		return
	}

	e.p.effects = C.int(effects)
}

func (e *EntVars) Gravity() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.gravity)
}

func (e *EntVars) SetGravity(gravity float32) {
	if !e.IsValid() {
		return
	}

	e.p.gravity = C.float(gravity)
}

func (e *EntVars) Friction() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.friction)
}

func (e *EntVars) SetFriction(friction float32) {
	if !e.IsValid() {
		return
	}

	e.p.friction = C.float(friction)
}

func (e *EntVars) LightLevel() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.light_level)
}

func (e *EntVars) SetLightLevel(lightLevel int) {
	if !e.IsValid() {
		return
	}

	e.p.light_level = C.int(lightLevel)
}

// Sequence animation sequence
func (e *EntVars) Sequence() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.sequence)
}

func (e *EntVars) SetSequence(sequence int) {
	if !e.IsValid() {
		return
	}

	e.p.sequence = C.int(sequence)
}

// GaitSequence movement animation sequence for player (0 for none)
func (e *EntVars) GaitSequence() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.gaitsequence)
}

func (e *EntVars) SetGaitSequence(gaitSequence int) {
	if !e.IsValid() {
		return
	}

	e.p.gaitsequence = C.int(gaitSequence)
}

// Frame % playback position in animation sequences (0..255)
func (e *EntVars) Frame() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.frame)
}

func (e *EntVars) SetFrame(frame float32) {
	if !e.IsValid() {
		return
	}

	e.p.frame = C.float(frame)
}

// AnymTime world time when frame was set
func (e *EntVars) AnymTime() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.animtime)
}

func (e *EntVars) SetAnimTime(animTime float32) {
	if !e.IsValid() {
		return
	}

	e.p.animtime = C.float(animTime)
}

// FrameRate animation playback rate (-8x to 8x)
func (e *EntVars) FrameRate() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.framerate)
}

func (e *EntVars) SetFrameRate(frameRate float32) {
	if !e.IsValid() {
		return
	}

	e.p.framerate = C.float(frameRate)
}

// Controller bone controller setting
func (e *EntVars) Controller() [4]byte {
	if !e.IsValid() {
		return [4]byte{}
	}

	return [4]byte{
		byte(e.p.controller[0]),
		byte(e.p.controller[1]),
		byte(e.p.controller[2]),
		byte(e.p.controller[3]),
	}
}

func (e *EntVars) SetController(controller [4]byte) {
	if !e.IsValid() {
		return
	}

	e.p.controller[0] = C.byte(controller[0])
	e.p.controller[1] = C.byte(controller[1])
	e.p.controller[2] = C.byte(controller[2])
	e.p.controller[3] = C.byte(controller[3])
}

// Blending blending amount between sub-sequences
func (e *EntVars) Blending() [2]byte {
	if !e.IsValid() {
		return [2]byte{}
	}

	return [2]byte{
		byte(e.p.blending[0]),
		byte(e.p.blending[1]),
	}
}

func (e *EntVars) SetBlending(blending [2]byte) {
	if !e.IsValid() {
		return
	}

	e.p.blending[0] = C.byte(blending[0])
	e.p.blending[1] = C.byte(blending[1])
}

func (e *EntVars) Scale() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.scale)
}

func (e *EntVars) SetScale(scale float32) {
	if !e.IsValid() {
		return
	}

	e.p.scale = C.float(scale)
}

func (e *EntVars) RenderMode() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.rendermode)
}

func (e *EntVars) SetRenderMode(renderMode int) {
	if !e.IsValid() {
		return
	}

	e.p.rendermode = C.int(renderMode)
}

func (e *EntVars) RenderAmt() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.renderamt)
}

func (e *EntVars) SetRenderAmt(renderAmt float32) {
	if !e.IsValid() {
		return
	}

	e.p.renderamt = C.float(renderAmt)
}

func (e *EntVars) RenderColor() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.rendercolor[0]),
		float32(e.p.rendercolor[1]),
		float32(e.p.rendercolor[2]),
	}
}

func (e *EntVars) SetRenderColor(renderColor vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.rendercolor[0] = C.float(renderColor[0])
	e.p.rendercolor[1] = C.float(renderColor[1])
	e.p.rendercolor[2] = C.float(renderColor[2])
}

func (e *EntVars) RenderFx() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.renderfx)
}

func (e *EntVars) SetRenderFx(renderFx int) {
	if !e.IsValid() {
		return
	}

	e.p.renderfx = C.int(renderFx)
}

func (e *EntVars) Health() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.health)
}

func (e *EntVars) SetHealth(health float32) {
	if !e.IsValid() {
		return
	}

	e.p.health = C.float(health)
}

func (e *EntVars) MaxHealth() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.max_health)
}

func (e *EntVars) SetMaxHealth(maxHealth float32) {
	if !e.IsValid() {
		return
	}

	e.p.max_health = C.float(maxHealth)
}

func (e *EntVars) Frags() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.frags)
}

func (e *EntVars) SetFrags(frags float32) {
	if !e.IsValid() {
		return
	}

	e.p.frags = C.float(frags)
}

func (e *EntVars) Weapons() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.weapons)
}

func (e *EntVars) SetWeapons(weapons int) {
	if !e.IsValid() {
		return
	}

	e.p.weapons = C.int(weapons)
}

func (e *EntVars) TakeDamage() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.takedamage)
}

func (e *EntVars) SetTakeDamage(takeDamage float32) {
	if !e.IsValid() {
		return
	}

	e.p.takedamage = C.float(takeDamage)
}

func (e *EntVars) DeadFlag() engine.DeadFlag {
	if !e.IsValid() {
		return 0
	}

	return engine.DeadFlag(int(e.p.deadflag))
}

func (e *EntVars) SetDeadFlag(deadFlag engine.DeadFlag) {
	if !e.IsValid() {
		return
	}

	e.p.deadflag = C.int(int(deadFlag))
}

func (e *EntVars) ViewOfs() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.view_ofs[0]),
		float32(e.p.view_ofs[1]),
		float32(e.p.view_ofs[2]),
	}
}

func (e *EntVars) SetViewOfs(viewOfs vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.view_ofs[0] = C.float(viewOfs[0])
	e.p.view_ofs[1] = C.float(viewOfs[1])
	e.p.view_ofs[2] = C.float(viewOfs[2])
}

func (e *EntVars) Button() engine.InButtonFlag {
	if !e.IsValid() {
		return 0
	}

	return engine.InButtonFlag(int(e.p.button))
}

func (e *EntVars) ButtonClear() {
	if !e.IsValid() {
		return
	}

	e.p.button = 0
}

func (e *EntVars) SetButton(button engine.InButtonFlag) {
	if !e.IsValid() {
		return
	}

	e.p.button = C.int(button)
}

func (e *EntVars) SetButtonBit(button engine.InButtonFlag) {
	if !e.IsValid() {
		return
	}

	e.p.button |= C.int(int(button))
}

func (e *EntVars) ButtonToggle(button engine.InButtonFlag) {
	if !e.IsValid() {
		return
	}

	e.p.button ^= C.int(int(button))
}

func (e *EntVars) ButtonHas(button engine.InButtonFlag) bool {
	if !e.IsValid() {
		return false
	}

	return e.p.button&C.int(int(button)) != 0
}

func (e *EntVars) ButtonClearBit(button engine.InButtonFlag) {
	if !e.IsValid() {
		return
	}

	e.p.button &= ^C.int(int(button))
}

func (e *EntVars) Impulse() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.impulse)
}

func (e *EntVars) SetImpulse(impulse int) {
	if !e.IsValid() {
		return
	}

	e.p.impulse = C.int(impulse)
}

func (e *EntVars) Chain() *Edict {
	if !e.IsValid() {
		return nil
	}

	return edictFromC(e.globalVars, e.p.chain)
}

func (e *EntVars) SetChain(chain *Edict) {
	if !e.IsValid() {
		return
	}

	e.p.chain = chain.p
}

func (e *EntVars) DmgInflictor() *Edict {
	if !e.IsValid() {
		return nil
	}

	return edictFromC(e.globalVars, e.p.dmg_inflictor)
}

func (e *EntVars) SetDmgInflictor(dmgInflictor *Edict) {
	if !e.IsValid() {
		return
	}

	e.p.dmg_inflictor = dmgInflictor.p
}

func (e *EntVars) Enemy() *Edict {
	if !e.IsValid() {
		return nil
	}

	return edictFromC(e.globalVars, e.p.enemy)
}

func (e *EntVars) SetEnemy(enemy *Edict) {
	if !e.IsValid() {
		return
	}

	e.p.enemy = enemy.p
}

func (e *EntVars) AimEnt() *Edict {
	if !e.IsValid() {
		return nil
	}

	return edictFromC(e.globalVars, e.p.aiment)
}

func (e *EntVars) SetAimEnt(aimEnt *Edict) {
	if !e.IsValid() {
		return
	}

	e.p.aiment = aimEnt.p
}

func (e *EntVars) Owner() *Edict {
	if !e.IsValid() {
		return nil
	}

	return edictFromC(e.globalVars, e.p.owner)
}

func (e *EntVars) SetOwner(owner *Edict) {
	if !e.IsValid() {
		return
	}

	e.p.owner = owner.p
}

func (e *EntVars) GroundEntity() *Edict {
	if !e.IsValid() {
		return nil
	}

	return edictFromC(e.globalVars, e.p.groundentity)
}

func (e *EntVars) SetGroundEntity(groundEntity *Edict) {
	if !e.IsValid() {
		return
	}

	e.p.groundentity = groundEntity.p
}

func (e *EntVars) SpawnFlags() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.spawnflags)
}

func (e *EntVars) SetSpawnFlags(spawnFlags int) {
	if !e.IsValid() {
		return
	}

	e.p.spawnflags = C.int(spawnFlags)
}

func (e *EntVars) Flags() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.flags)
}

func (e *EntVars) SetFlags(flags int) {
	if !e.IsValid() {
		return
	}

	e.p.flags = C.int(flags)
}

func (e *EntVars) FlagsHas(bit int) bool {
	if !e.IsValid() {
		return false
	}

	return e.p.flags&C.int(bit) != 0
}

func (e *EntVars) FlagsToggle(bit int) {
	if !e.IsValid() {
		return
	}

	e.p.flags ^= C.int(bit)
}

func (e *EntVars) SetFlagsBit(bit int) {
	if !e.IsValid() {
		return
	}

	e.p.flags |= C.int(bit)
}

func (e *EntVars) FlagsClearBit(bit int) {
	if !e.IsValid() {
		return
	}

	e.p.flags &= ^C.int(bit)
}

func (e *EntVars) Colormap() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.colormap)
}

func (e *EntVars) SetColormap(colormap int) {
	if !e.IsValid() {
		return
	}

	e.p.colormap = C.int(colormap)
}

func (e *EntVars) Team() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.team)
}

func (e *EntVars) SetTeam(team int) {
	if !e.IsValid() {
		return
	}

	e.p.team = C.int(team)
}

func (e *EntVars) TeleportTime() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.teleport_time)
}

func (e *EntVars) SetTeleportTime(teleportTime float32) {
	if !e.IsValid() {
		return
	}

	e.p.teleport_time = C.float(teleportTime)
}

func (e *EntVars) Armortype() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.armortype)
}

func (e *EntVars) SetArmortype(armorType float32) {
	if !e.IsValid() {
		return
	}

	e.p.armortype = C.float(armorType)
}

func (e *EntVars) Armorvalue() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.armorvalue)
}

func (e *EntVars) SetArmorvalue(armorValue float32) {
	if !e.IsValid() {
		return
	}

	e.p.armorvalue = C.float(armorValue)
}

func (e *EntVars) WaterLevel() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.waterlevel)
}

func (e *EntVars) SetWaterLevel(waterLevel int) {
	if !e.IsValid() {
		return
	}

	e.p.waterlevel = C.int(waterLevel)
}

func (e *EntVars) WaterType() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.watertype)
}

func (e *EntVars) SetWaterType(waterType int) {
	if !e.IsValid() {
		return
	}

	e.p.watertype = C.int(waterType)
}

func (e *EntVars) Target() string {
	if !e.IsValid() {
		return ""
	}

	return C.GoString(C.ReadString(e.globalVars, e.p.target))
}

func (e *EntVars) SetTarget(target string) {
	if !e.IsValid() {
		return
	}

	e.p.target = C.int(allocString(target))
}

func (e *EntVars) TargetName() string {
	if !e.IsValid() {
		return ""
	}

	return C.GoString(C.ReadString(e.globalVars, e.p.targetname))
}

func (e *EntVars) SetTargetName(targetName string) {
	if !e.IsValid() {
		return
	}

	e.p.targetname = C.int(allocString(targetName))
}

func (e *EntVars) NetName() string {
	if !e.IsValid() {
		return ""
	}

	return C.GoString(C.ReadString(e.globalVars, e.p.netname))
}

func (e *EntVars) SetNetName(netName string) {
	if !e.IsValid() {
		return
	}

	e.p.netname = C.int(allocString(netName))
}

func (e *EntVars) Message() string {
	if !e.IsValid() {
		return ""
	}

	return C.GoString(C.ReadString(e.globalVars, e.p.message))
}

func (e *EntVars) SetMessage(message string) {
	if !e.IsValid() {
		return
	}

	e.p.message = C.int(allocString(message))
}

func (e *EntVars) DmgTake() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.dmg_take)
}

func (e *EntVars) SetDmgTake(dmgTake float32) {
	if !e.IsValid() {
		return
	}

	e.p.dmg_take = C.float(dmgTake)
}

func (e *EntVars) DmgSave() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.dmg_save)
}

func (e *EntVars) SetDmgSave(dmgSave float32) {
	if !e.IsValid() {
		return
	}

	e.p.dmg_save = C.float(dmgSave)
}

func (e *EntVars) Dmg() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.dmg)
}

func (e *EntVars) SetDmg(dmg float32) {
	if !e.IsValid() {
		return
	}

	e.p.dmg = C.float(dmg)
}

func (e *EntVars) DmgTime() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.dmgtime)
}

func (e *EntVars) SetDmgTime(dmgTime float32) {
	if !e.IsValid() {
		return
	}

	e.p.dmgtime = C.float(dmgTime)
}

func (e *EntVars) Noise() string {
	if !e.IsValid() {
		return ""
	}

	return C.GoString(C.ReadString(e.globalVars, e.p.noise))
}

func (e *EntVars) SetNoise(noise string) {
	if !e.IsValid() {
		return
	}

	e.p.noise = C.int(allocString(noise))
}

func (e *EntVars) Noise1() string {
	if !e.IsValid() {
		return ""
	}

	return C.GoString(C.ReadString(e.globalVars, e.p.noise1))
}

func (e *EntVars) SetNoise1(noise1 string) {
	if !e.IsValid() {
		return
	}

	e.p.noise1 = C.int(allocString(noise1))
}

func (e *EntVars) Noise2() string {
	if !e.IsValid() {
		return ""
	}

	return C.GoString(C.ReadString(e.globalVars, e.p.noise2))
}

func (e *EntVars) SetNoise2(noise2 string) {
	if !e.IsValid() {
		return
	}

	e.p.noise1 = C.int(allocString(noise2))
}

func (e *EntVars) Noise3() string {
	if !e.IsValid() {
		return ""
	}

	return C.GoString(C.ReadString(e.globalVars, e.p.noise3))
}

func (e *EntVars) SetNoise3(noise3 string) {
	if !e.IsValid() {
		return
	}

	e.p.noise1 = C.int(allocString(noise3))
}

func (e *EntVars) Speed() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.speed)
}

func (e *EntVars) SetSpeed(speed float32) {
	if !e.IsValid() {
		return
	}

	e.p.speed = C.float(speed)
}

func (e *EntVars) AirFinished() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.air_finished)
}

func (e *EntVars) SetAirFinished(airFinished float32) {
	if !e.IsValid() {
		return
	}

	e.p.air_finished = C.float(airFinished)
}

func (e *EntVars) PainFinished() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.pain_finished)
}

func (e *EntVars) SetPainFinished(painFinished float32) {
	if !e.IsValid() {
		return
	}

	e.p.pain_finished = C.float(painFinished)
}

func (e *EntVars) RadsuitFinished() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.radsuit_finished)
}

func (e *EntVars) SetRadsuitFinished(radsuitFinished float32) {
	if !e.IsValid() {
		return
	}

	e.p.radsuit_finished = C.float(radsuitFinished)
}

func (e *EntVars) PContainingEntity() *Edict {
	if !e.IsValid() {
		return nil
	}

	return edictFromC(e.globalVars, e.p.pContainingEntity)
}

func (e *EntVars) SetPContainingEntity(pContainingEntity *Edict) {
	if !e.IsValid() {
		return
	}

	e.p.pContainingEntity = pContainingEntity.p
}

func (e *EntVars) PlayerClass() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.playerclass)
}

func (e *EntVars) SetPlayerClass(playerClass int) {
	e.p.playerclass = C.int(playerClass)
}

func (e *EntVars) MaxSpeed() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.maxspeed)
}

func (e *EntVars) SetMaxSpeed(maxSpeed float32) {
	if !e.IsValid() {
		return
	}

	e.p.maxspeed = C.float(maxSpeed)
}

func (e *EntVars) Fov() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.fov)
}

func (e *EntVars) SetFov(fov float32) {
	if !e.IsValid() {
		return
	}

	e.p.fov = C.float(fov)
}

func (e *EntVars) WeaponAnim() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.weaponanim)
}

func (e *EntVars) SetWeaponAnim(weaponAnim int) {
	if !e.IsValid() {
		return
	}

	e.p.weaponanim = C.int(weaponAnim)
}

func (e *EntVars) PushmSec() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.pushmsec)
}

func (e *EntVars) SetPushmSec(pushmSec int) {
	if !e.IsValid() {
		return
	}

	e.p.pushmsec = C.int(pushmSec)
}

func (e *EntVars) BInDuck() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.bInDuck)
}

func (e *EntVars) SetBInDuck(bInDuck int) {
	if !e.IsValid() {
		return
	}

	e.p.bInDuck = C.int(bInDuck)
}

func (e *EntVars) TimeStepSound() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.flTimeStepSound)
}

func (e *EntVars) SetTimeStepSound(timeStepSound int) {
	if !e.IsValid() {
		return
	}

	e.p.flTimeStepSound = C.int(timeStepSound)
}

func (e *EntVars) SwimTime() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.flSwimTime)
}

func (e *EntVars) SetSwimTime(swimTime int) {
	if !e.IsValid() {
		return
	}

	e.p.flSwimTime = C.int(swimTime)
}

func (e *EntVars) DuckTime() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.flDuckTime)
}

func (e *EntVars) SetDuckTime(duckTime int) {
	if !e.IsValid() {
		return
	}

	e.p.flDuckTime = C.int(duckTime)
}

func (e *EntVars) StepLeft() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.iStepLeft)
}

func (e *EntVars) SetStepLeft(stepLeft int) {
	if !e.IsValid() {
		return
	}

	e.p.iStepLeft = C.int(stepLeft)
}

func (e *EntVars) FallVelocity() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.flFallVelocity)
}

func (e *EntVars) SetFallVelocity(fallVelocity float32) {
	if !e.IsValid() {
		return
	}

	e.p.flFallVelocity = C.float(fallVelocity)
}

func (e *EntVars) GameState() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.gamestate)
}

func (e *EntVars) SetGameState(gameState int) {
	if !e.IsValid() {
		return
	}

	e.p.gamestate = C.int(gameState)
}

func (e *EntVars) OldButtons() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.oldbuttons)
}

func (e *EntVars) SetOldButtons(oldButtons int) {
	if !e.IsValid() {
		return
	}

	e.p.oldbuttons = C.int(oldButtons)
}

func (e *EntVars) GroupInfo() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.groupinfo)
}

func (e *EntVars) SetGroupInfo(groupInfo int) {
	if !e.IsValid() {
		return
	}

	e.p.groupinfo = C.int(groupInfo)
}

func (e *EntVars) IUser1() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.iuser1)
}

func (e *EntVars) SetIUser1(iUser1 int) {
	if !e.IsValid() {
		return
	}

	e.p.iuser1 = C.int(iUser1)
}

func (e *EntVars) IUser2() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.iuser2)
}

func (e *EntVars) SetIUser2(iUser2 int) {
	if !e.IsValid() {
		return
	}

	e.p.iuser2 = C.int(iUser2)
}

func (e *EntVars) IUser3() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.iuser3)
}

func (e *EntVars) SetIUser3(iUser3 int) {
	if !e.IsValid() {
		return
	}

	e.p.iuser3 = C.int(iUser3)
}

func (e *EntVars) IUser4() int {
	if !e.IsValid() {
		return 0
	}

	return int(e.p.iuser4)
}

func (e *EntVars) SetIUser4(iUser4 int) {
	if !e.IsValid() {
		return
	}

	e.p.iuser4 = C.int(iUser4)
}

func (e *EntVars) FUser1() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.fuser1)
}

func (e *EntVars) SetFUser1(fUser1 float32) {
	if !e.IsValid() {
		return
	}

	e.p.fuser1 = C.float(fUser1)
}

func (e *EntVars) FUser2() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.fuser2)
}

func (e *EntVars) SetFUser2(fUser2 float32) {
	if !e.IsValid() {
		return
	}

	e.p.fuser2 = C.float(fUser2)
}

func (e *EntVars) FUser3() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.fuser3)
}

func (e *EntVars) SetFUser3(fUser3 float32) {
	if !e.IsValid() {
		return
	}

	e.p.fuser3 = C.float(fUser3)
}

func (e *EntVars) FUser4() float32 {
	if !e.IsValid() {
		return 0
	}

	return float32(e.p.fuser4)
}

func (e *EntVars) SetFUser4(fUser4 float32) {
	if !e.IsValid() {
		return
	}

	e.p.fuser4 = C.float(fUser4)
}

func (e *EntVars) VUser1() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.vuser1[0]),
		float32(e.p.vuser1[1]),
		float32(e.p.vuser1[2]),
	}
}

func (e *EntVars) SetVUser1(vUser1 vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.vuser1[0] = C.float(vUser1[0])
	e.p.vuser1[1] = C.float(vUser1[1])
	e.p.vuser1[2] = C.float(vUser1[2])
}

func (e *EntVars) VUser2() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.vuser2[0]),
		float32(e.p.vuser2[1]),
		float32(e.p.vuser2[2]),
	}
}

func (e *EntVars) SetVUser2(vUser2 vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.vuser2[0] = C.float(vUser2[0])
	e.p.vuser2[1] = C.float(vUser2[1])
	e.p.vuser2[2] = C.float(vUser2[2])
}

func (e *EntVars) VUser3() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.vuser3[0]),
		float32(e.p.vuser3[1]),
		float32(e.p.vuser3[2]),
	}
}

func (e *EntVars) SetVUser3(vUser3 vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.vuser3[0] = C.float(vUser3[0])
	e.p.vuser3[1] = C.float(vUser3[1])
	e.p.vuser3[2] = C.float(vUser3[2])
}

func (e *EntVars) VUser4() vector.Vector {
	if !e.IsValid() {
		return vector.Vector{}
	}

	return vector.Vector{
		float32(e.p.vuser4[0]),
		float32(e.p.vuser4[1]),
		float32(e.p.vuser4[2]),
	}
}

func (e *EntVars) SetVUser4(vUser4 vector.Vector) {
	if !e.IsValid() {
		return
	}

	e.p.vuser4[0] = C.float(vUser4[0])
	e.p.vuser4[1] = C.float(vUser4[1])
	e.p.vuser4[2] = C.float(vUser4[2])
}

func (e *EntVars) EUser1() *Edict {
	if !e.IsValid() {
		return nil
	}

	return edictFromC(e.globalVars, e.p.euser1)
}

func (e *EntVars) SetEUser1(eUser1 *Edict) {
	if !e.IsValid() {
		return
	}

	e.p.euser1 = eUser1.p
}

func (e *EntVars) EUser2() *Edict {
	if !e.IsValid() {
		return nil
	}

	return edictFromC(e.globalVars, e.p.euser2)
}

func (e *EntVars) SetEUser2(eUser2 *Edict) {
	if !e.IsValid() {
		return
	}

	e.p.euser2 = eUser2.p
}

func (e *EntVars) EUser3() *Edict {
	if !e.IsValid() {
		return nil
	}

	return edictFromC(e.globalVars, e.p.euser3)
}

func (e *EntVars) SetEUser3(eUser3 *Edict) {
	if !e.IsValid() {
		return
	}

	e.p.euser3 = eUser3.p
}

func (e *EntVars) EUser4() *Edict {
	if !e.IsValid() {
		return nil
	}

	return edictFromC(e.globalVars, e.p.euser4)
}

func (e *EntVars) SetEUser4(eUser4 *Edict) {
	if !e.IsValid() {
		return
	}

	e.p.euser4 = eUser4.p
}

type TraceResult struct {
	AllSolid    bool          // if true, plane is not valid
	StartSolid  bool          // if true, the initial point was in a solid area
	InOpen      bool          // if true, the initial point was in empty space
	InWater     bool          // if true, the initial point was underwater
	Fraction    float32       // time completed, 1.0 = didn't hit anything
	EndPos      vector.Vector // final position
	PlaneDist   float32       // distance from the plane
	PlaneNormal vector.Vector // surface normal at impact
	Hit         *Edict        // entity the surface is on
	HitGroup    int           // 0 == generic, non-zero is specific body part
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
		result.EndPos = vector.Vector{
			float32(tr.vecEndPos[0]),
			float32(tr.vecEndPos[1]),
			float32(tr.vecEndPos[2]),
		}
	}

	if len(tr.vecPlaneNormal) == 3 {
		result.PlaneNormal = vector.Vector{
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

func (tr *TraceResult) ToC() *C.TraceResult {
	var result C.TraceResult

	if tr.AllSolid {
		result.fAllSolid = 1
	}

	if tr.StartSolid {
		result.fStartSolid = 1
	}

	if tr.InOpen {
		result.fInOpen = 1
	}

	if tr.InWater {
		result.fInWater = 1
	}

	result.flFraction = C.float(tr.Fraction)
	result.flPlaneDist = C.float(tr.PlaneDist)

	for i := 0; i < 3; i++ {
		result.vecEndPos[i] = C.float(tr.EndPos[i])
		result.vecPlaneNormal[i] = C.float(tr.PlaneNormal[i])
	}

	if tr.Hit != nil {
		result.pHit = tr.Hit.p
	}

	return &result
}

type Texture struct {
	p *C.texture_t
}

func textureFromC(t *C.texture_t) *Texture {
	return &Texture{
		p: t,
	}
}

func (t *Texture) Name() string {
	return C.GoString(&t.p.name[0])
}

func (t *Texture) Width() uint32 {
	return uint32(t.p.width)
}

func (t *Texture) Height() uint32 {
	return uint32(t.p.height)
}

func (t *Texture) AnimTotal() int {
	return int(t.p.anim_total)
}

func (t *Texture) AnimMin() int {
	return int(t.p.anim_min)
}

func (t *Texture) AnimMax() int {
	return int(t.p.anim_max)
}

func (t *Texture) AnimNext() *Texture {
	return textureFromC(t.p.anim_next)
}

func (t *Texture) AlternateAnims() *Texture {
	return textureFromC(t.p.alternate_anims)
}

func (t *Texture) PalOffset() uint32 {
	return uint32(t.p.paloffset)
}

func (t *Texture) toC() *C.texture_t {
	return t.p
}

type InfoBuffer struct {
	p *C.char
}

func infoBufferFromC(p *C.char) *InfoBuffer {
	return &InfoBuffer{
		p: p,
	}
}

func (ib *InfoBuffer) String() string {
	return C.GoString(ib.p)
}
