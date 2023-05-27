package ProcessScheduler

import (
	"fyne.io/fyne/v2/widget"
	"strings"
	"sync"
)

type ProcessScheduler struct {
	// Vars
	sourceDir, targetDir string
	progressbar          *widget.ProgressBar
	in                   chan string
	out                  chan bool
	worker               int
	numFiles             int
}

func New(sourceDir, targetDir string, progressbar *widget.ProgressBar) *ProcessScheduler {
	// Vars
	ps := ProcessScheduler{}
	ps.sourceDir = sourceDir
	ps.targetDir = targetDir
	ps.progressbar = progressbar
	ps.in = make(chan string)
	ps.out = make(chan bool)
	// TODO: Create function that will check how many files to be converted
	return &ps
}

// Start will start the ProcessScheduler
func (ps *ProcessScheduler) Start() {
	// Vars
	wg := sync.WaitGroup{}
	// Start the Workers
	for i := 0; i < ps.worker; i++ {
		go ps.Worker()
	}

	// There is only one Process that will catch the callsbacks from the Workers to increade the progressbar
	wg.Add(1)
	go ps.CallBack(&wg)
	wg.Wait()

}

func (ps *ProcessScheduler) Worker() {
	// Vars
	var file string
	// Loop
	for {
		// Get the file
		file = <-ps.in
		// Check if the file is HEIC
		if ps.isHEIC(file) {
			// Convert the file
			//ps.convert(file)
		}
	}
}

// CallBack will be called from the Workers to increase the progressbar
func (ps *ProcessScheduler) CallBack(wg *sync.WaitGroup) {
	files := 0
	for _ = range ps.out {
		// Increase the progressbar
		ps.progressbar.SetValue(ps.progressbar.Value + 1)
		files += 1
		// Check if all files are converted
		if files == ps.numFiles {
			// Close the in channel
			close(ps.in)
			// Close the out channel
			close(ps.out)
			// Break the loop
			break
		}
	}
	wg.Done()
}

// isHEIC if file is HEIC
func (ps *ProcessScheduler) isHEIC(file string) bool {
	if strings.ToLower(file[len(file)-5:]) == ".heic" {
		return true
	}
	return false
}
