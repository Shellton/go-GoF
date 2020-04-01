package builder

type StringBuilder struct {
}

func (sb *StringBuilder) Append(arg interface{}) *StringBuilder {
	return sb
}

func (sb *StringBuilder) Insert(str string) *StringBuilder {
	return sb
}
