package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHidden(t *testing.T) {
	assert.NotPanics(t, func() {
		_, err := Parse("../../test/data/dir1", "", false)
		assert.ErrorIs(t, err, EmptyErr)
	})
	assert.NotPanics(t, func() {
		_, err := Parse("../../test/data/dir2", "", false)
		assert.ErrorIs(t, err, EmptyErr)
	})
}
