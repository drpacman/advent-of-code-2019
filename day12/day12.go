package main

import (
	"fmt"
	"advent-of-code-2019/numbers"
)

type Dimensions struct {
	x, y, z int
}

type Moon struct {
	position Dimensions
	velocity Dimensions
}

func (moon *Moon) applyGravity(other Moon) {
	compare := func(a, b int) int {
		if a < b {
			return 1
		} else if a > b {
			return -1
		}
		return 0
	}
	moon.velocity.x += compare(moon.position.x, other.position.x)
	moon.velocity.y += compare(moon.position.y, other.position.y)
	moon.velocity.z += compare(moon.position.z, other.position.z)
}

func (moon *Moon) getTotalEnergy() int {
	potential := numbers.AbsInt(moon.position.x) + numbers.AbsInt(moon.position.y) + numbers.AbsInt(moon.position.z)
	kinetic := numbers.AbsInt(moon.velocity.x) + numbers.AbsInt(moon.velocity.y) + numbers.AbsInt(moon.velocity.z)
	return potential * kinetic
}

func runTimeStep(moons []*Moon) {
	for _, m1 := range moons {
		for _, m2 := range moons {
			if m1 != m2 {
				m1.applyGravity(*m2)
			}
		}
	}

	for _, moon := range moons {
		moon.position.x += moon.velocity.x
		moon.position.y += moon.velocity.y
		moon.position.z += moon.velocity.z
	}
}

func initMoons() []*Moon {
	return []*Moon{
		&Moon{Dimensions{-4, -9, -3}, Dimensions{0, 0, 0}},
		&Moon{Dimensions{-13, -11, 0}, Dimensions{0, 0, 0}},
		&Moon{Dimensions{-17, -7, 15}, Dimensions{0, 0, 0}},
		&Moon{Dimensions{-16, 4, 2}, Dimensions{0, 0, 0}}}
}

func greatestCommonDivisor(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func leastCommonMultiple(a, b int64) int64 {
	return a * b / greatestCommonDivisor(a, b)
}

func calculateLeastCommonMultiplier(x, y, z int64) int64 {
	return leastCommonMultiple(z, leastCommonMultiple(x, y))
}

func calculatePeriods(moons []*Moon) (int64, int64, int64) {
	xs := make([]int, len(moons))
	ys := make([]int, len(moons))
	zs := make([]int, len(moons))
	var periodX, periodY, periodZ int64 = 0, 0, 0
	for i, m := range moons {
		xs[i] = (*m).position.x
		ys[i] = (*m).position.y
		zs[i] = (*m).position.z
	}
	var step int64 = 1
	for periodX == 0 || periodY == 0 || periodZ == 0 {
		runTimeStep(moons)
		matchedX := true
		matchedY := true
		matchedZ := true
		for i, m := range moons {
			if xs[i] != (*m).position.x || (*m).velocity.x != 0 {
				matchedX = false
			}
			if ys[i] != (*m).position.y || (*m).velocity.y != 0 {
				matchedY = false
			}
			if zs[i] != (*m).position.z || (*m).velocity.z != 0 {
				matchedZ = false
			}
		}
		if matchedX && periodX == 0 {
			periodX = step
		}
		if matchedY && periodY == 0 {
			periodY = step
		}
		if matchedZ && periodZ == 0 {
			periodZ = step
		}
		step = step + 1
	}
	return periodX, periodY, periodZ
}

func part1() {
	moons := initMoons()
	for i := 0; i < 1000; i++ {
		runTimeStep(moons)
	}
	total := 0
	for _, m := range moons {
		total += m.getTotalEnergy()
	}
	fmt.Println("Part1: ", total)
}

func part2() {
	moons := initMoons()
	x, y, z := calculatePeriods(moons)
	fmt.Println("Part2:", calculateLeastCommonMultiplier(x, y, z))
}

func main() {
	part1()
	part2()
}
