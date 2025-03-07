package metamod_go

import "errors"

var (
	ErrLibNotLoaded      = errors.New("library is not loaded")
	ErrLibIsLoaded       = errors.New("library is already loaded, you must set the value before the library is loaded")
	ErrMetaQueried       = errors.New("meta is already queried, you must set the value before the Meta_Query function is called")
	ErrMetaIsNotAttached = errors.New("meta is not attached, you must get the value after the Meta_Attach function is called")
)
