package internal

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type DicomTheme struct {
}

func (m *DicomTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (m *DicomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return resourceEduSABeginnerRegularTtf
}

func (m *DicomTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m *DicomTheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {
	case theme.SizeNameText:
		return theme.DefaultTheme().Size(n) + 12
	default:
		return theme.DefaultTheme().Size(n)
	}
}
