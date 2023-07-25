package main

import (
	"dicom-viewer/internal"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&internal.DicomTheme{})
	w := a.NewWindow("")

	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("open", func() {
			})),
		fyne.NewMenu("About",
			fyne.NewMenuItem("version", func() {
			})),
	))

	split := container.NewHSplit(widget.NewLabel("Hello"), widget.NewLabel("World"))
	split.Offset = 0.25

	w.SetContent(split)
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	w.ShowAndRun()
}
