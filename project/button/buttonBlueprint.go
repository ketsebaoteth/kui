package button

import (
	"main/project/element"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Button is a struct that represents a button element
type Button struct {
	Pos            [2]int32
	Dim            [2]int32
	Text           string
	Fontfamily     rl.Font
	BorderColor    [4]int32
	Padding        [4]int32
	Margin         [4]int32
	BorderWidth    int
	BorderStyle    string
	Display        string
	Flexdirection  string
	AlignItems     string
	JustifyContent string
	Placeitems     string
	ClickCallBack  func()
	//borderradius
	//bgcolor
	//fgcolor
	//textdecoration
	//boxshadow
	//opacity
	//backdropfilter blur and invert
	//onhover callback
	//onmouseenter CALLBACK
	//onmouseleave CALLBACK
	//onmouseover validate
}

// checkes Intersection of a point with a button
func IntersectionChecker(b Button, x int32, y int32) bool {
	if x > b.Pos[0] && x < b.Pos[0]+b.Dim[0] && y > b.Pos[1] && y < b.Pos[1]+b.Dim[1] {
		return true
	}
	return false
}

// decalartion for the default button struct from defaultButton.go
var defbtn Button

// track if the first time the button is rendered
var firsttime = true

func (b Button) Render() {
	//only get the default button once
	if firsttime {
		defbtn = getDefaultButton()
		firsttime = false
	}
	//merge the default button with the given button
	//should be done every frame incase the default button changes
	b = mergeWithDefaults(b, defbtn)
	//draw basic button rectangle
	//Todo: handle rounded body
	drawButtonSkeleton(b)
	//Todo: handle rounded borders
	drawButtonBorder(b)
	drawButtonText(b)
}

func (b Button) UpdateState() {
	//handles different button states
	b.onClick()
}

func (b Button) onClick() {
	//excutes callback on button pressed
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		x := rl.GetMouseX()
		y := rl.GetMouseY()

		if IntersectionChecker(b, x, y) {
			b.ClickCallBack()
		}
	}
}

// Ensure Button implements element.Element
var _ element.Element = (*Button)(nil)
