/**
	Implement a simple cache as a key-value pair
**/

package utils

import (
	"gotest.tools/assert"
	"testing"
)

var (
	testCache = NewCache()
)

func TestBasic(t *testing.T) {
	expected := []string{"foo", "bar", "foobar"}
	testCache.Set(1, expected[0])
	testCache.Set("1", expected[1])
	testCache.Set(2.0, expected[2])

	assert.Equal(t, testCache.Get(1), expected[0])
	assert.Equal(t, testCache.Get("1"), expected[1])
	assert.Equal(t, testCache.Pop(2.0), expected[2])
	assert.Equal(t, testCache.Get(2.0), nil)
}
