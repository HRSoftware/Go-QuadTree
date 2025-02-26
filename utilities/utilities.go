package utilities

import (
	"errors"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

// linearSearch
func LinearSearch(slice []rl.Vector2, target rl.Vector2) (rl.Vector2, error) {
	for _, value := range slice {
		if value == target {
			return value, nil
		}
	}
	return rl.Vector2{}, errors.New("value not found")
}

func RandomColour() rl.Color {
	red := uint8(randomFloat32(0, 1) * 255)
	green := uint8(randomFloat32(0, 1) * 255)
	blue := uint8(randomFloat32(0, 1) * 255)

	return rl.Color{R: uint8(red), G: uint8(green), B: uint8(blue), A: 255}
}
