package world

import (
	"fmt"

	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
)

// Implementation for a Quad-Tree
type ChunkNode struct {
	topLeft  types.Vec2 // The top left corner of the ChunkNode
	botRight types.Vec2 // The bottom right corner of the ChunkNode

	chunk *Chunk
	depth int

	topLeftNode  *ChunkNode
	topRightNode *ChunkNode
	botLeftNode  *ChunkNode
	botRightNode *ChunkNode
}

func NewChunkNode(topLeft, botRight types.Vec2, chunkDepth int) *ChunkNode {
	return &ChunkNode{
		topLeft:  topLeft,
		botRight: botRight,
		depth:    chunkDepth,
	}
}

func (cn ChunkNode) StringR(depth int) string {
	var hasChunk string
	if cn.chunk != nil {
		hasChunk += "true"
	} else {
		hasChunk += "false"
	}
	var indentation string = util.Indent(depth)
	var tl string = ""
	if cn.topLeftNode != nil {
		tl += cn.topLeftNode.StringR(depth + 1)
	}
	var tr string = ""
	if cn.topRightNode != nil {
		tr += cn.topRightNode.StringR(depth + 1)
	}
	var bl string = ""
	if cn.botLeftNode != nil {
		bl += cn.botLeftNode.StringR(depth + 1)
	}
	var br string = ""
	if cn.botRightNode != nil {
		br += cn.botRightNode.StringR(depth + 1)
	}
	return fmt.Sprintf("ChunkNode{\n%s topLeft:%+v,\n%s topRight:%+v,\n%s hasChunk:%s,\n%s tl: %s,\n%s tr: %s,\n%s bl: %s,\n%s br: %s\n%s}", indentation, cn.topLeft, indentation, cn.botRight, indentation, hasChunk, indentation, tl, indentation, tr, indentation, bl, indentation, br, indentation)
}

func (cn ChunkNode) String() string {
	return cn.StringR(0)
}

func (cn *ChunkNode) GetChunkForPosition(pos types.Vec2) (*Chunk, error) {
	if !cn.Wraps(pos) {
		return nil, fmt.Errorf(
			"The given position %s is outside of the ChunkNode in the square between %s and %s",
			pos,
			cn.topLeft,
			cn.botRight,
		)
	}
	if cn.chunk != nil {
		return cn.chunk, nil
	}

	if cn.topLeftNode != nil && cn.topLeftNode.Wraps(pos) {
		return cn.topLeftNode.GetChunkForPosition(pos)
	} else if cn.topRightNode != nil && cn.topRightNode.Wraps(pos) {
		return cn.topRightNode.GetChunkForPosition(pos)
	} else if cn.botLeftNode != nil && cn.botLeftNode.Wraps(pos) {
		return cn.botLeftNode.GetChunkForPosition(pos)
	} else if cn.botRightNode != nil && cn.botRightNode.Wraps(pos) {
		return cn.botRightNode.GetChunkForPosition(pos)
	}

	return nil, fmt.Errorf("GetChunkForPosition: This Error Should Never Occur: %s, Position: %s", cn.String(), pos.String())
}

// Is currently not used!!!
// GetChukstouchingSquare: This function returns all the chunks that touch the given Rectangle
// The Idea is that later i can pass the current screen and render all the chunks that are currently
// on the screen
// func (cn *ChunkNode) GetChunksTouchingSquare(tl, br types.Vec2) []*Chunk {
// 	if !cn.TouchesSq(tl, br) {
// 		return []*Chunk{}
// 	}
// 	if cn.chunk != nil {
// 		return []*Chunk{cn.chunk}
// 	}

// 	chunks := []*Chunk{}
// 	chunks = append(chunks, cn.topLeftNode.GetChunksTouchingSquare(tl, br)...)
// 	chunks = append(chunks, cn.topRightNode.GetChunksTouchingSquare(tl, br)...)
// 	chunks = append(chunks, cn.botLeftNode.GetChunksTouchingSquare(tl, br)...)
// 	chunks = append(chunks, cn.botRightNode.GetChunksTouchingSquare(tl, br)...)

// 	return chunks
// }

// Searches for all chunknodes in the given rectangle.
// If the chunktree hasn't been created to this point,
// the function will create the missing chunks and add
// them to the chunktree. (hopefully -.-)
func (cn *ChunkNode) getChunkNodesTouchingSquare(tl, br types.Vec2) []*ChunkNode {
	if !cn.TouchesSq(tl, br) {
		return []*ChunkNode{}
	}
	if cn.depth == config.CHUNK_TREE_MAX_DEPTH {
		if cn.chunk == nil {
			// Create the chunk if it has not been created yet
			cn.chunk = NewChunk(cn)
		}
		return []*ChunkNode{cn}
	}

	if cn.topLeftNode == nil {
		cn.createChildren()
	}

	chunkNodes := []*ChunkNode{}
	chunkNodes = append(chunkNodes, cn.topLeftNode.getChunkNodesTouchingSquare(tl, br)...)
	chunkNodes = append(chunkNodes, cn.topRightNode.getChunkNodesTouchingSquare(tl, br)...)
	chunkNodes = append(chunkNodes, cn.botLeftNode.getChunkNodesTouchingSquare(tl, br)...)
	chunkNodes = append(chunkNodes, cn.botRightNode.getChunkNodesTouchingSquare(tl, br)...)

	return chunkNodes
}

func (cn *ChunkNode) createChildren() {
	if cn.depth == config.CHUNK_TREE_MAX_DEPTH {
		return
	}

	parentWidth := cn.botRight.X - cn.topLeft.X
	parentHeight := cn.botRight.Y - cn.topLeft.Y
	centerDisplacementWidth := parentWidth / 2
	centerDisplacementHeight := parentHeight / 2
	nextDepth := cn.depth + 1

	// Top Left
	cn.topLeftNode = NewChunkNode(
		cn.topLeft,
		types.Vec2{X: cn.topLeft.X + centerDisplacementWidth, Y: cn.topLeft.Y + centerDisplacementHeight},
		nextDepth,
	)
	// Top Right
	cn.topRightNode = NewChunkNode(
		types.Vec2{X: cn.topLeft.X + centerDisplacementWidth, Y: cn.topLeft.Y},
		types.Vec2{X: cn.botRight.X, Y: cn.topLeft.Y + centerDisplacementHeight},
		nextDepth,
	)
	// Bottom Left
	cn.botLeftNode = NewChunkNode(
		types.Vec2{X: cn.topLeft.X, Y: cn.topLeft.Y + centerDisplacementHeight},
		types.Vec2{X: cn.topLeft.X + centerDisplacementWidth, Y: cn.botRight.Y},
		nextDepth,
	)
	// Bottom Right
	cn.botRightNode = NewChunkNode(
		types.Vec2{X: cn.topLeft.X + centerDisplacementWidth, Y: cn.topLeft.Y + centerDisplacementHeight},
		cn.botRight,
		nextDepth,
	)
}

// Checks if the given Point is inside of the ChunkNode
func (cn *ChunkNode) Wraps(pos types.Vec2) bool {
	return util.IsInPointsSquare(pos, cn.topLeft, cn.botRight)
}

// Checks if the current chunk touches the given square defined by two points
func (cn *ChunkNode) TouchesSq(tl, br types.Vec2) bool {
	return util.SquareCollides(tl, br, cn.topLeft, cn.botRight)
}

func (cn *ChunkNode) Draw(screen *ebiten.Image, camPos types.Vec2) {
	if cn.chunk != nil {
		cn.chunk.Draw(screen, cn.topLeft.Sub(camPos))
	}
}
