package WindowManager

// This Package will create a window for the Application. It uses fyne.io/fyne/v2

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type WindowManager struct {
}

var WM WindowManager

// init initializes the WindowManager
func init() {
	WM = WindowManager{}
}

// CreateWindow creates a new Window
func (wm *WindowManager) CreateWindow() {
	// Create the Application
	a := app.New()

	// Set the Title of the Window
	w := a.NewWindow("HEIC2JPEG - a HEIC to JPEG Converter")

	// set size to 300 x 200
	w.Resize(fyne.NewSize(400, 170))

	// The app may not be resized
	w.SetFixedSize(true)

	// Create Labels and Buttons
	w.SetContent(createVBOX())

	// Center the window on the Screen
	w.CenterOnScreen()
	// Run the Application
	w.ShowAndRun()
}

// createVBOX creates a new VBOX
func createVBOX() *fyne.Container {
	// Create grid layout
	vbox := container.New(layout.NewVBoxLayout())

	// Create a Label with the Text "Choose a HEIC File to convert"
	label := widget.NewLabel("Choose a HEIC File to convert")

	// Add the label to the Window
	vbox.Add(label)

	// Under this label add the First Button to choose a HEIC File
	btn1 := widget.NewButton("Choose a Folder with HEIC Files", nil)

	// Add it to the VBOX
	vbox.Add(btn1)

	// Create a Label with the Text "Choose a HEIC File to convert"
	label2 := widget.NewLabel("Choose a Destination Folder")

	// Add the label to the Window
	vbox.Add(label2)

	// Under this label add the First Button to choose a HEIC File
	btn2 := widget.NewButton("Choose a Folder", nil)

	// Add it to the VBOX
	vbox.Add(btn2)
	return vbox
}
