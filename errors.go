package metamod_go

import "errors"

var (
	ErrLibIsLoaded = errors.New("library is already loaded, you must set the value before the library is loaded")
	ErrMetaQueried = errors.New("meta is already queried, you must set the value before the Meta_Query function is called")
)
