package progressbar

import (
	"errors"
	"fmt"
)

type ProgressBar struct {
	opts         Opts
	progressChan chan ProgressState
}

type Opts struct {
	Prefix string
}

type ProgressState struct {
	Progress uint8
	Done     bool
	//Failed   bool
	Msg string
}

func New(opts Opts) *ProgressBar {
	hideCursor()
	return &ProgressBar{}
}

func (p *ProgressBar) Update(state ProgressState) error {
	if state.Progress > 100 {
		return errors.New("progress out of range")
	}
	progressBar := fmt.Sprintf("%s %s %d%", p.opts.Prefix, state.Msg, state.Progress)
	clearLine()
	fmt.Print(progressBar)
	if state.Done {
		fmt.Print("\n")
		showCursor()
	}
	return nil
	//p.progressChan <- state
	//return nil
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

func clearLine() {
	fmt.Print("\033[2K\r")
}
