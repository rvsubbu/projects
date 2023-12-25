package problem

import (
	"gotest.tools/assert"
	"testing"
)

func TestMissingZero(t *testing.T) {
	test := []int{1,2,3,4,5}
	expected := 0
	testHelper(t, test, expected)
}

func TestMissingANumber(t *testing.T) {
	test := []int{0,1,3,4,5}
	expected := 2
	testHelper(t, test, expected)
}

func TestMissingANumberOutOfSequence(t *testing.T) {
	test := []int{1,0,3,4,5}
	expected := 2
	testHelper(t, test, expected)
}

func TestAllPresent(t *testing.T) {
	// Expect the first number not in array
	test := []int{0,1,2,3,4,5}
	expected := 6
	testHelper(t, test, expected)
}

func testHelper(t *testing.T, data []int, expected int) {
	assert.Equal(t, useMap(data), expected)
	assert.Equal(t, useSort(data), expected)
}
