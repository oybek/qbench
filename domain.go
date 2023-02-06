package main

import "math"

type queen struct {
	x int
	y int
}

func (q queen) attacks(q2 queen) bool {
	return q.x == q2.x ||
		q.y == q2.y ||
		math.Abs(float64(q.x-q2.x)) == math.Abs(float64(q.y-q2.y))
}

func attack(qs []queen, q queen) bool {
	for _, c := range qs {
		if c.attacks(q) {
			return true
		}
	}
	return false
}
