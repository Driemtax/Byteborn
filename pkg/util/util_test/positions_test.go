package util_test

import (
	"testing"

	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/Driemtax/Byteborn/pkg/util"
)

func TestIsInBounds(t *testing.T) {
	minX, minY, maxX, maxY := 0.0, 0.0, 100.0, 100.0
	inBoundsPos := types.Vector2D{X: 50, Y: 50}
	if !util.IsInBounds(inBoundsPos, minX, minY, maxX, maxY) {
		t.Error("Expected inBoundsPos to be in Bounds")
	}

	topOutside := types.Vector2D{X: 50, Y: -10}
	if util.IsInBounds(topOutside, minX, minY, maxX, maxY) {
		t.Error("Expected topOutside to be out of bounds")
	}

	rightOutside := types.Vector2D{X: 150, Y: 50}
	if util.IsInBounds(rightOutside, minX, minY, maxX, maxY) {
		t.Error("Expected rightOutside to be out of bounds")
	}

	bottomOutside := types.Vector2D{X: 50, Y: 110}
	if util.IsInBounds(bottomOutside, minX, minY, maxX, maxY) {
		t.Error("Expected bottomOutside to be out of bounds")
	}

	leftOutside := types.Vector2D{X: 50, Y: 110}
	if util.IsInBounds(leftOutside, minX, minY, maxX, maxY) {
		t.Error("Expected leftOutside to be out of bounds")
	}

	onTopBound := types.Vector2D{X: 50, Y: 0}
	if !util.IsInBounds(onTopBound, minX, minY, maxX, maxY) {
		t.Error("Expected onTopBound to be in bounds")
	}

	onTopLeftCorner := types.Vector2D{X: 0, Y: 0}
	if !util.IsInBounds(onTopLeftCorner, minX, minY, maxX, maxY) {
		t.Error("Expected onTopLeftCornert to be in bounds")
	}
}

func TestClamp(t *testing.T) {
	minX, minY, maxX, maxY := 0.0, 0.0, 100.0, 100.0

	onBoundary := types.Vector2D{X: 50, Y: 0}
	clampedPos := util.Clamp(onBoundary, minX, minY, maxX, maxY)
	if !onBoundary.Eq(clampedPos) {
		t.Error("Expected clampedPos to be unchanged")
	}

	outsideBoundary := types.Vector2D{X: 50, Y: -10}
	clampedPos2 := util.Clamp(outsideBoundary, minX, minY, maxX, maxY)
	if !clampedPos2.Eq(types.Vector2D{X: 50, Y: minY}) {
		t.Error("Expected clampedPos2 to be changed")
	}
}
