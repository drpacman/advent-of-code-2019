package main

import (
	"files"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
)

type LCGParams struct {
	a int64
	b int64
	n int64
}

func ReadInputAsLinearCongruentialGenerator(path string, n int64) LCGParams {
	lines := files.ReadLines(path)
	cut, _ := regexp.Compile("cut (-?[0-9]*)")
	inc, _ := regexp.Compile("deal with increment ([0-9]*)")

	// start with identity LCG
	instruction := LCGParams{1, 0, n}
	for _, line := range lines {
		switch {
		case cut.MatchString(line) == true:
			// corresponds to formula f(X) = x - val % n
			val, _ := strconv.Atoi(cut.FindStringSubmatch(line)[1])
			instruction = instruction.composeLCG(LCGParams{1, int64(-val), n})
		case inc.MatchString(line) == true:
			// corresponds to formula f(X) = val * x % n
			val, _ := strconv.Atoi(inc.FindStringSubmatch(line)[1])
			instruction = instruction.composeLCG(LCGParams{int64(val), 0, n})
		default:
			// corresponds to formula f(X) = -x- 1 % n
			instruction = instruction.composeLCG(LCGParams{-1, int64(-1), n})
		}
	}
	return instruction
}

func (l *LCGParams) composeLCG(params LCGParams) LCGParams {
	return LCGParams{l.a * params.a % l.n, (l.b*params.a + params.b) % l.n, l.n}
}

func (l *LCGParams) posForCard(x int64) int64 {
	result := big.NewInt(l.a)
	result.Mul(result, big.NewInt(x))
	result.Add(result, big.NewInt(l.b))
	result.Mod(result, big.NewInt(l.n))
	return result.Int64()
}

func (l *LCGParams) cardForPos(pos int64) int64 {
	// Inverse of F(X) = Ax + B % N
	// x = (F(X) - B) /  A % N
	result := big.NewInt(pos)
	result.Sub(result, big.NewInt(l.b))
	result.Mul(result, big.NewInt(modInverse(l.a, l.n)))
	result.Mod(result, big.NewInt(l.n))
	return result.Int64()
}

func modInverse(x, n int64) int64 {
	return modPow(x, n-2, n)
}

func modPow(x, n, m int64) int64 {
	if n == 0 {
		return int64(1)
	}
	xSquaredModM := new(big.Int)
	xSquaredModM.Mul(big.NewInt(x), big.NewInt(x))
	xSquaredModM.Mod(xSquaredModM, big.NewInt(m))

	if n%2 == 1 {
		result := big.NewInt(x)
		result.Mul(result, big.NewInt(modPow(xSquaredModM.Int64(), (n-1)/2, m)))
		result.Mod(result, big.NewInt(m))
		return result.Int64()
	} else {
		return modPow(xSquaredModM.Int64(), n/2, m) % m
	}
}

func (l *LCGParams) applyNTimes(t int64) LCGParams {
	// self application of ax + b % n t times is
	// l.a^t * x + (1-l.a^t)/(1-a) % l.n
	aToThePowerT := big.NewInt(modPow(l.a, t, l.n))

	invDenominator := big.NewInt(1)
	invDenominator.Sub(invDenominator, big.NewInt(l.a))
	invDenominator.ModInverse(invDenominator, big.NewInt(l.n))

	b := big.NewInt(1)
	b.Sub(b, aToThePowerT)
	b.Mul(b, big.NewInt(l.b))
	b.Mul(b, invDenominator)
	b.Mod(b, big.NewInt(l.n))
	return LCGParams{aToThePowerT.Int64(), b.Int64(), l.n}
}

func part1() {
	instruction := ReadInputAsLinearCongruentialGenerator("input", 10007)
	result := instruction.posForCard(2019)
	fmt.Printf("Part 1 - %v\n", result)
}

func part2() {
	instruction := ReadInputAsLinearCongruentialGenerator("input", 119315717514047)
	instruction = instruction.applyNTimes(101741582076661)
	result := instruction.cardForPos(2020)
	fmt.Printf("Part 2 - %v\n", result)
}

func main() {
	part1()
	part2()
}
