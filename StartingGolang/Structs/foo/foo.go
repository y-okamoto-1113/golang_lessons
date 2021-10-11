package foo

type T struct {
	Field1 int // 公開
	field2 int // 非公開
}

func (t *T) Method1() int {
	return t.Field1
}

func (t *T) method2() int {
	return t.field2
}
