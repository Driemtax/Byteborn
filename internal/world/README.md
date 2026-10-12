# World

The package manages the visible game world. It lazily generates required
chunks from Simplex noise, smoothly follows the player with the camera, and
draws only chunks that intersect the visible screen area.

## Structure

```text
world
├── world.go              World, camera, and rendering
├── generation
│   └── generation.go     Deterministic noise generators 
└── chunk
    ├── tree.go           QuadTree root and access
    ├── chunk_node.go     Spatial partitioning and lazy generation
    ├── chunk.go          16×16 tile chunk and rendering
    └── tile
        ├── tile.go       Tile types, noise mapping, and rendering
        └── tile_detail.go Tile name, abbreviation, and color
```

`World` calculates the camera position from the player target, clamps it to
the world boundaries, and uses the `Tree` to draw only relevant leaf nodes.
The `Tree` recursively divides the square world into four `Node`s. A `Chunk`
is created only at the maximum tree depth, so unseen areas remain
uninitialized. A `Chunk` is rendered from a flat tile array addressed by
`x + CHUNK_SIZE*y`. Each tile is selected from a noise value.

## Data Structures

- **`World`** groups the camera, world size, generation context, and QuadTree
  behind one central world interface.
- **`generation.Ctx`** stores all the needed noise generators. A shared context
  keeps terrain generation reproducible and avoids creating a generator per
  chunk.
- **`chunk.Tree` and `chunk.Node`** form the QuadTree. Rectangle queries find
  visible areas efficiently; pointer children and the optional `chunk` enable
  lazy generation and avoid allocating untouched areas.
- **`chunk.Chunk`** is a fixed-size (`16×16`) array instead of a nested matrix.
  This keeps the data compact and provides direct indexing.
- **`tile.Tile`** is a `uint8` enum. Small fixed values are sufficient for the
  tile types; `TILE_DETAILS` keeps name, abbreviation, and color data separate
  from each tile value and avoids repeating that data per tile.
- **`tile.TileDetail`** encapsulates rendering metadata, allowing terminal
  output and Ebiten rendering to use the same tile definitions.
