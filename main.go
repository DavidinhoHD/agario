package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"fmt"
	"math"
	"math/rand"
	"strconv"
)

type Player struct {
	radius      float32
	position    rl.Vector2
	speed       float32
	lastDrawPos rl.Vector2
	//direction rl.Vector2
}

type Food struct {
	radius     float32
	posX, posY int32
}

func (p *Player) DrawPlayer() {
	const drawSmoothing = 0.1
	newDrawX := p.position.X*drawSmoothing + p.lastDrawPos.X*(1-drawSmoothing)
	newDrawY := p.position.Y*drawSmoothing + p.lastDrawPos.Y*(1-drawSmoothing)

	drawX := int32(math.Round(float64(newDrawX)))
	drawY := int32(math.Round(float64(newDrawY)))
	rl.DrawCircle(drawX, drawY, p.radius, rl.Pink)

	p.lastDrawPos = rl.Vector2{X: newDrawX, Y: newDrawY}

}

func (p *Player) Move(mousPos rl.Vector2, speed float32) {
	// get the direction vector from player position - mouse position
	direction := rl.Vector2Subtract(mousPos, p.position)
	//direction := subVector2(mousPos, p.position)

	// calculate length of the direction vector
	length := float32(math.Sqrt(float64(direction.X*direction.X + direction.Y*direction.Y)))

	// get the unit vector
	if length != 0 {
		direction.X = direction.X / length
		direction.Y = direction.Y / length

		// add the unit vector to the current player position times the sped of the player
		p.position.X += direction.X * speed //rl.Clamp(direction.X, 0, 15000) * speed
		p.position.Y += direction.Y * speed //rl.Clamp(direction.Y, 0, 15000) * speed
	}
}

func spawnFood() Food {
	var f Food
	x := rand.Int31n(1920)
	y := rand.Int31n(1080)

	f.posX = x
	f.posY = y
	f.radius = 10.0

	return f
}

func main() {
	windowWidth := int32(1920)
	windowHeight := int32(1080)

	//worldWidth := int32(15000)
	//worldHeight := int32(15000)

	rl.InitWindow(windowWidth, windowHeight, "window title")
	defer rl.CloseWindow()

	// rl.KeyNull = 0
	rl.SetExitKey(rl.KeyNull)
	rl.SetTargetFPS(60)

	player := Player{
		radius:   30.0,
		position: rl.Vector2{X: 1920 / 2, Y: 1080 / 2},
		speed:    15,
	}

	foods := make([]Food, 0)
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())
	foods = append(foods, spawnFood())

	player.lastDrawPos = player.position

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
			// Camera offset (displacement from target)
			Offset: rl.Vector2{1920 / 2, 1080 / 2},
			// Camera target (rotation and zoom origin)
			Target: player.position,
			// Camera rotation in degrees
			Rotation: 0.0,
			// Camera zoom (scaling), should be 1.0f by default
			Zoom: 1,
		}

		rl.BeginMode2D(camera)

		for i := 0; i < 6; i++ {
			x := int32(320 * i)
			for j := 0; j < 6; j++ {
				y := int32(180 * j)
				rl.DrawTexture(backgroundImage, x, y, rl.White)
			}
		}

		rl.EndMode2D()

		for _, v := range foods {
			rl.DrawCircle(v.posX, v.posY, v.radius, rl.Green)
		}

		player.DrawPlayer()
		if rl.IsCursorOnScreen() {
			mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
			player.Move(mouseWorldPos, player.speed)
			fmt.Println(player.position)
		}

		if showFPS {
			fps := rl.GetFPS()
			s := strconv.FormatFloat(float64(fps), 'f', 0, 32)
			rl.DrawText(s, 5, 5, 30, rl.Green)
		}

		rl.EndDrawing()
	}
}
