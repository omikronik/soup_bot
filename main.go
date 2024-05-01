package main

import (
	"soup_bot/cmd"
)

func main() {
	cmd.Start()

	<-make(chan struct{})
}
