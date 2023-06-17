package main

import "HEIC2JPEG/WindowManager"

func main() {
	wm := WindowManager.WM
	wm.CreateWindow()
}

// Compile with: go build -ldflags="-H windowsgui" -o HEIC2JPEG.exe
