package main

import "testing"

func testCritiera(criteria func(int) bool, value int, expected bool, t *testing.T) {
	if criteria(value) != expected {
		t.Errorf("%v does not meet criteria value of %v", value, expected)
	}
}
func TestMeetsCriteriaPartA(t *testing.T) {
	testCritiera(meetsCriteriaPartA, 111111, true, t)
	testCritiera(meetsCriteriaPartA, 223450, false, t)
	testCritiera(meetsCriteriaPartA, 123789, false, t)
}

func TestMeetsCriteriaPartB(t *testing.T) {
	testCritiera(meetsCriteriaPartB, 112233, true, t)
	testCritiera(meetsCriteriaPartB, 123444, false, t)
	testCritiera(meetsCriteriaPartB, 111122, true, t)
	testCritiera(meetsCriteriaPartB, 777999, false, t)
}
