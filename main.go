package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"fmt"
	"math"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// World dimensions defined as multiples of the window size
const (
    worldWidth  = 1920 * 3 // 3 times the window width
    worldHeight = 1080 * 3 // 3 times the window height
    minSpeed    = 5.0     // Minimum player speed
    growthRate  = 1.0     // How much radius increases per food
    speedDecay  = 0.99    // Speed multiplier when eating food (0.99 = 1% slower)
)

// Player represents the user-controlled entity in the game
type Player struct {
    radius      float32    // Size of the player circle
    position    rl.Vector2 // Current position in the world
    speed       float32    // Movement speed
    lastDrawPos rl.Vector2 // Previous position for smooth rendering
}

// Food represents collectible items that increase player size
type Food struct {
    radius   float32    // Size of the food circle
    position rl.Vector2 // Position in the world
}

// GameState manages the game's mutable state with thread-safe access
type GameState struct {
    foods []Food    // Slice containing all food items
    mutex sync.Mutex // Mutex for thread-safe food operations
}

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

// spawnFood creates a new food item at a random position within the world bounds
func spawnFood() Food {
	var f Food
	x := rand.Int31n(worldWidth)
	y := rand.Int31n(worldHeight)

	f.position.X = float32(x)
	f.position.Y = float32(y)
	f.radius = 10.0

	return f
}

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

// foodPlayerCollision checks if a player has collided with a food item
func foodPlayerCollision(food Food, player Player) bool {
	return rl.CheckCollisionCircles(food.position, food.radius, player.position, player.radius)
}

func main() {
    // Initialize window and game settings
	windowWidth := int32(1920)
	windowHeight := int32(1080)

	//worldWidth := int32(15000)
	//worldHeight := int32(15000)

	rl.InitWindow(windowWidth, windowHeight, "window title")
	defer rl.CloseWindow()

	rl.SetExitKey(rl.KeyNull)
	rl.SetTargetFPS(60)

    // Create player at the center of the world
	player := Player{
		radius:   30.0,
		position: rl.Vector2{X: float32(worldWidth) / 2, Y: float32(worldHeight) / 2},
		speed:    15,
	}
	player.lastDrawPos = player.position

    // Initialize game state and food system
	gameState := &GameState{
        foods: make([]Food, 0),
    }

    // Spawn initial food items
    for i := 0; i < 6; i++ {
        gameState.foods = append(gameState.foods, spawnFood())
    }

    // Start automatic food spawning system
    spawnFoodTicker(gameState)

	backgroundImage := rl.LoadTexture("static/background.png")
	if backgroundImage.ID == 0 {
		fmt.Println("No background image")
	}
	defer rl.UnloadTexture(backgroundImage)

	showFPS := false

	// Game loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

        // Setup camera to follow player
		camera := rl.Camera2D{
			Offset:   rl.Vector2{X: float32(windowWidth) / 2, Y: float32(windowHeight) / 2},
			Target:   player.position,
			Rotation: 0.0,
			Zoom:     1.0,
		}

		rl.BeginMode2D(camera)

		// Calculate number of tiles needed for world coverage
		tilesX := worldWidth / 320
		tilesY := worldHeight / 180

		for i := int32(0); i < int32(tilesX); i++ {
			x := int32(320 * i)
			for j := int32(0); j < int32(tilesY); j++ {
			y := int32(180 * j)
			rl.DrawTexture(backgroundImage, x, y, rl.White)
			}
		}

        // Thread-safe food handling and collision detection
		gameState.mutex.Lock()
        for i := 0; i < len(gameState.foods); i++ {
            food := gameState.foods[i]
            rl.DrawCircle(int32(food.position.X), int32(food.position.Y), food.radius, rl.Green)
            if foodPlayerCollision(food, player) {
                // Smoother growth and speed adjustment
                player.radius += growthRate
                player.speed = float32(math.Max(float64(player.speed*speedDecay), minSpeed))
                gameState.foods = append(gameState.foods[:i], gameState.foods[i+1:]...)
                i--
            }
        }
        gameState.mutex.Unlock()

		rl.EndMode2D()

        // Player drawing and movement
		player.DrawPlayer()

		if rl.IsCursorOnScreen() {
			mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
			player.Move(mouseWorldPos, player.speed)
		}

		if showFPS {
			fps := rl.GetFPS()
			s := strconv.FormatFloat(float64(fps), 'f', 0, 32)
			rl.DrawText(s, 5, 5, 30, rl.Green)
		}

		rl.EndDrawing()
	}
}
