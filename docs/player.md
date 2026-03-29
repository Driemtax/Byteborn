# Player
## Basic Idea
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
$$x +=  & y = max(0,y-actualSpeed)$$
The other directions can be written as follows:
DOWN $$x += 0 & y = min(y+actualSpeed, HEIGHT)$$
LEFT $$x = max(0,x-actualSpeed) & y+=0$$
RIGHT $$x = min(x+actualSpeed, WIDTH) & y+=0$$

We further define a vector $\vec{v}$ for the velocity. This is calculated as the product of the direction vector and the actualSpeed. This results in a new vector, which holds information of the diretion and the speed the player is moving. The formulas for moving can now be combined to:
$$x = max(0,min(x+v_1, W))$$ where $v_1$ is the first component of the vector $\vec{v}$ and $W = \text{WIDTH} - \text{playerWidth}$. This is analogous for the y-Value:
$$y = max(0,min(y+v_2, H))$$ where $v_2$ is the second component of the vector $\vec{v}$ and $H = \text{HEIGHT} - \text{playerHeight}$. Therefore the vector $\vec{p}$ for the new Position of the player can be calculated quit effective.
### Checking for illegal moves
The last step is to check if the new position the player wants to move to is a legit position. The player cannot move through walls or walk on water. This can be done via the API of the map. Since this is currently not implemented, it will be added later on.
