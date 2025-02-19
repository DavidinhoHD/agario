package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	minSpeed = 5.0
	growthRate = 1.0
	speedDecay = 0.99
)

var WorldDimensions = rl.Vector2{
		X: 1920 * 3,
		Y: 1080 * 3,
}

type GameState struct {
	FoodPositions []Food
	PlayerPositions []Player
}

var GS GameState


func PlayerFoodCollision(f Food, p Player) bool {
	return rl.CheckCollisionCircles(f.Position, f.Radius, p.Position, p.Radius)
}