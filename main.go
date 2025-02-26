package main

import (
	"errors"
	"fmt"
	_ "fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type QuadTree struct {
	boundary  rl.Rectangle
	points    []rl.Vector2
	northWest *QuadTree
	northEast *QuadTree
	southWest *QuadTree
	southEast *QuadTree
}

func MakeQuadTree(aabb rl.Rectangle) *QuadTree {
	return &QuadTree{
		boundary:  aabb,
		points:    nil,
		northWest: nil,
		northEast: nil,
		southWest: nil,
		southEast: nil,
	}
}

func QuadTree_Destroy(tree *QuadTree) {
	if tree.northWest == nil || tree.northEast == nil || tree.southWest == nil || tree.southEast == nil {

		return
	}

	QuadTree_Destroy(tree.northWest)
	tree.northWest = nil

	QuadTree_Destroy(tree.northEast)
	tree.northEast = nil

	QuadTree_Destroy(tree.southWest)
	tree.southWest = nil

	QuadTree_Destroy(tree.southEast)
	tree.southEast = nil
}

func QuadTree_Insert(tree *QuadTree, point rl.Vector2) bool {
	if !rl.CheckCollisionPointRec(point, tree.boundary) {
		return false
	}

	if len(tree.points) < 4 && tree.northWest == nil {

		if tree.points == nil {
			tree.points = make([]rl.Vector2, 0)
		}

		tree.points = append(tree.points, point)
		fmt.Println("Appending point")
		return true
	}

	if tree.northEast == nil {
		fmt.Println("Subdividing")
		QuadTree_Subdivide(tree)
	}

	if QuadTree_Insert(tree.northWest, point) {
		return true
	}
	if QuadTree_Insert(tree.northEast, point) {
		return true
	}
	if QuadTree_Insert(tree.southWest, point) {
		return true
	}
	if QuadTree_Insert(tree.southEast, point) {
		return true
	}

	panic("Unable to subdivide any further, and unable to insert point into quad tree. Consider adjusting qtree parameters")
}

func QuadTree_Subdivide(tree *QuadTree) {
	halfBounds := rl.Vector2{tree.boundary.Width / 2, tree.boundary.Height / 2}

	tree.northWest = MakeQuadTree(rl.Rectangle{X: tree.boundary.X, Y: tree.boundary.Y, Width: halfBounds.X, Height: halfBounds.Y})
	tree.northEast = MakeQuadTree(rl.Rectangle{X: tree.boundary.X + halfBounds.X, Y: tree.boundary.Y, Width: halfBounds.X, Height: halfBounds.Y})
	tree.southWest = MakeQuadTree(rl.Rectangle{X: tree.boundary.X, Y: tree.boundary.Y + halfBounds.Y, Width: halfBounds.X, Height: halfBounds.Y})
	tree.southEast = MakeQuadTree(rl.Rectangle{X: tree.boundary.X + halfBounds.X, Y: tree.boundary.Y + halfBounds.Y, Width: halfBounds.X, Height: halfBounds.Y})

}

// QuadTree_Query returns all rectangles in a given quadtree
func QuadTree_Query(tree *QuadTree, rec rl.Rectangle) []rl.Vector2 {
	results := make([]rl.Vector2, 0)

	if !rl.CheckCollisionRecs(tree.boundary, rec) {
		return results
	}

	for _, point := range tree.points {
		if rl.CheckCollisionPointRec(rl.Vector2{X: point.X, Y: point.Y}, rec) {
			results = append(
				results,
				point,
			)
		}
	}

	if tree.northWest == nil {
		return results
	}

	if childResults := QuadTree_Query(tree.northWest, rec); len(childResults) > 0 {
		for _, point := range childResults {
			results = append(
				results,
				point,
			)
		}
	}

	if childResults := QuadTree_Query(tree.northEast, rec); len(childResults) > 0 {
		for _, point := range childResults {
			results = append(
				results,
				point,
			)
		}
	}

	if childResults := QuadTree_Query(tree.southEast, rec); len(childResults) > 0 {
		for _, point := range childResults {
			results = append(
				results,
				point,
			)
		}
	}

	if childResults := QuadTree_Query(tree.southWest, rec); len(childResults) > 0 {
		for _, point := range childResults {
			results = append(
				results,
				point,
			)
		}
	}

	return results
}

// QuadTree_Visualise returns a list of rectangles to
func QuadTree_Visualise(tree *QuadTree) []rl.Rectangle {
	rectList := make([]rl.Rectangle, 0)

	rectList = append(rectList, tree.boundary)

	if tree.northWest == nil {
		return rectList
	}

	if childRect := QuadTree_Visualise(tree.northWest); len(childRect) > 0 {
		for _, point := range childRect {
			rectList = append(rectList, point)
		}
	}

	if childRect := QuadTree_Visualise(tree.northEast); len(childRect) > 0 {
		for _, point := range childRect {
			rectList = append(rectList, point)
		}
	}

	if childRect := QuadTree_Visualise(tree.southWest); len(childRect) > 0 {
		for _, point := range childRect {
			rectList = append(rectList, point)
		}
	}

	if childRect := QuadTree_Visualise(tree.southEast); len(childRect) > 0 {
		for _, point := range childRect {
			rectList = append(rectList, point)
		}
	}

	return rectList
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func main() {
	var WIDTH int32 = 640
	var HEIGHT int32 = 480

	RECT_SIZE := rl.Vector2{X: 10., Y: 10.}

	rl.InitWindow(WIDTH, HEIGHT, "Quadtree Test")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	quadTree := MakeQuadTree(rl.Rectangle{X: 0, Y: 0, Width: float32(WIDTH), Height: float32(HEIGHT)})

	rectPositions := map[rl.Vector2]rl.Color{}

	randColours := []rl.Color{}

	for idx, _ := range randColours {
		randColours[idx] = randomColour()
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		mousePos := rl.GetMousePosition()

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			newRec := rl.Rectangle{X: mousePos.X, Y: mousePos.Y, Width: float32(WIDTH), Height: float32(HEIGHT)}
			rectPositions[rl.Vector2{X: newRec.X, Y: newRec.Y}] = randomColour()
			QuadTree_Insert(quadTree, rl.Vector2{X: newRec.X, Y: newRec.Y})
		}

		visualisation := QuadTree_Visualise(quadTree)

		for _, val := range visualisation {
			rl.DrawRectangleV(rl.Vector2{X: val.X, Y: val.Y}, rl.Vector2{X: val.Width, Y: val.Height}, rectPositions[rl.Vector2{X: val.X, Y: val.Y}])
		}

		rl.DrawRectangleV(mousePos, rl.Vector2{X: 20., Y: 20.}, rl.Red)

		queryResults := QuadTree_Query(quadTree, rl.Rectangle{X: mousePos.X, Y: mousePos.Y, Width: RECT_SIZE.X, Height: RECT_SIZE.Y})

		for key, _ := range rectPositions {
			_, err := linearSearch(queryResults[:], key)
			if err != nil {
				rl.DrawRectangleV(key, RECT_SIZE, rl.White)
			} else {
				rl.DrawRectangleV(key, RECT_SIZE, rl.Green)
			}
		}

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}

// linearSearch
func linearSearch(slice []rl.Vector2, target rl.Vector2) (rl.Vector2, error) {
	for _, value := range slice {
		if value == target {
			return value, nil
		}
	}
	return rl.Vector2{}, errors.New("value not found")
}

func randomColour() rl.Color {
	red := uint8(randomFloat32(0, 1) * 255)
	green := uint8(randomFloat32(0, 1) * 255)
	blue := uint8(randomFloat32(0, 1) * 255)

	return rl.Color{R: uint8(red), G: uint8(green), B: uint8(blue), A: 255}
}
