package element

// Element is the interface that all element types will implement
type Element interface {
	Render()
	UpdateState()
}
