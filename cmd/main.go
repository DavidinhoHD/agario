package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"fmt"
	//"math"
	//"math/rand"
	"strconv"
	//"sync"
	//"time"

	"github.com/DavidinhoHD/agario/internal/api"
	"github.com/DavidinhoHD/agario/internal/game"
)

/*
// World dimensions defined as multiples of the window size
const (
    worldWidth  = 1920 * 3 // 3 times the window width
    worldHeight = 1080 * 3 // 3 times the window height
    minSpeed    = 5.0     // Minimum player speed
    growthRate  = 1.0     // How much radius increases per food
    speedDecay  = 0.99    // Speed multiplier when eating food (0.99 = 1% slower)
)

*/

/*

// Player represents the user-controlled entity in the game
type Player struct {
    radius      float32    // Size of the player circle
    position    rl.Vector2 // Current position in the world
    speed       float32    // Movement speed
    lastDrawPos rl.Vector2 // Previous position for smooth rendering
}

*/

/*
// game.Food represents collectible items that increase player size
type game.Food struct {
    radius   float32    // Size of the food circle
    position rl.Vector2 // Position in the world
}

*/

/*
// GameState manages the game's mutable state with thread-safe access
type GameState struct {
    foods []game.Food    // Slice containing all food items
    mutex sync.Mutex // Mutex for thread-safe food operations
}

*/

/*

// DrawPlayer renders the player at the center of the screen
// The camera follows the player, making it appear stationary
func (p *Player) DrawPlayer() {
    // Draw player at the center of the screen
    centerX := float32(1920 / 2)
    centerY := float32(1080 / 2)
    rl.DrawCircle(int32(centerX), int32(centerY), p.radius, rl.Pink)
}

// Move updates the player's position based on mouse input
// Implements smooth movement and world boundary checking
func (p *Player) Move(mousePos rl.Vector2, speed float32) {
	// get the direction vector from player position - mouse position
	direction := rl.Vector2Subtract(mousePos, p.position)

	// calculate length of the direction vector
	length := float32(math.Sqrt(float64(direction.X*direction.X + direction.Y*direction.Y)))

	// Create a smooth transition for stopping
	const stopRadius = 15.0
	if length < stopRadius {
		speed *= length / stopRadius // Gradually reduce speed as we get closer to center
		if speed < 1 { // Minimum speed threshold
			return
		}
	}

	// get the unit vector
	if length != 0 {
		direction.X = direction.X / length
		direction.Y = direction.Y / length

		 // Calculate new position
        newX := p.position.X + direction.X*speed
        newY := p.position.Y + direction.Y*speed

        // Clamp position within world bounds
        p.position.X = rl.Clamp(newX, 0, float32(worldWidth))
        p.position.Y = rl.Clamp(newY, 0, float32(worldHeight))
	}
}

*/

/*
// spawnFood creates a new food item at a random position within the world bounds
func spawnFood() game.Food {
	var f game.Food
	x := rand.Int31n(worldWidth)
	y := rand.Int31n(worldHeight)

	f.Position.X = float32(x)
	f.Position.Y = float32(y)
	f.Radius = 10.0

	return f
}

*/

/*
// spawnFoodTicker periodically spawns new food items in a separate goroutine
// Uses mutex to ensure thread-safe modification of the foods slice
func spawnFoodTicker(state *GameState) {
    ticker := time.NewTicker(2 * time.Second) // Spawn food every 2 seconds
    go func() {
        for range ticker.C {
            state.mutex.Lock()
            state.foods = append(state.foods, spawnFood())
            state.mutex.Unlock()
        }
    }()
}

*/

/*
// foodPlayerCollision checks if a player has collided with a food item
func foodPlayerCollision(food game.Food, player game.Player) bool {
	return rl.CheckCollisionCircles(food.Position, food.Radius, player.Position, player.Radius)
}

*/

func main() {
	gameState, err := api.GetGameState()
	if err != nil {
		fmt.Println(err)
		return
	}

	player, err := api.GetNewPlayer()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize window and game settings
	rl.InitWindow(int32(game.WindowDimensions.X), int32(game.WindowDimensions.Y), "window title")
	defer rl.CloseWindow()

	player.LastDrawPos = player.Position
	/*
	       // Create player at the center of the world
	   	player := game.Player{
	   		Radius:   30.0,
	   		Position: rl.Vector2{X: float32(worldWidth) / 2, Y: float32(worldHeight) / 2},
	   		Speed:    15,
	   	}
	   	player.LastDrawPos = player.Position
	*/

	/*
		    // Initialize game state and food system
			gameState := &GameState{
		        foods: make([]game.Food, 0),
		    }

		    // Spawn initial food items
		    for i := 0; i < 6; i++ {
		        gameState.foods = append(gameState.foods, spawnFood())
		    }

	*/

	/*
	   // Start automatic food spawning system
	   spawnFoodTicker(gameState)

	*/

	backgroundImage := rl.LoadTexture(game.BackgroundImagePath)
	if backgroundImage.ID == 0 {
		fmt.Println("No background image")
		return
	}
	defer rl.UnloadTexture(backgroundImage)

	showFPS := true

	rl.SetExitKey(rl.KeyNull)
	rl.SetTargetFPS(60)

	// Game loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Setup camera to follow player
		camera := game.SetupCamera(player)
		/*
			camera := rl.Camera2D{
				Offset:   rl.Vector2{X: float32(game.WindowDimensions.X) / 2, Y: float32(game.WindowDimensions.Y) / 2},
				Target:   player.Position,
				Rotation: 0.0,
				Zoom:     1.0,
			}
		*/

		rl.BeginMode2D(camera)

		/*
			// Calculate number of tiles needed for world coverage
			tilesX := game.WorldDimensions.X / 320
			tilesY := game.WorldDimensions.Y / 180

			for i := int32(0); i < int32(tilesX); i++ {
				x := int32(320 * i)
				for j := int32(0); j < int32(tilesY); j++ {
				y := int32(180 * j)
				rl.DrawTexture(backgroundImage, x, y, rl.White)
				}
			}

		*/

		game.DrawWorld(backgroundImage)
		gameState, err = api.GetGameState()
		if err != nil {
			fmt.Println(err)
			return
		}
		gameState.SpawnFood()

		for _, v := range gameState.FoodPositions {
			if game.PlayerFoodCollision(v, player) {
				api.DeleteFoodById(v)
				player.ChangePlayerProperties()
			}
		}

		for _, v := range gameState.PlayerPositions {
			if game.PlayerPlayerCollision(player, v) {
				if player.Radius > v.Radius {
					api.DeletePlayerById(v)
				}
			}
		}

		/*
			        // Thread-safe food handling and collision detection
					gameState.mutex.Lock()
			        for i := 0; i < len(gameState.foods); i++ {
			            food := gameState.foods[i]
			            rl.DrawCircle(int32(food.Position.X), int32(food.Position.Y), food.Radius, rl.Green)
			            if foodPlayerCollision(food, player) {
			                // Smoother growth and speed adjustment
			                player.Radius += growthRate
			                player.Speed = float32(math.Max(float64(player.Speed*speedDecay), minSpeed))
			                gameState.foods = append(gameState.foods[:i], gameState.foods[i+1:]...)
			                i--
			            }
			        }
			        gameState.mutex.Unlock()
		*/
		rl.EndMode2D()

		// Player drawing and movement
		player.DrawPlayer()

		if rl.IsCursorOnScreen() {
			mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
			player.Move(mouseWorldPos, player.Speed)
		}

		if showFPS {
			fps := rl.GetFPS()
			s := strconv.FormatFloat(float64(fps), 'f', 0, 32)
			rl.DrawText(s, 5, 5, 30, rl.Green)
		}

		p := fmt.Sprintf("Points:%d", game.Points)
		rl.DrawText(p, 5, 1050, 30, rl.Green)
		rl.EndDrawing()
	}
}
