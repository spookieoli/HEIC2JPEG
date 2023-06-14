package ProcessScheduler

import (
	utils2 "HEIC2JPEG/utils"
	"fmt"
	"fyne.io/fyne/v2/widget"
	"github.com/jdeng/goheif"
	_ "github.com/jdeng/goheif"
	"image/jpeg"
	"os"
	"strings"
	"sync"
)

// The ProcessScheduler will schedule the conversion of the HEIC Files to JPEG Files
type ProcessScheduler struct {
	// Vars
	sourceDir, targetDir string
	progressbar          *widget.ProgressBar
	in                   chan string
	out                  chan bool
	worker               int
	numFiles             int
	Files                []string
	errors               int
}

func New(sourceDir, targetDir string, progressbar *widget.ProgressBar, worker int) *ProcessScheduler {
	// Vars
	ps := ProcessScheduler{}
	ps.sourceDir = sourceDir
	ps.targetDir = targetDir
	ps.progressbar = progressbar
	ps.in = make(chan string)
	ps.out = make(chan bool)
	ps.worker = worker
	ps.countHEICFiles()
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
	// There is only one Process that will catch the callbacks from the Workers to increase the progressbar
	wg.Add(1)
	go ps.CallBack(&wg)

	// Add all Files to the in channel
	for _, file := range ps.Files {
		ps.in <- file
	}
	// Wait for the final Worker to finish
	wg.Wait()
}

func (ps *ProcessScheduler) Worker() {
	for file := range ps.in {
		// Get the file

		// Check if the file is HEIC
		if ps.isHEIC(file) {
			// Convert the file
			ps.convert(file)
			// Increase the progressbar
			ps.out <- true
		}
	}
}

// CallBack will be called from the Workers to increase the progressbar
func (ps *ProcessScheduler) CallBack(wg *sync.WaitGroup) {
	files := 0
	for {
		// Increase the progressbar
		b := <-ps.out
		if b == true {
			ps.progressbar.SetValue(ps.progressbar.Value + 1)
		} else {
			ps.errors += 1
		}
		files += 1
		// Check if all files are converted
		if files == ps.numFiles {
			// Close the in channel, this will break the loop in the Workers and stop them
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
	if len(file) > 5 && strings.ToLower(file[len(file)-5:]) == ".heic" {
		return true
	}
	return false
}

// countHEICFiles will count the heicfiles in the given directory
func (ps *ProcessScheduler) countHEICFiles() {
	// open the sourcedir
	dirEntry, err := os.ReadDir(ps.sourceDir)
	if err != nil {
		panic(err)
	}
	// Loop through the files
	for _, file := range dirEntry {
		// Check if the file is HEIC
		if ps.isHEIC(file.Name()) {
			ps.numFiles += 1
			// Add the complete Filepath to the Files
			ps.Files = append(ps.Files, ps.sourceDir+"/"+file.Name())
		}
	}
}

// This Function will convert the given file
func (ps *ProcessScheduler) convert(file string) {
	// open the given file
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// Check if there is a EXIF in the file
	exif, err := goheif.ExtractExif(f)
	if err != nil {
		panic(err)
	}
	// Decode the HEIC
	img, err := goheif.Decode(f)
	if err != nil {
		panic(err)
	}
	// Create the Output writer
	o, err := os.OpenFile(ps.targetDir+"/"+file[len(ps.sourceDir)+1:len(file)-5]+".jpg", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer o.Close()

	// Create new utils Object
	utils, err := utils2.New(o, exif)
	// jpeg encode the image
	err = jpeg.Encode(utils, img, nil)
	if err != nil {
		panic(err)
	}
}
