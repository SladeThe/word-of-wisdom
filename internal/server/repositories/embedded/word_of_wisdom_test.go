package embedded

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordOfWisdom_OneRandom(t *testing.T) {
	r := NewWordOfWisdom()

	for i := 0; i < 10; i++ {
		word, err := r.OneRandom()
		require.NoError(t, err)
		require.NotEmpty(t, word.Text)
	}
}
