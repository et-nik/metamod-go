# [Metamod-Go](https://github.com/et-nik/metamod-go)

**Metamod-Go** is golang library for creating [Metamod](https://github.com/rehlds/metamod-r) plugins for GoldSource servers
(Counter-Strike 1.6, Half-Life, Team Fortress Classic, etc.).

Main purpose of this library is to provide easy way to create Metamod plugins using Go simplicity and power.
Second purpose is to avoid of legacy practices of Metamod plugins development.
Plugins written in Go can have all benefits of Go language, such as static typing, garbage collection, concurrency, etc.
Be aware of using concurrency in plugins, because GoldSource engine is single-threaded.

Plugins written in Go can be slower than plugins written in C++, because of Go runtime overhead.

## Active development notice

This project is currently in active development and is not ready for production use.
Many features are missing and the API is subject to change.

Some features did not test completely and may not work as expected.

### Known issues

#### `AddServerCommand` does not work with variable which is not in global scope. Example

**Does not work:**
```go
func MetaAttach(_ int) int {
	engineFuncs := metamod.GetEngineFuncs()
	
    engineFuncs.AddServerCommand("entinfo", func(argc int, argv ...string) {
        engineFuncs.ServerPrint("entinfo\n")
    }
}
```

**Works:**
```go

var engineFuncs *metamod.EngineFuncs

func MetaAttach(_ int) int {
    engineFuncs = metamod.GetEngineFuncs()
	
    engineFuncs.AddServerCommand("entinfo", func(argc int, argv ...string) {
        engineFuncs.ServerPrint("entinfo\n")
    }
}
```

Move vars which are used in `AddServerCommand` function to global scope.

## Plugins and Examples

List of plugins and examples:

* https://github.com/et-nik/metamod-go-example