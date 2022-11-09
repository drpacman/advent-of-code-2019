package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

type Image struct {
	layers [][]int
	width  int
	height int
}

func countEntries(layer []int, n int) int {
	entryCount := 0
	for _, e := range layer {
		if e == n {
			entryCount++
		}
	}
	return entryCount
}

func readImage(width, height int) Image {
	data, _ := ioutil.ReadFile("input")
	image := make([]int, len(data))
	for i, d := range data {
		image[i] = int(d) - 48
	}

	sliceLength := width * height
	layerCount := len(image) / sliceLength
	layers := make([][]int, layerCount)
	for j := 0; j < layerCount; j++ {
		layers[j] = image[j*sliceLength : (j+1)*sliceLength]
	}
	return Image{
		layers: layers,
		width:  width,
		height: height}
}

func (img *Image) getLayerWithLeast(n int) []int {
	minCount := math.MaxInt64
	var minLayer []int
	for _, layer := range img.layers {
		entryCount := countEntries(layer, n)
		if entryCount < minCount {
			minCount = entryCount
			minLayer = layer
		}
	}
	return minLayer
}

func (img *Image) printImage() {
	merged := make([]string, len(img.layers[0]))
	for i := 0; i < len(merged); i++ {
		for _, layer := range img.layers {
			if layer[i] != 2 {
				if layer[i] == 1 {
					merged[i] = "*"
				} else {
					merged[i] = " "
				}
				break
			}
		}
	}
	for i := 0; i < img.height; i++ {
		fmt.Println(merged[i*img.width : (i+1)*img.width])
	}
}

func part1(img Image) {
	layer := img.getLayerWithLeast(0)
	fmt.Printf("Part 1: %v\n", countEntries(layer, 1)*countEntries(layer, 2))
}

func part2(img Image) {
	fmt.Println("Part 2:")
	img.printImage()
}

func main() {
	image := readImage(25, 6)
	part1(image)
	part2(image)
}
