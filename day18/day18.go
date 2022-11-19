package main

import (
	"advent-of-code-2019/files"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type RuneSlice []rune

func (s RuneSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s RuneSlice) Len() int           { return len(s) }
func (s RuneSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type Point struct {
	x, y int
}

type Key struct {
	p     Point
	value rune
}

func (k Key) String() string {
	return string(k.value)
}

type Maze struct {
	contents [][]rune
}

func (m Maze) String() string {
	var sb strings.Builder
	for _, row := range m.contents {
		for _, value := range row {
			sb.WriteString(string(value))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func ReadMaze(filename string) (Maze, Point) {
	d := files.ReadLines(filename)
	contents := make([][]rune, len(d))
	var start Point
	for y, line := range d {
		contents[y] = []rune(line)
		for x, e := range contents[y] {
			if e == '@' {
				contents[y][x] = '.'
				start = Point{x, y}
			}
		}
	}
	m := Maze{contents}
	return m, start
}

func (m Maze) KeyLocations() []Key {
	var location []Key
	for y, row := range m.contents {
		for x, value := range row {
			if unicode.IsLower(value) {
				location = append(location, Key{
					p:     Point{x, y},
					value: value})
			}
		}
	}
	return location
}

type GuardedPath struct {
	path         []Point
	requiredKeys []rune
}

type Path struct {
	path          []Point
	collectedKeys []rune
}

func (m Maze) ProcessMaze(startPos Point) []Point {
	// get the shortest path between each pair of keys
	var pathsBetweenKeys = m.CalculatePathsBetweenKeys()

	// get the intial paths
	var paths = m.getInitialPaths(startPos)
	var allKeys = m.KeyLocations()
	var updatedPaths []Path
	var solutions []Path
	// try each of them
	for {
		for _, path := range paths {
			for _, updatedPath := range m.WalkPath(pathsBetweenKeys, path) {
				sort.Sort(RuneSlice(updatedPath.collectedKeys))
				// check if we found all the keys
				if len(updatedPath.collectedKeys) == len(allKeys) {
					solutions = append(solutions, updatedPath)
				} else {
					updatedPaths = append(updatedPaths, updatedPath)
				}
			}
		}
		if len(updatedPaths) == 0 {
			break
		}
		fmt.Printf("\nPaths %v", len(updatedPaths))
		var pathsMap = make(map[string]Path)
		for _, path := range updatedPaths {
			var currentPoint = path.path[len(path.path)-1]
			var currentKey = m.contents[currentPoint.y][currentPoint.x]
			var key = fmt.Sprintf("%v-%v", string(path.collectedKeys), currentKey)
			if existingPath, ok := pathsMap[key]; ok {
				if len(existingPath.path) > len(path.path) {
					pathsMap[key] = path
				}
			} else {
				pathsMap[key] = path
			}
		}
		updatedPaths = nil
		for _, path := range pathsMap {
			updatedPaths = append(updatedPaths, path)
		}
		paths = updatedPaths
		updatedPaths = nil
	}
	// walk the solutions and find the shortest
	var solution []Point
	for _, candidate := range solutions {
		if solution == nil || len(candidate.path) < len(solution) {
			solution = candidate.path
		}
	}
	return solution
}

func (m Maze) WalkPath(pathsBetweenKeys map[string][]GuardedPath, path Path) []Path {
	// get the shortest path between each pair of keys
	var tailPos = len(path.path) - 1
	var currentKey = m.contents[path.path[tailPos].y][path.path[tailPos].x]
	var keysRemaining []rune
	var paths []Path
	// see what keys still need collecting
	for _, targetKey := range m.KeyLocations() {
		var found = false
		for _, key := range path.collectedKeys {
			if targetKey.value == key {
				found = true
				break
			}
		}
		if !found {
			keysRemaining = append(keysRemaining, targetKey.value)
		}
	}
	for _, key := range keysRemaining {
		var entryKey = EntryKey(currentKey, key)
		var guardedPaths, ok = pathsBetweenKeys[entryKey]
		if !ok {
			panic(fmt.Sprintf("\nMissing path %v - current path is %v", entryKey, path))
		}
		// if we have all the keys already for those
		// required by the path then we can take it
		var haveAllRequiredKeys = true
		var chosenPath []Point
		for _, guardedPath := range guardedPaths {
			for _, requiredKey := range guardedPath.requiredKeys {
				var hasKey = false
				for _, collectedKey := range path.collectedKeys {
					if collectedKey == requiredKey {
						hasKey = true
						break
					}
				}
				if !hasKey {
					haveAllRequiredKeys = false
					break
				}
			}
			if haveAllRequiredKeys {
				if chosenPath == nil || len(guardedPath.path) < len(chosenPath) {
					chosenPath = guardedPath.path
				}
			}
		}

		if chosenPath != nil {
			// update the collected keys along with the extended path
			var combinedPath = make([]Point, len(path.path))
			copy(combinedPath, path.path)
			if path.path[len(path.path)-1] == chosenPath[0] {
				for i := 1; i < len(chosenPath); i++ {
					combinedPath = append(combinedPath, chosenPath[i])
				}
			} else {
				for i := len(chosenPath) - 2; i >= 0; i-- {
					combinedPath = append(combinedPath, chosenPath[i])
				}
			}
			newKeys := make([]rune, len(path.collectedKeys))
			copy(newKeys, path.collectedKeys)
			newKeys = append(newKeys, key)
			newPath := Path{combinedPath, newKeys}
			paths = append(paths, newPath)
		}
	}
	return paths
}

func EntryKey(keyA rune, keyB rune) string {
	if keyA < keyB {
		return fmt.Sprintf("%s-%s", string(keyA), string(keyB))
	}
	return fmt.Sprintf("%s-%s", string(keyB), string(keyA))
}

func (m Maze) getInitialPaths(startPos Point) []Path {
	var initialPaths []Path
	for _, key := range m.KeyLocations() {
		var paths = m.GetAllPathsToKey(startPos, key)
		// ignore path if it is blocked
		for _, path := range paths {
			var isBlocked = false
			for _, point := range path {
				if unicode.IsUpper(m.contents[point.y][point.x]) {
					isBlocked = true
					break
				}
			}
			if !isBlocked {
				// its reachable, add the path
				initialPaths = append(initialPaths, Path{path, []rune{key.value}})
			}
		}
	}
	return initialPaths
}

func (m Maze) CalculatePathsBetweenKeys() map[string][]GuardedPath {
	var pathsMap = make(map[string][]GuardedPath)
	// walk all the keys
	for i, keyA := range m.KeyLocations() {
		for j, keyB := range m.KeyLocations() {
			if i != j {
				var entryKey = EntryKey(keyA.value, keyB.value)

				// if we don't already have a shortest path for the key pair, calculate it
				if _, ok := pathsMap[entryKey]; !ok {
					fmt.Printf("\nFinding path between keys %v", entryKey)
					var paths = m.GetAllPathsToKey(keyA.p, keyB)
					fmt.Printf("\nFound %v paths between keys %v", len(paths), entryKey)
					var guardedPaths []GuardedPath
					for _, path := range paths {
						var requiredKeys []rune
						// gather path and get doors on that path
						for _, point := range path {
							if unicode.IsUpper(m.contents[point.y][point.x]) {
								requiredKeys = append(requiredKeys, unicode.ToLower(m.contents[point.y][point.x]))
							}
						}
						guardedPaths = append(guardedPaths, GuardedPath{path, requiredKeys})
					}
					pathsMap[entryKey] = guardedPaths
				}
			}
		}
	}

	return pathsMap
}

func (m Maze) GetAllPathsToKey(point Point, keyB Key) [][]Point {
	var updatedPaths [][]Point
	var paths = [][]Point{{point}}
	var allPaths [][]Point
	for len(paths) > 0 {
		for _, path := range paths {
			var extendedPaths, result = m.PathsToKey(path, keyB)
			if result != nil {
				allPaths = append(allPaths, result)
			} else {
				updatedPaths = append(updatedPaths, extendedPaths...)
			}
		}
		paths = updatedPaths
		updatedPaths = nil
	}
	return allPaths
}

func (m Maze) PathsToKey(segment []Point, key Key) ([][]Point, []Point) {
	var visited = make(map[Point]bool)
	for _, p := range segment {
		visited[p] = true
	}
	var currentPos = segment[len(segment)-1]
	var candidates = [4]Point{
		{currentPos.x - 1, currentPos.y},
		{currentPos.x + 1, currentPos.y},
		{currentPos.x, currentPos.y - 1},
		{currentPos.x, currentPos.y + 1}}
	var paths [][]Point
	for _, c := range candidates {
		if visited[c] {
			continue
		}
		var contents = m.contents[c.y][c.x]
		if contents != '#' {
			newSegment := make([]Point, len(segment))
			copy(newSegment, segment)
			if contents == key.value {
				// we found the shortest path, return it
				return nil, append(newSegment, c)
			} else {
				paths = append(paths, append(newSegment, c))
			}
		}
	}
	return paths, nil
}

func main() {
	m, start := ReadMaze("input")
	fmt.Println(m)
	var shortestPath = m.ProcessMaze(start)
	fmt.Printf("\nShortest Path is %v", len(shortestPath)-1)
}
