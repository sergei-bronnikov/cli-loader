package simplespinner

import (
	"fmt"
	"time"
)

const DefaultDelay = 300 * time.Millisecond

type SimpleSpinner struct {
	opts   Opts
	stopCh chan struct{}
}

type Opts struct {
	Delay   time.Duration
	Spinner []string
	Prefix  string
	Suffix  string
}

var DefaultOpts = Opts{
	Delay:   DefaultDelay,
	Spinner: Spinners.Default,
	Prefix:  "",
	Suffix:  "",
}

var Spinners = struct {
	Default []string
	Dots    []string
}{
	Default: []string{"|", "/", "-", "\\"},
	Dots:    []string{".", "..", "...", "....", "....."},
}

func New(opts Opts) *SimpleSpinner {
	return &SimpleSpinner{
		opts:   opts,
		stopCh: make(chan struct{}),
	}
}

func (s *SimpleSpinner) Start() {
	go s.doSpin()
}

func (s *SimpleSpinner) doSpin() {
	fmt.Print("\033[?25l")
	for {
		for _, char := range s.opts.Spinner {
			select {
			case <-s.stopCh:
				fmt.Print("\033[2K\r", s.opts.Prefix, "Done!", s.opts.Suffix, "\n")
				fmt.Print("\033[?25h")
				return
			default:
				fmt.Print("\033[2K\r", s.opts.Prefix, char, s.opts.Suffix)
				time.Sleep(s.opts.Delay)
			}
		}
	}
}

func (s *SimpleSpinner) Stop() {
	s.stopCh <- struct{}{}
}
