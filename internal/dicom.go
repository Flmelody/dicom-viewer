package internal

import (
	"github.com/suyashkumar/dicom/pkg/frame"
	"image"
	"image/color"
)

type Dicom struct {
	windowLevel       int16
	windowWidth       int16
	rows              int
	cols              int
	frameImage        *image.Image
	nativeFrame       *frame.NativeFrame
	isEncapsulated    bool
	encapsulatedFrame *frame.EncapsulatedFrame
	dicomData         *DicomData
}

type DicomData struct {
	Name string
}

func (d *Dicom) SetNativeFrame(nativeFrame *frame.NativeFrame) {
	d.nativeFrame = nativeFrame
	frameImage, err := d.nativeFrame.GetImage()
	if err != nil {

	}
	d.setBounds(frameImage)
}

func (d *Dicom) SetEncapsulatedFrame(encapsulatedFrame *frame.EncapsulatedFrame) {
	d.isEncapsulated = true
	d.encapsulatedFrame = encapsulatedFrame
	frameImage, err := d.encapsulatedFrame.GetImage()
	if err != nil {

	}
	d.setBounds(frameImage)
}

func (d *Dicom) setBounds(frameImage image.Image) {
	bounds := frameImage.Bounds()
	d.rows = bounds.Dy()
	d.cols = bounds.Dx()
	d.frameImage = &frameImage
}

func (d *Dicom) WindowLevel() int16 {
	return d.windowLevel
}

func (d *Dicom) SetWindowLevel(windowLevel int16) {
	d.windowLevel = windowLevel
}

func (d *Dicom) WindowWidth() int16 {
	return d.windowWidth
}

func (d *Dicom) SetWindowWidth(windowWidth int16) {
	d.windowWidth = windowWidth
}

func (d *Dicom) ColorModel() color.Model {
	return color.Gray16Model
}

func (d *Dicom) Bounds() image.Rectangle {
	if d.nativeFrame == nil && d.encapsulatedFrame == nil {
		return image.Rectangle{}
	}
	return image.Rect(0, 0, d.cols, d.rows)
}

func (d *Dicom) At(x, y int) color.Color {
	if d.nativeFrame == nil && d.encapsulatedFrame == nil {
		return color.Gray16{Y: 0}
	}
	if d.isEncapsulated {
		return (*d.frameImage).At(x, y)
	}
	windowMin := d.windowLevel - d.windowWidth/2
	windowMax := windowMin + d.windowWidth*2

	i := y*d.rows + x
	if i >= len(d.nativeFrame.Data) {
		return color.Black
	}

	raw := int16(d.nativeFrame.Data[i][0])

	if raw < windowMin {
		return color.Gray16{Y: 0}
	} else if raw >= windowMax {
		return color.Gray16{Y: 0xffff}
	}

	val := float32(raw-windowMin) / float32(d.windowWidth)
	return color.Gray16{Y: uint16(float32(0xffff) * val)}
}

func NewDicom(isEncapsulated bool, nativeFrame *frame.NativeFrame, encapsulatedFrame *frame.EncapsulatedFrame, windowLevel, windowWidth int16) *Dicom {
	if isEncapsulated {
		return NewEncapsulatedFrameDicom(encapsulatedFrame, windowLevel, windowWidth)
	} else {
		return NewNativeFrameDicom(nativeFrame, windowLevel, windowWidth)
	}
}

func NewNativeFrameDicom(frame *frame.NativeFrame, windowLevel, windowWidth int16) *Dicom {
	return &Dicom{nativeFrame: frame, windowLevel: windowLevel, windowWidth: windowWidth, dicomData: &DicomData{}}
}

func NewEncapsulatedFrameDicom(frame *frame.EncapsulatedFrame, windowLevel, windowWidth int16) *Dicom {
	return &Dicom{encapsulatedFrame: frame, isEncapsulated: true, windowLevel: windowLevel, windowWidth: windowWidth, dicomData: &DicomData{}}
}
