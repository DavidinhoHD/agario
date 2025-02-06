package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const (
    worldWidth  = 1920 * 3 // 3 times the window width
    worldHeight = 1080 * 3 // 3 times the window height
)

type Player struct {
	radius      float32
	position    rl.Vector2
	speed       float32
	lastDrawPos rl.Vector2
	//direction rl.Vector2
}

type Food struct {
	radius   float32
	position rl.Vector2
}

func (p *Player) DrawPlayer() {
    // Draw player at the center of the screen
    centerX := float32(1920 / 2)
    centerY := float32(1080 / 2)
    rl.DrawCircle(int32(centerX), int32(centerY), p.radius, rl.Pink)
}

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
func spawnFood() Food {
	var f Food
	x := rand.Int31n(worldWidth)
	y := rand.Int31n(worldHeight)

	f.position.X = float32(x)
	f.position.Y = float32(y)
	f.radius = 10.0

	return f
}

func spawnFoodTicker(foods *[]Food) {
	ticker := time.NewTicker(1 * time.Millisecond) // Spawn food every 1 second
	go func() {
		for range ticker.C {
			*foods = append(*foods, spawnFood())
		}
	}()
}

func foodPlayerCollision(food Food, player Player) bool {
	return rl.CheckCollisionCircles(food.position, food.radius, player.position, player.radius)
}

func main() {
	windowWidth := int32(1920)
	windowHeight := int32(1080)

	//worldWidth := int32(15000)
	//worldHeight := int32(15000)

	rl.InitWindow(windowWidth, windowHeight, "window title")
	defer rl.CloseWindow()

	rl.SetExitKey(rl.KeyNull)
	rl.SetTargetFPS(60)

	player := Player{
		radius:   30.0,
		position: rl.Vector2{X: float32(worldWidth) / 2, Y: float32(worldHeight) / 2},
		speed:    15,
	}
	player.lastDrawPos = player.position

	foods := make([]Food, 0)
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())

	backgroundImage := rl.LoadTexture("static/background.png")
	if backgroundImage.ID == 0 {
		fmt.Println("No background image")
	}
	defer rl.UnloadTexture(backgroundImage)

	showFPS := false

	// gameLoop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

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

		// draw foods and check for collisions
		for i := 0; i < len(foods); i++ {
			food := foods[i]
			rl.DrawCircle(int32(food.position.X), int32(food.position.Y), food.radius, rl.Green)
			if foodPlayerCollision(food, player) {
				player.radius += 5
				player.speed -= 1
				foods = append(foods[:i], foods[i+1:]...) // remove the eaten food
				i-- // adjust the index after removal
			}
		}

		rl.EndMode2D()

		// Draw player after camera transform
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
