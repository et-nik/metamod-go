# [Metamod-Go](https://github.com/et-nik/metamod-go)

**Metamod-Go** is golang library for creating Metamod plugins for Half-Life 1 mods.

## Active development notice

This project is currently in active development and is not ready for production use.
Many features are missing and the API is subject to change.


## Build using golang:1.23.3-bookworm:

```
dpkg --add-architecture i386
apt-get update
apt-get install gcc-multilib

CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -o experiments.so -buildvcs=false -buildmode=c-shared

```