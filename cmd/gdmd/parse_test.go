package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHidden(t *testing.T) {
	assert.NotPanics(t, func() {
		_, err := Parse("../../test/data/hidden", "", false)
		assert.ErrorIs(t, err, ErrEmpty)
	})
	assert.NotPanics(t, func() {
		_, err := Parse("../../test/data/hidden/dir", "", false)
		assert.ErrorIs(t, err, ErrEmpty)
	})
}

func TestParseReceiver(t *testing.T) {
	assert.NotPanics(t, func() {
		p, _ := Parse("../../test/data/receivers", "", false)
		assert.NotEmpty(t, p.Types)
		assert.NotEmpty(t, p.Types[0].Methods)
	})
}
