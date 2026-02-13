package iterseq

import "iter"

// Filter iter.Seqの要素を選別する
// - predicate: ひとつの要素を受け取り、それを残すべきかどうか判断する関数
// - seq: 任意の型のiter.Seq
func Filter[T any](predicate func(x T) bool, seq iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for x := range seq {
			if predicate(x) {
				if !yield(x) {
					return
				}
			}
		}
	}
}
