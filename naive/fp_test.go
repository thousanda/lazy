package naive

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	t.Run("フィルタリングできる", func(t *testing.T) {
		// Arrange
		list := []int64{1, 2, 3, 4, 5}
		isEven := func(x int64) bool {
			return x%2 == 0
		}

		// Act
		result := Filter(list, isEven)

		require.Equal(t, []int64{2, 4}, result)
	})
}
