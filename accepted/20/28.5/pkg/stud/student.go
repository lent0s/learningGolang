package stud

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent() *Student {
	return &Student{}
}
