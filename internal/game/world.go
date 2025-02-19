package game

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	minSpeed = 5.0
	growthRate = 1.0
	speedDecay = 0.99
)

var WorldDimensions = rl.Vector2{
		X: 1920 * 3,
		Y: 1080 * 3,
}