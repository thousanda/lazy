package iterseq

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	t.Run("フィルタリングできる", func(t *testing.T) {
		// Arrange
		isEven := func(x int64) bool {
			return x%2 == 0
		}
		list := []int64{1, 2, 3, 4, 5}
		// Act
		result := slices.Collect(Filter(isEven, slices.Values(list)))
		// Assert
		require.Equal(t, []int64{2, 4}, result)
	})
}
