# World

The job of the world package will be to render the world around the player and handle the camera movement.

The things that i want to achieve with the world package:

- Chunking
- Dynamic Camera movement
- The Camera follows the player slowly behind him
- The World receives the width of of the used screen and determines automatically how many tiles will be needed

# First Idea Chunks as Nodes

```go
type ChunkNode struct {
	currentChunk *Chunk
	north 			 *Chunk
	east 				 *Chunk
	south        *Chunk
	west      	 *Chunk
}
```

This Approach sadly does not work because starting from the root node. The root->north->west chunk and the root->west->north chunk are the same. So there is already a collision.

# Second Idea

Normal matrix but way to boring.... I wont go deeper into it

# Third Idea

Im using my first idea but this time im going to use the correct data structure a quad tree :). And using that it works

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
