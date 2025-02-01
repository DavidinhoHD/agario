package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"math/rand"
)

// adds 2 Vector2 objects
func addVector2(v1, v2 rl.Vector2) rl.Vector2 {
	var v rl.Vector2
	v.X = v1.X + v2.X
	v.Y = v1.Y + v2.Y

	return v
}

// subtracts 2 Vector2 objects (v1 - v2)
func subVector2(v1, v2 rl.Vector2) rl.Vector2 {
	var v rl.Vector2
	v.X = v1.X - v2.X
	v.Y = v1.Y - v2.Y

	return v
}

// divides 2 Vector2 objects (v1 / v2)
func divideVector2(v1, v2 rl.Vector2) rl.Vector2 {
	var v rl.Vector2
	v.X = v1.X / v2.X
	v.Y = v1.Y / v2.Y

	return v
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
