package slicex

// Map 对切片的各个元素进行批量操作
// 通过传入函数对切片元素进行计算并返回
func Map[T any](s []T, fn func(element T) T) []T {
	for i, v := range s {
		s[i] = fn(v)
	}
	return s
}

// Push 往切片末尾追加一个元素
func Push[T any](s []T, element T) []T {
	return append(s, element)
}

// Pop 切片弹出特定元素
func Pop[T any](s []T) ([]T, T) {
	l := len(s)
	return s[:l-1], s[l-1]
}
