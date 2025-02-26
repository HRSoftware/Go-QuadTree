package types

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"goQuadTree/utilities"
)

type QuadTree struct {
	boundary  rl.Rectangle
	points    []rl.Vector2
	northWest *QuadTree
	northEast *QuadTree
	southWest *QuadTree
	southEast *QuadTree
	Colour    rl.Color
}

func MakeQuadTree(aabb rl.Rectangle) *QuadTree {
	fmt.Println("Creating quadtree")
	return &QuadTree{
		boundary:  aabb,
		points:    nil,
		northWest: nil,
		northEast: nil,
		southWest: nil,
		southEast: nil,
		Colour:    utilities.RandomColour(),
	}
}

func (tree *QuadTree) Destroy() {
	if tree.northWest == nil || tree.northEast == nil || tree.southWest == nil || tree.southEast == nil {

		return
	}

	tree.northWest.Destroy()
	tree.northWest = nil

	tree.northEast.Destroy()
	tree.northEast = nil

	tree.southWest.Destroy()
	tree.southWest = nil

	tree.southEast.Destroy()
	tree.southEast = nil
}

// Insert is a recursive function that adds a point. Subdivision will happen if needed
func (tree *QuadTree) Insert(point rl.Vector2) bool {
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
		tree.Subdivide()
	}

	if tree.northWest.Insert(point) {
		return true
	}
	if tree.northEast.Insert(point) {
		return true
	}
	if tree.southWest.Insert(point) {
		return true
	}
	if tree.southEast.Insert(point) {
		return true
	}

	panic("Unable to subdivide any further, and unable to insert point into quad tree. Consider adjusting qtree parameters")
}

func (tree *QuadTree) Subdivide() {
	halfBounds := rl.Vector2{tree.boundary.Width / 2, tree.boundary.Height / 2}

	tree.northWest = MakeQuadTree(rl.Rectangle{X: tree.boundary.X, Y: tree.boundary.Y, Width: halfBounds.X, Height: halfBounds.Y})
	tree.northEast = MakeQuadTree(rl.Rectangle{X: tree.boundary.X + halfBounds.X, Y: tree.boundary.Y, Width: halfBounds.X, Height: halfBounds.Y})
	tree.southWest = MakeQuadTree(rl.Rectangle{X: tree.boundary.X, Y: tree.boundary.Y + halfBounds.Y, Width: halfBounds.X, Height: halfBounds.Y})
	tree.southEast = MakeQuadTree(rl.Rectangle{X: tree.boundary.X + halfBounds.X, Y: tree.boundary.Y + halfBounds.Y, Width: halfBounds.X, Height: halfBounds.Y})

}

// Query returns all rectangles in a given quadtree
func (tree *QuadTree) Query(rec rl.Rectangle) []rl.Vector2 {
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

	if childResults := tree.northWest.Query(rec); len(childResults) > 0 {
		for _, point := range childResults {
			results = append(
				results,
				point,
			)
		}
	}

	if childResults := tree.northEast.Query(rec); len(childResults) > 0 {
		for _, point := range childResults {
			results = append(
				results,
				point,
			)
		}
	}

	if childResults := tree.southEast.Query(rec); len(childResults) > 0 {
		for _, point := range childResults {
			results = append(
				results,
				point,
			)
		}
	}

	if childResults := tree.southWest.Query(rec); len(childResults) > 0 {
		for _, point := range childResults {
			results = append(
				results,
				point,
			)
		}
	}

	return results
}

type rectCol struct {
	Rect rl.Rectangle
	Col  rl.Color
}

// Visualise returns a list of rectangles to
func (tree *QuadTree) Visualise() []rectCol {
	rectList := make([]rectCol, 0)

	rectList = append(rectList, rectCol{tree.boundary, tree.Colour})

	if tree.northWest == nil {
		return rectList
	}

	if childRect := tree.northWest.Visualise(); len(childRect) > 0 {
		for _, point := range childRect {
			rectList = append(rectList, point)
		}
	}

	if childRect := tree.northEast.Visualise(); len(childRect) > 0 {
		for _, point := range childRect {
			rectList = append(rectList, point)
		}
	}

	if childRect := tree.southWest.Visualise(); len(childRect) > 0 {
		for _, point := range childRect {
			rectList = append(rectList, point)
		}
	}

	if childRect := tree.southEast.Visualise(); len(childRect) > 0 {
		for _, point := range childRect {
			rectList = append(rectList, point)
		}
	}

	return rectList
}
