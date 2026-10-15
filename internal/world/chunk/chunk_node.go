package chunk

import (
	"fmt"

	"github.com/Driemtax/Byteborn/internal/config"
	"github.com/Driemtax/Byteborn/internal/world/generation"
	"github.com/Driemtax/Byteborn/pkg/types"
	"github.com/Driemtax/Byteborn/pkg/util"
	"github.com/hajimehoshi/ebiten/v2"
)

// Implementation for a Quad-Tree
type Node struct {
	topLeft  types.Vec2 // The top left corner of the ChunkNode
	botRight types.Vec2 // The bottom right corner of the ChunkNode

	chunk *Chunk
	depth int // The depth of the chunknode inside the quadtree

	topLeftNode  *Node
	topRightNode *Node
	botLeftNode  *Node
	botRightNode *Node
}

func NewNode(topLeft, botRight types.Vec2, chunkDepth int) *Node {
	return &Node{
		topLeft:  topLeft,
		botRight: botRight,
		depth:    chunkDepth,
	}
}

func (n *Node) Draw(screen *ebiten.Image, camPos types.Vec2) {
	if n.chunk != nil {
		n.chunk.Draw(screen, n.topLeft.Sub(camPos))
	}
}

// A recursive String() function which will
// travel down the whole chunknode tree
func (n Node) stringR(depth int) string {
	var hasChunk string
	if n.chunk != nil {
		hasChunk += "true"
	} else {
		hasChunk += "false"
	}
	var indentation string = util.Indent(depth)
	var tl string = ""
	if n.topLeftNode != nil {
		tl += n.topLeftNode.stringR(depth + 1)
	}
	var tr string = ""
	if n.topRightNode != nil {
		tr += n.topRightNode.stringR(depth + 1)
	}
	var bl string = ""
	if n.botLeftNode != nil {
		bl += n.botLeftNode.stringR(depth + 1)
	}
	var br string = ""
	if n.botRightNode != nil {
		br += n.botRightNode.stringR(depth + 1)
	}
	return fmt.Sprintf("ChunkNode{\n%s topLeft:%+v,\n%s topRight:%+v,\n%s hasChunk:%s,\n%s tl: %s,\n%s tr: %s,\n%s bl: %s,\n%s br: %s\n%s}", indentation, n.topLeft, indentation, n.botRight, indentation, hasChunk, indentation, tl, indentation, tr, indentation, bl, indentation, br, indentation)
}

func (n Node) String() string {
	return n.stringR(0)
}

// Searches for all chunknodes in the given rectangle.
// If the chunktree hasn't been created to this point,
// the function will create the missing chunks and add
// them to the chunktree.
func (n *Node) getChunkNodesTouchingSquare(tl, br types.Vec2, genCtx *generation.Ctx) []*Node {
	if !n.touchesSq(tl, br) {
		return []*Node{}
	}
	if n.depth == config.CHUNK_TREE_MAX_DEPTH {
		if n.chunk == nil {
			// Create the chunk if it has not been created yet
			n.chunk = newChunk(n, genCtx)
		}
		return []*Node{n}
	}

	if n.topLeftNode == nil {
		n.createChildren()
	}

	chunkNodes := []*Node{}
	chunkNodes = append(chunkNodes, n.topLeftNode.getChunkNodesTouchingSquare(tl, br, genCtx)...)
	chunkNodes = append(chunkNodes, n.topRightNode.getChunkNodesTouchingSquare(tl, br, genCtx)...)
	chunkNodes = append(chunkNodes, n.botLeftNode.getChunkNodesTouchingSquare(tl, br, genCtx)...)
	chunkNodes = append(chunkNodes, n.botRightNode.getChunkNodesTouchingSquare(tl, br, genCtx)...)

	return chunkNodes
}

func (n *Node) createChildren() {
	if n.depth == config.CHUNK_TREE_MAX_DEPTH {
		return
	}

	parentWidth := n.botRight.X - n.topLeft.X
	parentHeight := n.botRight.Y - n.topLeft.Y
	centerDisplacementWidth := parentWidth / 2
	centerDisplacementHeight := parentHeight / 2
	nextDepth := n.depth + 1

	// Top Left
	n.topLeftNode = NewNode(
		n.topLeft,
		types.Vec2{X: n.topLeft.X + centerDisplacementWidth, Y: n.topLeft.Y + centerDisplacementHeight},
		nextDepth,
	)
	// Top Right
	n.topRightNode = NewNode(
		types.Vec2{X: n.topLeft.X + centerDisplacementWidth, Y: n.topLeft.Y},
		types.Vec2{X: n.botRight.X, Y: n.topLeft.Y + centerDisplacementHeight},
		nextDepth,
	)
	// Bottom Left
	n.botLeftNode = NewNode(
		types.Vec2{X: n.topLeft.X, Y: n.topLeft.Y + centerDisplacementHeight},
		types.Vec2{X: n.topLeft.X + centerDisplacementWidth, Y: n.botRight.Y},
		nextDepth,
	)
	// Bottom Right
	n.botRightNode = NewNode(
		types.Vec2{X: n.topLeft.X + centerDisplacementWidth, Y: n.topLeft.Y + centerDisplacementHeight},
		n.botRight,
		nextDepth,
	)
}

// Checks if the given Point is inside of the ChunkNode
func (n *Node) wraps(pos types.Vec2) bool {
	return util.IsInPointsSquare(pos, n.topLeft, n.botRight)
}

// Checks if the current chunk touches the given square defined by two points
func (n *Node) touchesSq(tl, br types.Vec2) bool {
	return util.SquareCollides(tl, br, n.topLeft, n.botRight)
}
