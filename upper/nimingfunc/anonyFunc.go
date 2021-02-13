package nimingfunc

func Anoy(n3 int,n4 int) int {
	res := func (n1 int,n2 int) int {
		return n1 + n2
	}(n3,n4)
	return res
}

func Anoy1(n3 int,n4 int) int {
	res := func (n1 int,n2 int) int {
		return n1 - n2
	}
	return res(n3,n4)
}