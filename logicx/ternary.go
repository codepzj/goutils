package logicx

// Ternary 三元表达式
// 如果b为true, 返回expr1
// 如果b为false, 返回expr2
func Ternary[T any](b bool, expr1 T, expr2 T) T {
	if b {
		return expr1
	}
	return expr2
}
