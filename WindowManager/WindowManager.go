package WindowManager

// This Package will create a window for the Application. It uses fyne.io/fyne/v2

import (
	"HEIC2JPEG/ConfigurationManager"
	"HEIC2JPEG/ProcessScheduler"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strings"
)

type WindowManager struct {
	w              fyne.Window
	cm             ConfigurationManager.ConfigurationManager
	sourceDir      string
	targetDir      string
	label1, label2 *widget.Label
	progressbar    *widget.ProgressBar
	convertButton  *widget.Button
	Errtxt         string
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

	// Add a menubar with a Data and an About Menu
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

	// Set the Labels to the last used values
	wm.label1.SetText(wm.cm.Config.SourceDir)
	wm.label2.SetText(wm.cm.Config.TargetDir)
	wm.sourceDir = wm.cm.Config.SourceDir
	wm.targetDir = wm.cm.Config.TargetDir

	// Run the Application
	wm.w.ShowAndRun()

}

// createVBOX creates a new VBOX
func (wm *WindowManager) createVBOX() *fyne.Container {
	// Create grid layout
	vbox := container.New(layout.NewVBoxLayout())

	// Create a Label with the Text "Choose a HEIC File to convert"
	wm.label1 = widget.NewLabel("Choose a HEIC File to convert")

	// Add the label to the Window
	vbox.Add(wm.label1)

	// Under this label add the First Button to choose a HEIC File
	btn1 := widget.NewButton("Choose a Folder with HEIC Files", wm.OpenSourceFileDiaglog)

	// Add it to the VBOX
	vbox.Add(btn1)

	// Create a Label with the Text "Choose a HEIC File to convert"
	wm.label2 = widget.NewLabel("Choose a Destination Folder")

	// Add the label to the Window
	vbox.Add(wm.label2)

	// Under this label add the First Button to choose a HEIC File
	btn2 := widget.NewButton("Choose a Folder", wm.OpenTargetFileDiaglog)

	// The ConvertButton
	wm.convertButton = widget.NewButton("Convert", wm.Convert)

	// Add it to the VBOX
	vbox.Add(btn2)

	// Create a Progressbar
	wm.progressbar = widget.NewProgressBar()

	// Add it to the VBOX
	vbox.Add(wm.progressbar)

	// Add this Button to the VBOX
	vbox.Add(wm.convertButton)

	return vbox
}

// AboutWindow will open a Popup with Information about the App
func (wm *WindowManager) AboutWindow() {
	dialog.ShowInformation("About HEIC2JPEG",
		"HEIC2JPEG is a HEIC to JPEG Converter. \nIt is written in Go and uses the fyne.io/fyne/v2 Framework. \nIt is licensed under the MIT License.", wm.w)
}

// DoneWindow will show a Popup when the Conversion is done
func (wm *WindowManager) DoneWindow() {
	dialog.ShowInformation("Done",
		"Conversion is done.", wm.w)
}

// ErrorViwe will show a Popup when an Error occurs
func (wm *WindowManager) ErrorView() {
	dialog.ShowInformation("Error",
		"An Error occured: "+wm.Errtxt, wm.w)
}

// setNormalSize sets the Window to the normal size
func (wm *WindowManager) setNormalSize() {
	wm.w.SetFixedSize(false)
	wm.w.Resize(fyne.NewSize(400, 270))
	wm.w.Content().Refresh()
	wm.w.SetFixedSize(true)
}

// setFileDialogSize sets the Window to the size of the File Dialog
func (wm *WindowManager) setFileDialogSize() {
	wm.w.SetFixedSize(false)
	wm.w.Resize(fyne.NewSize(600, 500))
	wm.w.Content().Refresh()
	wm.w.SetFixedSize(true)
}

// OpenSourceFileDiaglog opens a File Dialog and saves the selected Directory wm.sourceDir
func (wm *WindowManager) OpenSourceFileDiaglog() {
	wm.setFileDialogSize()
	dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
		if err == nil && list != nil {
			// remove file:// from the string
			wm.sourceDir = strings.Replace(list.String(), "file://", "", 1)
			wm.cm.Config.SourceDir = wm.sourceDir
			// If the sourcedir path is to long cut it to only 30 chars
			if len(wm.sourceDir) > 30 {
				wm.sourceDir = wm.sourceDir[len(wm.sourceDir)-30:]
			}
			wm.label1.SetText("... " + wm.sourceDir)
			wm.w.Content().Refresh()
		}
		wm.setNormalSize()
	}, wm.w)
}

// OpenTargetFileDiaglog opens a File Dialog and saves the selected Directory wm.targetDir
func (wm *WindowManager) OpenTargetFileDiaglog() {
	wm.setFileDialogSize()
	dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
		if err == nil && list != nil {
			// remove file:// from the string
			wm.targetDir = strings.Replace(list.String(), "file://", "", 1)
			wm.cm.Config.TargetDir = wm.targetDir
			if len(wm.sourceDir) > 30 {
				wm.targetDir = wm.targetDir[len(wm.targetDir)-30:]
			}
			wm.label2.SetText("..." + wm.targetDir)
			wm.w.Content().Refresh()
		}
		wm.setNormalSize()
	}, wm.w)
}

// Convert will start the conversion
func (wm *WindowManager) Convert() {
	if wm.sourceDir == "" || wm.targetDir == "" {
		wm.Errtxt = "Please choose a Source and a Target Folder"
		wm.ErrorView()
		return
	}
	wm.cm.WriteConfiguration()
	// Set the Progressbar to 0
	wm.progressbar.SetValue(0)
	// Create the ProcessScheduler
	ps := ProcessScheduler.New(wm.sourceDir, wm.targetDir, wm.progressbar, wm.cm.Config.Worker)
	// deactivate the buttons
	wm.convertButton.Disable()
	// Start the ProcessScheduler
	ps.Start()
	// after the ProcessScheduler has finished, activate the buttons
	wm.convertButton.Enable()
	wm.DoneWindow()
}
