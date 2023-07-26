package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type Viewer struct {
	win fyne.Window
}

func SetupViewer(a fyne.App) *Viewer {
	w := a.NewWindow("dicom viewer")

	viewer := &Viewer{
		win: w,
	}
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("open file", func() {
				d := dialog.NewFileOpen(func(f fyne.URIReadCloser, err error) {
					if f == nil || err != nil {
						return
					}
					// todo
				}, w)
				d.SetFilter(storage.NewExtensionFileFilter([]string{".dcm"}))
				d.Show()
			}),
			fyne.NewMenuItem("open folder", func() {
				d := dialog.NewFolderOpen(func(f fyne.ListableURI, err error) {
					if f == nil || err != nil {
						return
					}
					// todo
				}, w)
				d.Show()
			})),
		fyne.NewMenu("About",
			fyne.NewMenuItem("version", func() {
				d := dialog.NewInformation("dicom viewer", "version 0.1", w)
				d.Show()
			})),
	))

	split := container.NewHSplit(widget.NewLabel("Hello"), widget.NewLabel("World"))
	split.Offset = 0.25

	w.SetContent(split)
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	return viewer
}

func (viewer *Viewer) StartViewer() {
	viewer.win.ShowAndRun()
}
