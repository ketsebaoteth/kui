package elementtree

import (
	"main/project/element"
)

// ElementTree holds a list of elements
type ElementTree struct {
	Elements []element.Element
}

// AddElement adds a new element to the ElementTree
func (et *ElementTree) AddElement(e element.Element) {
	et.Elements = append(et.Elements, e)
}

// RenderTree renders all elements in the ElementTree
func (et *ElementTree) RenderTree() {
	for _, element := range et.Elements {
		element.UpdateState()
		element.Render()
	}
}
