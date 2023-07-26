package main

import (
	"dicom-viewer/internal"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&internal.DicomTheme{})
	viewer := internal.SetupViewer(a)
	viewer.StartViewer()
}
