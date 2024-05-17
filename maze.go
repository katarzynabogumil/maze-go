package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"slices"
	"strings"
	"unicode/utf8"
)

func main() {
	maze := setupMaze(20, 10, 50)
	start := generatePoint(&maze)
	end := generatePoint(&maze)

	path, err := shortestPath(&maze, start, end)
	if err != nil {
		fmt.Println(err)
		return
	}

	print(&maze, &path)
}

type Point struct {
	x          int
	y          int
	wall       bool
	distance   int
	neighbours []*Point
	prev       *Point
}

func shortestPath(maze *[][]*Point, start *Point, end *Point) ([]*Point, error) {
	q := []*Point{start}
	start.distance = 0
	visited := map[string]bool{}
	found := false

	for len(q) > 0 {
		point := q[0]
		q = q[1:]

		if point == end {
			found = true
			break
		}

		for _, neighbour := range point.neighbours {
			if !visited[strPoint(neighbour)] {
				visited[strPoint(neighbour)] = true
				q = append(q, neighbour)
				neighbour.distance = point.distance + 1
				neighbour.prev = point
			}
		}
	}

	if !found {
		return nil, fmt.Errorf("No path found")
	}

	path := []*Point{}
	point := end
	for point != start {
		path = append([]*Point{point}, path...)
		point = point.prev
	}

	return path, nil
}

func setupMaze(sizeX int, sizeY int, walls int) [][]*Point {
	var maze [][]*Point

	for j := 0; j < sizeY; j++ {
		var row []*Point

		for i := 0; i < sizeX; i++ {
			newPoint := Point{
				x:          i,
				y:          j,
				wall:       false,
				distance:   sizeX * sizeY,
				neighbours: nil,
				prev:       nil,
			}
			row = append(row, &newPoint)
		}
		maze = append(maze, row)
	}

	setWalls(&maze, walls)
	setNeighbours(&maze)

	return maze
}

func setWalls(maze *[][]*Point, walls int) {
	var flatMaze []*Point
	for _, row := range *maze {
		flatMaze = append(flatMaze, row...)
	}

	rand.Shuffle(len(flatMaze), func(i, j int) {
		flatMaze[i], flatMaze[j] = flatMaze[j], flatMaze[i]
	})

	for i := 0; i < walls; i++ {
		flatMaze[i].wall = true
	}
}

func setNeighbours(maze *[][]*Point) {
	for j, row := range *maze {
		for i, _ := range row {
			point := (*maze)[j][i]
			point.neighbours = getPointNeighbours(point, maze)
		}
	}
}

func getPointNeighbours(p *Point, maze *[][]*Point) []*Point {
	sizeX := float64(len((*maze)[0]))
	sizeY := float64(len(*maze))
	var neighbours []*Point

	for row := math.Max(float64(p.x-1), 0); row < math.Min(float64(p.x+2), sizeX); row++ {
		for col := math.Max(float64(p.y-1), 0); col < math.Min(float64(p.y+2), sizeY); col++ {
			point := (*maze)[int(col)][int(row)]
			if !point.wall && point != p && (int(row) == p.x || int(col) == p.y) {
				neighbours = append(neighbours, point)
			}
		}
	}
	return neighbours
}

func print(maze *[][]*Point, path *[]*Point) {
	for j := 0; j < len(*maze); j++ {
		row := ""
		for i := 0; i < len((*maze)[0]); i++ {
			point := (*maze)[j][i]
			icon, _ := utf8.DecodeRune([]byte{0xE2, 0xAC, 0x9C})
			if point.wall {
				icon, _ = utf8.DecodeRune([]byte{0xE2, 0xAC, 0x9B})
			}
			if path != nil {
				if slices.Contains(*path, point) {
					icon, _ = utf8.DecodeRune([]byte{0xE2, 0xAD, 0x95})
				}
				if point == (*path)[0] {
					icon, _ = utf8.DecodeRune([]byte{0xE2, 0xAD, 0x90})
				}
				if point == (*path)[len(*path)-1] {
					icon, _ = utf8.DecodeRune([]byte{0xF0, 0x9F, 0x8E, 0xAF})
				}
			}
			row += string(icon)
		}
		fmt.Println(row)
	}
}

func generatePoint(maze *[][]*Point) *Point {
	sizeX := len((*maze)[0])
	sizeY := len(*maze)
	var point *Point
	for {
		point = (*maze)[rand.IntN(sizeY)][rand.IntN(sizeX)]
		if !point.wall {
			break
		}
	}
	return point
}

func strPoint(p *Point) string {
	return strings.Join([]string{fmt.Sprint(p.x), fmt.Sprint(p.y)}, ",")
}
