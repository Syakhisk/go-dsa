package knights_tour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Run("1x3 board", func(t *testing.T) {
		assert.False(t, Run(0, 0, 3, 1))
	})

	t.Run("2x5 board", func(t *testing.T) {
		assert.False(t, Run(0, 0, 2, 5))
	})

	t.Run("4x4 board", func(t *testing.T) {
		assert.False(t, Run(0, 0, 4, 4))
	})

	t.Run("5x5 board", func(t *testing.T) {
		assert.True(t, Run(0, 0, 5, 5))
	})

	t.Run("6x6 board", func(t *testing.T) {
		assert.True(t, Run(0, 0, 6, 6))
	})
}
