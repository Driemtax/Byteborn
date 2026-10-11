package world

import (
	"math"

	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/Driemtax/Byteborn/pkg/types"
)

func CreateChunkNode(pNode *ChunkNode) *ChunkNode {
	if pNode.depth == config.CHUNK_TREE_MAX_DEPTH {
		pNode.chunk = NewChunk(pNode)
		return pNode
	}

	parentWidth := pNode.botRight.X - pNode.topLeft.X
	parentHeight := pNode.botRight.Y - pNode.topLeft.Y
	centerDisplacementWidth := parentWidth / 2
	centerDisplacementHeight := parentHeight / 2
	nextDepth := pNode.depth + 1

	// Top Left
	topLeftNode := NewChunkNode(
		pNode.topLeft,
		types.Vec2{X: pNode.topLeft.X + centerDisplacementWidth, Y: pNode.topLeft.Y + centerDisplacementHeight},
		nextDepth,
	)
	pNode.topLeftNode = CreateChunkNode(topLeftNode)
	// Top Right
	topRightNode := NewChunkNode(
		types.Vec2{X: pNode.topLeft.X + centerDisplacementWidth, Y: pNode.topLeft.Y},
		types.Vec2{X: pNode.botRight.X, Y: pNode.topLeft.Y + centerDisplacementHeight},
		nextDepth,
	)
	pNode.topRightNode = CreateChunkNode(topRightNode)
	// Bottom Left
	botLeftNode := NewChunkNode(
		types.Vec2{X: pNode.topLeft.X, Y: pNode.topLeft.Y + centerDisplacementHeight},
		types.Vec2{X: pNode.topLeft.X + centerDisplacementWidth, Y: pNode.botRight.Y},
		nextDepth,
	)
	pNode.botLeftNode = CreateChunkNode(botLeftNode)
	// Bottom Right
	botRightNode := NewChunkNode(
		types.Vec2{X: pNode.topLeft.X + centerDisplacementWidth, Y: pNode.topLeft.Y + centerDisplacementHeight},
		pNode.botRight,
		nextDepth,
	)
	pNode.botRightNode = CreateChunkNode(botRightNode)

	return pNode
}

func CreateChunkTree() *ChunkNode {
	maxChunkAmount := math.Pow(2, config.CHUNK_TREE_MAX_DEPTH)
	maxSize := maxChunkAmount * config.CHUNK_WIDTH
	root := NewChunkNode(types.Vec2{X: 0, Y: 0}, types.Vec2{X: maxSize, Y: maxSize}, 0)
	CreateChunkNode(root)
	return root
}
