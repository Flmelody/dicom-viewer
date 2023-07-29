package internal

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
	"strconv"
)

type Viewer struct {
	win                                              fyne.Window
	dicom                                            *Dicom
	canvasImage                                      *canvas.Image
	canvasID, canvasName, canvasSex, canvasBirthDate *widget.Label
}

func SetupViewer(a fyne.App) *Viewer {
	w := a.NewWindow("dicom viewer")
	dicomFile := NewDicom(false, nil, nil, 40, 400)

	canvasImage := canvas.NewImageFromImage(dicomFile)
	viewer := &Viewer{
		win:         w,
		dicom:       dicomFile,
		canvasImage: canvasImage,
	}
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("open file", func() {
				d := dialog.NewFileOpen(func(f fyne.URIReadCloser, err error) {
					if f == nil || err != nil {
						return
					}
					viewer.loadFile(f.URI().Path())
				}, w)
				d.SetFilter(storage.NewExtensionFileFilter([]string{".dcm"}))
				d.Resize(fyne.NewSize(1000, 800))
				d.Show()
			}),
			fyne.NewMenuItem("open folder", func() {
				d := dialog.NewFolderOpen(func(f fyne.ListableURI, err error) {
					if f == nil || err != nil {
						return
					}
					// todo
				}, w)
				d.Resize(fyne.NewSize(1000, 800))
				d.Show()
			})),
		fyne.NewMenu("About",
			fyne.NewMenuItem("version", func() {
				d := dialog.NewInformation("dicom viewer", "version 0.1", w)
				d.Show()
			})),
	))

	viewer.canvasID = widget.NewLabelWithStyle("none", fyne.TextAlignLeading, fyne.TextStyle{})
	viewer.canvasName = widget.NewLabelWithStyle("none", fyne.TextAlignLeading, fyne.TextStyle{})
	viewer.canvasSex = widget.NewLabelWithStyle("none", fyne.TextAlignLeading, fyne.TextStyle{})
	viewer.canvasBirthDate = widget.NewLabelWithStyle("none", fyne.TextAlignLeading, fyne.TextStyle{})
	wf := widget.NewForm()
	wf.Append("PatientID", viewer.canvasID)
	wf.Append("PatientName", viewer.canvasName)
	wf.Append("PatientSex", viewer.canvasSex)
	wf.Append("PatientBirthDate", viewer.canvasBirthDate)
	split := container.NewHSplit(wf, canvasImage)
	split.Offset = 0.25

	w.SetContent(split)
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	return viewer
}

func (viewer *Viewer) StartViewer() {
	viewer.win.ShowAndRun()
}

func (viewer *Viewer) loadFile(path string) {
	data, err := dicom.ParseFile(path, nil)
	if err != nil {
		dialog.ShowError(err, viewer.win)
		return
	}
	for _, elem := range data.Elements {
		if elem.Tag == tag.PixelData {
			frames := elem.Value.GetValue().(dicom.PixelDataInfo).Frames
			if len(frames) == 0 {
				panic("No images found")
			}
			if frames[0].IsEncapsulated() {
				return
			}
			viewer.dicom.SetNativeFrame(&frames[0].NativeData)
			viewer.refreshCanvas(viewer.canvasImage)
		} else if elem.Tag == tag.WindowCenter {
			str := fmt.Sprintf("%v", elem.Value.GetValue().([]string)[0])
			l, _ := strconv.Atoi(str)
			viewer.dicom.SetWindowLevel(int16(l))
		} else if elem.Tag == tag.WindowWidth {
			str := fmt.Sprintf("%v", elem.Value.GetValue().([]string)[0])
			l, _ := strconv.Atoi(str)
			viewer.dicom.SetWindowWidth(int16(l))
		} else if elem.Tag == tag.PatientID {
			viewer.canvasID.SetText(fmt.Sprintf("%v", elem.Value))
		} else if elem.Tag == tag.PatientName {
			viewer.canvasName.SetText(fmt.Sprintf("%v", elem.Value))
		} else if elem.Tag == tag.PatientSex {
			viewer.canvasSex.SetText(fmt.Sprintf("%v", elem.Value))
		} else if elem.Tag == tag.PatientBirthDate {
			viewer.canvasBirthDate.SetText(fmt.Sprintf("%v", elem.Value))
		}
	}
}

func (viewer *Viewer) refreshCanvas(co ...fyne.CanvasObject) {
	for _, object := range co {
		canvas.Refresh(object)
	}
}
