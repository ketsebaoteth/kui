package button

import rl "github.com/gen2brain/raylib-go/raylib"

//returns a default button instance
//Todo: customize the themeing
func getDefaultButton() Button {
	font := rl.LoadFont("./fonts/Roboto/Roboto-Regular.ttf")
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)
	defualtButton := Button{
		Pos:            [2]int32{10, 10},
		Dim:            [2]int32{10, 10},
		Text:           "Button",
		Fontfamily:     font,
		BorderColor:    [4]int32{45, 45, 45, 255},
		Padding:        [4]int32{10, 10, 10, 10},
		Margin:         [4]int32{2, 2, 2, 2},
		BorderWidth:    2,
		BorderStyle:    "solid",
		Display:        "flex",
		Flexdirection:  "row",
		AlignItems:     "center",
		JustifyContent: "flex-end",
		Placeitems:     "center",
		ClickCallBack: func() {
		},
	}

	return defualtButton
}
