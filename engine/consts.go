package engine

const (
	MaxEntLeafs = 48

	mipLevels = 4
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

type MoveType int

const (
	MoveTypeNone MoveType = iota
	MoveTypeAngleNoClip
	MoveTypeAngle
	MoveTypeWalk
	MoveTypeStep
	MoveTypeFly
	MoveTypeToss
	MoveTypePush
	MoveTypeNoclip
	MoveTypeFlyMissile
	MoveTypeBounce
	MoveTypeBounceMissile
	MoveTypeFollow
	MoveTypePushStep
)

// SolidType Some movetypes will cause collisions independent of SOLID_NOT/SOLID_TRIGGER when the entity moves
// SOLID only effects OTHER entities colliding with this one when they move - UGH!
type SolidType int

const (
	SolidTypeNot         SolidType = iota // no interaction with other objects
	SolidTypeTrigger                      // touch on edge, but not blocking
	SolidTypeBoundingBox                  // touch on edge, block
	SolidTypeBsp                          // bsp clip, touch on edge, block
)

type WalkMoveMode int

const (
	WalkMoveNormal WalkMoveMode = iota
	WalkMoveWorldOnly
	WalkMoveCheckOnly
)

type FindEntityField string

const (
	FindEntityFieldClassname   FindEntityField = "classname"
	FindEntityFieldModel       FindEntityField = "model"
	FindEntityFieldViewModel   FindEntityField = "viewmodel"
	FindEntityFieldWeaponModel FindEntityField = "weaponmodel"
	FindEntityFieldNetName     FindEntityField = "netname"
	FindEntityFieldTarget      FindEntityField = "target"
	FindEntityFieldTargetName  FindEntityField = "targetname"
	FindEntityFieldMessage     FindEntityField = "message"
	FindEntityFieldNoise       FindEntityField = "noise"
	FindEntityFieldNoise1      FindEntityField = "noise1"
	FindEntityFieldNoise2      FindEntityField = "noise2"
	FindEntityFieldNoise3      FindEntityField = "noise3"
	FindEntityFieldGlobalName  FindEntityField = "globalname"
)

type InButtonFlag int

const (
	InButtonAttack InButtonFlag = 1 << iota
	InButtonJump
	InButtonDuck
	InButtonForward
	InButtonBack
	InButtonUse
	InButtonCancel
	InButtonLeft
	InButtonRight
	InButtonMoveLeft
	InButtonMoveRight
	InButtonAttack2
	InButtonRun
	InButtonReload
	InButtonAlt1
	InButtonScore
)

type DeadFlag int

const (
	DeadFlagNo DeadFlag = iota
	DeadFlagDying
	DeadFlagDead
	DeadFlagRespawnable
	DeadFlagDiscardBody
)

// TraceHullType used by TraceHull
const (
	TraceHullPoint = 0
	TraceHullHuman = 1
	TraceHullLarge = 2
	TraceHullHead  = 3
)

const (
	TraceDontIgnoreMonsters = 0
	TraceIgnoreMonsters     = 1
	TraceMissile            = 2
)

const (
	MessageDestBroadcast     = iota // unreliable to all
	MessageDestOne                  // reliable to one (msg_entity)
	MessageDestAll                  // reliable to all
	MessageDestInit                 // write to the init string
	MessageDestPVS                  // Ents in PVS of org
	MessageDestPAS                  // Ents in PAS of org
	MessageDestPVSR                 // Reliable to PVS
	MessageDestPASR                 // Reliable to PAS
	MessageDestOneUnreliable        // Send to one client, but don't put in reliable stream, put in unreliable datagram ( could be dropped )
	MessageDestSpec                 // Sends to all spectator proxies
)

const (
	SvcTempEntity   = 23
	SvcIntermission = 30
	SvcCdTrack      = 32
	SvcWeaponAnim   = 35
	SvcRoomType     = 37
	SvcHLTV         = 50
)
