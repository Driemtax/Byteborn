package world

import "image/color"

type TileDetail struct {
	name         string
	abbreviation string // Max of 2 letters describing the block
	color        color.RGBA
}

func (td TileDetail) Abbreviation() string {
	return td.abbreviation
}

func (td TileDetail) Name() string {
	return td.name
}

func (td TileDetail) Color() color.RGBA {
	return td.color
}
