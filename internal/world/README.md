# World

The job of the world package will be to render the world around the player and handle the camera movement.

The things that i want to achieve with the world package:

- Chunking
- Dynamic Camera movement
- The Camera follows the player slowly behind him
- The World receives the width of of the used screen and determines automatically how many tiles will be needed

```go
type ChunkNode struct {
	topLeftPoint Vec2
	botRightPoint Vec2

	chunk *Chunk

	topLeftNode 	*ChunkNode
	topRightNode 	*ChunkNode
	botLeftNode 	*ChunkNode
	botRightNode 	*ChunkNode
}
```

I will define a fixed height until im accessing a chunk and determine a good depth by that. This way we will be able to get an open map.

# Definition

Each type that will be created inside of the world will get a String and a Display Function.
The String will always output something like: Chunk{...data...}
Display will be used for pretty printing. For example printing the map as a matrix on the terminal
