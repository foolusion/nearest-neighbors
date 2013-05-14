package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	qt "github.com/foolusion/quadtree"
)

type point struct {
	id int
	qt.XY
}

type nn struct {
	p1, p2   point
	distance float64
}

var file string = ""

func main() {
	points := ReadPointsFromFile(file)
	bounds := BoundingBox(points)
	root := qt.New(bounds)

	InsertAllPoints(root, points)

	results := NearestNeighbors(root)
	fmt.Println(results)
}

func ReadPointsFromFile(fn string) []*qt.XY {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	points := make([]*qt.XY, 100)

	done := false

	for !done {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		ptString := strings.Split(line, "\t")

		id, err := strconv.Atoi(ptString[0])
		if err != nil {
			log.Fatal(err)
		}

		x, err := strconv.ParseFloat(ptString[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.ParseFloat(ptString[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		pt := &point{id, qt.XY{x, y}}

		points = append(points, &pt.XY)
	}

	return points
}

func BoundingBox(points []*qt.XY) qt.AABB {
	return qt.AABB{}
}

func InsertAllPoints(root *qt.QuadTree, points []*qt.XY) {
}

func NearestNeighbors(root *qt.QuadTree) []nn {
	neighbors := make([]nn, 100)
	return neighbors
}
