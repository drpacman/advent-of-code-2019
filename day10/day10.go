package main

import (
	"advent-of-code-2019/files"
	"advent-of-code-2019/numbers"
	"fmt"
	"math"
	"sort"
)

type AsteroidMap struct {
	contents [][]int
}

type Coordinate struct {
	x int
	y int
}

type Targets struct {
	pos     Coordinate
	entries []Coordinate
}

func (c Coordinate) distance(other Coordinate) int {
	return numbers.AbsInt(c.x-other.x) + numbers.AbsInt(c.y-other.y)
}

func (t Targets) Len() int {
	return len(t.entries)
}

func (t Targets) Swap(i, j int) {
	t.entries[i], t.entries[j] = t.entries[j], t.entries[i]
}

func (t Targets) Less(i, j int) bool {
	t1 := t.entries[i]
	t2 := t.entries[j]
	return t1.distance(t.pos) < t2.distance(t.pos)
}

type Asteroid struct {
	location    Coordinate
	lineOfSight map[float64][]Coordinate
}

func (a *Asteroid) getTargets() []Coordinate {
	// get sorted keys and sort by distance
	var keys []float64
	var totalAsteroids = 0
	sortedMap := make(map[float64][]Coordinate, len(a.lineOfSight))
	for k := range a.lineOfSight {
		keys = append(keys, k)
		entries := a.lineOfSight[k]
		totalAsteroids += len(entries)
		// sort entries
		sort.Sort(Targets{a.location, entries})
		sortedMap[k] = entries
	}

	// sort keys (which are by angle)
	sort.Float64s(keys)
	targets := make([]Coordinate, totalAsteroids)
	targetPos := 0
	round := 0
	for targetPos < totalAsteroids {
		for _, key := range keys {
			entries, ok := sortedMap[key]
			if !ok || round >= len(entries) {
				continue
			}
			targets[targetPos] = entries[round]
			targetPos++
			if targetPos >= totalAsteroids {
				break
			}
		}
		round++
	}
	return targets
}

func readAsteroidMap(filepath string) AsteroidMap {
	lines := files.ReadLines(filepath)
	contents := make([][]int, len(lines))
	for i, rowContents := range lines {
		row := make([]int, len(rowContents))
		for j, _ := range rowContents {
			if rowContents[j] == "#"[0] {
				row[j] = 1
			}
		}
		contents[i] = row
	}
	return AsteroidMap{contents}
}

func (a *AsteroidMap) calculateLineOfSight(pos Coordinate) map[float64][]Coordinate {
	lineOfSight := make(map[float64][]Coordinate)
	for y := 0; y < len(a.contents); y++ {
		for x := 0; x < len(a.contents[y]); x++ {
			if a.contents[y][x] == 1 && !(pos.x == x && pos.y == y) {
				angle := angle(pos.x, pos.y, x, y)
				entry, ok := lineOfSight[angle]
				if ok == false {
					entry = make([]Coordinate, 0)
				}
				lineOfSight[angle] = append(entry, Coordinate{x, y})
			}
		}
	}
	return lineOfSight
}

func (a *AsteroidMap) findBestAsteroid() Asteroid {
	var target Asteroid
	var max = 0
	for y := 0; y < len(a.contents); y++ {
		for x := 0; x < len(a.contents[y]); x++ {
			if a.contents[y][x] == 1 {
				location := Coordinate{x, y}
				los := a.calculateLineOfSight(location)
				if len(los) > max {
					max = len(los)
					target = Asteroid{
						location:    location,
						lineOfSight: los}
				}
			}
		}
	}
	return target
}

// treats "down" as zero and rotates anti-clockwise to reflect upside down grid
func angle(x1, y1, x2, y2 int) float64 {
	stepX, stepY := float64(x2-x1), float64(y2-y1)
	if stepX == 0 {
		if stepY < 0 {
			return 0
		} else {
			return math.Pi
		}
	}

	if stepX < 0 {
		if stepY <= 0 {
			return math.Pi*3/2 + math.Atan(stepY/stepX)
		} else {
			return math.Pi + math.Atan(-stepY/stepX)
		}
	} else {
		if stepY > 0 {
			return math.Pi + math.Atan(stepY/-stepX)
		} else {
			return math.Pi/2 + math.Atan(stepY/stepX)
		}
	}
}

func (a AsteroidMap) String() string {
	result := ""
	for _, v := range a.contents {
		result += fmt.Sprintf("%v\n", v)
	}
	return result
}

func part1() {
	asteroidMap := readAsteroidMap("input")
	asteroid := asteroidMap.findBestAsteroid()
	fmt.Println("Part 1: ", len(asteroid.lineOfSight)+1)
}

func part2() {
	asteroidMap := readAsteroidMap("input")
	asteroid := asteroidMap.findBestAsteroid()
	target := asteroid.getTargets()[199]
	fmt.Println("Part2: ", target.x*100+target.y)
}

func main() {
	part1()
	part2()
}
