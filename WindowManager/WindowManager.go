package WindowManager

// This Package will create a window for the Application. It uses fyne.io/fyne/v2

import (
	"HEIC2JPEG/ConfigurationManager"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type WindowManager struct {
	w  fyne.Window
	cm ConfigurationManager.ConfigurationManager
}

var WM WindowManager

// init initializes the WindowManager
func init() {
	WM = WindowManager{}
	WM.cm = *ConfigurationManager.New()
}

// CreateWindow creates a new Window
func (wm *WindowManager) CreateWindow() {
	// Create the Application
	a := app.New()

	// Set the Title of the Window
	wm.w = a.NewWindow("HEIC2JPEG - a HEIC to JPEG Converter")

	// set size to 300 x 200
	wm.setNormalSize()

	// The app may not be resized
	wm.w.SetFixedSize(true)

	// Create Labels and Buttons
	wm.w.SetContent(wm.createVBOX())

	// Add a menubar with a Data and a About Menu
	wm.w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("Data",
			fyne.NewMenuItem("Options", nil), // Will open a new Window with Options
		),
		fyne.NewMenu("About",
			fyne.NewMenuItem("About HEIC2JPEG", wm.AboutWindow), // Will open a new Window with Information about the App
		),
	))

	// argh fyne cannot Set the position of the window - this is bad... all you can do is center it
	wm.w.CenterOnScreen()

	// Run the Application
	wm.w.ShowAndRun()
}

// createVBOX creates a new VBOX
func (wm *WindowManager) createVBOX() *fyne.Container {
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

	// The ConvertButton
	convertButton := widget.NewButton("Convert", nil)

	// Add it to the VBOX
	vbox.Add(btn2)

	// Create a Progressbar
	progressBar := widget.NewProgressBar()

	// Add it to the VBOX
	vbox.Add(progressBar)

	// Add this Button to the VBOX
	vbox.Add(convertButton)

	return vbox
}

// AboutWindow will open a Popup with Information about the App
func (wm *WindowManager) AboutWindow() {
	dialog.ShowInformation("About HEIC2JPEG",
		"HEIC2JPEG is a HEIC to JPEG Converter. \nIt is written in Go and uses the fyne.io/fyne/v2 Framework. \nIt is licensed under the MIT License.", wm.w)
}

// setNormalSize sets the Window to the normal size
func (wm *WindowManager) setNormalSize() {
	wm.w.Resize(fyne.NewSize(400, 270))
}
