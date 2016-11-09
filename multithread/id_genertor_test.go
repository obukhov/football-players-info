package multithread

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIdGenerator_Current(t *testing.T) {
	generator := NewIdGenerator(1, 6, 10)

	assert.Equal(t, 1, generator.Current())
	assert.True(t, generator.GenerateNext())

	assert.Equal(t, 7, generator.Current())
	assert.False(t, generator.GenerateNext())
}
