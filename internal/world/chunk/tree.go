package chunk

import (
	"github.com/Driemtax/Byteborn/internal/world/generation"
	"github.com/Driemtax/Byteborn/pkg/types"
)

type Tree struct {
	root   *Node
	genCtx *generation.Ctx
}

func NewTree(worldSize float64, generationCtx *generation.Ctx) *Tree {
	depth := 0
	return &Tree{
		root: NewNode(
			types.Vec2{X: 0, Y: 0}, types.Vec2{X: worldSize, Y: worldSize},
			depth,
		),
		genCtx: generationCtx,
	}
}

func (n *Tree) GetChunkNodesTouchingSquare(tl, br types.Vec2) []*Node {
	return n.root.getChunkNodesTouchingSquare(tl, br, n.genCtx)
}
