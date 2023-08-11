package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandomUppercase(t *testing.T) {
	output := RandomUppercase(5)

	require.Equal(t, 5, len(output))
}
