package aggregator

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayersCollection(t *testing.T) {
	playerCollection := NewPlayersCollection()

	buf := new(bytes.Buffer)

	playerCollection.Add("Harry Potter", 21, "Barcelona")
	playerCollection.Add("Alex Klar", 25, "Barcelona")
	playerCollection.Add("Boris Blade", 19, "Barcelona")
	playerCollection.Add("Boris Blade", 19, "Spain")

	playerCollection.Output(buf)

	expected := `1. Alex Klar; 25; Barcelona
2. Boris Blade; 19; Barcelona, Spain
3. Harry Potter; 21; Barcelona
`
	assert.Equal(t, expected, buf.String())
}
