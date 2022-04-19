package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAbsolute(t *testing.T) {

	t.Run("sub negative test case", func(t *testing.T) {
		c := Absolute(-5)
		assert.Equal(t, 5, c)
	})

	t.Run("sub positive test case", func(t *testing.T) {
		c := Absolute(5)
		assert.Equal(t, 5, c, "expect 5 got %d", c)
	})

}

func TestAdd_Table(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{name: "negative and negative", a: -1, b: -2, expected: -3},
		{name: "positive and positive", a: 1, b: 2, expected: 3},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Add(tc.a, tc.b)
			assert.Equal(t, tc.expected, c)
		})

	}
}

func TestAdd(t *testing.T) {
	//if short test mode argument passeed, skip this test
	if testing.Short() {
		t.Skip()
	}
	<-time.After(5 * time.Second)
	t.Run("sub n + n test case", func(t *testing.T) {
		c := Add(-1, -2)
		assert.Equal(t, -3, c)
	})
	t.Run("sub p + p test case", func(t *testing.T) {
		c := Add(1, 2)
		assert.Equal(t, 3, c)
	})
}
