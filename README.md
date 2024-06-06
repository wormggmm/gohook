# gohook

[![Build Status](https://github.com/wormggmm/gohook/workflows/Go/badge.svg)](https://github.com/wormggmm/gohook/commits/master)
[![CircleCI Status](https://circleci.com/gh/wormggmm/gohook.svg?style=shield)](https://circleci.com/gh/wormggmm/gohook)
![Appveyor](https://ci.appveyor.com/api/projects/status/github/wormggmm/gohook?branch=master&svg=true)
[![Go Report Card](https://goreportcard.com/badge/github.com/wormggmm/gohook)](https://goreportcard.com/report/github.com/wormggmm/gohook)
[![GoDoc](https://godoc.org/github.com/wormggmm/gohook?status.svg)](https://godoc.org/github.com/wormggmm/gohook)
<!-- This is a work in progress. -->

## Requirements (Linux):

[Robotgo-requirements-event](https://github.com/go-vgo/robotgo#requirements)

## Install:

With Go module support (Go 1.11+), just import:

```go
import "github.com/wormggmm/gohook"
```

## Examples:

```Go
package main

import (
	"fmt"

	hook "github.com/wormggmm/gohook"
)

func main() {
	add()

	low()
}

func add() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press w---")
	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("w")
	})

	s := hook.Start()
	<-hook.Process(s)
}

func low() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("hook: ", ev)
	}
}

```

Based on [libuiohook](https://github.com/kwhat/libuiohook).