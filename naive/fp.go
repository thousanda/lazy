package naive

// Filter リストの要素を選別する
// - list: 任意の型のスライス
// - predicate: listの要素を受け取り、それを残すべきかどうか判断する関数
func Filter[T any](list []T, predicate func(x T) bool) []T {
	result := make([]T, 0, len(list))
	for _, x := range list {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return result
}
