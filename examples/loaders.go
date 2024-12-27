package main

import (
	"github.com/sergei-bronnikov/cli-loader/simple-spinner"
	"time"
)

func main() {
	spinner := simplespinner.New(simplespinner.Opts{
		Delay:   simplespinner.DefaultDelay,
		Spinner: simplespinner.Spinners.Dots,
		Prefix:  "In progress: ",
		Suffix:  "",
	})
	spinner.Start()
	time.Sleep(5 * time.Second)
	spinner.Stop()
}
