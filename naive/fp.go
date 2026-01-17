package naive

// Filter リストの要素を選別する
// - predicate: ひとつの要素を受け取り、それを残すべきかどうか判断する関数
// - list: 任意の型のスライス
func Filter[T any](predicate func(x T) bool, list []T) []T {
	result := make([]T, 0, len(list))
	for _, x := range list {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return result
}

// Map リストのすべての要素に同じ関数を適用する
func Map[S any, T any](apply func(x S) T, list []S) []T {
	result := make([]T, 0, len(list))
	for _, x := range list {
		result = append(result, apply(x))
	}
	return result
}
