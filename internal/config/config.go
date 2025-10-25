package config

const (
	WINDOW_WIDTH  = 400
	WINDOW_HEIGHT = 400

	// -----------------
	// --- W O R L D ---
	// -----------------
	TILE_SIZE    = 32 // The width and height of a single tile in pixels
	CHUNK_SIZE   = 16 // The width and height of a single chunk in tiles
	CHUNK_WIDTH  = TILE_SIZE * CHUNK_SIZE
	CHUNK_HEIGHT = TILE_SIZE * CHUNK_SIZE
	// Chunk Matrix
	CHUNK_MATRIX_X_AMOUNT = 16
	CHUNK_MATRIX_Y_AMOUNT = 16
	CHUNK_MATRIX_WIDTH    = CHUNK_WIDTH * CHUNK_MATRIX_X_AMOUNT
	CHUNK_MATRIX_HEIGHT   = CHUNK_HEIGHT * CHUNK_MATRIX_Y_AMOUNT
	// Chunk Quad-Tree
	CHUNK_TREE_MAX_DEPTH = 8
	// Camera
	CAMERA_SPEED = 6.0
)
