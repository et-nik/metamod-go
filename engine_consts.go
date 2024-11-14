package main

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
	FindEntityFieldViewmodel   FindEntityField = "viewmodel"
	FindEntityFieldWeaponmodel FindEntityField = "weaponmodel"
	FindEntityFieldNetname     FindEntityField = "netname"
	FindEntityFieldTarget      FindEntityField = "target"
	FindEntityFieldTargetname  FindEntityField = "targetname"
	FindEntityFieldMessage     FindEntityField = "message"
	FindEntityFieldNoise       FindEntityField = "noise"
	FindEntityFieldNoise1      FindEntityField = "noise1"
	FindEntityFieldNoise2      FindEntityField = "noise2"
	FindEntityFieldNoise3      FindEntityField = "noise3"
	FindEntityFieldGlobalname  FindEntityField = "globalname"
)
