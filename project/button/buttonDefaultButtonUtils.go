package button

import "strings"

// mergeWithDefaults merges the given Button with default values
func mergeWithDefaults(b Button, defbtn Button) Button {
	if b.Pos == [2]int32{0, 0} {
		b.Pos = defbtn.Pos
	}
	if b.Dim == [2]int32{0, 0} {
		b.Dim = defbtn.Dim
	}
	if b.Text == "" {
		b.Text = defbtn.Text
	}
	if b.Fontfamily.Texture.ID == 0 {
		b.Fontfamily = defbtn.Fontfamily
	}
	if b.BorderColor == [4]int32{0, 0, 0, 0} {
		b.BorderColor = defbtn.BorderColor
	}
	if b.Padding == [4]int32{0, 0, 0, 0} {
		b.Padding = defbtn.Padding
	}
	if b.Margin == [4]int32{0, 0, 0, 0} {
		b.Margin = defbtn.Margin
	}
	if b.BorderWidth == 0 {
		b.BorderWidth = defbtn.BorderWidth
	}
	if b.BorderStyle == "" {
		b.BorderStyle = defbtn.BorderStyle
	}
	if strings.TrimSpace(b.Display) == "" {
		b.Display = defbtn.Display
	}
	if b.Flexdirection == "" {
		b.Flexdirection = defbtn.Flexdirection
	}
	if b.AlignItems == "" {
		b.AlignItems = defbtn.AlignItems
	}
	if b.JustifyContent == "" {
		b.JustifyContent = defbtn.JustifyContent
	}
	if b.Placeitems == "" {
		b.Placeitems = defbtn.Placeitems
	}
	if b.ClickCallBack == nil {
		b.ClickCallBack = defbtn.ClickCallBack
	}
	return b
}
