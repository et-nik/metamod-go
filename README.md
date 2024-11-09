

# Build using golang:1.23.3-bookworm:

```
dpkg --add-architecture i386
apt-get update
apt-get install gcc-multilib

CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -o experiments.so -buildvcs=false -buildmode=c-shared

```