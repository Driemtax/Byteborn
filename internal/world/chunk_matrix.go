package world

import (
	"fmt"

	"github.com/Driemtax/Byteborn/internal/config"
)

// An Idea of using a matrix as set of chunks but this approach is pretty boring and does not scale
// So i don't like it and won't use it. I'm going to implement a quad tree instead... make it more interesting :)

type ChunkMatrix [config.CHUNK_MATRIX_X_AMOUNT * config.CHUNK_MATRIX_Y_AMOUNT]*Chunk

func NewChunkMatrix() *ChunkMatrix {
	newChunkMatrix := ChunkMatrix{}
	for i := range config.CHUNK_MATRIX_X_AMOUNT * config.CHUNK_MATRIX_Y_AMOUNT {
		newChunkMatrix[i] = NewChunk()
	}
	return &newChunkMatrix
}

func (cm *ChunkMatrix) GetChunk(x, y int) (*Chunk, error) {
	if y*config.CHUNK_MATRIX_X_AMOUNT+x < config.CHUNK_MATRIX_X_AMOUNT*config.CHUNK_MATRIX_Y_AMOUNT {
		return cm[y*config.CHUNK_MATRIX_X_AMOUNT+x], nil
	}
	return nil, fmt.Errorf("Desired Chunk x:%d, y:%d is out of bounds}", x, y)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
