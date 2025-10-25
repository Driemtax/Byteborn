package world

import (
	"math"

	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/Driemtax/Byteborn/pkg/types"
)

func CreateChunkNode(pNode *ChunkNode, depth int) *ChunkNode {
	if depth == config.CHUNK_TREE_MAX_DEPTH {
		pNode.chunk = NewChunk()
		return pNode
	}

	parentWidth := pNode.botRight.X - pNode.topLeft.X
	parentHeight := pNode.botRight.Y - pNode.topLeft.Y
	centerDisplacementWidth := parentWidth / 2
	centerDisplacementHeight := parentHeight / 2

	// Top Left
	topLeftNode := NewChunkNode(
		pNode.topLeft,
		types.Vec2{X: pNode.topLeft.X + centerDisplacementWidth, Y: pNode.topLeft.Y + centerDisplacementHeight},
	)
	pNode.topLeftNode = CreateChunkNode(topLeftNode, depth+1)
	// Top Right
	topRightNode := NewChunkNode(
		types.Vec2{X: pNode.topLeft.X + centerDisplacementWidth, Y: pNode.topLeft.Y},
		types.Vec2{X: pNode.botRight.X, Y: pNode.topLeft.Y + centerDisplacementHeight},
	)
	pNode.topRightNode = CreateChunkNode(topRightNode, depth+1)
	// Bottom Left
	botLeftNode := NewChunkNode(
		types.Vec2{X: pNode.topLeft.X, Y: pNode.topLeft.Y + centerDisplacementHeight},
		types.Vec2{X: pNode.topLeft.X + centerDisplacementWidth, Y: pNode.botRight.Y},
	)
	pNode.botLeftNode = CreateChunkNode(botLeftNode, depth+1)
	// Bottom Right
	botRightNode := NewChunkNode(
		types.Vec2{X: pNode.topLeft.X + centerDisplacementWidth, Y: pNode.topLeft.Y + centerDisplacementHeight},
		pNode.botRight,
	)
	pNode.botRightNode = CreateChunkNode(botRightNode, depth+1)

	return pNode
}

func CreateChunkTree() *ChunkNode {
	maxChunkAmount := math.Pow(2, config.CHUNK_TREE_MAX_DEPTH)
	maxSize := maxChunkAmount * config.CHUNK_WIDTH
	root := NewChunkNode(types.Vec2{X: 0, Y: 0}, types.Vec2{X: maxSize, Y: maxSize})
	CreateChunkNode(root, 0)
	return root
}
