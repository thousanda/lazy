package naive

import (
	"strconv"
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
		result := Filter(isEven, list)
		// Assert
		require.Equal(t, []int64{2, 4}, result)
	})
}

func TestMap(t *testing.T) {
	t.Run("すべての要素に関数を適用できる", func(t *testing.T) {
		// Arrange
		double := func(x int64) int64 {
			return 2 * x
		}
		list := []int64{1, 2, 3}
		// Act
		result := Map(double, list)
		// Assert
		require.Equal(t, []int64{2, 4, 6}, result)
	})
	t.Run("入力と出力の型が異なっていてもいい", func(t *testing.T) {
		// []int -> []string
		// Arrange
		toString := strconv.Itoa
		list := []int{1, 2, 3}
		// Act
		result := Map(toString, list)
		// Assert
		require.Equal(t, []string{"1", "2", "3"}, result)
	})
}
