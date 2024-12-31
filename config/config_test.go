package config

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	c, err := Load()
	require.NoError(t, err)
	require.NotNil(t, c)
	assert.Equal(t, "3200", c.HttpConfig.Port)
}
