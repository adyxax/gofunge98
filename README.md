# GoFunge98 : a Funge-98 interpreter written in go

This repository contains code for a go program that can interpret a valid [Funge-98](https://github.com/catseye/Funge-98/blob/master/doc/funge98.markdown) program. It passes the [mycology test suite](https://github.com/Deewiant/Mycology).

Current limitations are :
- currently does not implement any fingerprints
- does not implement concurrent execution with the `t` command
- does not implement file I/O with the `i` and `o` commands
- does not implement system execution with the `=` command

## Contents

- [Dependencies](#dependencies)
- [Quick install](#quick-install)
- [Usage](#usage)
- [Building](#building)

## Dependencies

go is required. Only go version >= 1.20 on linux amd64 (Gentoo) is being regularly tested.

## Quick Install

```
go install git.adyxax.org/adyxax/gofunge98/cmd/gofunge98@latest
```

## Usage

Launching the gofunge98 is as simple as :
```
gofunge98 -f something.b98
```

The interpreter will then load and execute the specified Funge-98 program until the program normally terminates or is interrupted or killed.

## Building

To run tests, use :
```
go test -cover ./...
```

For a debug build, use :
```
go build ./cmd/gofunge98/
```

For a release build, use :
```
go build -ldflags '-s -w -extldflags "-static"' ./cmd/gofunge98/
```
