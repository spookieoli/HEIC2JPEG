package ProcessScheduler

import (
	"fyne.io/fyne/v2/widget"
	"strings"
)

type ProcessScheduler struct {
	// Vars
	sourceDir, targetDir string
	progressbar          *widget.ProgressBar
	in                   chan string
	out                  chan bool
}

func New(sourceDir, targetDir string, progressbar *widget.ProgressBar) *ProcessScheduler {
	// Vars
	ps := ProcessScheduler{}
	ps.sourceDir = sourceDir
	ps.targetDir = targetDir
	ps.progressbar = progressbar
	ps.in = make(chan string)
	ps.out = make(chan bool)
	return &ps
}

// Check if file is HEIC
func (ps *ProcessScheduler) isHEIC(file string) bool {
	if strings.ToLower(file[len(file)-5:]) == ".heic" {
		return true
	}
	return false
}
