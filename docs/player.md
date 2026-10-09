# Player
## Basic Idea
The player structs holds the following attributes.
## Move
The Function `Move()` handels one move of the player. It takes the following as an argument:
- dir: The Direction Vector, the player wants to move

The Direction Vector is a 2D-Vector, which has can have the values described in [[Vector2D.md]]. 

The first thing we need to do is to calculate the actual speed, because the player could be currently running. This could later be extended to swiming etc., where the player has a different speed, depending on the movement it is doing. 
### Handling Diagonal Movement
First important thing is to normalize the direction vector. If we move a single direction e.g. up, than we move exactly $1$ Unit. If we move diagonally e.g. UP/LEFT, than we move $1$ Unit UP and $1$ Unit to the LEFT. The distance can simply be calculated with the [https://en.wikipedia.org/wiki/Pythagorean_theorem](Pythagorian) Theorem and is therefore:
$$1² + 1² = 2$$ The Square Root of this is the actual distance, we moved, which is $\sqrt{2}$, hence approximatly $1.44...$. This would result in a faster diagonal movement then lateral movement. Therefore we normalize the direction vector before applying the movement. The result is, that we travel less then $1$ Unit UP and less then $1$ Unit LEFT and exactly $1$ Unit diagonally.
### General Formula to move
Next, I derived a formula that applies to every direction of motion. The basic idea was to not have a long `if-else` chain. Firstly i had a look at all possible moving directions i needed. Thus beeing UP, DOWN, LEFT, RIGHT. The diagonal movement could later be a sum of two lateral moves.

For moving UP the x-Value of the player position would not need to change. The y-Value of the player position would decrease of the value of the actual moving speed, since the upper, left corner of the window has the coordinates $0,0$. As we dont want to go higher then the window boundaries, the value cannot go lower then $0$. To ensure this, i used the `max` function. The UP-Movement can therefore be written as:

$$x += 0,\quad y = max(0,y-actualSpeed)$$

The other directions can be written as follows:

DOWN 

$$x += 0, \quad y = min(y+actualSpeed, HEIGHT)$$

LEFT 

$$x = max(0,x-actualSpeed), \quad y+=0$$

RIGHT 

$$x = min(x+actualSpeed, WIDTH),\quad y+=0$$

We further define a vector $\vec{v}$ for the velocity. This is calculated as the product of the direction vector and the actualSpeed. This results in a new vector, which holds information of the diretion and the speed the player is moving. The formulas for moving can now be combined to:

$$x = max(0,min(x+v_1, W))$$ 

where $v_1$ is the first component of the vector $\vec{v}$ and $W = \text{WIDTH} - \text{playerWidth}$. This is analogous for the y-Value:

$$y = max(0,min(y+v_2, H))$$ 

where $v_2$ is the second component of the vector $\vec{v}$ and $H = \text{HEIGHT} - \text{playerHeight}$. Therefore the vector $\vec{p}$ for the new Position of the player can be calculated quit effective.
### Checking for illegal moves
The last step is to check if the new position the player wants to move to is a legit position. The player cannot move through walls or walk on water. This can be done via the API of the map. Since this is currently not implemented, it will be added later on.
## Draw
The `draw()` function draws the player sprite to the screen.

As we have a spritesheet containing all player sprites we need a way to only print the sub image to the screen containing the current player sprite. This depens on the look direction of the player aswell as of an animation counter. When the player is moving there are 3 sprites continiouslly being drawn to the screen, so it looks like the player is walking. 

Firstly we need to create to points $p_1$ and $p_2$, which define the rectangle containing the current player sprite. $p_1$ defines the upper left corner of the rectangle, whereas $p_2$ defines the lower right corner. $p_1$ can be defined as:

$$x_0 = \text{frameCount} * \text{playerSizeX}, \quad y_0 = \text{lookDirection} * \text{playerSizeY}$$

This holds because, the lookDirection defines the row of the spriteSheet we need, where as one row is exactly the size of the player itself, in this case 32 pixel. The frameCount tells us which column we currently need. One row of the spridesheet contains 3 columns with the 3 sprites we need for the walking animation. This is indexed by the frameCount times the playerSize, in this case 32 pixel aswell.

We then cut exactly this sub image from the original spritesheet. Important is to translate the new sub image to the player position on the screen. Otherwise we would always print the player to the top left corner of the window, despite its actual position.
## UpdateAC
The `UpdateAC` function updates the `animationCount` variable on every game tick. The config file defines how fast the animation of the player walking should be played. 

We use the build-in delta time feature of ebiten here by calling the `ebiten.Tick()` function. This returns the count of ticks that have ellapsed since the game has started. Ebiten ensures, that the configured Ticks per second hold on every hardware. If your computer is too slow, you will have a lower tick rate. But if your pc is faster, you will **not** have a higher tick rate, which is quit important, so your walking animation is not played to fast. We use the modulo operation to check if enough ticks have ellapsed to update our animation. The `animationCount` varibale defines the index of the column of the spritesheet we draw on the screen. The bitwise AND operation is used here, because it is much faster then the modulo operator. The following code 

```go
p.animationCount = 1 + (p.animationCount & 1)
```

can be interpreted as a $% 2$ operation. This ensures, that the index for the spritesheet always toggles between $1$ and $2$, since those are the columns of the walking animation of the player in the spritesheet. If the player is not walking, then we simply set `animationCount` $=0$. This column holds the sprite for a standing player.
## UpdateLookDirection()
The `UpdateLookDirection` function updates the look direction of the player based on its prior move. If the player does not move, the `lastLookDirection` variable is not updated. On a player move it will be updated. For updating the look direction we simply use switch statement. The `lookDirection` holds the index of the row of the spritesheet we need to address to draw the correct sprite to the screen.

Firstly the approach for diagonal movement was to simply add 4 cases to the switch statement for every diagonal move. Moves that move the player to the upper left or lower left corner should result in a left look direction. This is analogous for the right direction. A switch statement is very performant, because the compiler can use a hashed jump table. The problem was, that i had to create 8 new 2D vectors, that use memory. This looked something like this:

```go
	switch p.LastDirection {
		case types.NewVector2D(0, -1):
			p.lookDirection = int(UP)
		case types.NewVector2D(0, 1):
			p.lookDirection = int(DOWN)
		case types.NewVector2D(-1, 0):
			p.lookDirection = int(LEFT)
		case types.NewVector2D(-1,1):
			p.lookDirection = int(LEFT)
		case types.NewVector2D(-1,1):
			p.lookDirection = int(LEFT)
		case types.NewVector2D(1, 0):
			p.lookDirection = int(RIGHT)
		case types.NewVector2D(1,1):
			p.lookDirection = int(RIGHT)
		case types.NewVector2D(1,-1):
			p.lookDirection = int(RIGHT)
	}
```

So i decided to use a tagless switch statement, so that i dont need to create any new struct instances. The compiler cannot longer use the efficient way of a jump table, under the hood this is just a if-else chain. But it is not that long, having 4 cases. As you can see, this is also more readable then the old version:

```go
	switch {
		case p.LastDirection.X == -1:
			p.lookDirection = int(LEFT)
		case p.LastDirection.X == 1:
			p.lookDirection = int(RIGHT)
		case p.LastDirection.Y == -1:
			p.lookDirection = int(UP)
		case p.LastDirection.Y == 1:
			p.lookDirection = int(DOWN)
	}

```
