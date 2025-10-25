package world

import (
	"math"

	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	maxWorldSize int
	chunkTree    *ChunkNode
	cameraPos    types.Vec2
	cameraTarget types.Vec2
	cameraSpeed  float64
}

func NewWorld() *World {
	maxChunkAmount := math.Pow(2, config.CHUNK_TREE_MAX_DEPTH)
	maxWorldSize := maxChunkAmount * config.CHUNK_WIDTH
	return &World{
		maxWorldSize: int(maxWorldSize),
		chunkTree:    CreateChunkTree(),
		cameraSpeed:  config.CAMERA_SPEED,
	}
}

func (w *World) UpdateCamera(camTarget types.Vec2, dt float64) error {
	// Calculate the target position for the camera
	// The camera will always try to keep the target in the center of the screen
	// Therefore, the top left corner of the camera = playerPos - half screen size
	// is the correct target for our camera
	w.cameraTarget.X = camTarget.X - float64(config.WINDOW_WIDTH)/2
	w.cameraTarget.Y = camTarget.Y - float64(config.WINDOW_HEIGHT)/2

	// Clamp the camera target to the map boundaries
	w.cameraTarget = util.Clamp(
		w.cameraTarget,
		0,
		0,
		float64(w.maxWorldSize-config.WINDOW_WIDTH),
		float64(w.maxWorldSize-config.WINDOW_HEIGHT),
	)

	// --- Smoothly move the camera towards the target position (Linear Interpolation - Lerp)

	// Calculate the vector from the current camera position to the target camera position
	diff := w.cameraTarget.Sub(w.cameraPos)
	moveStep := diff.Mul(w.cameraSpeed * dt)
	moveStepLength := moveStep.Len()
	if moveStepLength > diff.Len() || moveStepLength < 0.1 {
		w.cameraPos = w.cameraTarget
	} else {
		w.cameraPos = w.cameraPos.Add(moveStep)
	}

	// Now also clamp the camera position to the map boundaries for safety
	w.cameraPos = util.Clamp(
		w.cameraPos,
		0,
		0,
		float64(w.maxWorldSize-config.WINDOW_WIDTH),
		float64(w.maxWorldSize-config.WINDOW_HEIGHT),
	)

	return nil
}

func (w *World) Draw(screen *ebiten.Image) {
	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()
	chunkNodes := w.chunkTree.GetChunkNodesTouchingSquare(
		w.cameraPos,
		types.Vec2{X: w.cameraPos.X + float64(screenWidth), Y: w.cameraPos.Y + float64(screenHeight)},
	)
	for _, node := range chunkNodes {
		node.Draw(screen, w.cameraPos)
	}
}
