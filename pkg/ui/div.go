package ui

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Div struct {
	Box
	components []Component

	isWireframed bool
}

func NewDiv(isWireFramed bool) *Div {
	return &Div{
		components:   []Component{},
		isWireframed: isWireFramed,
	}
}

func (d *Div) AddChild(c Component) {
	d.components = append(d.components, c)
}

// This for now only works for the button div in the main menu
func (d *Div) OnKeyPress(key ebiten.Key) bool {
	for _, c := range d.components {
		if h, ok := c.(KeyPressHandler); ok && h.OnKeyPress(key) {
			return true
		}
	}
	return false
}

// Adapting the layout of a div also adapts the position and size of all its child components. Previous scale is secured.
func (d *Div) Layout(bounds image.Rectangle) {
	parentBounds := d.bounds
	d.bounds = bounds

	// Check if the div has child components. If not just return here
	childCount := len(d.components)
	if childCount <= 0 {
		return
	}

	if childCount == 1 {
		d.components[0].Layout(bounds)
	}

	// If there are child components save the relative size of each child in comparison to the parent div and reapply it afterwards
	oldParentWidth := parentBounds.Bounds().Dx()
	if oldParentWidth == 0 {
		oldParentWidth = bounds.Dx()
	}

	oldParentHeight := parentBounds.Bounds().Dy()
	if oldParentHeight == 0 {
		oldParentHeight = bounds.Dy()
	}
	// The difference in parent Position is needed to calculate the new position of each children.
	parentXDiff := bounds.Min.X - parentBounds.Min.X
	parentYDiff := bounds.Min.Y - parentBounds.Min.Y

	// so the idea here is to calculate the relative width and height of the child in comparison to the old parents width and height.
	// and then reapply this factor to the new width and height of the parent, so the child gets scaled correctly.
	for _, c := range d.components {
		childBounds := c.Bounds()
		childWidth := childBounds.Dx()
		childHeight := childBounds.Dy()
		childWidthRelativeFactor := float32(childWidth) / float32(oldParentWidth)
		childHeightRelativeFactor := float32(childHeight) / float32(oldParentHeight)

		newChildX := childBounds.Min.X + parentXDiff
		newChildY := childBounds.Min.Y + parentYDiff
		newChildWidth := int(float32(bounds.Dx()) * childWidthRelativeFactor)
		newChildHeight := int(float32(bounds.Dy()) * childHeightRelativeFactor)
		childNewBounds := image.Rect(newChildX, newChildY, newChildX+newChildWidth, newChildY+newChildHeight)

		c.Layout(childNewBounds)
	}

}

func (d *Div) Draw(screen *ebiten.Image) {
	if d.isWireframed {
		// Draw a white border for the whole div to make it visible
		vector.StrokeRect(screen, float32(d.bounds.Min.X), float32(d.bounds.Min.Y), float32(d.bounds.Dx()), float32(d.bounds.Dy()), 2.0, color.White, false)
	}

	for _, c := range d.components {
		c.Draw(screen)
	}
}
