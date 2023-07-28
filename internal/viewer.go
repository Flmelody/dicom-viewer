package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
)

type Viewer struct {
	win         fyne.Window
	dicom       *Dicom
	canvasImage *canvas.Image
}

func SetupViewer(a fyne.App) *Viewer {
	w := a.NewWindow("dicom viewer")
	dicom := NewDicom(false, nil, nil, 40, 400)

	canvasImage := canvas.NewImageFromImage(dicom)
	viewer := &Viewer{
		win:         w,
		dicom:       dicom,
		canvasImage: canvasImage,
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

	split := container.NewHSplit(layout.NewSpacer(), layout.NewSpacer())
	split.Offset = 0.25

	w.SetContent(split)
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	return viewer
}

func (viewer *Viewer) StartViewer() {
	viewer.win.ShowAndRun()
}
