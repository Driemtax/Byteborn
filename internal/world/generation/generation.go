package generation

import "github.com/kelindar/noise"

type Ctx struct {
	TerrainGenerator *noise.Simplex
}

// Creates a new generation context
// each generator will be initialized
// by a seed derived from the given seed
func NewCtx(seed uint32) *Ctx {
	return &Ctx{
		TerrainGenerator: noise.NewSimplex(seed),
	}
}
