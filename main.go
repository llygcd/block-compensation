package main

import (
	"github.com/llygcd/block-compensation/cmd"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	cmd.Execute()
}
